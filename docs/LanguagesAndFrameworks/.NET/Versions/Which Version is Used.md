# Which .NET Version is Used?

The policies used by the .NET tools, SDK, and runtime for selecting versions provide a balance between running applications using the specified versions and enabling ease of upgrading both developer and end-user machines. These policies enable:

* Easy and efficient deployment of .NET, including security and reliability updates.
* Use the latest tools and commands independent of target runtime.

## SDK

The SDK uses the latest installed version on the machine, even if:

* The project targets an earlier version of the .NET runtime.
* The latest version of the .NET SDK is a preview version.

On rare occasions you may need to specify which version of .NET the SDK should use. You can specify that in the file `global.json`.

## Target Framework Monikers

You build your project against APIs defined in a **Target Framework Moniker** (TFM). You specify the target framework in the project file.

```csproj
<TargetFramework>net5.0</TargetFramework>
```

You may build your project against multiple TFMs. Setting multiple target frameworks is more common for libraries but can be done with applications as well:

```csproj
<TargetFrameworks>net5.0;netcoreapp3.1;net47</TargetFrameworks>
```

## Self-contained deployments include the selected runtime

This approach bundles the .NET runtime and libraries with your application. Self-contained deployments don't have a dependency on runtime environments. Runtime version selection occurs at publishing time, not run time.


