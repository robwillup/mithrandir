# Command to open the Kubernetes dashboard
```PowerShell
az aks browse -g $resourceGroup -n $clusterName
```

# Command to solve permission errors when viewing the Kubernetes dashboard
```PowerShell
kubectl create clusterrolebinding kubernetes-dashboard -n kube-system --clusterrole=cluster-admin --serviceaccount=kube-system:kubernetes-dashboard
```

Virtual Kubelet allows us to elastically scale the K8S cluster
