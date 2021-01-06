# Storing and Retrieving Data in Azure Redis Cache

**Pricing Tiers**

* Basic
  * Minimal feature set, less throughput and higher latency
  * no SLA
  * no for production
  * up to 53 GB
* Standard
  * 2 Replicated nodes
  * SLA: 99.9% availability
  * 53 GB of memory
  * 20,000 clients
* Premium
  * Redis cluster
  * low latency
  * 99.5% availability
  * 100 GB of memory
  * 40,000 clients
* Enterprise
  * Full Redis feature set
  * 99.9% availability
* Enterprise Flash
  * fast non-volatile storage


You can scale up, but you cannot scale down.

## Exploring Caching Patterns

* Cache-aside pattern
  * store most frequently access data in Redis
  * 1. Does the data exist in the cache?
  * 2. if not, retrieve from the data store
  * 3. Store a copy in the cache
* Content Cache Pattern
  * Cache static content
    * images
    * templates
    * style sheets
  * Reduced server loads
  * Redis output cache provider for ASP.NET
* User Session Caching Pattern
  * Maintain application state
    * shopping cart
  * session cookies or local storage
    * limited data storage
    * slow performance
  * Best option: store the data in Redis and point to it in the app

**Advanced Patterns**

* Job and Message Queuing
* Distributed Transactions

## Default Time to Live

* Azure CDN: 7 days
* Azure FrontDoor: between 1 and 3 days
* Azure Cache for Redis: no default expiry, items persist until deleted
* 