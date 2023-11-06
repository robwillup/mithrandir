Azure API Management (APIM) can be used to provide and API to customers. 
It is possible to implement response caching for the APIM gateway. The caching mechanism can detect the user ID of the client that accesses data for a given location and cache the response for that user ID. To do so, the following policies should be added to the policies file:

- a set-variable policy to store the detected user identity
- a cache-lookup-value policy
- a cache-store-value policy
- a find-and-replace policy to update the response body with the user profile information

See to which policy section these policies belong:

| Policy | Policy section |
|:------:|:--------------:|
| set-variable | Inbound |
| cache-lookup-value | inbound |
| cache-store-value | Outbound |
| Find-and-replace | Outbound |