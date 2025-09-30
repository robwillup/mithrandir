# What's coming in ASP.NET Core v10

## Security

- Support the latest security standards & best practices
  - WebAuthN & Passkey authentication
- Make authentication easier to set up & use
  - Authentication scaffolding
  - Identity & authentication docs improvements

### WebAuthN & Passkeys

- Passkeys are cryptographic credentials that replace traditional passwords.
- They are phishing-resistant, easy to use, and secure by design.
- How Passkeys work:
  - A passkey consists of a public-private key pair, scoped to a specific account and origin.
  - The private key is stored in an Authenticator
  - The Authenticator requires the user verification before allowing authentication.
- Add built-in support for Passkeys to ASP.NET Core Identity
  - built into the shared framework, but inspired by fido2-net-lib
- Project templates & scaffolding updated to add Passkey support
- Existing projects using ASP.NET Identity can add Passkey support
  - will require a database schema migration

## App observability and diagnostics

- More metrics
  - memory pool diagnosability metrics
  - Authentication and authorization metrics
  - Blazor specific metrics
- Add activities to Blazor Server for distributed tracing
- Diagnostics for Blazor WebAssembly

## Performance

- Releasing memory from Kestrel memory pool
- API JSON deserialization performance
- Investigate performance of antiforgery
- Blazor WebAssembly startup improvements

