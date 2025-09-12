# Dynamic Data Management with Linked Lists

Unlike arrays that require a fixed size and contiguous memory, linked lists
allow you to manage memory in a more flexible manner by linking elements via
pointers.

## Linked List

> A linear data structure where elements (nodes) are linked using pointers.
> Does not require contiguous memory allocation.

Each node in a list consists of data, which is the actual value stored, and
a pointer, which is a reference to the next node. This makes it ideal
for dynamic memory allocation where the number of elements is not known
beforehand. 

The final pointer points to null indicating that we're at the end of the list.

> The runtime complexity is O(1) for adding/removing at the beginning or middle.
> It's O(n) for searching.

So it's ideal for insertion and deletion, but not for apps where you have to
search frequently.

```c
struct Node {
    int data;
    struct Node* next;
};

struct Node* head = (
    struct Node*
)malloc(sizeof(struct Node));

head->data = 10;
head->next = NULL;

printf("First node data: %d", head->data);

// Inserting a node at the beginning

void insertAtBeginning(struct Node** head, int newData) {
    struct Node* newNode = (
        struct Node*)malloc(sizeof(struct Node)
    );

    newNode->data = newData;
    newNode->next = *head;
    *head = newNode;
}

// Deleting a node

void deleteNode(struct Node** head, int key) {
    struct Node* temp = *head, *prev;
    if (temp != NULL && temp->data == key) {
        *head = temp-next;
        free(temp);

        return;
    }

    while (temp != NULL && temp->data != key) {
        prev = temp;
        temp = temp->next;
    }

    if (temp == NULL) return;
    prev->next = temp->next;
    free(temp);
}

```


