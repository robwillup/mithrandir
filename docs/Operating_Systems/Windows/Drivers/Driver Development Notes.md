# Driver Development Notes

These are notes I'm taking as I'm learning driver development.

## Setting up the development environment

These are some of the tools used to driver development on Windows.

### Visual Studio

I downloaded Visual Studio and installed the "Desktop Development with C++" module.

**Missing "Kernel Mode Driver, Empty (KMDF)" template**

When I tried to create a new project, the templace "Kernel Mode Driver, Empty (KMGF)"
was missing.
I found some information only that I needed to downdoad the Windows Driver Kit which
I actually already had installed on my system, but I had installed it using the .iso
file. I went on to uninstall it so that I could reinstall using the .exe installer.
As I was uninstalling the Windows Driver Kit from my machine, I found out that I
actually had three different versions installed. So I removed all versions.

After reinstalling WDK the template is still not available in VS.

