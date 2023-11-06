# Introducing Streams

>"... a stream is a sequence of bytes that you can use to read from and write to a backing store, which can be one of several storage mediums"

*Microsoft Documentation*

https://docs.microsoft.com/en-us/dotnet/standard/io


`Application Code`

↑

Write bytes and Read bytes

↓

`Stream` == Abstraction of a sequence of bytes

↑

↓

`Backing Store` == An specific type of stream knows how to talk to its backing store

Another abstraction between the `Application Code` and the `Stream` is the `Stream reader / writer`. This allows our applications to write strings or integers, for example, instead of having to manipulate bytes directly.

### Examples of Backing Stores

* Files
* input/output device/hardware
* TCP/IP sockets


## Benefits of Using Streams

* Incremental data processing
* Abstraction of backing store
* Flexibility / control
* Random access / seeking
* Composability / pipelines

## Stream Classes

`Stream` == Abstract base class for all streams

↑

`FileStream` == backing store: file | `MemoryStream` == backing store: memory | Others




