# What is a hash table?

**References:**

* [https://khalilstemmler.com/blogs/data-structures-algorithms/hash-tables/]

A hash table is a data structure that you can use to store data in key-value format with direct
access to its items in constant time = O(1).

Hash tables are said to be associative, which means that for each key, data occurs at most once.

> The above sentence probably means that you will not find different values with the same key.

Hash tables let us implement things like phone books or dictionaries; in them we store the
association between a value and its key.

## Why use hash tables?

The most valuable aspect of a hash table over abstract data structures is its speed to perform
insertion, deletion, and search operations. Hash tables can do them all in constant time O(1).
Hash tables can perform nearly all methonds, except list, very fast in O(1).

Because of this efficiency, you'll find hash tables to be pretty dang useful for many use cases.
And if you look carefully, you'll notice that they're actually implemented in a variety of places
throughout your tools like your databases, caches, data-fetching libraries, etc.

## How do hash tables work?

There are four distinct aspects to discuss how hash tables work:

### Storage

A hash table is an abstract data type that relies on using a more primitive data type (such as
an array or an object) to store the data. There are slight implementation implications depending on
which underlying data type is used.

### Key-value pairs

To store data as a value in a hash table, we need some way to identify it uniquely. We need a key.
Sometimes data contains an individual property that can very naturally assume responsibility for the
key. For example, if we assume that a `user` has an `email` and a `userId`, then either the `email`
or the `userId` could be used as they key if they are guaranteed to be unique.

```json
{
	email: 'rob@test.com',
	name: 'Robson',
	userId: 1,
	city: 'Curitiba'
}
```

Other times, we're working with data where uniqueness is a little more ambiguous. For example,
a dictionary may contain two different entries for the word `chair`, the furniture and the person.

An option is to organize the data that we want to store in the hash table a little differently,
for example in an object.

### Hash functions

Once we have the key-value pair, we pass them to the hash table to store the data for later
retrieval. Hash tables need a hash function to determine how the table should store the data,
and this is one of the standard hash table operations.

The hash function requires both key and the value.

The key contains the logic that determines what index the value will live at withing the underlying
data structure (array or object).

* When we use objects as storage, the index is often just some string or integer version of the key
* When we use a fixed-size array as storage, we need to generate a hash code: a value that maps the key to a valid index/position in the array.

After deciding on an index, depending on how we've implemented it, the hash function can insert or merge
the value at the position specified by the index.

For example, suppose our dictionary hash table had saved only one of the possible definitions for the word
`chair` so far. In that case, we could configure the hash function to merge additional non-duplicate
definitions to the `chair` value's definition when we try to hash it.

Alternatively, we could also configure it to throw an error if the hash function encounters a duplicate.
In implementations where we're not merging data, this is generally referred to as a collision.

### Methods/operations (what can it do?)

We know that one of the operations a hash table should perform is to hash an item - that is, to create
the unique index based on a key. What other things should it be capable of?

At the base minimum, a hash table should be able to do the following:

* add - add a key-value pair
* get - get a value by key
* remove - remove a value by key
* list - get all the keys or key-value pairs
* count - count the number of items in the table.

## How to choose the key in a hash table

To determine what should be the key in the key-value pair, it must be:

* Unique
* known by the consumers so that data can be retrieved in constant time.

## Basic hash table examples

Most general-purpose languages have few different ways to implement a hash table.

### Object implementation

Create a plain old JavaScript object:

```javascript
let table = {};
```

There are a few ways the API allows us to add a key-value pair to our table. The first is to use the
dot property:

```javascript
table.yellow = 'color between green an orange';
console.log(table);
```

Another way is to use object brackets (which is useful when our key is numeric):

```javascript
table['green'] = 'color between blue and yellow';
console.log(table['green']);
console.log(table.green);
```

To see if a value is present, we use the `hasOwnProperty` method that's a part of every object in JavaScript:

```javascript
table.hasOwnProperty('green'); // true
```

When you're solving one-off algorithmic problems, you can use objects to get the functionality you
need from a hash table quickly. However, there are better ways to design hash tables to be easier
to use, and one such route is through encapsulation.

## Encapsulated hash table

When you have related functionality, it's generally a good idea to encapsulate it withing a cohesive
object or a collection of functions. Here, we create a `HashTable`