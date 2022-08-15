# GRPC Demo


## Using gRPCurl

[gRPCurl](https://github.com/fullstorydev/grpcurl) is like `curl` but for gRPC.

It's a very useful tool for interacting with gRPC servers.

- Install: https://github.com/fullstorydev/grpcurl#installation
- List all exposed services on `host:port`, e.g.:
    ```
    grpcurl --plaintext localhost:50051 list
    ```
- List all methods of a specified service, e.g.:
    ```
    grpcurl --plaintext localhost:50051 list grpc.demo.Calculator
    ```
- Describe a service or message, e.g.:
    ```
    grpcurl --plaintext localhost:50051 describe grpc.demo.Calculator.NthFibonacci
    ```
    ```
    grpcurl --plaintext localhost:50051 describe grpc.demo.FibonacciParams
    ```
- Invoke an RPC, e.g.:
    ```
    grpcurl --plaintext -d '{"n": 4}' localhost:50051 grpc.demo.Calculator.NthFibonacci
    ```
    ```
    grpcurl --plaintext -d '{"numbers": [6,0,3,9]}' localhost:50051 grpc.demo.Calculator.Sum
    ```
    ```
    grpcurl --plaintext -d '{"min": 10, "max": 20, "count": 1000}' localhost:50051 grpc.demo.Calculator.RandomStream
    ```
