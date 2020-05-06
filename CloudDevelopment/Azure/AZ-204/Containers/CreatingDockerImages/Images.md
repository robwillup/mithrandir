# Images
## Creating Docker images with Dockerfile

Dockerfile = Contains instructions for how to create a Docker image.

```Dockerfile
FROM microsoft/dotnet:aspnetcore-runtime
WORKDIR /app
COPY ./out .
ENTRYPOINT ["dotnet", "samplewebapp.dll"]
```

To build an image from the Dockerfile above we can just run this command:

```PowerShell
docker build -t samplewebapp .
```

## Multi-stage Dockerfiles

Single stage Dockerfile
- Copies in a pre-built application

Use Docker to build your application
- No developer SDK's needed locally

Multi-stage Dockerfile
- Stage 1: build the application in a container with the SDKs
- Stage 2: copy the published application into a container with the runtime

```Dockerfile
FROM microsoft/dotnet:sdk AS build-env
WORKDIR /app

COPY *.csproj ./
RUN dotnet restore

COPY . ./
RUN dotnet publish -c Release -o out

FROM microsoft/dotnet:aspnetcore-runtime
WORKDIR /app
COPY --from=build-env /app/out .
ENTRYPOINT ["dotnet", "samplewebapp.dll"]
```

We run dotnet restore separately to create a distinct layer for the dependencies. That's for efficiency.

