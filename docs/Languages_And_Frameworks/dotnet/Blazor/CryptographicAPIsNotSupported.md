# Not Possible To Cryptographically Sign A JWT

Today I have learned that the APIs in `System.Security.Cryptography` namespace
are not supported in Blazor WebAssembly.

So, for example, you cannot do this in your Blazor WebAssembly app:

```csharp
RSA rsa = RSA.Create();
```

To me this means that I'm not going to be able to generate a GitHub JWT and sign it
in my Blazor WebAssembly app.
