# Using Configuration in XUnit

This is how we can use IConfiguration in xUnit:

```csharp
var conf = ConfigurationBuilder().AddJsonFile("appsettings.json").Build();

string myValue = conf["key"];
```