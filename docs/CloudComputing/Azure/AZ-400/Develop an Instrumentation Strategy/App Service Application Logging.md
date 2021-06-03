# Enable and configure App Service application logging

## What are app logs?

App logs are the output of runtime trace statements in app code.

> Trace statement is a message that is output during execution of a debug session.

For example, you might want to check some logic in your code
by adding a trace to show when a particular function is being
processed, or you might only want to see logged message when
a particular level of error has occurred.

*App logging is primarily for apps in pre-production and for
troublesome issues*, because excessive logs can carry a performance
git quickly consume storage; for this reason, logging to the
file system is automatically disabled after 12 hours.

### --> Highlights!

> * App logs are the output of runtime trace statements in your code.
> * Trace statement is a message that is output during execution of debug session.
> * App logging is primarily for apps in **pre-production**.
> * Excessive logs have a performance hit and consume storage
> * Logging to the file system is automatically disabled after 12 hours.

App logging has scale limitations, primary because **files**
are being used to save the logged output. If you have multiple
instances of an app, and the same storage used is shared acros