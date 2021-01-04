# Azure Queue Storage

**Requires an Azure Storage account** (general-purpose v2)

**Queues are created withing a single storage account**

**Supports messages up to 64 KiB in size**

**Messages exists within a single queue**

**Supports a configurable timi-to-live per message** (with the default at 7 days)

## Interacting with Queues using the CLI

* Create a queue
  * az storage queue create --name myqueue
* Delete a queue
  * az storage queue delete --name myqueue
* View a message in a queue
  * az storage message peek --queue-name myqueue
* Delete all messages in a queue
  * az storage message clear --queue-name myqueue
 