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


### Centralized Version Control Systems

The next major issue that people encounter is that they need to collaborate with developers on other systems. To deal with this problem, Centralized Version Control Systems (CVCSs) were developed. These systems (such as CVS, Subversion, and Performance) have a single server that contains all the versioned files, and a number of clients that check out files from that central place. For many years, this has been the standard for version control.

```diag
____________             __________________________
Computer A |             |  Central VCS Server    |
File       | <---------- |  ____________________  |
____________             |  | Version Database |  |
                         |  |                  |  |
____________             |  |    Version 3     |  |
Computer B |             |  |        |         |  |
File       | <---------- |  |    Version 2     |  |
____________             |  |        |         |  |
                         |  |    Version 1     |  |
                         |  |                  |  |
                         __________________________
```

This setup offers many advantages, especially over local VCSs. For example, everyone knows to a certain degree what everyone else on the project is doing. Administrators have fine-grained control over who can do what, and it's far easier to administer a CVCS than it is to deal with local databases on every client.

However, this setup also has some serious downsides. The most obvious is the single point of failure that the centralized server represents. If that server goes down for an hour, then during that hour nobody can collaborate at all or save versioned changes to anything they're working on. If the hard disk the central database is on becomes corrupted, and proper backups haven't been kept, you lose absolutely everything - the entire history of the project except whatever single snapshots people happen to have on their local machines. Local VCSs suffer from this same problem - whenever you have the entire history of the project in a single place, you risk losing everything.

### Distributed Version Control Systems

This is where Distributed Version Control Systems (DVCSs) step in. In a DVCS (such as Git, Mercurial or Darcs), clients don't just check out the latest snapshot of the files; rather, they fully mirror the repository, including its full history. Thus, if any server dies, and these systems were collaborating via that server, any of the client repositories can be copied back up to the server to restore it. Every clone is really a full backup of all the data.

```diag
                              __________________________
                              |     Server Computer    |
                              | ______________________ |
                              | |  Version Database  | |
                              | |                    | |
                              | |      Version 3     | |
                              | |          |         | |
                              | |      Version 2     | |
                              | |          |         | |
                              | |      Version 1     | |
                              | |____________________| |
                              |________________________|
                              ^                      ^
                             /                        \
                            /                          \
__________________________ /                            \__________________________
|        Computer A      |                               |        Computer B      |
|           File         |                               |           File         |
|            ^           |                               |            ^           |
| ___________|__________ |                               | ___________|__________ |
| |  Version Database  | |                               | |  Version Database  | |
| |                    | | <---------------------------> | |                    | |
| |     Version 3      | |                               | |     Version 3      | |
| |         |          | |                               | |         |          | |
| |     Version 2      | |                               | |     Version 2      | |
| |         |          | |                               | |         |          | |
| |     Version 1      | |                               | |     Version 1      | |
| |____________________| |                               | |____________________| |
|________________________|                               |________________________|

```

Furthermore, many of these systems deal pretty well with having several remote repositories they can work with, so you can collaborate with different groups of people in different ways simultaneously within the same project. This allows you to set up several types of workflows that aren't possible in centralized systems, such as hierarchical models.
