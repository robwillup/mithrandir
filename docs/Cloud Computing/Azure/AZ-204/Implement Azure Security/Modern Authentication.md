# Modern Authentication

## Legacy Authentication

* Basic
* NTLM
* Kerberos

Kerberos worked really nice On Premises, but it does not scale to the cloud.

## Modern Authentication

* WS-* and SAML
* OAuth (delegation protocol, pseudo-authentication protocol)
* OpenID Connect

(distributed identity might be the next big thing in modern auth)

## OpenID Connect (Apps)

* SPA
  * PKCE
* Native
  * AuthCode without secret (PKCE)
* Web
  * AuthCode with secret
* Daemon
  * Client Credential Flow
* Limited UI
  * Device Code Flow


