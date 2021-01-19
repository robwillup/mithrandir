# Temperature Converter in C

A program always consists of basically 3 steps:

* input
* process
* output

Declaring a variable in C means informing the compiler about some fundamental properties about the variable.

# Naming Variables

* Letters
* Digits
* underscore

Variable names must start with either an underscore or a letter, digits are not allowed as the first character.

## Common Naming Styles

* snake_case, recent_files
* camelCase, userName

## Reading user input

```c
scanf("%f", &temperatureF);
```
> Reminder: the ampersand & in front of a variable gets its address.

## Complete Code

```c
#include <stdio.h>

int main(void){
    printf("Enter temperature in Fahrenheit: ");
    
    float temperatureF;
    
    scanf(%f, &temperatureF);

    float temperatureC = (temperatureF - 32.0) * 5.0 / 9.0;

    printf("The corresponding temperature in Celsius is: %f", temperatureC);

    return 0;
}
```