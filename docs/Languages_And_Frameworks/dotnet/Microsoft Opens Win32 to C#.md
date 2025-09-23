# Microsoft Opens Up Old Win32 APIs to C#

This will let coders use C# (and soon other languages) instead of default
C/C++ option or individual workarounds.

Before you would need to use workarounds, such as P/Invoke, but these are
cumbersome.

In January 2021, Microsoft introduced its **win32metadata** project,
along with a couple of initial Win32 language projections. A programming
language projection is a subsystem -- variously described as a set of
wrappers or an adapter -- that fosters development using the APIs of a
platform (Win32 in this case) in a natural and familiar way for the
target language.

The metadata project eases the creation of language projections by
providing a complete description of the Win32 API surface so it can be
automatically projected to any language, which Microsoft said improves
correctness and minimizes maintenance. Thus these metadata descriptions
aren't meant to be consumed directly by developers, who instead will use
the language projections that in turn consume the metadata and project
the APIs into the natural patterns of specific languages.
