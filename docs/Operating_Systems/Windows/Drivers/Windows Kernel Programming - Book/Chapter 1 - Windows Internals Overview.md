# Chapter 1: Windows Internals Overview

In this chapter the most important concepts in the internal workings
of Windows are described.

## Process

A process is a containment and management object that represents a
running instance of a program. I want to let this sink in. A process
is a representation of a running instance of a program.
A program by itself might not be running. It's static and in the
Windows world it may be an ".exe" file. When that program starts
running Windows keeps track of that running instance with a process.

The term "process runs" which is used fairly often, is inaccurate.
Processes don't run - processes manage. Threads are the ones that
execute code and technically run. From a high-level perspective, a
process owns the following:

* An executable program, which contains the initial code and data
used to execute code within the process.
* A private virtual address space, used for allocating memory for
whatever purposes the code within the process needs it.
* A primary token, which is an object that stores the default
security context of the process, used by threads executing code
within the process (unless a thread assumes a different token
by using impersonation.)
* A private handle table to execute object, such as events,
semaphores, and files.
* One or more threads of execution. A normal user mode process is
created with one thread (executing the classic main/WinMain function).
A user mode process without threads is mostly useless and under normal
circumstances will be destroyed by the kernel.

So a process has
|_________________> an executable image/file (contains code and data)
|_________________> a primary token (object storing default security context)
|_________________> a private virtual address space
|_________________> a handle table
|_________________> one or more threads of execution

A process is uniquely identified by its Process ID, which remains unique
as long as the kernel process object exists. Once it's destroyed, the
same ID may be reused for a new process. It's important to realize that
the executable file itself is not a unique identifier of the process.
For example, there may be five instances of notepad.exe running at the
same time. Each process has its own address space, its own threads, its
own handle table, its own unique process ID, etc. All those five
processes are using the same image file (notepad.exe) as their initial
code and data.

### Virtual Memory

Every process has its own virtual, private, linear address space. This
address spaces starts out empty (or close to empty, since the executable
image and NtDll.Dll are the first to be mapped, followed by more
subsystem DLLs). Once execution of the main (first) thread begins,
memory is likely to be allocated, more DLLs loaded, etc. This address
space is private, which means other processes cannot access it directly.
The address space range starts at zero (although technically the first
64KB of address cannot be allocated or used in any way), and goes all
the way to a maximum which depends on the process "bitness" (32 or 64 bit)
and the operating system "bitness" as follows:

* For 32-bit processes on 32-bit Windows systems, the process address
space size is 2 GB by default.
* For 32-bit processes on 32-bit Windows systems that use the increase
user address space setting (LARGEADDRESSWARE flag in the Portable
Executable header), that process address space size can be as large as
3 GB (depending on the exact setting). To get the extended address
space range, the executable from which the process was created must
have been marked with the LARGEADDRESSWARE linker flag in its header.
If it was not, it would still be limited to 2 GB.
* For 64-bit processes (which can only run on 64-bit Windows systems),
the address space size is 8 TB (Windows 8 and earlier) or 128 TB
(Windows 8.1 and later).
* For 32-bit processes on a 64-bit Windows system, the address space
size is 4 GB if the executable image is linked with the LARGEADDRESSWARE
flag. Otherwise, the size remains at 2 GB.

> INFO: The requirement of the LARGEADDRESSWARE flag stems from the fact
that a 2 GB address range requires 31 bits only, leaving the most
significant bit (MSB) free for application use. Specifying this flag
indicates that the program is not using bit 31 for anything and so setting
that bit to 1 (which would happen for address larger than 2 GB) is not an
issue.

Each process has its own address space, which makes any process address
relative, rather than absolute. For example, when trying to determine
what lies in address 0x20000, the address itself is not enough; the process
to which this address relates must be specified.

The memory itself is called virtual, which means there is an indirect
relationship between an address range and the exact location where it's
found in physical memory (RAM). A buffer within a process may be mapped
to physical memory, or it may temporarily reside in a file (such as a
page file). The term virtual refers to the fact that from an execution
perspective, there is no need to know if the memory about to be
accessed is in RAM or not; if the memory is indeed mapped to RAM, the
CPU will access the data directly. If not, the CPU will raise a page
fault exception that will cause the memory manager's page fault handler
to fetch the data from the appropriate file, copy it to RAM, make the
required changes in the page table entries that map the buffer, and
instruct the CPU to try again.

So a process' virtual memory address space may be mapping to RAM, or to
disk. There is an inderect relationship. I understand that this means
that the virtual memory is an abstraction of memory for the processes.

The unit of memory management is called a page. Every attribute related
to memory is always at a page's granularity, such as its protection.
Normal (sometimes called small) page size is 4 KB on all Windows
supported architectures.

Apart from the normal (small) page size, Windows also supports large
pages. The size of a large page is 2 MB (x86/x64/ARM64) and 4 MB
(ARM). This is based using the Page Directory Entry (PDE) to map to
map the large page without using a page table. This results in
quicker translation, but most importantly better use the Translation
Lookaside Buffer (TLB) - a cache of recently translated pages
maintained by the CPU. In the case of a large page, a single TLB entry
is able to map significanly more memory than a small page.
