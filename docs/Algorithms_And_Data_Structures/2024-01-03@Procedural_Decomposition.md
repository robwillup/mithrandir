## Procedural Decomposition

> Reference: [Embedded in Academia](https://blog.regehr.org/archives/942)

There is a pattern where students who are having trouble getting their code to work often have not done enough procedural decomposition. That is, they have
failed to appropriately break up a large, difficult programming problem into a collection of smaller ones.
I've come to believe that procedural decomposition for hard problems is probably one of those skills that we can keep getting better at for as long as we care to keep writing code.

The tricky thing is recognizing that it's time to think about breaking up the problem rather than attacking it all at once. `The temptation to keep banging on a function that seems to be almost correct is strong.`

We don't do enough procedural decomposition. When I catch myself writing poorly abstracted code, I try to analyze what I'm thinking and usually it boils down to a mixture of `laziness` ("if I can just get this stupid function to work, I won't have to figure out how to break it up into pieces") and overconfidence in my ability to deal with complex case analysis all at once.

This thing we are here calling "procedural decomposition" is sometimes also referred to as "procedural abstraction", "functional abstraction", or "functional decomposition".

### It's not just for reuse

A lot of us were taught that the main reason we break code into functions is to facilitate reuse. It's usually not too hard to notice when this kind of decomposition is necessary: we just get tired of writing quicksort from scratch after the second or third time. In contract, procedural decomposition of a single problem is more difficult because it's not as obvious where to introduce abstraction.
