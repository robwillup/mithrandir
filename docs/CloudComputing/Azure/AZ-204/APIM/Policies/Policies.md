# Policies in Azure API Management

[Reference](https://docs.microsoft.com/en-us/azure/api-management/api-management-howto-policies)

>In Azure API Management (APIM), policies are a powerful capability of the system that allow the publisher to **change the behavior** of the API through **configuration**. Policies are a **collection of Statements** that are executed sequentially on the request or response of an API. Popular Statements include format conversion from XML to JSON and call rate limiting to restrict the amount of incoming calls from a developer. Many more policies are available out of the box.

* What are Polices?
  * a capability of APIM
* What can we do with Policies?
  * Publishers can modify the behavior of APIs
* How can publishers change APIs behavior?
  * They can do that through configuration.
* What is an advantage of being able to modify an API behavior through configuration?
  * One advantage is the flexibility to change behavior without having to change the actual source code.
* What constitutes Policies?
  * They are a collection of Statements
* How are the Statements executed?
  * The Statements are executed sequentially on the request or response of an API.
* What are some examples of Statements?
  * Format conversion from XML to JSON.
  * call rate limiting

> Policies are applied inside the gateway which sits between the API consumer and the managed API. The gateway receives all requests and usually forwards them unaltered to the underlying API. However, a policy can apply changes to both the inbound request and outbound response.

* Where are Policies applied?
  * Policies are applied inside the gateway.
* Where does the gateway reside?
  * The gateway sits between the API consumer and the managed API.
* What does the gateway do?
  * It receives all requests and usually forwards them unaltered to the underlying API.
* How are Policies related the gateway?
  * A policy can be configured to apply changes to both inbound request and outbound responses that come and go through the gateway.

Policy expressions can be used as attribute values or text values in any of the API management policies, unless the policy specifies otherwise. Some policies such as the Control flow and Set variable policies are based on policy expressions.

## Understanding policy configuration

> The **policy definition** is a simple **XML document** that **describes** a **sequence** of **inbound** and **outbound** **statements**. The XML can be edited directly in the definition window. A list of statements is provided to the right and statements applicable to the current scope are enabled and highlighted.
> Clicking an enabled statement will add the appropriate XML at the location of the cursor in the definition view. [similar to the "assistant" in Azure DevOps to add tasks to the YAML]
> The configuration is divided into `inbound`, `backend`, `outbound`, and `on-error`. The series of specified policy statements is executed in order for a request and a response.

* What is the policy definition?
  * It is an XML document that describes a sequence of inbound and outbound statements.
* How is the configuration divided?
  * It is divided into:
    * inbound
    * backend
    * outbound
    * on-error
* How are the series of specified policy statements executed?
  * They are executed in order for a request and response.

Example:

```xml
<policies>
    <inbound>
        <!-- statements to be applied to the request go here -->
    </inbound>
    <backend>
        <!-- statements to be applied before the request is forwarded to  the backend service go here -->
    </backend>
    <outbound>
        <!-- statements to be applied to the response go here -->
    </outbound>
    <on-error>
        <!-- statements to be applied if there is an error condition go here -->
    </on-error>
</policies>
```

If there is an error during the processing of a request, any remaining steps in the `inbound`, `backend`, or `outbound` sections are skipped and execution jumps to the statements in the `on-error` section.