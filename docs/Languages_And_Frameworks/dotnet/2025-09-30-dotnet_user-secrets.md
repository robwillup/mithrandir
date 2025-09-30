# Safe storage of app secrets in development in ASP.NET Core

[Reference](https://learn.microsoft.com/en-us/aspnet/core/security/app-secrets?view=aspnetcore-9.0&tabs=windows)

This is a study about how to manage sensitive data for an ASP.NET Core app on a development machine. Never store
passwords or other sensitive data in source code or configuration files. Production secrets shouldn't be used for
development or test. Secrets shouldn't be deployed with the app. Production secrets should be accessed through a
controlled means like Azure Key Vault.

## How to use dotnet user-secrets in a console app

This is the most basic example showing how to use `dotnet user-secrets` in a console app.

```C#
using static System.Console;
using Microsoft.Extensions.Configuration;
using Microsoft.Extensions.DependencyInjection;
using Microsoft.Extensions.Hosting;

using IHost host = Host.CreateDefaultBuilder(args).UseEnvironment("Development").Build();

IConfiguration config = host.Services.GetRequiredService<IConfiguration>();

int keyOneValue = config.GetValue<int>("keyOne");
bool keyTwoValue = config.GetValue<bool>("keyTwo");

WriteLine($"Key one = {keyOneValue}");
WriteLine($"Key two = {keyTwoValue}");
```

And this example shows how to use user-secrets in a single file app:

```C#
#:package Microsoft.Extensions.Configuration.UserSecrets@9.0.9
#:package Microsoft.Extensions.Hosting@9.0.9

using static System.Console;
using Microsoft.Extensions.Configuration;
using Microsoft.Extensions.Configuration.UserSecrets;
using Microsoft.Extensions.DependencyInjection;
using Microsoft.Extensions.Hosting;

using IHost host = Host.CreateDefaultBuilder(args)
    .UseEnvironment("Development")
    .ConfigureAppConfiguration((hostingContext, config) =>
        {
        if (hostingContext.HostingEnvironment.IsDevelopment())
        {
            config.AddUserSecrets("daf44188-0d10-4241-a7e2-6860cdec8c4d");
        }
    })
    .Build();

IConfiguration config = host.Services.GetRequiredService<IConfiguration>();

int keyOne = config.GetValue<int>("keyOne");
bool keyTwo = config.GetValue<bool>("keyTwo");

WriteLine($"Value one: {keyOne}");
WriteLine($"Value two: {keyTwo}");
```

For these to work, you first need to initialize the user secrets for the C# project. You can do that by navigating into
that first example which should have a `.csproj` file, and in there you can run:

```shell
dotnet user-secrets init
dotnet user-secrets set keyOne 7
dotnet user-secrets set keyTwo true
```

## Environment variables

Environment variables are used to avoid storage of app secrets in code or in local configuration files. Environment
variables override configuration values for all previously specified configuration sources.

Consider an ASP.NET Core web app in which **Individual Accounts** security is enabled. A default database connection
string is included in the project's `appsettings.json` file with the key `DefaultConnection`. The default connection
string is for LocalDB, which runs in user mode and doesn't require a password. During app deployment, the
`DefaultConnection` key value can be overriden with an environment variable's value. The environment variable may
store the complete connection string with sensitive credentials.

> **WARNING!**
> Environment variables are generally stored in plain, unencrypted text. If the machine or process is compromised,
> environment variables can be accessed by untrusted parties.

The `:` separator doesn't work with environment variable hierarchical keys on all platforms. For example, the `:`
separator is not supported by [Bash](https://linuxhint.com/bash-environment-variables/). The double underscore, `__`,
is:

- supported by all platforms
- automatically replaced by a colon, `:`.

## Secret Manager

The Secret Manager tool stores sensitive data during application development. In this context, a piece of sensitive
data is an app secret. App secrets are stored in a separate location from the project tree. The app secrets are
associated with a specific project or shared across several projects. The app secrets aren't checked into source
control.

> **WARNING!**
> The Secret Manager tool doesn't encrypt the stored secrets and shouldn't be treated as a trusted store. It's for
> development purposes only. The keys and values are stored in a JSON configuration file in the user profile directory.

### How the Secret Manager tool works

