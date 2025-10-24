# ASP.NET Core Microservices

In a Microservices architecture, code is separated into domains.
But the domains can have `bounded contexts`. What are those?

Bounded context are subdomains that are bound through one or more entities that they contain.

According to the instrctor of this course, binding contexts should be a backend concern, and not a frontend problem.
The reason for that is that if the frontend should only be concerned with the presentation, if it has to worry about
making multiple calls to complete the data for one requests, it's inefficient and there's some kind of business logic
there.

## Service to service communication

One way to have the backend bind contexts to return all the data that frontend needs is through service to service
communication. When the frontend makes a request to a backend service, that backend will process it and if it needs
more data, it can itself call another backend service to aggregate that data and then return it to the frontend.

Of course there are caveats with this approach. The other service becomes yet another dependency that can add
latency and chances of errors.



