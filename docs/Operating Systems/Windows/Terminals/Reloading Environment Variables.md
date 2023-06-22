# Reloading Environment Variables

With PowerShell you can reload environment variables without having to restart
your terminal.

```powershell
$Env:Path = [System.Environment]::GetEnvironmentVariable("Path","Machine") + ";" + [System.Environment]::GetEnvironmentVariable("Path","User")
```
