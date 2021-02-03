# Variables

* Named data storage
* Strongly typed
  * int dataValue;
* Variable Naming
  * use only letters and numbers
* Style Names Using Camel Case
  * Start each word after the first with upper case
  * All other letter are lower case
  * int studentCount;
* Variables can be declared final
  * use final modifier:
    * final int maxStudents = 25;
  * Value cannot be changed once set
  * Helps avoid errors caused by inadvertent variable changes

## Primitive Data Types

* Built into the language
* Foundation of all other types
* Four categories
  * integer
    * byte: bits = 8 | min value = -128 | max value = 127 | 0
    * short: bits = 16 | min value = -32768 | max value = 32767 | 0
    * int: bits = 32 | min value = -2147483648 | 2147483647 | 0
    * long: bits = 64 | min value = -9223372036854775808 | max value = 9223372036854775808 | 0L
  * floating point
    * float: bits = 32 | min = 1.4 * 10<sup>-45</sup> | max = 3.4 * 10<sup>38</sup> | literal form = 0.0f
    * double: bits = 64 | min = 4.9 * 10 <sup>-324</sup>| max = 1.7 * 10 <sup>308</sup>| literal form = 0.0 or 0.0d
  * character
    * stores a single unicode character
    * literal values placed within single quotes
    * For unicode code points, use \u followed by 4-digit hex value
    * char accentedU = '\u00DA'; // Ú
  * boolean
    * boolean iLoveCSharp = true;

Primitive Types are stored by Value