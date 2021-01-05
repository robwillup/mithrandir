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

## Three different types of SAS

* User delegation SAS
  * Using AAD, use specific scope named "user impersonation" so that user's permissions matter when you try to access Azure Storage
* Service SAS
  * Secured with the Storage Account key. Delegates access to a resource in only one of Storage Account services, e.g. either Blob, queue, table or file.
* Account SAS
  * Secured with the Storage Account key. Delegates access to resources to one or more of Storage Account services.

## A Typical SAS token (ad-hoc token)

```
https://robstorage123.blob.core.windows.net/? -> URL
sv=2019-12-12&                                -> signedVersion
ss=bfqt&                                      -> signedServices
srt=s&                                        -> signedResourceType
sp=rwdlacupx&                                 -> signedPermission 
se=2022-10-19T12:50:12Z&                      -> signedExpiry
st=2022-10-19T04:50:12Z&                      -> signedStart                     
spr=https&                                    -> signedProtocol
sig=dXxS3I%2F1LdlNzu9oLUOixzgESdlVhXNXlg...   -> signature
```

## Kinds of SAS

* Ad-Hoc
  * everything you need, all the information that you need is available in the token itself, the SAS URI
* Service SAS with Stored Access Policy
  * Reused by multiple SAS
  * Defined on Resource Container
  * Permission & Validity Period
  * Service level SAS only


## A Typical Stored Access Policy token

```
https://robstorage123.blob.core.windows.net/? -> URL
sr=c&
si=mypolicy&
sig=dXxS3I%2F1LdlNzu9oLUOixzgESdlVhXNXlg...   -> signature
```