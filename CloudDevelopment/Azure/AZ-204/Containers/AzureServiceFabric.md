# Challenges of Microservices

* Deployment
* Health monitoring
* Scaling out to multiple instances
* Service to service communication
* Upgrades
* Recover from hardware failures
  * Azure Service Fabric
  
  # Azure Service Fabric
  * An "Application platform"
    * Scalable and reliable microservices
  * Hosting options
    * On-premises or other cloud providers
    * Development laptop
    * Azure
  * Cluster
    * Monitors service health
    
  ## Programming Models
  * Stateful service
    * Co-locate compute and data
  * Stateless services
    * Web APIs or executables
  * Linux and Windows containers
    * Constrain RAM and CPU allocation
    * Docker Compose YAML support
  
  ## Service Fabric Benefits
  * Powers many key Azure services
    * e.g. Cortana, Skype, Cosmos DB & Power BI
  * Why choose Service Fabric?
    * Microservice applications
    * Windows containers 
    * Ability to deploy outside Azure
    * Orchestration features
 
  ## Service Fabric Mesh
  * Based on Service Fabric
  * Simplified deployment model
    * Container focused
    * Serverless - no need to pre-provision infractructure
    * Just specify resources required per service
    * Deployment model based on ARM
    * YAML format also available
    
  ```PowerShell
  # add az cli mesh extension
  az extension add --name mesh  
  #deploy mesh
  az mesh deployment create -g $resGroup --template-file $templateFile
  ```
