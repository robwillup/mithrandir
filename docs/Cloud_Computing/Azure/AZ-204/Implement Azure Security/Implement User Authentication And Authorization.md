# Implement User Authentication and Authorization

* Azure AD App Manifests
  * appRoles
  * groupMembershipClaims
  * optionalClaims
  * oauth2AllowImplicitFlow - :/
  * oauth2Permissions
  * signInAudience
* Azure RBAC
  * Security Principal
  * Role Definition
  * Scope
  * Role Assignment
* Azure Storage SAS
  * user delegation
  * Service
  * Account
  * Best practices
    * Always use HTTPS
    * use User delegation SAS whenever possible
    * Define a stored access policy for a service, specific SAS
    * Use near-term expiration on ad hoc, service, or account SAS
    * Follow least-privilege access
* Mutual TLS
  * not supported on free or shared tiers
  * certificate is the X-ARR-ClientCert header
  * certificate value is Base64 encoded
  * App code is required to validate certificate

## Ways to Secure Azure Storage

When thinking about securing storage accounts we need to think about planes.

* Management Plane: As the name suggests this regards who has access to manage the storage account, who can change permissions, who can add other users, etc.
* Data Plan: who can access the data in the storage account

Encryption is another area to consider.

**Management Plane: RBAC**

There are 3 things at the heart of this:

* Security Principal: somebody or something you want to grant access to. It can be a user, groups of users can also be security principal. Sometimes we have a headless process (daemon) which is also a security principal. Manage Identity is also a security principal whose credentials are managed by Microsoft Azure.

* Role definition: certain actions you can perform.
* Scope: set of resources that you want to apply certain access to.

These 3 things come together with **Role Assignment**. 

Role Assignment lets you attach role definition to a security principal on a scope.

Example:

Rob (security principal) is attached "Storage account contributor" (role definition) to "storage account robwill123" (scope).

Multiple role assignments are additive: example: Owner+contributor

Deny assignments can block access.

**Data Place**

There are 3 ways to protect access to Storage Accounts:

* Keys
  * when you create a new Storage Account, two Keys are provided and they are equivalent. The reason there are 2 is that you can use the one while the other is being renewed, so it is for continuity purposes. The keys have **root** access, therefore these keys must be secured. Also, these keys should be rotated frequently.
* Shared Access Signature
* Azure AD