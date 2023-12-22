## Notes on Code Signing

Right off the bat:

1. Always store code signing keys in a secure location and not on your local machine.
2. Define & limit the number of people who sign code.
3. Require approval before releasing code.

### What is code signing?

Code signing is the process of digitally signing programs and scripts to verify the software author and ensure that the code remains unchanged an trustworthy after it is signed.

Code signing eliminates "Unknown Publisher" warning and displays the Publisher's or organization's name, establishing their identity.

The code signing certificate enables validation of code sign with an authentic root certificate. Utilizing code sign enables authentication of code, cryptographic protection, and software/code author validation.

### What is a Code Signing Certificate?

Code Signing Certificates are digital certificates offered by reputable Certificate Authorities such as DigiCert and it encompasses extensive information ensuring an entity's thorough identification.

The digital certificate authority uses public key infrastructure and robust authentication practices to sign a code for applications, software, and drivers.

### How Code Signing Works?

The architecture behind code signing has 4 components:

* Code Signing System (CSS)
* Certificate Authority (CA)
* Time Stamp Authority (TSA)
* Verifier

Code signing requires a digital certificate and a signing tool to sign executable files and scripts locally. The process is supported by public key infrastructure (PKI) technology.

The signature is encrypted with private key & time stamped. This signature can then be decrypted by public key and the hash can be compared to the original to verify authenticity.

### Types of Code Signing Certificates

* Rigorous Extended Validation: Enterprises
* Light Business Validation: smaller companies
* Light Validation: individual

### Risks

If the private key, which is used for signing, is not kept securely, there is a high risk of breach activity.