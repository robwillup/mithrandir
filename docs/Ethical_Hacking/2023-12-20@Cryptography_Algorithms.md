### Common Cryptography Algorithms

#### Digital Signature Algorithm (DSA)

Federal Information Processing Standard for digital signatures.
This signature works by creating a 320-bit digital signature with 512-1024-bit security.
This is a crypto system that is utilized for both private and public keys.
There are two processes involved in DSA as it employs a signature generation and signature verification process. The private key is used to know who signed it, and the public key is used to verify whether the digital signatures is genuine.

#### Rivest Shamir Adleman (RSA)

This is a public-key crypto system. It uses two large prime numbers for its basis. It is
employed for internet encryption and authentication, applied in popular operating systems. It is utilized in networking cards, smart cards and in hardware-secured phones.

RSA is an asymmetric encryption algorithm. Keys are different but they are mathematically
linked.

Obvious but... the public key is... well, public. You can share it. The private key on the other hand is really private, you must keep it secret.

The public key let's people decrypt the message that you encrypted with your private key. However, if a message is encrypted with your public key, then guess who is going to be the only one who can decrypt that message? You, by using your private key. So this mechanism is
secure because the private key is secret and it's THE ONLY key that can decrypt a message encrypted with the corresponding public key. So to illustrate this, you can think of the public key more like a lock, and the private key is what can unlock it.

With RSA, you can only encrypt data whose size is equal or less then your RSA key size, which is typically 2048 bits, or 256 bytes.

##### How RSA Works

Two large prime numbers are taken (a and b), then the product of a and b is determined (c=ab, where c is called the modulus).

RSA chooses e (less than c) and relatively prime to (a-1) (b-1) e and (a-1) (b-1) have no common factor except 1.

RSA chooses f (ef-1) is divisible by (a-1) (b-1).

The value of e and f are the public and private exponents.

The public key is the pair (c,e) and the private key is the pair (c,f).

It is difficult to obtain the private key (c,f) from the public key (c,e) unless someone can factor c into a and b, then that person can decipher the private key (c,f).

RSA is very secure.

#### Diffie-Hellman

Allows two parties to establish a shared secret key over an insecure channel.
Does not provide key exchange authentication and is vulnerable to a variety of cryptographic assaults.

#### YAK

Public key based authentication key exchange. It's very fast and efficient for providing
mutual authentication and integrity protection. It's resistant to man-in-the-middle attacks. Utilizes public key pairs and requires PKI to distribute authentic public keys.
