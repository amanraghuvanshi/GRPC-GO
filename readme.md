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


