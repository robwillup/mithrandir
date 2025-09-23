# Linked List

A linked list is a container where data is stored in nodes consisting of a single data item and a reference to the next
node.

Some say this is one of the most important data structures in computer science.

Some people compare a linked list to a chain. Each node in the list would be a link in the chain.
There are no shortcuts, to get to an item in the middle of the chain you need to go from one link to the next.


## The Node

Each link in the chain is called a node.

A node contains a single value.
Second it contains a reference that refers to the next node in the list.

_________
|       |
| value |
---------
| next  |
---------

```CSharp
class Node
{
    public Node Next;
    public int Value;

    public Node(int value)
    {
        this.Value = value;
        this.Next = null;
    }
}
```

### Connecting Nodes into a list

```CSharp
Node head = new Node(1);
head.Next = new Node(2);
head.Next.Next = new Node(3); // head.Next.Next.Next at this point is null
```

### Singly Linked List

A linked list that provides forward iteration from the start to the end of the list.

### Doubly Linked List

A linked list that provides forward iteration from the start to the end of the list, and reverse iteration, from end to
start.

```CSharp
class Node
{
    public Node Previous;
    public Node Next;
    public int Value;

    public Node(int value)
    {
        this.Value = value;
        this.Previous = null;
        this.Next = null;
    }
}
```

### The Add method for a doubly linked list

I was wondering, where does the DoublyLinkedList class keep the nodes? Does it use an array?
Well, what do you know?
It just creates objects out in the wild. It's the next/previous fields that takes care of store the address!

```CSharp
public void AddHead(T value)
{
    DoublyLinkedListNode<T> adding = new DoublyLinkedListNode<T>(value, null, head);

    if (head != null)
    {
        head.Previous = adding
    }

    head = adding;

    if (tail == null)
    {
        tail = head;
    }

    Count++;
}
```

