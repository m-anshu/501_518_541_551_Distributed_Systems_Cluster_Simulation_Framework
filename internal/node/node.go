package node

import (
    "fmt"
    "context"
    "time"
    "log"
    "github.com/google/uuid"
    "github.com/docker/docker/api/types/container"
    "github.com/docker/docker/client"
)

// Node status constants
const (
    NodeStatusRunning  = "Running"
    NodeStatusStopped  = "Stopped"
    NodeStatusUnhealthy = "Unhealthy"
)

// Node structure to store node information
type Node struct {
    ID       string    `json:"id"`
    CPUs     int       `json:"cpus"`
    UsedCPUs int       `json:"used_cpus"`
    Status   string    `json:"status"`
    Pods     []string  `json:"pods"`
    CreatedAt time.Time `json:"created_at"`
}

// Function to create a new node container
//Name of the container is the node id
func CreateNodeContainer(cpus int) (string, error) {
    cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
    if err != nil {
        return "", err
    }
    containerName := fmt.Sprintf("node_container_%s", uuid.New().String())

    resp, err := cli.ContainerCreate(
        context.Background(),
        &container.Config{
            Image: "python:3.8-slim",
            Cmd:   []string{"sh", "-c", "while true; do sleep 30; done"},
        },
        nil, nil, nil, containerName)
    if err != nil {
        return "", err
    }

    err = cli.ContainerStart(context.Background(), resp.ID, container.StartOptions{})
    if err != nil {
        return "", err
    }

    // Verify container is running
    inspect, err := cli.ContainerInspect(context.Background(), containerName)
    if err != nil {
        return "", fmt.Errorf("failed to verify container status: %v", err)
    }

    if !inspect.State.Running {
        return "", fmt.Errorf("container failed to start")
    }

    return containerName, nil
}
func DeleteNodeContainer(nodeID string) (error){

    cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
    if err != nil {
        return err
    }
    ctx := context.Background()

    // Attempt to stop the container (if not already stopped).  Force stop if needed.
    if err := cli.ContainerStop(ctx, nodeID, container.StopOptions{}); err != nil {
        log.Printf("Error stopping container %s: %v", nodeID, err)
        // Continue even if stopping fails.
    }
    // Remove the container.
    if err := cli.ContainerRemove(ctx,nodeID, 
        // Force remove the container so it gets cleaned up.
        container.RemoveOptions{Force: true}); err != nil {
        return err
    }
    return nil
}

func StopNodeContainer(nodeID string) error {
    cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
    if err != nil {
        return err
    }
    ctx := context.Background()

    // Get current container state
    inspect, err := cli.ContainerInspect(ctx, nodeID)
    if err != nil {
        return fmt.Errorf("failed to inspect container: %v", err)
    }

    // Only stop if running
    if inspect.State.Running {
        // Use graceful stop with default timeout
        if err := cli.ContainerStop(ctx, nodeID, container.StopOptions{}); err != nil {
            log.Printf("Error stopping container %s: %v", nodeID, err)
            return err
        }
        log.Printf("Container %s stopped gracefully", nodeID)
    } else {
        log.Printf("Container %s is already stopped", nodeID)
    }

    return nil
}

// Function to create a new node container with the same id as the failed node
// Function to restart a node container while preserving its ID and data
func RestartNodeContainer(nodeID string) error {
    cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
    if err != nil {
        return fmt.Errorf("failed to create Docker client: %v", err)
    }

    ctx := context.Background()

    // First stop the container
    if err := cli.ContainerStop(ctx, nodeID, container.StopOptions{}); err != nil {
        return fmt.Errorf("failed to stop container: %v", err)
    }

    // Wait a moment to ensure the container is fully stopped
    time.Sleep(2 * time.Second)

    // Start the same container again
    if err := cli.ContainerStart(ctx, nodeID, container.StartOptions{}); err != nil {
        return fmt.Errorf("failed to start container: %v", err)
    }

    // Wait a moment to ensure the container is fully started
    time.Sleep(2 * time.Second)

    // Verify the container is running
    inspect, err := cli.ContainerInspect(ctx, nodeID)
    if err != nil {
        return fmt.Errorf("failed to inspect container: %v", err)
    }

    if !inspect.State.Running {
        return fmt.Errorf("container is not running after restart")
    }

    return nil
}
