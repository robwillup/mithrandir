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

### Algorithms and Data Structures - Guide

We agree with decades of common wisdom that familiarity with
common algorithms and data structures is one of the most
empowering aspects of a computer science education.
This is also a great place to train one's general problem-solving abilities,
which will pay off in every other area of study.

There are hundreds of books available, but our favorite is
[The Algorithm Design Manual](https://www.amazon.com/Algorithm-Design-Manual-Steven-Skiena/dp/1848000693/?pldnSite=1)
by Steven Skiena. He clearly lovels algorithmic problem
solving and typically succeeds in fostering similar enthusiasm among his students
and readers.

For those who prefer video lectures, [Skiena generously provides his online](https://www3.cs.stonybrook.edu/~skiena/373/videos/).
We also really like Tim Roughgarden's course, available
[on Coursera](https://www.coursera.org/specializations/algorithms) and [elsewhere](http://timroughgarden.org/videos.html).

For practice, our preferred approach is for students to solve problems on [Leetcode](https://leetcode.com/).
These tend to be interesting problems with decent accompanying
solutions and discussions. They also help you test progress
against questions that are commonly used in technical interviews
at the more competitive software companies. We suggest solving
around 100 random leetcode problems as part of your studies.

Finally, we strongly recommend *[How to Solve It](https://www.amazon.com/How-Solve-Mathematical-Princeton-Science/dp/069116407X/?pldnSite=1)*
as an excellent and unique guide to general problem solving;
it's as applicable to computer science as it is to mathematics.

> I have only one method that I recommend extensively-
> it's called think before you write.
-- *Richard Hamming*

### Mathematics for Computer Science - Guide

In some ways, computer science is an overgrown branch of applied mathematics.
While many software engineers try-and to varying degrees succeed-
at ignoring this, we encourage you to embrace it with direct study.
Doing so successfully will give you an enormous competitive
advantage over those who don't.

The most relevant area of math for CS is broadly called "discrete mathematics",
where "discrete" is the opposite of "continuous" and is
loosely a collection of interesting applied math topics
outside of calculus. Given the vague definition, it's
not meaningful to try to cover the entire breadth of
"discrete mathematics". A more realistic goal is to build
a working understanding of logic, combinatorics and probability,
set theory, graph theory, and a little of the number theory
informing cryptography. Linear algebra is an additional
worthwhile area of study, given its importance in computer
graphics and machine learning.

Our suggested starting point for discrete mathematics is the set
of [lecture notes by László Lovász](http://www.cs.elte.hu/~lovasz/dmbook.ps).
Profesor Lovász did a good job of making the content approachable
and intuitive, so this servers as a better starting point
than more formal texts.

For a more advanced treatment, we suggest [Mathematics for Computer Science](https://ocw.mit.edu/courses/electrical-engineering-and-computer-science/6-042j-mathematics-for-computer-science-fall-2010/).
That course's video lectures are also [freely available](https://ocw.mit.edu/courses/electrical-engineering-and-computer-science/6-042j-mathematics-for-computer-science-fall-2010/video-lectures/),
and are our recommended video lectures for discrete math.

> If people do not believe that mathematics is simple,
> it is only because they do not realize how complicated life is.
-- *John von Neumann*

### Operating Systems - Guide

[Operating Systems Concepts](https://www.amazon.com/dp/1118063333/)
(the "Dinosaur book") and [Modern Operating Systems](https://www.amazon.com/dp/013359162X/)
are the "classic" books on operating systems. Both have attracted criticism for their
lack of clarity and general student unfriendliness.

*Operating Systems: Three Easy Pieces* is a good alternative
that's [freely available online](https://pages.cs.wisc.edu/~remzi/OSTEP/).
We particularly like the structure and readability of the book, and feel that
the exercises are worthwhile.

After OSTEP, we encourage you to explore the design decisions
of specific operating systems, through "{OS name} Internals"
style books such as
[Lion's commentary on Unix](https://www.amazon.com/Lions-Commentary-Unix-John/dp/1573980137/)
[The Design and Implementation of the FreeBSD Operating System](https://www.amazon.com/Design-Implementation-FreeBSD-Operating-System/dp/0321968972/)
and [Mac OS X Internals](https://www.amazon.com/Mac-OS-Internals-Systems-Approach/dp/0321278542/).
For Linux, we suggest Robert Love's fantastic
[Linux Kernel Development](https://www.amazon.com/Linux-Kernel-Development-Robert-Love/dp/0672329468).

A great way to consolidate your understanding of operating
system is to read the code of a small kernel and add features.
One choice is [xv6](https://pdos.csail.mit.edu/6.828/2016/xv6.html)
, a port of Unix V6 to ANSI C and x86, maintained for a courses at MIT.
OSTEP has an appendix of potential [xv6 labs](https://pages.cs.wisc.edu/~remzi/OSTEP/lab-projects-xv6.pdf)
full of great ideas for potential projects.

### Computer Networking - Guide

Given that so much of software engineering is on web servers and clients,
one of the most immediately valuable areas of computer science
is computer networking.
Our self-taught students who methodically study networking find that they finally
understand terms, concepts and protocols they'd been surrounded
by for years.

Our favorite book on the topic is
[Computer Networking: A Top-Down Approach](https://www.amazon.com/Computer-Networking-Top-Down-Approach-7th/dp/0133594149/?pldnSite=1).
The small projects and exercises in the book are well worth doing,
and we particularly like the "Wireshark labs",
which they have [generously provided online](https://gaia.cs.umass.edu/kurose_ross/wireshark.htm).

For those who prefer video lectures, we suggest Stanford's
[Introduction to Computer Networking course](https://www.youtube.com/playlist?list=PLvFG2xYBrYAQCyz4Wx3NPoYJOFjvU7g2Z)
previously available via Stanford's MOOC platform Lagunita,
but sadly now only available as unofficial playlists on Youtube.

> You can't gaze in the crystal ball and see the future.
> What the Internet is going to be in the future i what
> society makes it.
-- *Bob Kahn*

### Databases - Guide

It takes more work to self-learn about database systems than
it does with most other topics.
It's a relatively new (i.e. post 1970s) field of study with
strong commercial incentives for ideas to stay behind closed doors.
Additionally, many potentially excellent textbook authors
have preferred to join or start companies instead.

Given the circumstances, we encourage self-learners to generally
avoid textbooks and start with [recordings of CS 186](https://www.youtube.com/user/CS186Berkeley/videos),
Joe Hellerstein's database course at Berkeley, and to progress
to reading papers after.

One paper particularly worth mentioning for new students is
"[Architecture of a Database System](https://dsf.berkeley.edu/papers/fntdb07-architecture.pdf)",
which uniquely provides a high-level view of how relational
database management systems (RDBMS) work.
This will serve as a useful skeleton for further study.

*Readings in Database Systems*, better known as
[the databases "Red Book"](http://www.redbook.io/),
is a collection of papers compiled and edited by Peter Bailis,
Joe Hellerstein and Michael Stonebraker.
For those who have progressed beyond the level of the CS 186 content,
the Red Book should be your next step.

If you're adamant about using an introductory textbook,
we suggest [Database Management Systems](https://www.amazon.com/Database-Management-Systems-Raghu-Ramakrishnan/dp/0072465638/?pldnSite=1)
by Ramakrishnan and Gehrke. For more advanced students,
Jum Gray's classic
[Transaction Processing: Concepts and Techniques](https://www.amazon.com/Transaction-Processing-Concepts-Techniques-Management/dp/1558601902)
is worthwhile, but we don't encourage using this as a first resource.

Finally, data modeling is a neglected and poorly taught aspect
of working with databases. Our suggested book on the topic is
[Data and Reality: A Timeless Perspective on Perceiving and Managing Information
in Our Imprecise World](https://www.amazon.com/Data-Reality-Perspective-Perceiving-Information/dp/1935504215).

### Languages and Compilers - Guide

Most programmers learn languages, whereas most computer
scientists learn *about* languages. This gives the computer
scientist a distinctive advantage over the programmer,
even in the domain of programming! Their knowledge generalizes;
they are able to understand the operation of a new language
more deeply and quickly than those who have merely learned
specific languages.

Our suggested introductory text is the excellent [Crafting Interpreters](https://craftinginterpreters.com/contents.html)
by Bob Nystrom, available for free online. It's well organized,
highly entertaining, and well suited to those whose primary goal
is simply to better understand their languages and language
tools. We suggest taking the time to work through the whole thing, attempting
whichever of the "challenges" sustains your interest.

A more traditional recommendation is
[Compilers: Principles, Techniques & Tools](https://www.amazon.com/Compilers-Principles-Techniques-Tools-2nd/dp/0321486811?pldnSite=1),
commonly called "the Dragon Book". Unfortunately, it's not
designed for self-study, but rather for instructors to pick
out 1-2 semesters worth of topics for their courses.

If you elect to use the Dragon Book, it's almost essential
that cherry-pick the topics, ideally with the help of a
mentor. In fact, our suggested way to utilize the Dragon Book,
if you so choose, is as a supplementary reference for a video
lecture series. Our recommended one is [Alex Aiken's, on edX](https://www.edx.org/course/compilers).

> Don't be a boilerplate programmer. Instead, build tools
> for users and other programmers. Take historical note
> of textile and steel industries: do you want to build
> machines and tools, or do you want to operate those machines?
-- *Ras Bodik at the start of his compilers course*

### Distributed Systems - Guide

As computers have increased in number, they have also *spread*.
Whereas businesses would previously purchase larger and larger
mainframes, it's typical now for even very small applications
to run across multiple machines.
Distributed systems is the study of how to reason about the
trade-offs involved in doing so.

Our suggested book for self-study is Martin Kelppmann's
[Designing Data-Intensive Applications](https://www.amazon.com/Designing-Data-Intensive-Applications-Reliable-Maintainable-ebook/dp/B06XPJML5D/?pldnSite=1).
Far better than a traditional textbook, DDIA is a highly
readable book designed for practitioners, which somehow
avoids sacrificing depth or rigor.

For those seeking a more traditional text, or who would
prefer one that's available for free online, we suggest
Maarten van Steen and Andrew Tanenbaum's
[Distributed Systems, 3rd Editions](https://www.distributed-systems.net/index.php/books/ds3/).

For those who prefer video, an excellent course with videos
available online is [MIT's 6.824](https://www.youtube.com/watch?v=cQP8WApzIQQ&list=PLrw6a1wE39_tb2fErI4-WkMbsvGQk9_UB),
a graduate course taught by Robert Morris with readings
available [here](https://pdos.csail.mit.edu/6.824/schedule.html).

No matter the choice of textbook or other secondary resources,
study of distributed systems absolutely mandates reading papers.
A good list is [here](http://dsrg.pdos.csail.mit.edu/papers/),
and we would highly encourage attending your local
[Papers We Love](https://paperswelove.org/) chapter.

## Other Topics - AI/graphics/pet-topic-X?

We've tried to limit our list to computer science topics that
we feel *every practicing software engineer* should know,
irrespective of specialty or industry, but with a focus on systems.
In our experience, thse will be the highest ROI topics for the
overwhelming majority of self-taught engineers and bootcamp grads,
and provide a solid foundation for further study.
Subsequently, you'll be in a much better position to pick up
textbooks or papers and learn the core concepts without much guidance.
Here are our suggested starting points for a couple of common "electives":

* For artificial intelligence: do
[Berkeley's intro to AI course](http://ai.berkeley.edu/home.html) by watching the
videos and completing the excellent Pacman projects.
As a textbook, use Russell and Norvig's
*Artificial Intelligence: a Modern Approach*.
* For machine learning: do Andrew Ng's Coursera course.
Be patient, and make sure you understand the fundamentals
before racing off to shiny new topics like deep learning.
* For computer graphics: work through
[Berkeley's CS 184](https://inst.eecs.berkeley.edu//~cs184/fa12/onlinelectures.html)
material, and use
[Computer Graphics: Principles and Practice](https://www.amazon.com/Computer-Graphics-Principles-Practice-3rd/dp/0321399528)
as a textbook.

## Don't Give Up

We're very confident that you could teach yourself everything above.
