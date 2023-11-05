# Review Monitor and Logging

Areas of Focus

* Enabling App Service Logging
  * Configure web server to log to the FileSystem
    * az webapp log config...
  * Configure app logging to Azure Blob storage (windows only)
    * az webapp log config
  * Configure container logging to the file system (Linux only)
    * az webapp log config ... --docker-container-logging filesystem
  * Tail logs from app service app
    * az webapp log tail...
* Transient Faults
  * application should log transient faults
  * a retry strategy should be in place where needed
  * retry logic is already built into most SDK interactions
  * Implement architectural patterns that help with transient faults
* Configuring Docker Containers
  * env vars
* Web Test Alerts
  * URL Ping
  * Multi-step web
  * Custom