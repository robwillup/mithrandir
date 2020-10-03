# Learning Cosmos DB

Topics:

* Global distribution
* Horizontal partinioning
* Proviioning throuput
* Multiple APIs and data models

## What is NoSQL?

What is the common problem that the solutions below are trying to solve?
* Cosmos DB
* mongoDB
* cassandra
* Apache HBASE
* Amazon DynamoDB
* OrientDB
* ArangoDB

*NoSQL* is a misnomer. But regardless of that, NoSQL databases are trying to work with the 3 Vs:

* Volume
* Variety
* Velocity

These 3 Vs define what is called "Big Data".

Fun Fact: Every 60 seconds there are:
* 98K + tweets
* 695K status updates in Facebook
* 11 million instant messages
* 698,445 Google searches
* 168 million+ emails sent
* 1,820TB of data created
* 217 new mobile web users

Relational databases were not created for that huge amount of data. They were created in a time where scalability would not grow so fast, and they would scale by getting more memory periodically. So NoSQL came to solve this problem and scale on demand for volume, variety and velocity. 

NoSQL is just a buzword to address the challenges of the 3 Vs and so a NoSQL database is built to be distributed from the ground up. They also scale by horizontal partitioning and scalle throughput. And they are "schema-free" which means that there is no schema management. 

Today's applications are required to be **highly responsive** and **always online**. To achieve low latency (highly responsive) and high availability (always online), instances of these applications need to be deployed in datacenters that are close to their users. Applications need to respond in real time to large changes in usage at peak hours, store ever increasing volumes of data, and make this data available to uers in milliseconds.

## What is Cosmos DB?

* Evolution of DocumentDB
* Virtually unlimited scale
* Turnkey global distribution
* Multi-model / Multi-API

Azure Cosmos DB is Microsoft's globally distributed, multi-model database service.
