# Azure App Service

* Web application hosting platform
* Custom domains & SSL certificates
* Scale up or scale out
* Staging slots
* Manage environment variables
* Rich monitoring experience
* Automated deployment
* IP address whitelisting, AD authentication
* Enabling CDN

### App Service Plan
* Can host multiple web apps
* Several pricing tiers available
  * basic, standard, premium
* Supports many platforms
  * ASP.NET, Node.js, Java, PHP, Python
* Also supports containers

## Why Web App for Containers?
* Consistent deployment model
* Support more frameworks
* Multi-container support
  * Docker-compose or Kubernetes YAML
* Sandboxed environment
* Trigger deployments from a container registry

## Creating a Web App for Containers

```PowerShell
#Create the web app
az webapp create -n $appName -g $resourceGroup --plan $planName -i $dockerRepo

#Get the fully qualified domain named
$wordpressDbHost = (az mysql server show -g $resourceGroup -n $mysqlServerName --query "fullyQualifiedDomainName" -o tsv)

#Set the web app settings
az webapp config appsettings set -n $appName -g $resourceGroup --settings WORDPRESS_DB_HOST=$wordpressDbHost WORDPRESS_DB_USER="$adminUser@$mysqlServerName" WORDPRESS_DB_PASSWORD="$adminPassword"



```
