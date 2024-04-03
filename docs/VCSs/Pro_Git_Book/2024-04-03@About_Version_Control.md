## About Version Control

Version control is a system that records changes to a file or set of files over
time so that you can recall specific versions later.

It's not uncommon to think that the main purpose of Git is to simply save your
work to a remote server, but as a version control system the main purpose is to
keep a history of changes to a file or files allowing you to recall these specific changes later.

A Version Control System (VCS) works with nearly any type of file on a computer. So a graphic designer, for example, can revert selected files back to a previous state, revert the entire project back to a previous state, compare changes over time, see who last modified something that might be causing a problem, who introduced an issue and when, and more. Using a VCS also generally means that if you mess up or lose files, you can easily recover.

### Local Version Control Systems

Many people's version-control method of choice is to copy files into another directory (perhaps a time stamped directory, if they're clever). This approach is very common because it is so simple, but it is also incredibly error prone. It is easy to forget which directory you're in and accidentally write to the wrong file or copy over files you don't mean to.

To deal with this issue, programmers long ago developed local VCSs that had a simple database that kept all the changes to files under revision control.

```diag
Local Computer

Checkout   |   Version Database
           |
File ------|------ Version 3
           |           |
           |       Version 2
           |           |
           |       Version 1
```

One of the most popular VCS tools was a system called RCS, which is still distributed with many computers today. RCS works by keeping patch sets (that is, the differences between files) in a special format on disk; it can then re-create what any file looked like at any point in time by adding up all the patches.
