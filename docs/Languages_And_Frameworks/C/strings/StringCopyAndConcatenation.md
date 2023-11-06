# String Copy and Concatenation

In C we cannot simply "assign" a string to a variable to copy it. So the following is not available:

```c
new_str = my_str;
```

Instead we need to use methods provided by the standard library. For example:

```c
/* Destination string*/

chat dest[100];

/* Copy source to destination */
/* dest = source */

strcpy(dest, source);
```

## Concatenating

Again, the following syntax is not available in C:

```c
new_str = str1 + str2;
```

Instead, we need to:

```c
/* Append source to dest */
/* dest = dest + source */
strcat(dest, source);
```
The string manipulation functions above and others are found in the `#include <string.h>` 

## Beware of Buffer Overflow

Use safe string functions:

* strcpy_s
* strcat_s

These safer versions require an additional parameter that represent the destination buffer size. 

```c
strcpy_s(msg, sizeof(msg), name);
```