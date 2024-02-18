# **Go Language GRPC Server and Client**

**This is a basic gRPC server and client written in Go. It is based on the gRPC documentation**


<hr>

# **Proof of Concept**

    * Client -- > Server (Regular Architecture)
    * This is synchoronus and slow and can't be scaled (in case lots of server and mircoservices)


# **Concept of gRPC**

    * Client --> Server (Remote Calls)
    * This helps in calling the function directly from client to the server. It's like we have server and the client at the same place
    * Here instead of JSON, we use protobuf. Due to having small size the communication is fast

<hr>

<hr>

# **Client ===== Server**
* Streaming 
* Could be scaled to quite extent
* Continuous flow of data
* Asynchoronus that means there is no waiting for the server to resond back and hence it's fast
* Microservices and the BlockChain peers communication

<hr>

<hr>

# **Types of Communication that can happen?**

* **1. Regular Request and Response (Unary API)**
* **2. Server Streaming**
    * The client makes a request to the server and then there is a stream of message

* **Client Streaming**
    * The client sends a stream of requests to the server and then the server by the response

* **Bi-Directional Streaming**
    * Client and Server both communicate using streams. It completely asynchoronus in the nature. There is no order while sending that either client should send first or the server should.

<hr>

# **ProtoBuf File for the configuration**

**Here we have all the streaming configuration done and ready for the communication between the server and client. The direction of the stream and their returns has been declared. We now have to work on code, to generate some code.**

**Here we also need to create the protobuf file for the communication and the same could be done using**

**`protoc --go_out=. --go-grpc_out=.`**


**If in case the code is not running we need to download some dependencies**
    * go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
    * go get -u google.golang.org/protobuf/cmd/protoc-gen-go
    * go install google.golang.org/protobuf/cmd/protoc-gen-go
    * go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc
    * go install google.golang.org/grpc/cmd/protoc-gen-go-grpc

<hr>

**Once we are done with creation of the .pb file, run the `tidy` command in the kernel to get all the dependencies**

<hr>

<hr>

# **Talks on the API**

* In the server as well as the client for the unary we have seperate relation, where the client will send one request and the server will respond with the HelloResponse**

* We are sending a request with no params but we are expecting some response and the server should respond
