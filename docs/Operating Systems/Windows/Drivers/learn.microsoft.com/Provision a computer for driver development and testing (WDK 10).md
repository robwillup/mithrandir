# Provision a computer for driver development and testing (WDK 10)

> WARNING!
> This poioka does NOT work! I spend a day trying to get this to
> work and thoroughly followed all steps. No luck.
> If you want, read on, but it may only be a waist of time.

Reference: [https://learn.microsoft.com/en-us/windows-hardware/drivers/gettingstarted/provision-a-target-computer-wdk-8-1]

Provisioning a target or test computer is the process of configuring a
computer for automatic driver deployment, testing, and debugging. To
provision a computer, use Microsoft Visual Studio.

A testing and debugging environment has two computers: the host
computer and the target computer. The target computer is also called
the test computer. You develop and build your driver in Visual Studio
on the host computer. The debugger runs on the host computer and is
available in the Visual Studio user interface. When you test and debug
a driver, the driver runs on the target computer.

The host and target computers must be able to ping each other by name.
This might be easier if both computers are joined to the same
workgroup or the same network domain. If your computers are in a
workgroup, we recommend that you connect the computers with a router
rather than a hub or switch.

It is recommended that both the host and target run the same version
of Windows.

## Prepare the target computer for provisioning

1. On the target computer, install the operating system that you'll
use to run and test your driver.

2. [Install the WDK](https://learn.microsoft.com/en-us/windows-hardware/drivers/download-the-wdk).
You do not need to install Visual Studio, however, unless you plan on
doing driver development on the target computer.

3. If Secure Boot is enabled on an x86 or x64 target computer, disable
it. To check if Secure Boot is enabled in a target computer with
any of these architectures, check the following:

a. Click on the Start button and search for "System Information"
b. On the right panel, in the "Item" column, look for "Secure Boot State"

If it is "On", and your target machine is a Hyper-V VM, you can disable
it by doing the following:

a. Click on the VM and select "Settings"
b. Under "Security" uncheck "Enable Secure Boot"

4. On the target computer, run the WDK Test Target Setup MSI that
matches the platform of the target computer. You can find the MSI in
the Windows Driver Kit (WDK) installation directory under Remote.

> Example: C:\Program Files (x86)\Windows Kits\10\Remote\x64\
> WDK Test Target Setup x64-x64_en-us.msi

## Provision the target computer

Now you're ready to provision the target computer from the host
computer in Visual Studio.

1. On the host computer, in Visual Studio, select the **Extensions**
menu, point to **Driver**, point to **Test**, and select **Configure
Devices**.

2. In the **Configure Devices** dialog, select **Add new device.**

3. For **Network host name**, enter the name of local IP address
of your target computer. Select **Provision device and choose
debugger settings**.

4. Select **Next**

## Result

It always fails with this error:

> ERROR: Task "Installing desktop driver test framework" failed to complete successfully. Look at the logs in the driver test group explorer for more details on the failure.

