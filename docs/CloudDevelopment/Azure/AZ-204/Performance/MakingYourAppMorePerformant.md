# Making your app more performant

After your application is up and running in Azure, you want it to be as 
performant as possible. Azure provides a range of services that can 
help you with that.

## Azure Traffic Manager

When it comes to global applications, one of the most obvious problems we have to deal with is latency, which is the time it takes for a signal 
or a request to travel to a user. The farther away users are from your 
application, the more latency they experience.

**Azure Traffic Manager** scales across regions, helping to reduce latency and provides users with a performant experience, regardless of where they are. Azure Traffic Manager is an intelligent routing mechanism that you put in front of your Web Apps applications. Web Apps acts as an endpoint, which Azure Traffic Manager monitors for health and performance.

When users access your application, Azure Traffic Manager routes them to the Web Apps application that is most performant in their proximity. 

Including Azure Traffic Manager in your architecture is a great way to improve the performance of your applications.

## Azure Front Door

Your users may be spread across the globe and, at times, they may be traveling. This can make it difficult to ensure that they have a performant experience and that your application is available and secure, regardless of location.

**Azure Front Door** can help. This service can route traffic from users to the most performant application endpoint for them to improve performance. Azure Front Door can route to available endpoints, while avoiding endpoints that are down.

Azure Traffic Manager does this well, but in a different manner to Azure Front Door. Azure Front Door works at **OSI layer 7** or the HTTP/HTTPS layer, while Azure Traffic Manager works with DNS. In other words, Azure Front Door works on the application level, and Azure Traffic Manager works on the network level. This is a fundamental different that determines the capabilities of the services.

Because of this difference, Azure Front Door does a lot more than route users to available and performant endpoints. 

Azure Front Door allows you to author custom web application firewall (WAF) rules for access control to protect your HTTP/HTTPS workload from exploitation based on client IP addresses, country codes, and HTTP parameters.

Additionally, AFD enables you to create rate-limiting rules to battle malicious bot traffic. These are just some of the unique capabilities of Azure Front Door.