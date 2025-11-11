# Systems vs Applications

## Systems are not Applications

### Applications

* An application has a single executable and runs on a single machine
* Usually has a single source of information
* Applications don't know about "connectivity"

### Systems

* A system can be made up of multiple executable elements on multiple machines
* Usually has multiple sources of information
* System must deal with "connectivity"
* Each executable in a system is not an application because it must deal with connectivity

> Trying to abstract away the network is where the problems of distributed programs start.

## Connectivity - The Network Matters

The fallacies of distributed systems.

3 more are:

* The system is atomic/monolithic
* The system is finished
* Business logic can and should be centralized

# Fallacies of Network application

## Fallacy 1: The network is reliable

* Solutions
  * Retry & Ack / Store & Forward / Transactions
    * Don't roll your own - too many edge cases
  * Use reliable messaging infrastructure
    * RabbitMQ
    * etc.

But these don't have a request/response synchronous method-centric model.

## Fallacy 2: Latency isn't a problem

* Time to cross the network in one direction. Serialization/deserialization does not count.

* Solutions
  * Don't cross the network if you don't have to
  * Inter-object chit-chat shouldn't cross the network
  * If you have to cross the network, take all the data you might need with you.

