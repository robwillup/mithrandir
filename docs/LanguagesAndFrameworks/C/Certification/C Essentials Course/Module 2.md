# Module 2

## Floating-point numbers

It's time to talk about another type, witch is designed to represent and store numbers
that (as a mathematician would say) have a **non-empty decimal fraction**.

You can **omit** zero when it's the only digit in front of or after the decimal point.
(But you still need to use the point `.`)

You can say that the **point makes a double**.

When you want to use any numbers that are very large or very small, you can use
**scientific notation**. Take, for example, the speed of light, expressed in meters
per second. Written directly it would look like this:

```code
300000000
```

To avoid tediously writing so many zeros, physics textbooks use an abbreviated form,
which you have probably already seen:

```code
3*10^8
```

It reads: "three times ten to the power of eight"

In the C language, the same effect is achieved in a slightly different form -
take a look:

```c
3E8
```

The letter *E*(you can also use the lower case letter e - it comes from the word
exponent) is a concise version of the phrase "times ten to the power of".

Note:

* the exponent **must** be an integer.
* the base **can** be an integer.

Let's see how we use this convention to record numbers that are very small
(in the sense of their absolute value, which is close to zero).
A physical constant called *Planck's constant* (and denoted as *h*) has,
according to the textbooks, the value of:

6.62607 x 10<sup>-34</sup>

If you would like to use it in a program, you would write it this way:

```c
6.62607E-34
```

Declaring floating-point variables:

```c
float PI, Field;
```

```c
PI = 3.14;
```

The transformation from type int into float is always possible and feasible,
but in some case can cause a loss of accuracy.

The opposite situation will result in loss of accuracy.
The other aspect of this operation is that a *float* cannot always be converted
into an *int*. Integer variables (like floats) have a **limited capacity**.

In such cases, it isn't clear what will happen during the assignment.
Certainly, a **loss of accuracy** will occur, but the value assigned to the
variable is not known in advance. In some systems, it may be
the maximum permissible *int* value, while in others an error occurs,
and in others still the value assigned may be completely random.

This is what we call an **implementation dependent issue**. It's the second
(and uglier) face of **software portability**.

## Operators

An **operator** is a symbol of the programming language which is able to
**operate** on the values.

### Multiplication

An asterisk (*) is a **multiplication** operator.

### Division

A slash (/) is a **division** operator. The value in front of the slash is the
**dividend**, the value behind the slash, a **divisor**.

#### Division by zero

As you've probably guessed, division by zero is strictly **forbidden**, but the
penalty for violating this rule will come to you at different moments.
If you date to write something like that, a judgement will be issued by the
compiler - you may get a **compilation error** and you won't be able to run your
program.

### Addition

The **addition** operator is the `+` (plus) sign, which is one that we already
know from mathematics.

### Subtraction

The **subtraction** operator is obviously the `-` (minus) sign. This operator also
has another meaning, it can change the sign of a number.

This is a great opportunity to show you a very important distinction between
**unary** and **binary** operators

#### Unary minus

In "subtracting" applications, the minus operator expects two arguments:
the left (a **minuend** in arithmetic terms) and the right (a **subtrahend**).
For this reason, the subtraction operator is considered one of the binary operators,
just like the addition, multiplication and division operators.
But the minus operator may also be used in a different way:

```c
int i, j;

i = -100;
j = -1;

// j = 100
```

#### Unary plus

Unary plus, which could be used to preserve the sign, is syntactically correct,
but using it doesn't make much sense.

### Remainder

The **remainder** operator (or **modulus** as it is often called) is quite a
peculiar operator, because it has no equivalent among traditional arithmetic operators.
Its graphical representation in the C language is the following character: `%` (percent),
which may look a bit confusing. It's a binary operator and both arguments cannot
be floats (don't forget this!).

```c
int i, j, k;

i = 13;
j = 5;

k = i % j;
```

The `k` variable is set to the value of `3`, because 2*5+3=13.

### Prefix and Postfix Operators

Hand-written notes.

### Shortcut operators

```c
i = i * 2;
```

In the C language, there's a shortened way to write this operation:

```c
i *= 2;
```

|||
|:-:|:-:|
|i = i + 2 * j;| i += 2 * j; |
| Var = Var / 2; | Var /= 2; |
| Rem = Rem % 10; | Rem %= 10; |
| j = j - (i + Var + Rem);|j -= (i + Var + Rem);|

## The *char* Type

To store and manipulate characters, the C language provides a special
type of data. This type is called a **char**.

```c
char Character;
```

### ASCII code

Computers stores characters as numbers. Every character used by a computer
corresponds to a number, and vice versa. This system of assignments includes more
characters than you would probably expect. Many of them are invisible to humans
but essential for computers.

Some of these characters are called **white spaces**, while others are named
**control characters**, because their purpose is to control the input/output devices.

Out of the need to introduce a universal and widely accepted standard implemented
by (almost) all computers and operating systems all over the world, **ASCII**
(which is a short for *American Standard Code for Information Interchange*)
was created and it is the most widely used system in the world, and it's safe to
assume that nearly all modern devices (like computers, printers,
mobile phones, tablets, etc.) use this code.

The code provides space for 256 different characters.

### Literal

The literal is a symbol which **uniquely identifies its value**.
Some prefer to use a different definition: the **literal means itself**.

* `Character` - this is not a literal; you cannot guess what value is currently assigned
* `'A'` - this is a literal; you can immediately guess its value.
* `100` - this is a literal too, of the `int` type.

### Character Literals

The `\` character (called *backslash*) acts as an **escape character**,
because by using the `\` we can escape from the normal meaning of the character
that follows the slash.

### Escape characters

The C language allows us to escape in other circumstances too. Let's start with
those that denote literals representing white spaces.

`\n` - denotes a **transition to a new line** and is sometimes called an
**LF (Line Feed)**, as printers react to this character by pulling out the paper
by one line of text.

`\r` - denotes the **return to the beginning of the line** and is sometimes called
a **CR (Carriage Return** - "carriage" was the synonym of a "print head" in the
typewriter era); printers respond to this character as if they are told to re-start
printing from the left margin of the already printed line.

`\a` (as in **alarm**) is a relic of the past when teletypes were often
used to communicate with computers; sending this character to a teletype turn on
its ringer; hence the character is officially called **BEL** (as in bell);

`\0` - called **nul** (from the Latin word **nullus** - none) is a
character that **does not represent any character**; despite first impressions
it could be very useful.

### Char values are int values

There's an assumption in the C language that may seem surprising at first glance:
the `char` type is treated as a special kind of `int` type. This means that:

* You can always assign a `char` value to an `int` variable;
* You can always assign an `int` value to a `char` variable, but if the
value exceeds 255 (the top-most character code in ASCII), you may expect a loss
of value;
* The value of the `char` type can be subject to the same operators as the data
of type `int`.

## Logical operators

**is equal to?**

```c
1 == 0; // false
```

**is not equal to?**

```c
1 != 0; // true
```

**is greater than?**

```c
1 > 0; // true
```

**is greater than or equal to?**

```c
1 >= 0; // true
```

**is less than or equal to?**

```c
1 < 0; // false

1 <= 0; // false
```

### How to use the answers we get?

There are at least two possibilities:
First, we can store it in a variable and make use of it later.

```c
int answer, value1, value2;
answer = value1 >= value2;
```

If the answer is true, the computer will assign `1` to `answer`. Otherwise,
the variable `answer` will be assigned `0`.

The second way is by using `if`.