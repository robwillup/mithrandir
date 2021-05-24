# Computer Science Self-study

This study is based on the website <https://teachyourselfcs.com>

You can give yourself a world-class CS education without
investing years and a small fortune in a degree program.

You need answers to these questions:

* **Which subjects** should you learn, and why?
* What is the **best book or video lecture series** for each
subject?

<h2>TL;DR</h2>

Study all nine subjects below, in roughly the presented
order, using either the suggested textbook or video
series, but ideally both. Aim for 100-200 hours of study
of each topic, then revisit favorites throughout your career.

## Computer Programming

Don't be the person who "never quite understood" something like recursion.

Best book: *Structure and Interpretation of Computer Programs*

Best video lectures: Brian Harvey's Berkeley CS 61A

## Computer Architecture

If you don't have a solid mental model of how a computer actually works,
all of your higher-level abstractions will be brittle.

Best book: *Computer Systems: A Programmer's Perspective*

Best video lectures: Berkeley CS 61C

## Algorithms and Data Structures

if you don't know how to use ubiquitous data structures like
queues, stacks, trees, and graphs, you won't be able to
solve challenging problems.

Best book: *The Algorithms Design Manual*

Best video lectures: Steven Skiena's lectures

## Math for CS

CS is basically a runaway branch of applied math, so learning
math will give you a competitive advantage.

Best book: *Mathematics for Computer Science*

Best video lectures: tom Leighton's MIT 6.042J

## Operating Systems

Most of the code you write is run by an operating system,
so you should know how those interact.

Best book: *Operating Systems: Three Easy Pieces*

Best video lectures: Berkeley CS 162

## Computer Networking

The Internet turned out to be a big deal: understand
how it works to unlock its full potential.

Best book: *Computer Networking: A Top-Down Approach*

Best video lectures: Stanford CS 144

## Databases

Data is at the heart of most significant programs, but few
understand how database systems actually work.

Best book: *Readings in Database Systems*

Best video lectures: Joe Hellerstein's Berkeley CS 186

## Languages and Compilers

If you understand how languages and compilers actually work, you'll write better
code and learn new languages more easily.

Best book: *Crafting Interpreters*

Best video lectures: Alex Aiken's course on edX

## Distributed Systems

These days, *most* systems are distributed systems.

Best book: *Designing Data-Intensive Applications* by Martin Kleppmann.

Best video lectures: MIT 6.824

## Why learn computer science?

To become the type of software engineer who understands computer science in depth,
which brings the ability to do challenging, innovative work.

This type of software engineer will progress toward more fulling and well-remunerated
work over time.

These engineers find ways to learn computer science in depth,
whether through conventional mean or by relentlessly learning
throughout their careers.

## Subject Guides

### Computer Programming - Guide

"Introduction" to computer programming is beneficial even to more advanced
programmers who may have missed important concepts and
programming models while first learning to code.

The standard recommendation is the classic
*Structure and Interpretation of Computer Programs*,
which is available for free both as [a book](https://mitpress.mit.edu/sites/default/files/sicp/full-text/book/book.html),
and as a set of [MIT video lectures](https://ocw.mit.edu/courses/electrical-engineering-and-computer-science/6-001-structure-and-interpretation-of-computer-programs-spring-2005/video-lectures/).
While those lectures are great, our video suggestion is
actually [Brian Harvey's SICP lectures](https://archive.org/details/ucberkeley-webcast-PL3E89002AA9B9879E?sort=titleSorter)
(for the 61A course at Berkeley) instead. These are more
refined and better targeted at new students than the MIT lectures.

We recommend working through at least the first three chapters
of SICP and doing the exercises. For additional practice,
work through a set of small programming problems like
those on [exercism](https://exercism.io/).

Why SICP? Because the book is unique in its ability-
at least potentially-to alter your fundamental beliefs
about computers and programming. Not everybody will experience this.
Some will hate the book, others won't get past the first
few pages.
But the potential reward makes it worth trying.

### Computer Architecture - Guide

Computer Architecture-sometimes called "computer systems"
or "computer organization"-is an important firts look at
computing below the surface of software.
In our experience, it's the most neglected area among
self-taught software engineers.

Our favorite introductory book is *[Computer Systems: A Programmer's Perspective](http://csapp.cs.cmu.edu/3e/home.html)*
and a typical introductory computer architecture course
using the book [would cover](http://csapp.cs.cmu.edu/3e/courses.html)
most of chapters 1-6.

We love CS:APP for the practical, programmer-oriented approach.
While there's much more to computer architecture than
what is covered in the book, it serves as a great starting
point for those who'd like to understand computer systems
primarily in order to write faster, more efficient and more
reliable *software*.

For those who'd prefer a both a gentler introduction to the
topic and a balance of hardware and software concerns, we
suggest *The Elements of Computing Sytems*, also known as
"Nand2Tetris". This is an ambitious book attempting to give
you a cohesive understanding of how everything in a computer works.
Each chapter involves building a small piece of the overall
system, from writing elementary logic gates in HDL,
through a CPU and assembler, all the way to an application
the size of a Tetris game.

We recommend reading through the six first chapters of the book
and completing the associated projects. This will develop
your understanding of the relationship between the architecture
of the machine and the software that runs on it.

The first half of the book (and all of its projects) are available
for free from [the Nand2Tetris website](https://www.nand2tetris.org/).
It's also available as a [Coursera course with accompanying videls](https://www.coursera.org/learn/build-a-computer).

In seeking simplicity and cohesiveness, Nand2Tetris trades off depth.
In particular, two very important concepts in modern computer architecture are
pipelining and memory hierarchy, but both are mostly absent from the text.

Once you feel comfortable with the content of Nand2Tetris, we suggest either
returning to CS:APP, or considering Patterson and Hennessy's 
[Computer Organization and Design](https://www.amazon.com/Computer-Organization-Design-Fifth-Architecture/dp/0124077269?pldnSite=1),
and excellent and now classic text. Not every section in the book is essential,
we suggest following Berkeley's [CS61C course](https://inst.eecs.berkeley.edu//~cs61c/sp15/)
"Great Ideas in Computer Architecture" for specific readings.
The lecture notes and labs are available online, and past
lectures are [on the Internet Archive](https://archive.org/details/ucberkeley-webcast-PL-XXv-cvA_iCl2-D-FS5mk0jFF6cYSJs_).

> Hardware is the platform
-- *Mike Acton, Engine Director at Insomniac Games*

([watch his CppCon talk](https://www.youtube.com/watch?v=rX0ItVEVjHc))
