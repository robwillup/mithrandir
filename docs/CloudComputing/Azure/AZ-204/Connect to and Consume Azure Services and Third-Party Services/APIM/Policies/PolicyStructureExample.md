# Policy Structure Example

```xml
<policies>
    <inbound>
        <rate-limit calls="5" renewal-period="10" />
        <cache-lookup vary-by-developer="false" vary-by-developer-groups="false" must-revalidate="true" downstream-caching-type="none" caching-type="internal" />
        <base />
        <backend>
            </base>
        </backend>
        <outbound>
            <cache-store duration="60"/>
            </base>
        </outbound>
        <on-error>
            </base>
        </on-error>
    </inbound>
</policies>
```