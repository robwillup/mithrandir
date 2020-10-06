## Using environment variables

[Reference](https://docs.microsoft.com/en-us/nuget/reference/nuget-config-file#using-environment-variables)

You can use environment variables in `nuget.config` values (NuGet 3.4+) to apply settings at run time.

For example, if the `HOME` environment variable on Windows is set to `C:\Users\Username`, then the value 
of %HOME\NuGetRepository% in the configuration file resolses to `C:\Users\Username\NuGetRepository`.

Note that you have to use Windows-style environment variables (starts and ends with %) even on Mac/Linux.
Having `$HOME/NuGetRepository` in a configuration file will not resolse. On Mac/Linux the value of 
`%HOME%/NuGetRepository` will resolse to `/home/myStuff/NuGetRepository`.

If an environment variable is not found, NuGet uses the literal value from the configuration file.
For example `%MY_UNDEFINED_VAR%/NuGetRepository` will be resolved as 
`path/to/current_working_dir/$MY_UNDEFINED_VAR/NuGetRepository`.

