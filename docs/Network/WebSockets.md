# WebSockets

WebSocket is a computer communication protocol providing full-duplex communication channels over a single TCP connection.

WebSocket is distinct from HTTP. Both protocols are located at layer 7 in the OSI model and depend on TCP at layer 4. WebSocket is compatible with HTTP.

The WebSocket protocol enables interaction between a client application and a web server with lower overhead than half-duplex alternatives such as HTTP polling, facilitating real-time data transfer from and to the server. This is made possible by providing a standardized way for the server to send content to the client without being first requested by the client, and allowing messages to be passed back and forth while keeping the connection open. In this way, a two-way ongoing conversation can take place between the client and the server. The communications are usually done over TCP port number 443, which is beneficial for environments that block non-web Internet connections using a firewall.

Unlike HTTP, WebSocket provides full-duplex communication. Additionally, WebSocket enables streams of message on top of TCP. TCP alone deals with streams of bytes with no inherent concept of a message.

The WebSocket protocol specification defines `ws` (WebSocket) and `wss` (WebSocket Secure) as two new uniform resource identifier schemes that are used for unencrypted and encrypted connections, respectively. 