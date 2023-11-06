# Container Registries

Docker images can be stored locally

Use container registries to share images

Docker Hub hosts public images
- https://hub.docker.com/
- Also supports private image hosting

Azure Container Registry (ACR)
- Store containers in Azure
- Can build images automatically


This allows us to run our own container registry running in the same data center where you plan to host your applications, this has the benefit of reducing network latency, and minimizing data ingress and egress costs.

It's also possible to replicate the registry across multiple Azure regions.

## Creating Azure Container Registries

Creating a resource group and an ACR:

```Powershell
az group create -n "MyResourceGroup" -l brazilsouth

az acr create -n "MyACR" -g "MyResourceGroup" --sku Basic --admin-enabled true
```

## Container Registry Security
Images not accessible to the public

Can restrict access rights

Only allow trusted images

Automatically update to use latest security updates

## Pushing an image to ACR

```PowerShell
#Login to the registry so that Docker can have access to push the image

> az acr login -n $registryName
Login Succeeded

#We need to tag our images in a special way before pushing them, and that requires us to know the fully qualified name of our ACR server. We can run the following command to get that name:

$loginServer = az acr show -na $registryName --query loginServer -o tsv

$loginServer
pluralsightacr.azurecr.io

#Tag our image with the fully qualified ACR server name:
docker tag samplewebapp:v2 $loginServer/samplewebapp:v2

#push the image
docker push $loginServer/samplewebapp:v2

#View all the images in our ACR
az acr repository list -n $registryName -o table
Result
------------
samplewebapp

#View all tags
az acr repository show-tags -n $registryName --repository samplewebapp - table
Result
------------
v2

# Deleting images
az acr repository delete -n $registryName -t samplewebapp:v2

```

