# Azure Container Registry authentication with service principals.

You can use an Azure Active Directory service principal to provide container image `docker push` and `pull` access to your container registry. By using a service principal, you can provide access to "headless" services and applications.

# What is a service principal?

A service principal is a user identity *for a service*, where "service" is any application, service, or platform that needs to access the resources. You can configure a service principal with access rights scoped only to those resources you specify. Then, configure your application or service to use the service principal's credentials to access those resources.

In the context of Azure Container Registry, you can create an Azure AD service principal with pull, push and pull, or other permissions to your private registry in Azure.

# Why use a service principal?

By using an Azure AD service principal, you can provide scoped access to your private container registry. Create different service principals for each of your applications or services, each with tailored access rights to your registry. And, because you can avoid sharing credentials between services and applications, you can rotate credentials or revoke access for only the service principal (and thus the application) you choose.

# When to use a service principal

You should use a service principal to provide registry access in **headless scenarios**. That is, any application, service, or script that must push or pull container images in an automated or otherwise unattended manner. For example:

  * Pull: Deploy containers from a registry to orchestration systems including Kubernetes, DC/OS, and Docker Swarm.
  * Push: Build container images and push them to a registry using continuous integration and deployment solutions such as Azure pipelines or Jenkins.
  
# Create a service principal

This [link](https://docs.microsoft.com/en-us/azure/container-registry/container-registry-auth-service-principal#create-a-service-principal) to the official documentation contains a script demonstrating how to create a new service principal and assign it an ACR role.

  
# Use an existing service principal

This [link](https://docs.microsoft.com/en-us/azure/container-registry/container-registry-auth-service-principal#use-an-existing-service-principal) to the official documentation contains a script demonstrating how to assign an ACR role to an existing service principal.

# Authenticate with the service principal

Once you have a service principal that you've granted access to your container registry, you can configure its credentials for access to "headless" services and applications, or enter them using the `docker login` command. User the following values:

 * User name - service principal application ID (also called client ID)
 * Password - service principal password (also called client secret)
 
