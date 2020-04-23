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

### General Purpose & Performant
The stack should be applicable to a broad class of use-cases while sacrificing little in performance when compared to a use-case specific stack.

### Layered
Key facets of the stack must be able to evolve independently. A revision to the wire-format should not disrupt application layer bindings.

### Payload Agnostic
Different services need to use different message types and encodings such as protocol buffers, JSON, XML, and Thrift; the protocol and the implementations must allow for this. Similarly the need for payload compression varies by use-case and payload type: the protocol should allow for pluggable compression mechanisms.

### Streaming
Storage systems rely on streaming and flow-control to express large data-sets. Other services, like voice-to-text or stock-tickers, rely on streaming to represent temporally related message sequences.

### Blocking & Non-Blocking
Support both asynchronous and synchronous processing of the sequence of messages exchanged by a client and server. This is critical for scaling and handling streams on certain platforms. 



