# Configuring Continous Delivery

```PowerShell
# Configuring a Web App for Container to use an image from ACR
az webapp config container set -n $appName -g $resourceGroup -c "$acrLoginServer/samplewebapp:latest" -r "https://$acrLoginServer" -u $acrUserName -p $acrPassword

# Creating a deployment slot
az webapp deployment slot create -g $resourceGroup -n $appName -s staging --configuration-source $appName

# Finding the domain of the slot
az webapp show -n $appName -g $resourceGroup -s stagin --query "defaultHostName" -o tsv

# Enabling CD
az webapp deploymenht container config -g $resourceGroup -n $appName -s staging --enable-cd true
# The command above generate a WebHook URL that is going to be used by the ACR to notify the web app whenever a new image is pushed to that registry

# Get the WebHook URL
$cicdurl = az webapp deployment container show-cd-url -s staging -n $appName -g $resourceGroup --query CI_CD_URL -o tsv

# Configure ACR to call that WebHook whenever a new image is pushed
az acr webhook create --registry $acrName --name myacrwebhook --actions push --uri $cicdurl

# Performing a SWAP
az webapp deployment slot swap -g $resourceGroup -n #appName --slot staging --target-slot production
```

