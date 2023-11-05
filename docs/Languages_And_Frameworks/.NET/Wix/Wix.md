# Wix

The WiX Toolset is a set of tools that build Windows installation
packages from XML source code.

WiX is an open source project, originally developed by Microsoft and
maintained by Rob Mensching.

The toolset is written in C# and requires the .NET Framework to run.
However, this only applies to the toolset itself. The installation
packages you create with the toolset do not require any extra framework
or software to be installed on the target computer.

## Introduction

The Windows Installer moved to a declarative description, it specifies
the state the target machine should be left in after various phases of
installation and uninstallation. This declarative approach makes it
possible to cope with unexpected conditions, differing target machine
environments, aborted installations, shared resources. It is of paramount
importance for setup developers to make sure that whatever happens during
the process, the target machine should be left in a known, stable state,
without introducing any detrimental side effects.

Most setup tools only offer a graphical interface that allows the
developers to collect the files and other related tasks making up the
installation process manually. WiX uses a different approach, it is much
more like a programming language. It uses a text file (XML) to describe
all the elements of the installation process. The toolset has a compiler
and a linker that will create the setup program just like oiur usual
compiler creates our application from the source files. Therefore, WiX
can be made part of any automated application build process very easily,
be that based either on the classical technology of makefiles or the
similar features of contemporary integrated development environments.

Traditionally, setup programs were only written when the main application
had already been finished; often even by different developers. This
approach requires a tedious and error prone process of collecting
information about all the resources making up the application. While the
files themselves are usually obvious, registry entries, services and most
forms of inter-source dependencies are often hard to reconstruct in a
later stage.

In an integrated application/setup development environment, the original
developer should modify the WiX source files in sync with the application
development.

WiX has a steep learning curve and requires exposure to the internal
details and intricacies of the underlying Windows Installer technology.

WiX features:

* declarative approach
* unrestricted access to Windows Installer functionality
* source code instead of GUI-based
* complete integration into application build processes
* possible integration with application development
* free, open source

## Getting Started

The `.msi` file we want to distribute our application with is not a setup
application but an installation database. The programming logic, the
knowledge about thow to install an application, how to modify registry
keys, how to create shortcuts, users and network shares, how to
manipulate web directories or services resides in Windows Installer. Our
setup file only describes *what* we expect Windows Installer to do and
provides the files to be deployed (as well as interface elements used in
the process).

This database approach means that our WiX source files are not built like
regular programs. There is no notion of sequential execution in WiX, the
first source line is not supposed to be executed prior to the second one.
There will be no declarations that need to precede the references.
Various elements might be described in different places and, wherever a
link is required between them, one will refer to the other using unique
identifiers we need to provide.

WiX is a "comfortable", XML-style way to describe your installation
requirements that gets translated into Windows Installer .msi databases
by its compiler and linker. WiX is a relatively thin wrapper around
Windows Installer technology.

To start with, you'll need two GUIDs, one for your product and one for
the installation package (actually, for any real world project, you'll
also need an `UpgradeCode` GUID; be sure to checkout the other lessons
before you ship anything). While the other two have to be kept on file
because you will probably need to refer to them later, `Package` GUIDs
need to be different with each package you create. To make it easier and
less likely to forget to issue a new one, we can instruct WiX to
auto generate one by typing an asterisk---but remember, this only applies
to package GUIDs: all other GUIDs will need to stay unique and kept
recorded for times to come. This, with all other textual information
about the product will go into the very first part of our `*.wxs` file:

```xml
<?xml version='1.0' encoding='windows-1252'?>
<Wix xmlns='http://schemas.microsoft.com/wix/2006/wi'>
    <Product Name='Foobar 1.0' Manufacturer='Acme Ltd.'
        Id='YOURGUID-86C7-4D14-AEC0-86416A69ABDE' 
        UpgradeCode='YOURGUID-7349-453F-94F6-BCB5110BA4FD'
        Language='1033' Codepage='1252' Version='1.0.0'>
    <Package Id='*' Keywords='Installer' Description="Acme's Foobar 1.0 Installer"
        Comments='Foobar is a registered trademark of Acme Ltd.' Manufacturer='Acme Ltd.'
        InstallerVersion='100' Languages='1033' Compressed='yes' SummaryCodepage='1252' />
```
