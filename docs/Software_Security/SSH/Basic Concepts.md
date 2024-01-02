## Basic SSH Concepts

Describes basic SSH concepts.

[Source](https://schh.medium.com/ssh-for-dummies-ea168e6ff547)

### TL;DR

* SSH Stands for Secure Shell.
* Purpose: Remove server administration.
* Communication is encrypted.
* Lets a user authenticate to a remote server.
* Lets the client (where the user operates) send input to the host.
* Sends the output from the host (remote server) back to client.

SSH stands for `Secure Shell` and it is a remote administration
protocol that allow users to control and modify their remote servers
over the internet. It uses cryptographic techniques to ensure that all
communication to and from the remote server happens in an encrypted
manner. It provides a mechanism for authenticating a remote user,
transferring inputs from the client to the host, and relaying the
output back to the client. It uses port 22 and communicates over TCP
protocol.

In an SSH command, you will find three parts:

```bash
ssh {user}@{host}
```

The part `{user}` represents the account you want to access.

The part `{host}` refers to the computer you want to access. This can
be an IP address or a domain name.

For example:

```bash
ssh pi@192.0.0.1
```

SSH protocol uses symmetric encryption, asymmetric encryption and
hashing in order to secure transmission of information. The SSH
connection between the client and the server happens in three stagers:

1. Verification of the server by the client.
2. Generation of a session key to encrypt all the communication.
3. Authentication of the client.

### Verification of the Server by the Client

Server authentication is a process that allows client applications to
validate a server's identity. If the server fails the SSH host key
authentication process, it's possible that the server's host key was
simply changed by the admin which is basically not a big problem.
However, it could also mean that someone has carried out a spoofing
or man-in-the-middle attack and, therefore, the client is likely on
the verge of connecting to a malicious server. That is a serious problem.

If a user unknowingly, logs in to a malicious server, whoever has
control of that server could easily acquire that user's login
credentials and then use those to gain access to the legitimate server.
Secondly, if the unwitting user uploads files to the malicious server,
those files will surely fall into the wrong hands. Lastly, if a user
downloads files from the server, that user could end up downloading malware.

Server authentication helps prevent these from happening because if the
authentication process fails, the client will be given an appropriate
warning.

There are two important files to perform ssh authentication:
`known-hosts` and `authorized_keys`.
