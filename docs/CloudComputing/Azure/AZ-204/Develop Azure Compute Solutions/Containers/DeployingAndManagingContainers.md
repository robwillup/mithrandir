# Docker Basics

* Images 
    * application + dependencies
    * built from Dockerfile
    * Tagged = rwill/voi:7
    * Publish to a "container registry"
         * Docker Hub
         * Azure Container Registry
* Container
    * instance based on the image
    * run on a Docker host
    * works the same everywhere
    * provides memory and CPU
    * Publish ports
    * Disk access (image only)
    * can be stopped and restarted
    * multiple containers at the same time 
* Data Storage
    * Layers are read-only
    * Layers are shared
    * Use volumes to persist data 
    * Mount volumes to a path in the container - e.g. /var/lib/mysql
