## Clean Code Tips

Your code should be kept as short and simple as possible. This means that unnecessarily complex logic should not be used just for the sake of making
the code look more sophisticated. That will actually prove the opposite. Break down your code into smaller functions or methods that perform
specific tasks.

It is also important to study and know the conventions of the programming
language you're using. Other developers will be looking for these
conventions in your code which would help them understand the code more
quickly. These conventions are expected and should be followed. So you need
to become more familiar with the style guids and code standards and strive
to follow them diligently.

Variables and functions should have clear and descriptive names. Cryptic
abbreviations or single-letter names must be avoided. The name should convey
the purpose of that function or variable.

In a clean codebase comments should be only used to explain the intent and
reasoning for some decisions. For example, there might be a particular
function that could have been implemented in a different way, but due to
certain constraints you had to use an alternative method. That is a good
opportunity to add a concise and informative comment exmplaining the rationale.

Code formatting is important and whitespaces should be used to improve
readability. Utilize whitespaces effectively by adding line breaks and
indentation to separate logical sections, loops, and conditionals.

Avoid unnecessary code and eliminate redundancy by removing commented-out
sections, unused variables and functions. This will keep your codebase
lean and free of clutter as well as reduce the chances of bugs.

I love code refactoring. In fact, I think I like refactoring more than
writing the original code. That's because it feels like going over a
video game level or act and improving over what you did before, finishing
the level faster, taking less hits, and getting the most prizes and
collectibles to achive a higher score.
Essentially, that is what code refactoring is all about, improving that
function or code block so that it becomes more performat, has less bugs,
and becomes cleaner and clearer.

If you have played Sonic 2 on the Mega Drive or Genesis, or Sonic Mania, you
will know what is the "1.5 player mode". In that mode, the second player is
controlled by the CPU and it will help you with certain things at times, however, you are still in controll of the 1st player. This is a good analogy for
a linter. Automated tools such as linters can be invaluable allies in your
quest for clean code. Linters analyze your code for potential issues,
style violations, and command programming mistakes. They provide actionable
suggestions to improve code quality and maintain consistency. Integrate a
linter into your development workflow and configure it according to your
project's specific requirements.

Tests are essential and your code needs proper test coverage. Writing
automated tests helps you verify that your code behaves as intended and
guards agains regressions. Adopt a test-driven development (TDD) approach.
Well-designed tests serve as a living documentation and provide confidence
when making changes or refactoring code.

Feedback from others is important and so having other developers look at
your code is great. This can be done by developers at any level. More
experienced developers will chanllenge your code and spot parts for
improvement, and less experienced developers can do that too as well as
demonstrate how hard your code is to understand for onboarding devs.
