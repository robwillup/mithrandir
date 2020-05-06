# Volumes

Storing container data
* Containers are disposable
* Keep data in Docker volumes
* Mount a volume

```PowerShell
docker run -d -p 5432:5432 -v postgres-data:/var/lib/postgresql/data --name postgres1 postgre
```

