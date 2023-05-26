# Examples of HTTP requests in Go

## How to make an HTTP POST request behind a proxy

```go
package main

import (
        "net/url"       // used to parse URL strings and to create the form data
        "net/http"      // used to create the HTTP client and request
        "strings"       // used to read the request body into bytes
        "strconv"       // used to convert an integer to ASCII
        "io/ioutil"     // used to read the response.Body into byte array
        "encoding/json" // used to deserialize
)

var (
        tokenURLStr = "https://login.microsoftonline.com/<tenant-id>/oauth2/v2.0/token"
        proxyStr    = "http://server:port"
)

func main() {
        data := url.Values{
                "client_id":     {"client-id"},
                "client_secret": {"client-secret"},
                "grant_type":    {"client_credentials"},
                "scope":         {"api://the-service/.default"},
        }
        
        proxyURL, err := url.Parse(proxyStr)
        if err != nil {
                print(err)
        }
        
        tokenURL, err := url.Parse(tokenURLStr)
        if err != nil {
                print(err)
        }
        
        transport := &http.Transport{
                Proxy: http.ProxyURL(proxyURL),
        }
        
        client := &http.Client{
                Transport: transport,
        }
        
        request, err := http.NewRequest("POST", tokenURL.String(), strings.NewReader(data.Encode()))
        ir err != nil {
                print(err)
        }
        request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
        request.Header.Add("Content-Length", strconv.Itoa(len(data.Encoded())))
        
        response, err := client.Do(request)
        if err != nil {
                print(err)
        }
        if response.StatusCode != httpStatusOK {
                print(response.StatusCode)
        }
        
        bodyBytes, err := ioutil.ReadAll(response.Body)
        if err != nil {
              print(err)
        }
        responseData := make(map[string]interface{})
        err = json.Unmarshal(bodyBytes, &responseData)
        print(reponseData["access_token"].(string))
}
```
