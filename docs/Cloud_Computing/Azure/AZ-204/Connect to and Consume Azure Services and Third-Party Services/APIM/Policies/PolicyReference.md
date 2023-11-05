# API Management Policies

[Reference](https://docs.microsoft.com/en-us/azure/api-management/api-management-policies)

* Advanced Policies
  * Set variable - Persist a value in a named context variable for later access.
  The `set-variable` policy declares a context variable and assigns it a value specified via an expression or a string literal. If the expression contains a literal it will be converted to a string and the type of the value will be System.String. 

## Policy Statement

```xml
<set-variable name="variable name" value="Expression | String literal" />
```
* Caching policies
  * Get value from cache - Retrieves a cached item by key.
  Use the `cache-lookup-value` policy to perform cache lookup by key and return a cached value. The key can have an arbitrary string value and its typically provided using a policy expression. 
  
  **NOTE** This policy must have a corresponding **Store value in cache** policy.