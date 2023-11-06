# About API Management

API Management (APIM) is a way to create *consistent* and *modern* API gateways for existing back-end services.

API Managements helps organizations publish APIs to external, partner, and internal developers to unlock the potential of their data and services. API Managements provides the core competencies to ensure a successful API program through developer engagement, business insights, analytics, security, and protection. You can use Azure API Management to take any backend and launch a full-fledged API program based on it.

## Overview

To use APIM, administrators create APIs. Each API consists of one or more operations, and each API can be added to one or more products. To use an API, developers subscribe to a product that contains that API, and then they can call the API's operation, subject to any usage policies that may be in effect. Common scenarios include:

* **Securing mobile infrastructure** by gating access with API keys, preventing DOS attacks by using throttling, or using advanced security policies like JWT token validation.
*  **Enabling ISV partner ecosystems** by offering fast partner onboarding through the developer portal and building and API facade to decouple from internal implementations that are not ripe for partner consumption.
*  **Running an internal API program** by offering a centralized location for the organization to communicate about the availability and latest changes to APIs, gating access based on organizational accounts, all based on a secured channel between the API gateway and the backend.

## Policies

[Policies](./Policies.md)

## Areas of Focus for the Exam

* service tiers
  * consumption tier: pay per call to APIM
  * developer tier: testing and validation. No SLA
  * basic tier: sla
  * standard: ad integration
  * premium: increased SLA, multi region deployments
  * isolated: completely isolated
* caching
  * Internal: cache provided within APIM service
    * limited in size based on APIM tier.
    * Not available on the consumption tier
  * External: Redis compatible cache outside of APIM such as Azure Cache for Redis
  * both types of caching are configured in the APIM policies
* Access restriction and authentication
  * access restriction: check for http header, limit call rate by subscription, restrict by IP address, usage quota per key, validate JWT
  * Basic Auth, client cert, managed identity auth
* Policy definition
  * 

