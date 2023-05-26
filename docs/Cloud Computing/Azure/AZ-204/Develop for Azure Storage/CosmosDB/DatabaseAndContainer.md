# Creating a new container

In many cases you may only have one container in a database. All the intems in a container can be 
of different types and can have different schemas. Containers are **not** equivalent to tables in
relational databases. 

Partitionaing is how CosmosDB provides virtually unlimitted storage capacity for a container, based 
on the partinion key that you define for it. You need to come up with a partition key, some property 
in your data that will be used internally to physically store multiple documents together inside partions
within the container.

Throughput is what you request the performance levels that you need this container to deliver. This is
expressed in the form of RU (request units).
