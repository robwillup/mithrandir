# Azure KeyVault

This is a service which allows you to securely store and access secrets.

**Types os Secrets**

* keys
  * cryptographic keys used in other Azure services such as Azure Disk Encryption
* secrets
  * any sensitive information including connection strings, passwords, client_secrets, etc.
* certificates
  * x509 certificates used in HTTPS/TLS communications (encryption in transit)

## Azure KeyVault Pricing Tiers

* Standard
  * Software-protected
* Premium
  * Standard + HSM-protected

## Three options to authenticate to Azure KeyVault

* Azure AD App Registration
* Use Managed Identities
* Use Azure KeyVault References

## Creating Azure KeyVault

* Azure CLI
* PowerShell
  * Provision a KeyVault resource:
    New-AzKeyVault -VaultName 'kv-demo' -ResourceGroupName 'group-demo' -Location 'Brazil South'

## Azure KeyVault Soft-delete

Allows recovery of the deleted vaults and key vault objects (keys, secrets and certificates).

Soft delete is enabled by default for all new KeyVaults

## Azure KeyVault Purge Protection

When purge protection is enabled, a vault or an object in the deleted state cannot be purged until the retention period has passed.

