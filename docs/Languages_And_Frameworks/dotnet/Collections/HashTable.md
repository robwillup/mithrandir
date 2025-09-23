# HashTable

In C# you do not need to code your own hash table. Instead, you can use the existing implementation
from `System.Collections.Hashtable`.

A `Hashtable` collections stores a (key, value) pair and uses the Key to hash and obtain the storage
location. The Key is immutable and cannot have duplicate entries in the Hashtable. This sample uses
several instances of a simple Person class to store in a Hashtable. The last name is used as the key.