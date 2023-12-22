## Message Digest

A message digest is also referred to as a hash value. This cryptography algorithmic type takes a block of data a produces a unique fingerprint or hash value of that data. This enables authentication to ensure the data has not been altered.

### Message Digest Algorithms

MD2, MD4, MD5, MD6 each one of these can be used in a digital signature program to compress data securely before it's signed by a private key. The resulting message digest always has a size of 128 bits.

MD5 (and others) is a one-way hashing algorithm that generates a 128-bit hash. It's used in digital signature applications, file integrity check, and storing passwords. But on the other hand, MD5 is not collision resistant. So it's better to use the most recent algorithms such as MD6, SHA-2, and SHA-3.

SHA-1 (Secure Hash Algorithm) was introduced back in 1993 by the NSA. It's function takes an input and produces a 160-bit hash value that is then turned into hexadecimal, which is about 40-characters long when it's converted. It's typically used with other protocols like PGP (Pretty Good Privacy), TLS (Transport Layer Security), SSH (Secure Shell), SSL (Secure Socket Layer).

(SHA-0 was so bad it was pulled because of a major flaw).

SHA-2 is similar to SHA-1 but it uses 256-bit.

SHA-3 is completely different. Data is absorbed and the result is squeezed out. Message blocks are XORed into initial bits of state or a subset which is transformed into a permutation function.

CHAP - Challenge Handshake Authentication Protocol - is an authentication mechanism used with PPP (Point to Point Protocol) and a three-way handshake. But it used a shared key, this is a downside.

EAP - Extensible Authentication Protocol - was designed for point-to-point communications and it's used as an alternative to CHAP. Supports different authentication mechanisms.