# Review Secure Cloud Solutions

* Managed Identities
  * system-assigned
    * widely supported across Azure
    * automatically attached to a single Azure resource
    * deleted when attached Azure resource is deleted
    * Azure resources can have a single system-assigned identity
  * user-assigned
    * supported by a growing list of services on Azure
    * created as a standalone Azure resource
    * Must be deleted manually
    * Azure resources can have multiple user-assigned identities
* Azure Key Vault
  * create a key vault using PowerShell
    * New-AzKeyVault -Name 'sample' -ResourceGroupName 'samplegroup' -Location 'East US'
  * create a key vault using Azure CLI
    * az keyvault create --name "sample" -g "samplegroup" --location eastus