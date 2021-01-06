# Introduction to Azure Virtual Machines

* on-demand, scalable computing resource
* typically, a VM will be chosen if you need more control over the computing environment than the choices such as App Service or Cloud Services offer.
* Azure VMs provide you with:
    * Operating System
    * storage
    * networking capabilities
* VMs are part of IaaS

## Azure Virtual Machine creation checklist
There are always a lot design considerations when you build out an application infrastructure in Azure. These aspects of a VM are important to think about before you start:

* The names of your application resources
* The location where the resources are stored
* The size of the VM
* The maximum number of VMs that can be created
* The operating system that the VM runs
* The configuration of the VM after it starts
* The related resources that the VM needs

### Naming
The name of your VM can be up to 15 characters.
If you use Azure to create the operating system disk, the computer name and the virtual machine name are the same.

### Location
All resources created in Azure are distributed across multiple geographical regions around the world. Usually, the regions is called location when you create a VM. Form a VM, the location specifies where the virtual hard disks are stored.

Here are some of the ways you can get a list of available locations:

| Method | Description|
|:-:|:-:|
| Azure Portal | Select a location from the list when you create a VM |
| Azure PowerShell | Use the Get-AzLocation command |
| REST API | Use the list locations operations |
| Azure CLI | Use the az account list-locations operation |

### VM Size
The size of the VM that you use is determined by the workload that you want to run. The size that you choose then determines factors such as processing power, memory, and storage capacity. 