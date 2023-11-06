# Linux Command Line Basics

## Linux Help Resources

### $ man <COMMAND>

```bash
$ man curl
```

### $ info or $ info <COMMAND>

The `info` program is a stand-alone program, which is used to view Info files on a text terminal.


### /usr/share/docs/<COMMAND>

### <COMMAND> --help

### type <COMMAND>

tells us how a specified word will be interpreted by Bash.

# The Linux Terminal

The *thing* that you open on your desktop is a **Terminal Window**. When you open that **terminal window**, a new **Terminal Session** is created for you.

A shell session automatically places us in our `home` directory.

* ls -a => lists all files including hidden (.file)
* less => prints a specified file one screen at a time
* ssh => used to log in to remote sessions
  * ssh ubuntu@10.0.70.110


# Linux Command Syntax Patterns and Shortcuts

Invoking a command always starts with the command name. In some cases, as with `ls`, just the name on its own is enough to return a result. However, you will usually have to enter some combination of arguments and parameter to get the most out of a tool. In most cases, command line tools will have their parameters and arguments with `--` for the long form, and `-` for the short Alias form.

* ls -l => lists the content of a directory in their long form, displaying object permissions, ownership, size and age details.
* ls -lh => gives us the same as above but the file size in human-readable format.
* ls -lht => organizes the listing in chronological order.
* cd => when running `cd` without arguments takes us back to our home directory.
* history => lists the last 1000 executed commands
