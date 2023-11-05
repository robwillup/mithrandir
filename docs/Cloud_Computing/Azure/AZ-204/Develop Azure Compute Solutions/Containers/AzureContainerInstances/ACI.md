# Introducing Azure Container Instances

## What is ACI?
The quickest and easiest way to run a container in Azure

"Serverless" - you don't need to provision infrastructure

Great for quick experiments

Pay only while your container is running
- Per-second billing model

## ACI User Cases
- NOT for permanently running containers
- Continuous Integration build agents
  - Scale out to support multiple concurrent builds
- Short-lived experiments
  - e.g. try out MongoDB, or perform a load test
- Batch jobs
  - Run for a few hours overnight
- Elastic scale for K8S clusters
  - Virtual Kubelet

## ACI Features
Easy to create and manage
- Azure CLI, PowerShell, C# SDK, ARM

Networking
- Public IP address
- Domain name prefix
- Expose ports

Windows or Linux containers

Restart policy

Mount volumes
- Azure file share
- Secrets, Git repositories

Specify the command to run

Configure environment variables

Access container logs

"Container groups"
- One or more containers
- Run on the same server and share resources 