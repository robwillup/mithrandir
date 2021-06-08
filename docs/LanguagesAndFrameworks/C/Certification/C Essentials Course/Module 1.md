# C Essentials - Part 1: Module 1

In this module, you will learn about:

* languages: natural and artificial
* machine languages
* high-level programming languages
* obtaining the machine code: compilation process
* writing simple programs
* variables
* integer values in real life and in C
* integer literals
* comments

## Natural languages vs programming language

Language is a tool for expressing and recording human thoughts.

The languages we use to communicate with other people are
called natural languages. They were created over many centuries
and still are subject to change. If we ignore languages that
have been created artificially, such as Esperanto or even
**Quenya** (the language used by the elves in Tolkien's world),
we can say that natural languages have evolved almost independently,
naturally; an evolution over which we have little or no control.

However, there are languages whose creation and development
were and continue to be dictated by specific needs, and their
development is fully subject to control by wide groups of people,
such as international committees or work groups.

The shapes of these languages are defined by international
standards, and although they are understood by many people,
the exchange of thought between human-beings is not their most
important application.

Such languages are, among others, programming languages. A
programming language is defined by a set of certain rigid rules,
much more inflexible than any natural language.

For example, these rules determine which symbols (letters,
digits, punctuation marks, and so on) could be used in the
language. This part of the definition of the language is called **lexicon**.

Another set of rules determines the appropriate ways of collating
the symbols - this is the **syntax** of the language.

We would also like to be able to recognize the meaning of every
statement expressed in the given language - and this is what
we call **semantics**.

Any program we write must be correct in these three ways:
lexically, syntactically, and semantically, otherwise it
will neither run not produce any acceptable results. You can
be sure that in the coure of your programming career, you will
experience all of these errors: to err is human, and these
humans write computer programs.

The expressive power of programming languages is much, much
weaker than those offered by natural languages. We cannot
(although we can try to) use a programming language to express
human emotions, and it's hard to imagine a declaration of love
encoded in it. It's a simple fact that the message embedded inside
a computer is not intended for a human, but for a machine.

A computer, even the most technically sophisticated one, is
devoid of even a trace of intelligence. It responds only to
a predetermined **set of known commands**. These recognized
commands are very simple. These recognized commands are very simple.
We can imagine that the computer responds to orders like
"take that number, add to another and save the result".
A complete set of well-known commands is called an
**instruction list**, sometimes abbreviated to **IL**.
Different types of computer may vary depending on the size
of their ILs and the instructions themselves could differ
completely from one model to the next.

The IL is in fact the alphabet of a language, commonly known
as a **machine language**. This is the simplest and most
primary language we can use to give commands to our computer.
We could say that it's the computer's mother tongue.

Computer programming is the act of composing selected
commands (instructions) in the proper order so that a
desired effect is produced.

It is possible, and often used in practice, for a computer
program to be coded directly in machine language using
elementary instructions (orders). This kind of programming
is tedious, time consuming and highly prone to a
programmer's mistakes. At the early stages of computer
technology, it was the only available method of programming
and it very quickly revealed some serious flaws. Firstly,
programming in machine language requires an
exhaustive knowledge of the computer's hardware design and
its internal structure. This also means that replacing the
computer with one that differs in design can make the
programmer's entire knowledge unusable. Also, the old
programs could become completely useless if the new
computer "used" a different IL. Secondly, programs written
in machine language are very difficult for humans to
understand, including experienced programmers. It also takes
a long time to develop programs in machine language, and it's
very costly and cumbersome too.

All these circumstances led to a need for some kind of
**bridge** between the human language (natural language)
and the computer language (machine language). That bridge
is also a language - an intermediate common language for
humans and computer to work together. Such languages are
often called **high-level programming languages**.

A high-level programming language is at least somewhat similar
to a natural language; it uses symbols, words and conventions
readable to humans. This language enables humans to express
complex commands for computers.

You might ask how we make computers understand programs
written in this way. This works by **translating** the
program into machine language. The translation can be done
by a computer, making the whole process fast and efficient.

The programs written in high-level language could be
translated into any number of different machine languages
and thus make them usable on many different computers.
This feature of high-level programming languages is called
**portability**.

## Compilation

The translation we are referring to is made by a specialized
computer program called a **compiler**. The process of
translating from a high-level language into machine language is called **compilation**.

To create a program you need to write it in accordance with
the rules of the chosen programming language. Such a program
(which in fact is just text) is called the source code, or
simply source, while the file which contains the source is
called the **source file**.

To write the source code you need a text editor that allows
you to manipulate text without any formatting information.
This code is placed in a file and the name of the file should
give you some clue as to its content. For example, it's
common for a file containing the source code in the "C"
language to have its name ending with the suffix "**.c**".

Next, your source code needs to be compiled. To do this you
run a compiler, instructing it where you stored the source
code that you want to be translated into the machine language.
The compiler reads your code, does some complex analysis
and its first goal is to determine whether or not you made
any errors during the coding. These analyses are very
insightful, but remember that they are made by a machine,
not a human, and you shouldn't expect too much from them.

If your mistake was that you tried to add up two numbers
using "#" instead of a "+", the compiler will kindly inform
you of your error.
However, if you typed a "-" instead of a "+", the compiler
will no longer be able to guess that your intention was to
add two numbers, rather than to subtract them. **Do not
expect the compiler to think for you**. But there's no
reason to be sad about it - thanks to that, developers are
still needed.

If the compiler doesn't notice any mistakes in your source,
the result of its work will be a file containing your
program translated into machine language. That
file is commonly called an **executable file**. The name of
the file depends on the compiler you use and the OS you're
working with. For example, most compilers designed for the
Unix/Linux system create an output file named "a.out" by
default. Compilers designed for use in Windows can give
this file the same name as the source file, only changing
the suffix from ".c" to ".exe".

We must admit that the whole process is actually a bit more
complicated. Your source code might be comprehensive and
divided among several or even dozen of source files. It may
also happen that the program was not written by you alone,
but by a team, in which case the division of sources into
multiple files is simply a must. In such cases, the compiling
splits into two phases - a **compilation** of your source, in
order to translate it into machine language, and a joining
(or gluing) of your executable code with the executable code
derived from other developers into a single and unified
product. The phase of "gluing" the different executables
codes is commonly known as **linking** while the program
which conducts the process is called a **linker**.

## A few words about C

The "C" language is one of a huge number of programming
languages currently in use, and one of the oldest. It was
created in the early seventies by Dennis Ritchie while he was
working in Bell Laboratories. Some say that "C" was a
by-product of a project which led to the very first version
of the Unix operating system.