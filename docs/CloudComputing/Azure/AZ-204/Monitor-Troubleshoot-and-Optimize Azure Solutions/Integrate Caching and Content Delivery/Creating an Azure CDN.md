# Creating an Azure CDN

Creating an Azure CDN is simple and similar to creating other Azure resources. In the portal, just search for CDN and create a new one providing a name, and selecting a subscription, resource group, location and pricing tier.

After the CDN is created, you need to add an endpoint to it, and that's where you get to select the Origin type, (web app for example). Origin hostname is your app service for example.

## How Caching Works

**What is a Content Delivery Network?**
* Globally distributed network
* Reduced asset load times
* Reduced hosting bandwidth
* Increased availability and redundancy
* Protection against DDoS

**Content Types**

* Static Content
  * images
  * CSS files
  * JS files
* Dynamic content
  * changes on user interaction
  * dashboards
  * Query results


CDN allows the distribution of content to regions that are geographically closer to users. The first time a user makes a request the web server will respond with the content and the CDN will cache it (during time to live). Next time that user makes this request, this content will be retrieved from the PoP (point of present) closed to him/her.

## Configuring Azure CDN

**Pricing Tiers**

* Standard Microsoft
* Standard Akamai
* Standard Verizon
* Premium Verizon

**HTTP Headers**

Cache directive headers

* Control caching using headers
  * Cache-control
  * Expires
* Cache-control takes precedence over expires
* Standard Akamai
  * Max-age
  * No-cache
  * No-store

**Configuration Options**

* Custom DNS Domain
* Compression
* Caching Rules
* Geo Filtering
* Optimization
* Dynamic Site Acceleration

Microsoft recommends that if you have a site with lots of dynamic content then you should use Azure FrontDoor instead of Azure CDN.

## Azure CDN in Code

NuGet packages:

* Microsoft.Azure.Management.Cdn

