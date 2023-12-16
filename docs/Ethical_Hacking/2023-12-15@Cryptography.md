## Cryptography For Beginners

In my first years as a software developers, I saw this word a lot: `cryptography`.
I even used it in parts of my work. But at that time my understanding of it was still very superficial and unstructured.
In this very short post I attempt to offer a simple explanation of cryptography for others who may also feel like they need to understand it better.

### What does the word cryptography mean?

Each word caries a world of meaning and history and getting to know how they have been formed and what each part means goes a long way in structuring our understanding of things. Let's take a look at the etymology and morphology of the word `cryptography`.

According to [etymonline.com](https://www.etymonline.com/word/cryptography), cryptography is the *art of writing in secret characters* and it's formed with the Greek works *kryptos* which means "hidden" and *graphia* "to write".

Based on this, cryptography is a technique to hide the meaning of a piece of data.

### What is the purpose of cryptography?

To put it simply: protect data.

The purpose is to help ensure that a message or other piece of data should only be available to the intended recipient.

Cryptography provides *confidentiality*. When done properly, it will make it harder for others to spy on data that they should not have access to.

Another way that cryptography helps protect data is through *authentication*. Since only the intended party can access that data, it's possible to determine the identity of whoever legitimately access that data.

And a third aspect of this data protection is something they call *non-repudiation*. This means that cryptography provides a way to prevent someone from denying that they have interacted with a piece of data when they in fact did.

### How have I used cryptography software development?

The two most common ways I have used cryptography in my day-to-day work is through *hashing* and *encryption*.

I have used hashing for password protection. When a password was stored in a system it was not
kept *verbatim*, i.e., as plain text, because if the system was ever to be compromised those passwords would be exposed too. Instead, the passwords were hashed producing an unintelligible string, and that's what was stored. This hashing process is "one-way-only", this means that the hash cannot be reverted back to the original input. So if someone with ill intentions gets the hash for your password, it would be useless because they would not be able to revert it to the plain text password.

Encryption has also been pretty common in several systems and protocols, an one example that comes to mind is when I used encryption to protected an authentication key that had to be kept on the users device. This key had to be encrypted so that third-parties would not be able to read it, but it had to be decrypted before used by my application. So encryption is a two-way type of cryptography.

### What's more to know?

There is a lot more to learn about cryptography. And I mean *a lot*. But to keep this post short I'm stopping here. In future posts about software security and cryptography other aspects and more details will be covered.

![Cryptography](https://github.com/robwillup/mithrandir/blob/main/assets/Ethical_Hacking/2023-12-15@Cryptography_simple_flow.jpeg?raw=true)
