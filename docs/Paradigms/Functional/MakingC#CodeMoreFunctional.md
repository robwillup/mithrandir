## Making C# Code More Functional

### Pure Functions

* Understanding Side Effects

Side effect is the primary obstacle to achieving functional design.

Side effect is anything that happens when a function executes, which makes it possible for us to tell that the function did execute.

Changing a system's state is a typical side effect. Like calling a function in an object and then noticing a change made in field in the object. Execution of any function could depend on system state and there's nothing odd about that. Odd things begin to happen when a function with side effects is called again. Then in the next call it will face the modified system state. In general we could expect that the repeated execution would be affected by previous executions if the function had side effects.

Therefore calling methods in a typical object-oriented code has another aspect: the order in which methods execute. The order of invocations affects the outcome. That makes it harder for us programmer to ponder about our code, to think about the ways our application will run.

* Pure functions

Results only depends on inputs.

Produces no observable side effects.

The concert is state changes that could affect subsequent function calls
