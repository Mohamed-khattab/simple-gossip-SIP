package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"math/rand"
	"os"
	"strconv"
	"sync"
	"time"
)

// Node represents a gossip node
type Node struct {
	ID        int
	Data      string
	Peers     map[int]*Node
	mutex     sync.Mutex
	gossipLog []string
}

// Gossip sends data to all peers and handles unresponsive peers
func (n *Node) Gossip() {
	n.mutex.Lock()
	defer n.mutex.Unlock()

	// Create a wait group to wait for all gossip messages to be sent
	var wg sync.WaitGroup

	for peerID, peer := range n.Peers {
		// Increment the wait group counter
		wg.Add(1)

		go func(peerID int, peer *Node) {
			defer wg.Done()

			// Simulate network delay
			n.simulateNetworkDelay()

			// Attempt to send data to the peer
			err := n.sendDataToPeer(peerID, peer)
			if err != nil {
				// Handle unresponsive peer
				fmt.Printf("Node %d encountered an error sending data to Node %d: %v\n", n.ID, peerID, err)
				// Notify other nodes about the unresponsive peer
				n.notifyPeersAboutUnresponsiveNode(peerID)
			}
		}(peerID, peer)
	}

	// Wait for all gossip messages to be sent
	wg.Wait()
}

// sendDataToPeer sends data to a peer
func (n *Node) sendDataToPeer(peerID int, peer *Node) error {
	// Send data to the selected peer
	fmt.Printf("Node %d sent data to Node %d: %s\n", n.ID, peerID, n.Data)
	peer.ReceiveData(n.Data)
	n.gossipLog = append(n.gossipLog, fmt.Sprintf("Sent to Node %d", peerID))
	return nil
}

// ReceiveData receives data from a peer
func (n *Node) ReceiveData(data string) {
	// Simulate processing delay
	n.simulateProcessingDelay()

	// Process received data
	fmt.Printf("Node %d received data: %s\n", n.ID, data)

	// Update own data
	n.mutex.Lock()
	n.Data = data
	n.gossipLog = append(n.gossipLog, "Received")
	n.mutex.Unlock()
}

func (n *Node) simulateNetworkDelay() {
	// Simulate a random network delay
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(100)))
}

func (n *Node) simulateProcessingDelay() {
	// Simulate a random processing delay
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(100)))
}

func (n *Node) notifyPeersAboutUnresponsiveNode(unresponsiveNodeID int) {
	// Notify other nodes about the unresponsive peer
	for peerID, peer := range n.Peers {
		if peerID != unresponsiveNodeID {
			// Simulate network delay
			n.simulateNetworkDelay()
			// Notify the peer about the unresponsive node
			fmt.Printf("Node %d notifying Node %d about unresponsive Node %d\n", n.ID, peerID, unresponsiveNodeID)
			peer.HandleUnresponsiveNode(unresponsiveNodeID)
		}
	}
}

func (n *Node) HandleUnresponsiveNode(unresponsiveNodeID int) {
	// Handle the notification about an unresponsive node
	fmt.Printf("Node %d received notification about unresponsive Node %d\n", n.ID, unresponsiveNodeID)
	// Implement your logic to handle the unresponsive node (e.g., mark it as inactive)
}

func (n *Node) writeLogToFile() {
	logFileName := fmt.Sprintf("node%d.log", n.ID)
	file, err := os.Create(logFileName)
	if err != nil {
		log.Fatalf("Error creating log file for Node %d: %v", n.ID, err)
	}
	defer file.Close()

	logger := log.New(file, "", log.LstdFlags)

	for _, entry := range n.gossipLog {
		logger.Println(entry)
	}
}

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Get the number of nodes from the environment variable
	numNodesStr := os.Getenv("NUM_NODES")
	numNodes, err := strconv.Atoi(numNodesStr)
	if err != nil {
		log.Fatal("Error converting NUM_NODES to an integer")
	}

	// Create nodes
	nodes := make(map[int]*Node)
	for i := 1; i <= numNodes; i++ {
		nodes[i] = &Node{
			ID:        i,
			Data:      fmt.Sprintf("Initial data from Node %d", i),
			Peers:     make(map[int]*Node),
			gossipLog: make([]string, 0),
		}
	}

	// Establish random peer connections
	for _, node := range nodes {
		for id, peer := range nodes {
			if id != node.ID {
				node.Peers[id] = peer
			}
		}
	}

	// Simulate the network for 10 seconds with a gossip interval of 1 second
	duration := 10 * time.Second
	gossipInterval := 1 * time.Second
	endTime := time.Now().Add(duration)

	for time.Now().Before(endTime) {
		// Select a random node to initiate gossip
		nodeID := rand.Intn(numNodes) + 1
		go nodes[nodeID].Gossip()

		// Wait for the next gossip interval
		time.Sleep(gossipInterval)
	}

	// Write gossip logs to external files for each node
	for _, node := range nodes {
		node.writeLogToFile()
	}
}
