# Git Basics - Tagging

Git has the ability to tag specific points in a repository's history as being important. This functionality is typically used to mark release points (v1.0, v2.0 and so on).

## Listing Your Tags

```bash
$ git tag
v1.0
v2.0
```

This command lists the tags in alphabetical order; the order in which they are displayed has no real importance.

## Creating Tags

Git supports two types of tags: `lightweight` and `annotated`.

A lightweight tag is very much like a branch that doesn't change - it's just a pointer to a specific commit.

Annotated tags, however, are stored as full objects in the Git database. They're checksummed; container tagger name, email, and data; have a tagging message; and can be signed and verified with GNU Privacy Guard (GPG). It's generally recommended that you create annotated tags so you can have all this information; but if you want a temporary tag or for some reason don't want to keep the other information, lightweight tags are available too.

## Annotated Tags

Creating an annotated tags is simple. The easiest way is to specify `-a` when you run the `tag` command:

```bash
$ git tag -a v1.4 -m "my version 1.4"
$ git tag
v1.4
```

The `-m` specifies a tagging message, which is stored with the tag. If you don't specify a message for an annotated tag, Git launches your editor so you can type it in.

You can see the tag data along with the commit that was tagged by using the `git show` command:

```bash
$ git show v0.1.0
tag v0.1.0
Tagger: Robson William <r.willup@hotmail.com>
Date:   Thu Dec 17 21:27:08 2020 -0300

First released version of Quest
```

## Sharing Tags

By default, the `git push` command doesn't transfer tags to remote servers. You have to explicitly push tags to a shared server after you have created them. This process is just like sharing remote branches - you can run `git push origin <tagname>`.