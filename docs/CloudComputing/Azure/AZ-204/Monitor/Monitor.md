# az monitor metrics alert

## Example

Create a high CPU usage alert on a VM with no action

```bash
az monitor metrics alert create -n alert1 -g "resource group" --scopes "VM ID" --condition "avg Percentage CPU > 90" --description "High CPU"
```

