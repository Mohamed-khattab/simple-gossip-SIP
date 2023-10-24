# Gossip Algorithm Simulation in Go

## Overview

This Go program simulates a gossip algorithm for communication between nodes in a distributed system. The implementation, inspired by `system design principles book`, introduces features such as dynamic gossip intervals, simulated network delays, and handling unresponsive nodes.

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
## Handling Unresponsive Nodes
The system can handle unresponsive nodes by notifying other peers and implementing a method to handle notifications.

```go
func (n *Node) notifyPeersAboutUnresponsiveNode(unresponsiveNodeID int) {
    // ... (see code)
}

func (n *Node) HandleUnresponsiveNode(unresponsiveNodeID int) {
    // ... (see code)
}
```
## Main Function
The main function orchestrates the simulation, loading environment variables, creating nodes, establishing peer connections, and simulating gossiping over a specified duration.
```go 
func main() {
    // ... (see code)
} ```

## Logging
The writeLogToFile method writes the gossip log for each node to an external file.

```go
func (n *Node) writeLogToFile() {
    // ... (see code)
}
```
## Acknowledgment
This implementation is part of a learning project based on concepts from system design literature. The code is written in Go for familiarity with the language.
