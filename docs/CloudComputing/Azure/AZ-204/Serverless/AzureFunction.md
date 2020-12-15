# An introduction to Azure Functions

Azure Functions allow you to run small pieces of code (called "functions") without worrying about application infrastructure. With Azure Functions, the cloud infrastructure provides all the up-to-date servers you need to keep your applications running at scale. 

A function is "triggered" by a specific type of event. Supported triggers include responding to changes in data, responding to messages, running on a schedule, or as the result of an HTTP request. 

While you can always code directly against a myriad or services, integrating with other services is streamlined by using bindings. Bindings give you declarative access to a wide variety of Azure and third-party services.

## Features

Some key features of Azure Functions include:

* **Serverless applications:** Functions allow you to develop `serverless` applications on Microsoft Azure.
* **Choice of language:** Write functions using your choice of `C#, Java, JavaScript, Python, and PowerShell`.
* **Pay-per-use pricing model:** Pay only for the time spent running your code. See the Consumption hosting plan option in the pricing section.
* **Bring your own dependencies:** Functions supports NuGet and NPM, giving you access to your favorite libraries.
* **Integrated security:** Protect HTTP-triggered functions with OAuth providers such as Azure Active Directory, Facebook, Google, Twitter, and Microsoft Account.
* **Simplified integration:** Easily integrate with Azure services and software-as-a-service (SaaS) offerings.
* **Flexible development:** Set up continuous integration and deploy your code through GitHub, Azure DevOps Services, and other supported development tools.
* **Stateful serverless architecture:** Ochestrate serverless applications with Durable Functions.
* **Open-source:** The Functions runtime is open-source and available on GitHub.

## Azure Functions Core Tools
* Command line tools
  * Creating new Function App
  * Adding functions
  * Deploying to Azure

Includes the Azure Functions runtime

Azure Storage Emulator
* Windows only
* Use a real Storage Account instead

You are developing an application that uses Azure Blob storage. The application must read the transaction logs of all the changes that occur to the blobs and the blobs metadata in the storage account for auditing purposes. The changes must be in the order in which they occurred, include only create, update, delete, and copy operations and be retained for compliance reasons. You need to process the transactions logs asynchronously. What should you do? 

A. Process all Azure Blob storage events by using Azure Event Grid with a subscriber Azure Function app.

B. Enable the change feed in the Storage Account and process all changes for available events.

C. Process all Azure Storage Analytics logs for successful blob events.

D. Use the Azure Motinor HTTP Data Collector API and scan the request body for successful blob events.

**Correct answer: B**
