## Architectural Styles and the Design of Network-based Software Architectures by Roy Thomas Fielding

Software architecture research investigates methods for determining how best to partition a system, how components identify and comunicate with each other, how information is communicated, how elements of a system can evolve independently, and how all of the above can be described using formal and informal notations.

An architectural style is a named, coordinated set of architectural constraints.

Representational State Transfer (REST) emphasizes scalability of component interactions, generality of interfaces, independent deployment of components, and intermidiary components to reduce interaction latency, enforce security, and encapsulate legacy systems.

- *The lines below are impressive. This was written in 2000* - Robson

As predicted by Perry and Wolf, software architecture has been a focal point for software engineering research in the 1990s. The complexity of modern software systems have necessitated a greater emphasis on componetized system, where the implementation is partitioned into independent components that communicate to perform a desired task. Software architecture research investigates methods for determining how best to partion a system, how components identify and communicate with each other, how information is communicated, how elements of a system can evolve independently, and how all of the above can be described uing formal and informal notations.

The guideline "form follows function" comes from hundred years of experience with failed building projects, but is often ignored by software practitioners.

How often we see software projects begin with adoption of the lates fad in architectural design, and only later discover whether or not the system requirements call for such an architecture. Design-by-buzzword is a common occurrence. At least some of this behavior within the software industry is due to a lack of understanding of why a given set of architectural constraints is useful. In other words, the reasoning behind good software architectures is not apparent to designers when those architectures are selected for reuse.

## Software Architecture

### 1.1 Run-time Abstraction

    A **software architecture* is an abstraction of the run-time elements of a software system during some phase of its operation. A system may be composed of many levels of abstraction and many phases of operation, each with its own software architecture.
