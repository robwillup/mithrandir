# Shared Access Signature (SAS)

Secure, delegated access, without sharing the key.

Control what the clients access, for how long, etc.

With SAS you have granular control over how a client can access your data. The keys are too much power, and that's why SAS was created, to provide the least-privilege option. With SAS you can control:

* What resources the client may access
* What permissions they have to those resources
* How long the SAS is valid.

## What is SAS?

A shared access signature is a *temporary* token that can be attached to the URI to delegate access to a specific storage account resources without exposing your entire account.

This token contains a lot of information such as which resources can be accessed, lifetime by setting the start and expiration data of the signature, and the permissions been granted (reading, adding, updating and deleting resource such as container blobs, queue messages, and table entities). It gives you granular control over the type of access you grant, for example, you can give a web application the permission to write message to queues and give another function the rights to process and delete those messages.