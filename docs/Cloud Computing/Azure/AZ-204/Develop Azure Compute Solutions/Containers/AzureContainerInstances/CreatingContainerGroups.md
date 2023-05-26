# Creating Container Groups with the Azure CLI

az container create

-n mycontainergroup -g myresourcegroup

--image someimage:sometag

--ip-address public

--dns-name-label mysite # mysite.eastus.azurecontainer.io

--port 80

--os-type Windows # default is Linux

--cpu 1 --memory 1.5

-e name=value

--restart-policy never # always, onfailure

--azure-file-volume... # credentials, mount path, share-name

```PowerShell
$resourceGroup = "AciGhostDemo"
$location = "brazilsouth"
az group create -n $resourceGroup -l $location
$containerGroupName = "ghost-blog1"
az container create -g $resourceGroup -n $containerGroupName --image ghost --ports 2368 --ip-address public --dns-name-label ghostaci1

az group delete -n $resourceGroup
```

## Using ACR and Mounting Volumes

```PowerShell
$acrServer = az acr show -n $acrName --query loginServer --output tsv
$acrPassword = az acr credential show -n $acrName --query "passwords[0].value" -o tsv

az storage account create -g $resourceGroup -n $storageAccountName --sku Standand_LRS

 $storageConnectionString = az storage account show-connection-string -n $storageAccountName -g $resourceGroup --query connectionString -o tsv

$env:AZURE_STORAGE_CONNECTION_STRING = $storageConnectionString
$shareName = "aciShare"
az storage share create -n $shareName

$storageKey=$(az storage account keys list -g $resourceGroup --account-name $storageAccountName --query "[0].value" --output tsv)

$containerGroupName = "aci-acr"
az container create -g $resourceGroup -n $containerGroupName --image $acrServer/samplewebapp:v2 --cpu 1 --memory 1 --registry-username $acrName --registry-password $acrPassword --azure-file-volume-account-name $storageAccountName --azure-file-volume-account-key &storageKey --azure-file-volume-share-name $shareName --azure-file-volume-mount-path "/home" -e TestSetting=FromAzCli TestFileLocation=/home/message.txt --dns-name-label "aciacr1" --port 80
```