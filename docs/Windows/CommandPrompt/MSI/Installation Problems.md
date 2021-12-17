# Software Installation Issues in Windows

If you ever try to install and application in Windows with an MSI installer and get an error saying
that either the resource is in a different network and is unavailable or an error similar to this:

> the file ... is not a valid installation package for the product "..." . try to find
> the installation package in a folder from which you can install ...

This could mean that something messed with the Windows Installer properties.
When I had this issue, it happened because I cleared the `C:\Windows\Installer` folder.

After several attempts trying to solve this finally I learned that I had to delete some Registry entries
related to the software I was trying to install. Those Registry keys are here:

> Computer\HKEY_CLASSES_ROOT\Installer\Products

Under `Products`, you need to search for the key related to the software you're trying to install and delete them.

> NOTE: It might be a good idea to export your current Registry state with "File >> Export"