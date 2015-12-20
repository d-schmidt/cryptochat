# cryptochat
Rule 1: never build your own crypto  

Seriously: Don't assume this program provides a safe way to chat!  

## What is this then?
This is a little project to learn how to use Websockets with different languages.  
The client frontend uses AngularJS to connect to a local Go webserver which connects to an embedded Jetty Java server. All connections are websockets.  
  
The cryptography is just an added bonus after I got all the different websockets and servers running.  
It uses Public-Key-Cryptography RSA and AES to create an End-to-End encrypted tunnel between two Golang client servers. The Jetty server could MITM these connection. Multiple 1:1 connections are supported at the same time.  

### Server
build Java server with `mvn clean build`

### Client
build the Golang client-server with:  
`go get golang.org/x/crypto/sha3`  
`go get golang.org/x/net/websocket`  
`go get github.com/toqueteos/webbrowser`  
`go build`  

### Usage
- start the Java server
- start a client and register with a name
- start a second client and register with another name
- connect the clients with each others names
- chat
