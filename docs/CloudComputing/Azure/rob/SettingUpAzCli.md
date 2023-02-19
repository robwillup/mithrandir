# Setting Up Az CLI

In this document I am journaling the things I've learned while setting up AZ CLI.

## Installation

I have downloaded the AZ CLI using [winget](https://learn.microsoft.com/en-us/windows/package-manager/winget/). I found the instructions to do that in this page:

[Install Azure CLI on Windows](https://learn.microsoft.com/en-us/cli/azure/install-azure-cli-windows?tabs=winget)

What is the AZ CLI:
> The Azure Command-Line Interface (CLI) is a cross-platform command-line tool that can be
> installed locally on Windows computers. You can use it to connect to Azure and execute
> administrative commands on Azure resources.

What is WinGet:
> The winget command line tool enables users to discover, install, upgrade, remove and configure
> applications on Windows 10 and 11. This tool is the client interface to the Windows Package
> Manager Service.

And this is the command that I used to install AZ CLI on my computer:

```pwsh
> winget install -e -id Microsoft.AzureCLI

This application is licensed to you by its owner.
Microsoft is not responsible for, nor does it grant any licenses to, third-party packages.
Downloading https://azcliprod.azureedge.net/msi/azure-cli-2.45.0.msi
  ██████████████████████████████  52.5 MB / 52.5 MB
Successfully verified installer hash
Starting package install...
Successfully installed
```

After doing that I of course had to restart my terminal so that it would have the updated
PATH environment variable so that I could start running the `az` command:

```pwsh
➜ az

Welcome to Azure CLI!
---------------------
Use `az -h` to see available commands or go to https://aka.ms/cli.

Telemetry
---------
The Azure CLI collects usage data in order to improve your experience.
The data is anonymous and does not include commandline argument values.
The data is collected by Microsoft.

You can change your telemetry settings with `az configure`.


     /\
    /  \    _____   _ _  ___ _
   / /\ \  |_  / | | | \'__/ _\
  / ____ \  / /| |_| | | |  __/
 /_/    \_\/___|\__,_|_|  \___|


Welcome to the cool new Azure CLI!

Use `az --version` to display the current version.
Here are the base commands:
...
```

## Login

The next thing I wanted to do was to log in, so I simply ran `az login` and that opened up a
page in my browser. I signed in there and got this in my terminal:

```pwsh
[
  {
    "cloudName": "AzureCloud",
    "homeTenantId": "1d188************************",
    "id": "8589b7****************************",
    "isDefault": true,
    "managedByTenants": [],
    "name": "Pay-As-You-Go",
    "state": "Enabled",
    "tenantId": "1d18*****************************",
    "user": {
      "name": "r***********************",
      "type": "user"
    }
  }
]
```
