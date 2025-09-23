# GetAdaptersAddresses function

The **GetAdaptersAddresses** function retrieves the addresses associated
with the adapters on the local computer.

## Parameters

`Family`

The address family of the address to retrieve:

|Value|Meaning|
|:-|:-|
|AF_UNSSPEC / 0| Return both IPv4 and IPv6 addresses associated with adapters with IPv4 or IPv6 enabled|
|AF_INET / 2| Return only IPv4 addresses associated with adapters with IPv4 enabled|
|AF_INET6 / 23 |Return only IPv6 addresses associated with adapters with IPv6 enabled|

`Flags`

The type of addresses to retrieve. The possible values are defined in the *Iptypes.h* header file.

## Remarks

The **GetAdaptersAddress** function can retrieve information for IPv4
and IPv6 addresses.

Addresses are returned as a linked list of `IP_ADAPTER_ADDRESSES`
structures in the buffer pointed to by the *AdapterAddresses*
parameter. The application that calls the **GetAdaptersAddresses**
function must allocate the amount of memory needed to return the
**IP_ADAPTER_ADDRESSES** structures pointed to by the *AdapterAddresses*
parameter. When these returned structures are no longer required, the
application should free the memory allocated. This can be done in C# by
calling the `Marshal.AllocHGlobal` function to allocate memory and later
calling the `Marshal.FreeHGlobal` function to free the allocated memory.

**GetAdaptersAddresses** is implemented only as a synchronous function
call. The **GetAdaptersAddresses** requires a significant amount of
network resources and time to complete since all of the low-level network
interface tables must be traversed.

One method that can be used to determine the memory needed to return the
`IP_ADAPTER_ADDRESSES` structures pointed to by the *AdapterAddresses*
parameter is to pass too small a buffer size an indicated in the
*SizePointer* parameter in the first call to the **GetAdaptersAddresses**
function, so the function will fail with **ERROR_BUFFER_OVERFLOW**. When
the return value is **ERROR_BUFFER_OVERFLOW**, the *SizePointer*
parameter returned points to the required size of the buffer to hold the adapter information. Note that it is possible for the buffer size
required for the `IP_ADAPTER_ADDRESSES` structures pointed to by the
*AdapterAddresses* parameter to change between subsequent calls to the
**GetAdaptersAddresses** function if an adapter address is added or
removed. However, this method of using the **GetAdaptersAddresses**
function is strongly discouraged. This method requires calling the
**GetAdaptersAddresses** multiple times.

The recommended method of calling the **GetAdaptersAddresses** function
is to pre-allocate a 15KB working buffer pointed to by the
*AdapterAddresses* parameter. On typical computers, this dramatically
reduces the changes that the **GetAdaptersAddresses** function returns
**ERROR_BUFFER_OVERFLOW**, which would require calling
**GetAdaptersAddresses** function multiple times.
