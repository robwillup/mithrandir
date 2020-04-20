# gRPC Motivation and Design Principles

## Motivation
Google has been using a single general-purpose RPC infrastructure called Stubby `to connect the large number of microservices` running within and across our data centers for over a decade. Our internal systems have long embraced the microservice architecture gaining popularity today. 

With the advent of SPDY, HTTP/2, and QUIC, many of these same features have appeared in public standards, together with other features that Stubby does not provide. It became clear that it was time to rework Subby to take advantage of this standardization, and to extend its applicability to mobile, IoT, and Cloud use-cases.

## Principles & Requirements

### Services not Objects, Messages not References
Promote the microservices design philosophy of coarse-grained message exchange between systems while avoiding the pitfalls of distributed objects and the fallacies of ignoring the network. 

### Coverage & Simplicity
The stack should be available on every popular development platform and easy for someone to build for their platform of choice. It should be viable on CPU and memory-limited devices.

### Free & Open
Make the fundamental features free for all to use. Release all artifacts as open-source efforts with licensing that should facilitate and not impede adoption.

### Interoperability & Reach
The wire protocol must be capable of surviving traversal over common internet infrastructure. 