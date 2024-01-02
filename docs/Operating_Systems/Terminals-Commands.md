## Running a command multiple times

for i in {< START >..< FINISH >}; do < COMMAND > ; done

```bash
for i in {1..10}; do dotnet test ; done
```

### Linux

•  cd: Changes the current directory:

```bash
cd ..
cd /
```

•  chmod: Changes the protection of a file or directory

```bash
chmod 777
```

•  chown: Changes the owner of a file or directory
•  chgrp: Changes the group of a file or directory
•  cmp: Compares two files
•  comm: Selects or rejects common lines of two selected files
•  cp: Copy files
•  crypt: Encrypts or decrypts files
•  diff: Compares the content of two files ASCII
•  file: Determines the type of the file
•  grep: Searches a file for a pattern, which is a very useful filter:

```bash
cat a.txt | grep ola 
```

•  gzip: Compresses or expands a file
•  ln: Creates a link to a file
•  ls: Lists the content of a directory
•  lsof: Lists open files
•  mkdir: Creates a directory
•  mv: Moves or renames files and directories
•  pwd: Displays the full path of the directory where we are
•  quota: Displays disk usage and limits
•  rm: Deletes files
•  rmdir: Deletes directories
•  stat: Displays the status of a file
•  sync: Flushes the buffers of the file system, syncs disk data with memory
•  sort: sorts, joins or compares text, it can be used to extract information from text files or to sort data from other commands
•  tar: Creates or extracts files
•  tee: Copies the input to the standard output or to other files
•  tr: Translates characters
•  umask: Changes file protection
•  uncompress: Restores a compressed file
•  uniq: Reports or deletes duplicated lines in a file
•  wc: Count lines, words and even characters in a file

### MS-DOS

DATE - Command to update the operating system’s Date

```cmd
C:\>date
```

TIME - Command to update the operating system's Time

```cmd
C:\>time
```

VER - Command to display the version number of the operating system that's in use.

```cmd
C:\>ver
```

DIR - Command that lists the content of a folder

```cmd
C:\>dir
```

CLS - Command to clear the screen

```cmd
C:\>cls
```

MKDIR or MD - Command to create a new directory

```cmd
C:\>mkdir test
```

CHDIR or CD - Command to navigate to a different directory

```cmd
C:\>cd ..
```

RMDIR or RD - Command to remove a directory

```cmd
C:\>rd test
```

TREE - Command that graphically displays the directory tree from the root directory to show the user a hierarchical view of the organization of their disk.

```cmd
C:\>tree
```

CHKDSK - Command that checks the integrity and the specifications of the disk displaying information about it

```cmd
C:\>chkdsk
```

RENAME or REN - Command to allow users to rename a file or directory

```cmd
C:\>ren test test2
```

Copy - Command to copy a file or a group of files from one folder to another

```cmd
C:\>copy test\file.txt test2
```

XCOPY - Copy files and directory trees based on certain criteria

MOVE - Rename directories or move files from one folder to another

```cmd
C:\>move test2 test3
```

TYPE - Show the content of a file

```cmd
C:\> type file1.txt
```

DEL or DELETE - Deletes files

```cmd
C:\del file.txt
```

UNDELETE - This commands allows retrieving one or more deleted files, when possible

DELTREE - This command eliminates one or more subdirectories
