# Message and Message Queues

This section describes messages and message queues and how to use them in your
applications.

## About Messages and Message Queues

Unlike MS-DOS-based applications, Windows-based applications are event-driven.
They do not make explicit function calls (such as C run-time library calls) to
obtain input. Instead, they wait for the system (OS) to pass input to them.

The system passes all input for an application to the various windows in the
application. Each window has a function, called a window procedure, that the
system calls whenever it has input for the window.
