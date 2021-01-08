# Design and Implement Logging

Loggin enables a centralized view and retention of required information.

Organizations may also require granular access to information even when in a central store along with rich analytics.

## Log Aggregation

Some components have their own native log store but can be difficult to work with due to various reasons

* Storage location and format
* Retention
* Permissions
* Integration for analysis, alerting

Ideally logs we care about should be aggregated to a common store

## Azure Metrics and Logs

* Azure has a large number of metrics and logs
* These vary by resource type
* Azure resource metrics are stored in the native store for 90 days and can optionally be sent to a target
* Azure resource logs must be sent to a target

## Azure Diagnostic Targets

* Azure Storage
  * Good for long term, cheap retention
* Even Hub
  * Good to sent to external solutions
* Log Analytics
  * Good for storage and analytical analysisrr