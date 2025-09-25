# Installing C Compilers

## Debian

This is how you install the GCC (GNU Compiler Collection) compiler on Debian. It supports C and other languages, along
with utilities like "make" that are essential for building programs on Debian.

1. Update the package list:

```bash
sudo apt update
```

2. Install the build-essential package:

```bash
sudo apt install build-essential
```

3. (Optional) Install manual pages for development:

```bash
sudo apt-get install manpages-dev
```

4. Verify that GCC is installed by checking its version:

```bash
gcc --version
```
