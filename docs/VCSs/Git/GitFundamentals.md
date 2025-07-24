# Git Fundamentals

The goal here is to obtain a solid foundation to understanding the fundamentals of Git.

## Thinking in Git

Git is by far the most widely used version control system in the world.
Let's start thinking in Git by understanding its definition.

> Git is a free and open source distributed version control system designed to handle
> everything from small to very large projects with speed and efficiency.

Git gives you the ability to go back in time at any point during your development and
see the work at that specific point in time. To do this, Git tracks changes made to
files in a project and creates a unique identifier for each change so that it's
easily discoverable if you ever want to go back and retrieve a version of the project
or revert a change that you've made. You can also think of this as an autosave at an
specific checkpoint that you can easily reference for your project. The great thing
about Git is that you can version control nearly any type of file, not just source
code files. Students often use Git to keep track of different versions of an essay or
a research paper. A designer can keep versions of an image without worrying about
modifying the only version of that image. And version control gives you the freedom
to explore new ideas within your project with the assurance that you can always
revert back to a previous version.

### Types of version control systems

**Local Version Control** is simply keeping different versions of files on your local machine.
This is the most basic form of version control, where you manually save different versions of your files.
One of the limitations of this approach is that others do not have visibility into your changes.
Another limitation is that you cannot easily collaborate with others on the same project.
Additionally, if you lose your local machine, you lose all your versions.

**Centralized Version Control** systems are based on the idea that there is a single central
copy of your project somewhere, probably on a server, and team members will make their changes
to this central copy. This allows for multiple clients to checkout files from that centralized
location at the same time, make their changes, and then send them back to the centralized server
for others to see.
This is a big upgrade from local version control, but it also has its flaws.
The biggest problem is that everything is being store on a centralized server. So if that server
goes down, nobody can access the project. Similar to local version control, if that single central
copy is lost, and there are no backups, then all the versions of the project are lost.

**Distributed Version Control** systems copy the complete project, including its history and metadata,
to every user's machine. This eliminates the sole reliance on a central server.

### The Characteristics of Git

Git is not just distributed, it's an advanced versions control system that is incredibly robust with
lot's of commands and options for you to perform that give you a lot of control over how your project
is being tracked and maintained.

#### Snapshots, Not Diffs

Unlike other version control systems, Git stores versions as snapshots, not diffs.
Git does not track version changes as just a list file-based changes between one version to another
by only representing the changes to a base version of a file. These version differences are commonly
referred to as the deltas or diffs. Instead, Git stores each saved version of your project as a complete
snapshot of the entire project's file system at that specific point in time, and not just the files being
modified and their deltas/diffs.

This snapshot captures everything in the project, including its history and all of its metadata. You can
reference old snapshots whenever you need to and new snapshots are created when your project is modified
and the modifications are committed to the repository. 

#### Optimized for Local Development

As a distributed version control system, and so we have a complete copy of the project with all of its
history and metadata on our own local machine. So if the main version of a project is being maintained on
a server which will be the case for nearly every Git project, you can work disconnected from the main
server since Git doesn't need a constant connection for you to work in the project. This gives you the
ability to work completely offline and still be able to navigate the project and use Git, then when you're
online again and you have access to the remote server, you can push your changes so that everyone can see,
and also pull changes that other team members have made.

#### Git is Explicit

Git won't perform any action without you explicitly telling it to do so. Git won't sync your changes
automatically, it won't commit your changes automatically, and it won't push your changes to the remote
server automatically. At first this may seem like a bottleneck because you will have to explicitly tell
Git what to do, but once you get the hang of it, you'll find that this level of interaction gives you
complete freedom to tell Git when to do things, how to do things, and where to do it.

#### Git is Designed for Non-Linear Development

Git enables you to diverge your work by creating parallel timelines, giving you the freedom to explore
and test out new ideas, without the fear of messing up the main timeline of your project. Now, this
concept of creating alternate timelines is called branching, and it's a foundational features to how
people work using Git. Branching in Git is lightweight and considered inexpensive.

This is because when you create a branch in Git you are creating a pointer to a specific snapshot or
saved version of your work in the project. This pointer marks the beginning of a new parallel timeline
at that specific point in time and then allows you to try out new ideas or incorporate new features
without affecting the main development timeline. Then when you're ready to incorporate those changes into
the main timeline you can merge your branch into the main development timeline and those new changes will
not be incorporated.

### The Statuses of Your Files

In a Git project, a file can be in one of two states: tracked or untracked.

When you add a new file to an existing Git project, Git won't automatically start tracking it. You will
have to tell Git to track it. This is by design and gives you more control over how and when to save
your files.

Tracked files are files that were in the last snapshot of the project and are being tracked by Git.

Once a file has been tracked, it can be in one of three states:
1. **Unmodified**: The file has not changed since the last snapshot.
2. **Modified**: The file has changed since the last snapshot, but those changes have not been staged for commit.
3. **Staged**: The file has changed since the last snapshot and those changes have been staged for commit.

### The 3 Directories of Git

1. **.git (A.K.A. Local History)**: This is the directory where Git stores all of its metadata and object database.
   It contains all the information about the repository, including its history, branches, and configuration settings.
   It is a hidden directory at the root of your project and is created when you initialize a new Git repository.
2. **Working Directory**: This is the directory where you have your project files checked out. It's the area where
   you make changes to your files and where Git tracks those changes.
3. **Staging Area (A.K.A. Index or Cache)**: This is a temporary area where you can place changes that you want to
   include in your next commit. It's like a buffer between the working directory and the repository.

### Introduction to Git Commands

Git Repository: a repository contains all of your project's files and their history.

#### Git Basics

* git init
* git add
* git status
* git commit
* git config
* git log
* git diff

#### Git Branches

* git branch - list, create, or delete branches
* git checkout - switch branches, create new branches
* git merge - merge branches

#### Remote Repositories

* git clone - copies an entire repository from a remote server to your local machine
* git remote - create and show linked repositories
* git push - send updates to associated repositories
* git pull - Retrieves and integrates changes from other repositories
* git fetch - Retrieves but doesn't integrate changes from other repositories.

#### Undoing Changes

* git revert - create a new commit that undoes a previous commit
* git reset - removes files from the staging area.

## Start Using Git

* git init - Initializes a new Git repository in the current directory.
* git branch -m <name> - Renames the current branch to <name>.
* git status - Displays the status of the working directory and staging area.

### Git Configuration Levels

* System Level: Configuration settings that apply to all users on the system.
* User Level: Configuration settings that apply to a specific user.
* Repository Level: Configuration settings that apply to a specific repository.