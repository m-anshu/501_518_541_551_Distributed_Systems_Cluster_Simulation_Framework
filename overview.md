# KuberKlone - A Kubernetes-like Cluster Simulator

KuberKlone is a lightweight, simulation-based distributed system that mimics core Kubernetes cluster management functionalities. It provides a simplified yet comprehensive platform for demonstrating key distributed computing concepts.

## Features

### Core Components

1. **API Server (Central Control Unit)**
   - Node Manager: Tracks registered nodes and their statuses
   - Pod Scheduler: Assigns pods to available nodes
   - Health Monitor: Receives health signals from nodes

2. **Node Management**
   - Add nodes to the cluster with specified CPU cores
   - List all nodes and their health status
   - Automatic node health monitoring
   - Node failure detection and recovery

3. **Pod Management**
   - Launch pods with specific CPU requirements
   - Automatic pod scheduling based on resource availability
   - Pod rescheduling in case of node failures

4. **Health Monitoring**
   - Periodic heartbeat signals from nodes
   - Automatic failure detection
   - Node health status tracking
   - Automatic recovery mechanisms

## Project Structure

```
KuberKopy/
├── cmd/              # Command-line interface
├── config/           # Configuration files
├── docker/           # Docker-related files
├── internal/         # Core implementation
│   ├── api/         # API server implementation
│   ├── health/      # Health monitoring
│   ├── node/        # Node management
│   ├── pod/         # Pod management
│   └── scheduler/   # Pod scheduling
├── tests/           # Test files
└── main.go          # Application entry point
```

## Prerequisites

- Go 1.16 or later
- Docker
- Make (optional, for using Makefile commands)

## Installation

1. Clone the repository:
```bash
git clone <repository-url>
cd KuberKlone
```

2. Install dependencies:
```bash
go mod download
```

## Running the Application

1. Start the API server:
```bash
go run main.go [port]
```
The default port is 8080 if not specified.

2. The API server will be available at `http://localhost:8080`

## API Endpoints

### Node Operations
- `POST /add_node` - Add a new node to the cluster
- `GET /nodes` - List all nodes and their status
- `PUT /restart_node` - Restart a specific node
- `DELETE /delete_node` - Remove a node from the cluster

### Pod Operations
- `POST /add_pod` - Launch a new pod with specified CPU requirements

## Example Usage

1. Add a node:
```bash
curl -X POST http://localhost:8080/add_node -H "Content-Type: application/json" -d '{"cpus": 4}'
```

2. List all nodes:
```bash
curl http://localhost:8080/nodes
```

3. Launch a pod:
```bash
curl -X POST http://localhost:8080/add_pod -H "Content-Type: application/json" -d '{"cpus": 2}'
```

## Health Monitoring

The system automatically:
- Monitors node health every 10 seconds
- Detects node failures
- Attempts to restart failed nodes
- Reschedules pods from failed nodes to healthy ones

## Contributing

1. Fork the repository
2. Create a feature branch
3. Commit your changes
4. Push to the branch
5. Create a Pull Request

## License

[Specify your license here]

## Acknowledgments

- This project is a simplified implementation of Kubernetes concepts
- Designed for educational purposes to understand distributed systems
