# Example Scenarios

## Scenario 1

Sylvia's company is in the process of moving multiple web apps to Azure. The web applications themselves are deployed as containers. Application demand varies, and they have struggled with uptime in the past.

What is the most cost effective approach they could take?

**Azure App Service Web Apps for Containers, Standard (S) tier with Linux Runtime**

## Scenario 2

Edward has created a document processing service for his company. After his app uploads a document to blob storage, it calls an API. The API triggers the document processing on a VM.

Is this the most efficient and cost effective approach for this solution?

**No** He could use an Azure Function with the blob storage as the input binding. The function could use an ACI or Batch service to process the document.

## Scenario 3

Cindy's company provides a digital assets management SaaS solution. They are trying to find more cost effective ways to process large videos. Cindy has read about Durable Functions and believes this could be a solution.

Is this problem solved by using Durable Functions?

**No. processing of large files is not an identified use case**

## Scenario 4

William's company currently runs a fantasy epic game platform. Currently they perform multiple actions on a VM when a new user is added. William is afraid to move it to a single Azure function due to a possible timeout.

Is this problem solved by using Durable Functions?

**Yes. Function chaining is a valid Durable Functions use case**

## Scenario 5

Oscar's company is deploying a new web app using App Service. Oscar will be deploying the app using the CLI. Oscar will be deploying into a brand new account that is currently empty.

What is the correct order of the steps Oscar will need to follow?

* **az group create**
* **az appservice plan create**
* **az webapp create**

## Scenario 6

James's company has multiple Windows VM's deployed in a VNet. They need high-speed communication to analyze shared streaming data. Currently they are experiencing higher than desired lag between their VM's.

Which solution could reduce the latency between VM's?

**Accelerated Networking**