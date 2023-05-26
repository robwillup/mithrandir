# AZ CLI - Static Web App

For my [personal website](../../../LanguagesAndFrameworks/.NET/Blazor/Personal%20Website.md),
I wanted to host it in [Azure Static Web Apps](https://azure.microsoft.com/en-us/products/app-service/static)
which seemed to be a good choice for Blazor WASM.

## Creating a Resource Group

```pwsh
> az group create --location westus --resource-group Rob

{
  "id": "/subscriptions/8589b7a1-8087-4163-9b87-0f2e15fccf80/resourceGroups/Rob",
  "location": "westus",
  "managedBy": null,
  "name": "Rob",
  "properties": {
    "provisioningState": "Succeeded"
  },
  "tags": null,
  "type": "Microsoft.Resources/resourceGroups"
}
```

## Creating a Static Web App

```pwsh
> az staticwebapp create -n Rob -g Rob

{
  "allowConfigFileUpdates": true,
  "branch": null,
  "buildProperties": null,
  "contentDistributionEndpoint": "https://content-dm1.infrastructure.2.azurestaticapps.net",
  "customDomains": [],
  "defaultHostname": "happy-stone-0ae01dd10.2.azurestaticapps.net",
  "enterpriseGradeCdnStatus": "Disabled",
  "id": "/subscriptions/8589b7a1-8087-4163-9b87-0f2e15fccf80/resourceGroups/Rob/providers/Microsoft.Web/staticSites/Rob",
  "identity": null,
  "keyVaultReferenceIdentity": "SystemAssigned",
  "kind": null,
  "linkedBackends": [],
  "location": "Central US",
  "name": "Rob",
  "privateEndpointConnections": [],
  "provider": "None",
  "publicNetworkAccess": null,
  "repositoryToken": null,
  "repositoryUrl": null,
  "resourceGroup": "Rob",
  "sku": {
    "capabilities": null,
    "capacity": null,
    "family": null,
    "locations": null,
    "name": "Free",
    "size": null,
    "skuCapacity": null,
    "tier": "Free"
  },
  "stagingEnvironmentPolicy": "Enabled",
  "tags": null,
  "templateProperties": null,
  "type": "Microsoft.Web/staticSites",
  "userProvidedFunctionApps": null
}
```