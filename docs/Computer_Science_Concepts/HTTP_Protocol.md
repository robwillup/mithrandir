# About HTTP Protocol

It's important to have a very clear understanding of HTTP. For a long time
as a software developer I simply worked with the web frameworks without
worrying too much about having a more foundational knowledge about the
protocol all these frameworks use.

## What is HTTP?

`The protocol of the web`

It's an acronym that stands for `Hyper Text Transfer Protocol`. It's used for
communication between web servers and web clients. It's known as the
"protocol of the web".

> Every time you open up your browser and you visit a web page, or submit
> a form, or you click a button that sends some form of AJAX request, you're
> using HTTP and you're going through the request/response cycle. You make
> a request and you get some response back.

## HTTP is stateless

Every request is completely independent and it has no context of other
requests. Each request can be understood as a single transaction.

## HTTPS

This is `Hyper Text Transfer Protocol Secure` and it is used to send data
encrypted using SSL/TLS. The certificate is installed on web host.

## HTTP Methods

* GET
  * Retrieves data from the server
* POST
  * Submit data to the server
* PUT
  * Update data already on the server
* DELETE
  * Deletes data from the server

## HTTP Header Fields

Headers are metadata of a request and response.

## Status Code

* 1xx: Informational
  * Request recieved / processing
* 2xx: Success
  * Successfully received, understood and accepted
  * 200: OK
  * 201: Created
* 3xx: Redirect
  * Further action must be taken / redirect
  * 301: Moved to new URL
  * 304: Not modified (visit page that's cached version)
* 4xx: Client error
  * Request does not have what it needs
  * 400: bad request
  * 401: unauthorized
  * 404: not found
* 5xx: Server error
  * Server failed to fulfil an apparent valid request
  * 500: internal server error

## HTTP/2

* Major revision of HTTP
* Under the hood changes (applications don't need to change how they use HTTP)
* Respond with more data
* Reduce latency by enabling full request and response multiplexing
* Fast, efficient & secure
