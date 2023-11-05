# Types of Windows Drivers

There are two basic types of Microsoft Windows drivers:

* *User-mode drivers* execute in user mode, and they tipically provide
an interface between a Win32 application and kernel-mode drivers or
other operating system components.

For example, in Windows Vista, all printer drivers execute in user
mode.

* *Kernel-mode drivers* execute in kernel mode as part of the
executive, which consists of kernel-mode operating system components
that manage I/O, Plug and Play memory, processes and threads,
security, and so on. Kernel-mode drivers are typically layered.
Generally, higher-level drivers typically receive data from
applications, filter the data, and pass it to a lower-level driver
that supports device functionality.

Some kernel-mode drivers are also *WDM drivers*, which conform to
the Windows Driver Model. All WDM drivers support Plug and Play,
and power managements. WDM drivers are source-compatible (but not
binary-compatible) across Windows 98/Me and Windows 2000 and later
operating systems.

Like the operating system itself, kernel-mode drivers are implemented
as discrete, modular components that have a well-defined set of
required functionalities. All kernel-mode drivers supply a set of
system-defined standard driver routines.

This is a division of kernel-mode drivers:

* Highest-level drivers:
    * File system drivers (FAT, NTFS, CDFS)
* Intermediate drivers
    * Legacy virtual disk or mirror drivers
    * Legacy class drivers
    * Optional PnP upper-level filter drivers
    * PnP function drivers WDM class/miniclass driver pair
    * Optional PnP lower-lever filter drivers
    * PnP software bus drivers
    * WDM software bus drivers
* Lowest-level drivers
    * Legacy device drivers
    * PnP hardware bus drivers

As shown in the figure, theree are three basic types of kernel-mode
drivers in a driver stack: highest-level, intermediate, and
lowest-level. Each type differs only slightly in structure but
greatly in functionality:

1. Highest-level drivers include file system drivers include file
system drivers (FSDs) that support file systems, such as:

* NTFS
* File allocation table (FAT)
* CD-ROM file system (CDFS)

Highest-level drivers always depend on support from underlying
lower-level drivers, such as intermediate-level function drivers and
lowest-level hardware bus drivers.

2. Intermediate drivers, such as virtual disk, mirror, or device-
type-specific class driver. Intermediate drivers depend on support
from underlying lower-level drivers. Intermediate drivers are
subdivided further as follows:

* Function drivers: control specific peripheral devices on an I/O bus.
* Filter drivers insert themselves above or below function drivers.
* Software bus drivers present a set of child devices to which still
higher-level class, function, or filter drivers can attach themselves.

For example, a driver that controls a multifunction adapter with an
on-board set of heterogeneous devices is a software bus drivers.

* Any system-supplied class driver that exports a system-defined
class/miniclass interface is, in effect, an intermediate driver
with one or more linked miniclass drivers (sometimes called
minidrivers). Each linked class/minidrivers pair provides
functionality that is equivalent to that of a function driver or
a software bus drivers.

3. Lowest-level drivers control an I/O bus to which peripheral
devices are connected. Lowest-level drivers do not depend on
lower-level drivers.

* Hardware bus drivers are system-supplied and usually control
dynamically configurable I/O buses.

Hardware bus drivers work with Plug and Play manager to configure
and reconfigure system hardware resources, for all child devices
that are connected to I/O buses that the driver controls. These
hardware resources include mappings for device memory and interrupt
requests (IRQs).

* Legacy drivers that directly control a physical device are
lowest-level drivers.
