# What is a Driver?

A driver lets the operating system and a device communicate with each other.

A driver is any software component that observes or participates in the
communication between the operating system and a device.

Application <-> OS <-> Driver <-> Device

* not all drivers have to written by the companies that designed the device.
* not all drivers communicate directly with a device. Some drivers just
manipulate the request and pass it along.
* Some filter drivers observe and record information about I/O requests but
do not actively participate in them.
* There are also *software drivers* which are not associated with a hardware
device. Software drivers always run in kernel mode. The main reason for
writing a software driver is to gain access to protected data that is
available only in kernel mode.