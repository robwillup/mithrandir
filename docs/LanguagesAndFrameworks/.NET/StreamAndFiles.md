# Streams and Files in C#

## System.IO Namespace

The System.IO namespace contains types that allow reading and writing to files and data streams,
and types that provide basic file and directory support.

## Recap: Streams in C#

> `Stream` is the abstract base class of all streams. 
> **A stream is an abstraction of a sequence of bytes**, 
> such as a file, an input/output device, an interprocess communication pipe, or a TCP/IP socket.

Streams involve three fundamental operations

* You can read from streams. Reading is the transfer of data from a stream into a data structure, 
such as an array of bytes.
* You can write to streams. Writing is the transfer of data from a data structure into a stream.
* Streams can support seeking. Seeking refers to querying and modifying the current position within
a stream. Seeking capability depends on the kind of backing store a stream has. For example, network
streams have no unified concept of a current position, and therefore typically do not support seeking.

## The File class

The File class is contained in the System.IO namespace and it provides static methods for the creation,
 copying, deletion, moving, and opening of a single file, and aids in the creation of `FileStream` objects.
 
 ## References
 
 [C# File class documentation](https://docs.microsoft.com/en-us/dotnet/api/system.io.file?view=netcore-3.1)
 [C# System.IO namespace documentation](https://docs.microsoft.com/en-us/dotnet/api/system.io.file?view=netcore-3.1)
 
 
