# Pointers and Arrays

An address is just a number which points to the first byte in memory holding the value of a variable.
Address of operator = &, e.g.: &myVariable

```C
int main()
{
    int i = 1234;

    int * p = &i;

    int j = *p;

    *p = 2345;
}
```