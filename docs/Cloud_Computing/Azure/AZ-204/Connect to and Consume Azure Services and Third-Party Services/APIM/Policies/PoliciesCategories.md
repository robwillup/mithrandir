# Access Restriction Policies

* Limit call rate by key
  * Prevent API usage by limiting call rate, on a per key basis
* Validate JWT tokens
  * enforces existence and validity of a JWT token in header or query parameter
* Set usage quota by key
  * enforces a renewable or lifetime call volume and/or bandwidth quota
* Check HTTP header presence
  * Enforces existence and/or value of an HTTP Header
* Limit call rate by subscription
  * Prevents API usage by limiting call rate, on a per subscription basis

# Advanced Policies

* Mock response
  * Returns a mocked response directly to the caller
* Forward request
  * Forwards the request to the backend service
* Retry
  * Retries execution of a request at the specified time interval
* Set request method
  * Allows changing the HTTP method for a request
* Trace
  * Adds custom tracing into the API inspector output or Application Insights

# Transformation Policies

* Convert XML to JSON
* Convert JSON to XML
* Find and replace string in body
* Set backend service
* Set query string parameter

# Caching Policies

* Store to cache
  * Caches response according to the specified cache control configuration
* Get from cache
  * Perform cache lookup and return a valid cached response when available
* remove value from the cache
  * remove an item from the cache by key

