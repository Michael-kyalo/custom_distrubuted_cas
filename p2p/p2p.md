# PEER TO PEER NETWORK

contains interfaces and classes that describe our peer to peer communication

## Files

- ### transport:

this file contains the interface definitions relavant to creating communication channels and nodes 

### interfaces:
    
- Peer : defines a node in the p2p network

```go
type Peer interface {
}
```
- Transport : defines a transport channel (handler) in the p2p network
```go
type Transport interface {}
```
