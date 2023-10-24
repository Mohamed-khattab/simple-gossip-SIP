# Gossip Algorithm Simulation in Go

## Overview

This Go program simulates a gossip algorithm for communication between nodes in a distributed system. The implementation, inspired by system design principles, introduces features such as dynamic gossip intervals, simulated network delays, and handling unresponsive nodes.

## Node Structure

The `Node` structure encapsulates essential properties of a gossip node:

```go
type Node struct {
    ID        int
    Data      string
    Peers     map[int]*Node
    mutex     sync.Mutex
    gossipLog []string
}
```
## Gossiping

 The `Gossip` method initiates the gossip process, asynchronously sending data to all peers, simulating network delays, 
 and handling unresponsive peers.
 ```go  
 func (n *Node) Gossip() {
     // ... (see code)
 } 
 ```
##  Simulating Delays
 Methods like simulateNetworkDelay and simulateProcessingDelay introduce random delays to simulate network and processing delays.
 
``` go
 func (n *Node) simulateNetworkDelay() {
     // ... (see code)
 }
 
 func (n *Node) simulateProcessingDelay() {
     // ... (see code)
 }
 ```
