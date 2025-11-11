## Clean Code Tips

Your code should be kept as short and simple as possible. This means that unnecessarily complex logic should not be
used just for the sake of making the code look more sophisticated. That will actually prove the opposite.
Break down your code into smaller functions or methods that perform specific tasks.

It is also important to study and know the conventions of the programming language you're using. Other developers will
be looking for these conventions in your code which would help them understand the code more quickly. These conventions
are expected and should be followed. So you need to become more familiar with the style guides and code standards and
strive to follow them diligently.

Variables and functions should have clear and descriptive names. Cryptic abbreviations or single-letter names must be
avoided. The name should convey the purpose of that function or variable.

In a clean code base comments should be only used to explain the intent and reasoning for some decisions. For example,
there might be a particular function that could have been implemented in a different way, but due to certain
constraints you had to use an alternative method. That is a good opportunity to add a concise and informative comment
explaining the rationale.

Code formatting is important and white spaces should be used to improve readability. Utilize white spaces effectively
by adding line breaks and indentation to separate logical sections, loops, and conditionals.

Avoid unnecessary code and eliminate redundancy by removing commented-out sections, unused variables and functions.
This will keep your code base lean and free of clutter as well as reduce the chances of bugs.

I love code refactoring. In fact, I think I like refactoring more than writing the original code. That's because it
feels like going over a video game level or act and improving over what you did before, finishing the level faster,
taking less hits, and getting the most prizes and collectibles to achieve a higher score. Essentially, that is what
code refactoring is all about, improving that function or code block so that it becomes more performant, has less
bugs, and becomes cleaner and clearer.

If you have played Sonic 2 on the Mega Drive or Genesis, or Sonic Mania, you will know what is the "1.5 player mode".
In that mode, the second player is controlled by the CPU and it will help you with certain things at times, however,
you are still in control of the 1st player. This is a good analogy for a linter. Automated tools such as linters can be
allies in your quest for clean code. Linters analyze your code for potential issues, style violations, and command
programming mistakes. They provide actionable suggestions to improve code quality and maintain consistency. Integrate
a linter into your development workflow and configure it according to your project's specific requirements.

Tests are essential and your code needs proper test coverage. Writing automated tests helps you verify that your code
behaves as intended and guards against regressions. Adopt a test-driven development (TDD) approach.
Well-designed tests serve as a living documentation and provide confidence when making changes or refactoring code.

Feedback from others is important and so having other developers look at your code is great. This can be done by
developers at any level. More experienced developers will challenge your code and spot parts for improvement, and less
experienced developers can do that too as well as demonstrate how hard your code is to understand for on-boarding devs.

### functions

Functions should either do something or answer something, but not both.

#### Small

Functions should be as small as possible. The code inside `if`, `else`, and loops should ideally be function calls.

The indent level of functions should not be more than one or two.

#### Do ONE thing

If a function does only steps that are one level below the stated name of the function, then the function is doing
one thing only.

> Another way to know if a function is doing more than one thing is if you can extract another function from it with
> a name that is not merely a restatement of its implementation.

> Mixing levels of abstractions in functions is always confusing.

##### The Stepdown Rule

Your code should be read from top to bottom from higher levels of abstractions to lower levels. It should be possible
to read it like a series of TO paragraphs.

*To print this page you need to call the RenderPage() function
To render a page, call the FetchData() function...*

##### Arguments

From best to worse, the ideal number of arguments:

- niladic - zero arguments
- monadic - one arguments
- dyadic  - two arguments
- triadic - should be avoided
- polyadic - requires very special justification - but still shouldn't be used.

###### Flag Arguments

Flag arguments are ugly. Passing a Boolean into a function is a truly terrible practice. It loudly proclaims that the
function does more then one thing.

##### Have no Side Effects

Side effects are lies. Your function promises to do one thing, but it does other hidden things as well.

##### Output Arguments

> Anything that forces you to check the function signature is equivalent to a double-take. It's a cognitive break and
> should be avoided.

Output arguments should be avoided. If your function needs to change the state of something, have it change the state
of its owning object.

#### How do you write functions like this?

When I write functions, they come out long and complicated. They have lots of indenting and nested loops. They are long
argument lists. The names are arbitrary, and there is duplicated code. But I also have a suite of unit tests that
cover every one of those clumsy lines of code.

Then I massage and refine the code, splitting out functions, changing names, eliminating duplication. I shrink the
methods and reorder them. Sometimes I break out whole classes, all the while keeping the tests passing.

In the end, I wind up with functions that follow the rules in this chapter. I don't write them that way to start.
I don't think anyone could.

### Switch Statements

Switch statements by definition will span across multiple lines. So they must be buried in a low-level class and never
repeated. We do this with polymorphism.



