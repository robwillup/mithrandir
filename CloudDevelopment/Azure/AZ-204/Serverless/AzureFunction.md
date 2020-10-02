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

