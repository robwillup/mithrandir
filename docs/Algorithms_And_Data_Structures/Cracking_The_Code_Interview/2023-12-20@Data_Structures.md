# Data Structures

> Notes taken while reading the book [Cracking the Coding Interview](https://www.google.com/books/edition/_/jD8iswEACAAJ?hl=en)

## Hash Tables

A hash table is a data structure that maps keys to values for highly efficient lookup. So this is similar to dictionaries.
Hash tables benefit from fast data retrieval and are foundational to standard tools and techniques like caching and database indexing.
There are a number of ways of implementing this. Here let's see a common but simple interpretation.
In this implementation an array of linked lists and a hash code function are used. To insert a key - which
might be a string or essentially any other data type - and value, this is what is done:

1. First, compute the key's hash code, which will usually be an int or long. Note that two different keys
could have the same hash code, as there may be an infinite number of keys and a finite number of ints.
2. Then, map the hash code to an index in the array. This could be done with something like `hash (key) % array_length`. Two different hash codes could, of course, map to the same index.
3. At this index, there is a linked list of keys and values. Store the key and value in this index. We must
use a linked list because of collisions: you could have two different keys with the same hash code, or two
different hash codes that map to the same index.

To retrieve the value by its key, you repeat this process. Compute the hash code from the key, and then
compute the index from the hash code. Then, search through the linked list for the value with this key.

If the number of collisions is very high, the worse case runtime is O(N), where N is the number of keys. However, we generally assume a good implementation that keeps collisions to a minimum, in which case the
look up time is O(1).

Alternatively, we can implement the hash table with a balanced binary search tree. This gives us an O(log N) lookup time. The advantage of this is potentially using less space, since we no longer allocate a large
array. We can also iterate through the keys in order, which can be useful sometimes.

## ArrayList & Resizable Arrays

In some languages, arrays (often called lists in this case) are automatically resizable. The array or list
will grow as you append items. In other languages, like Java, arrays are fixed length. The size is defined when you create the item.

When you need an array-like data structure that offers dynamic resizing, you would usually use an ArrayList. An ArrayList is an array that resizes itself as needed while still providing O(1) access. A
typical implementation is that when the array is full, the array doubles in size. Each doubling takes O(n),
but happens so rarely that its amortized insertion time is still O(1).

```java
ArrayList<String> merge(String[] words, String[] more) {
	ArrayList<String> sentence = new ArrayList<String>();
	for (String w : words) sentence.add(w);
	for (String w : more) sentence.add(w);
	return sentence;
}
```

This is an essential data structure for interviews. Be sure you are comfortable with dynamically resizable
arrays/lists in whatever language you will be working with. Note that the name of the data structure as well as the "resizing factor" (which is 2 in Java) can vary. In C#, the data structure is also called
`ArrayList` and it also doubles every time capacity is exceeded.

```csharp
ArrayList resizableArray = new(1); // The '1' passed to the constructor is the initial capacity. Optional.
resizableArray.Add(1);
Console.WriteLine($"Capacity: {resizableArray.Capacity} - Count: {resizableArray.Count}");
resizableArray.Add(2);
Console.WriteLine($"Capacity: {resizableArray.Capacity} - Count: {resizableArray.Count}");
resizableArray.Add(2);
Console.WriteLine($"Capacity: {resizableArray.Capacity} - Count: {resizableArray.Count}");
Console.WriteLine(resizableArray[2]);

// Result
// Capacity: 1 - Count: 1
// Capacity: 2 - Count: 2
// Capacity: 4 - Count: 3
// 2
```

Why is the amortized insertion runtime O(1)?

Suppose you have an array of size N. We can work backwards to compute how many elements we copied at each
capacity increase. Observe that when we increase the array to K elements, the array was previously half
that size. Therefore, we need to copy K/2 elements.

final capacity increase : n/2 elements to copy
previous capacity increase: n/4 elements to copy
...
second capacity increase: 2 elements to copy
first capacity increase: 1 element to copy

Therefore, the total number of copies to insert N elements is roughly n/2 + n/4 + ... + 2 + 1, which
is just less than N.

If the sum of this series isn't obvious to you, imagine this: Suppose you have a kilometer-long walk to
the store. You walk 0.5 kilometers, and then 0.25 kilometers, and then 0.125 kilometers, and so on. You
will never exceed one kilometer (although you will get very close to it).

Therefore, inserting N elements takes O(N) work total. Each insertion is O(1) on average, even though
some insertions take O(N) time in the worst case.

## StringBuilder

Imagine you were concatenating a list of strings, as shown below. What would the runtime of this code
be? For simplicity, assume that the strings are all the same length (call this X) and that there are n
strings.

```java
String joinWords(String[] words) {
	String sentence = "";
	for (String w : words) {
		sentence = sentence + w;
	}
	return sentence;
}
```

On each concatenation, a new copy of the string is created, and the two strings are copied over,
character by character. The first iteration requires us to copy x characters. The second iteration
requires copying 2x characters. The third iteration requires 3x, and so on. The total time therefore is
O(x + 2x + ... + nx).

Why is it O(xn²)?Because 1 + 2 + ... + n equals n(n + 1) / 2, or O (N²).

StringBuilder can help you avoid this problem. StringBuilder simply creates a resizable array of all the
strings, copying them back to string only when necessary.

```java
String joinWords(String[] words) {
	StringBuilder sentence = new StringBuilder();
	for (String w : words) {
		sentence.append(w);
	}
	return sentence.toString();
}
```

A good exercise to practice strings, array, and general data structures is to implement your own
version of StringBuilder, HashTable and ArrayList. (I've done this in C#.)

## Interview Questions

1.1 **Is Unique:** Implement an algorithm to determine if a string has all unique characters. What if you
cannot use additional data structures?

**SOLUTION**:

You should first ask your interviewer if the string is an ASCII string or a Unicode string. Asking this
question will show an eye for detail and a solid foundation in computer science.

Every time you need to implement an algorithm to deal with strings think about ENCODING! Is it ASCII?
Is it Unicode?

The enconding information is used when converting a byte array into a string, because you need to know
how certain bytes should be represented. For example, the same byte array could be either one of the
following strings depending on whether it was decoded using ASCII or UTF-8:

```csharp
byte[] bytes = File.ReadAllBytes("Pt.txt");

string portuguese = System.Text.Encoding.UTF8.GetString(bytes, 0, bytes.Length);

Console.WriteLine(portuguese);

// ASCII
// Foz do Igua??u,
// S??o Jos?? dos Pinhais.

// UTF-8
// Foz do Iguaçu,
// São José dos Pinhais.
```

When solving an algorithm always check if there are alternatives to know the result before extra
processing.
For example, if you need to check if an array of characters has only unique characters, when the array
length is larger than the number of available characters for a particular encoding, then there is no need for extra processing, you
already know that there are repeated characters and you can return the result.

Example in C# solving the IsUnique problem:

```csharp
Console.WriteLine("Checking if a string has only unique characters: \n");

byte[] bytes = System.Text.Encoding.ASCII.GetBytes(File.ReadAllText("Pt.txt"));

if (bytes.Length > 128)
{
	Console.WriteLine(false);
	Environment.Exit(0);
}

for(int i = 0; i < bytes.Length; i++)
{
	for (int j = i + 1; j < bytes.Length; j++)
	{
		if (bytes[i] == bytes[j])
		{
			Console.WriteLine(false);
			Environment.Exit(0);
		}
	}
}
Console.WriteLine(true);
```

**1.2 Check Permutation:** Given two strings, write a method to decide if one is a permutation of the
other.

* Do both strings have the same encoding?
* Is the comparison case sensitive?
* Is whitespace significant?

My implementation in C#:

```csharp
if (s.Length != s2.Length)
{
    Console.WriteLine(false);
    Environment.Exit(0);
}

Hashtable hashtable = new(s.Length);

for (int i = 0; i < s.Length; i++)
{
    hashtable[s[i]] = (int)(hashtable[s[i]] ?? 0) + 1;
}

for (int i = 0; i < s2.Length; i++)
{
    hashtable[s2[i]] = (int)(hashtable[s2[i]] ?? 0) - 1;
    if ((int)(hashtable[s2[i]] ?? 0) < 0)
    {
        Console.WriteLine(false);
        Environment.Exit(0);
    }
}

for (int i = 0; i < hashtable.Count; i++)
{
    if ((int)(hashtable[i] ?? 0) != 0)
    {
        Console.WriteLine(false);
        Environment.Exit(0);
    }
}

Console.WriteLine(true);
```

**1.3 URLify:** Write a method to replace all spaces in a string with `%20`. You may assume that the
string has sufficient space at the end to hold the additional characters, and that you are given the
"true" length of the string.

**1.4 Palindrome Permutation:** Given a string, write a function to check if it is a permutation of
a palindrome. A palindrome is a word or phrase that is the same forwards and backwards. A permutation
is a rearrangement of letters. The palindrome does not need to be limited to just dictionary words.
