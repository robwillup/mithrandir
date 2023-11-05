# Implement Manage Identities for Azure Resources

Client Resources - e.g.: Azure Function, App Service, etc. - that want to access Target Resources - e.g.: Azure KeyVault, SQL server, Service Bus, etc. - can use a Managed Identity. The Target Resources are configured to allow access to the Managed Identity of the Client Resources.

**Azure AD Managed Identities**
>Provides Azure services with an automatically managed identity. use this identity to authenticate to any service that supports Azure AD authentication.

**Services that support Managed Identities**

| Client Services | Target Services |
|:---------------:|:---------------:|
| Azure Virtual Machines and Scale Sets | Azure KeyVault |
| App Service, Function Apps | Azure SQL |
| Logic Apps | Azure Service Bus, Azure Event Hubs |
| Azure Cognitive Search | Azure Storage (blob & queues) |
| Azure Data Factory V2 | Azure Data Lake |
| Azure Container Instances | Azure Resource Manage |
| Azure API Management | Azure Analysis Services |


## Types of Managed Identities 

| system-assigned | user-assigned |
|:---------------:|:---------------:|
| enabled directly on the Azure service instance | created as a standalone Azure resource |
| One per each Azure service instance | can be assigned to one or more Azure service instances |
| gets cleaned up if Azure service instance is deleted | its lifecycle is separate from the lifecycle of Azure service to which it's assigned |
| widely supported by Azure resources | might be in preview for some resources |

