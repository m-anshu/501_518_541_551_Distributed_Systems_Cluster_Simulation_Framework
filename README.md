# Step-by-Step Guide for KuberKopy

## 1. Prerequisites Installation

### Windows
1. Install Go:
   - Download Go from [https://golang.org/dl/](https://golang.org/dl/)
   - Run the installer and follow the instructions
   - Verify installation by opening PowerShell and running:
     ```powershell
     go version
     ```

2. Install Docker:
   - Download Docker Desktop for Windows from [https://www.docker.com/products/docker-desktop](https://www.docker.com/products/docker-desktop)
   - Run the installer and follow the instructions
   - Start Docker Desktop
   - Verify installation by running:
     ```powershell
     docker --version
     ```

3. Install Git:
   - Download Git from [https://git-scm.com/download/win](https://git-scm.com/download/win)
   - Run the installer and follow the instructions
   - Verify installation by running:
     ```powershell
     git --version
     ```

## 2. Project Setup

1. Clone the repository:
   ```powershell
   git clone <repository-url>
   cd <into repo>
   ```

2. Install project dependencies:
   ```powershell
   go mod download
   ```

3. Build the CLI:
   ```powershell
   go build -o cluster-cli.exe .\cmd\cli.go
   ```

4. Pull Docker Image for containers
```powershell
   docker pull python:3.8-slim
   ```

## 3. Running the Project

1. Start the API server:
   ```powershell
   go run main.go
   ```
   The server will start on port 8080 by default.

2. In a new terminal window, verify the server is running:
   ```powershell
   curl http://localhost:8080/nodes
   ```
   You should see an empty list of nodes `[]` if this is your first run.

## 4. Using the CLI Interface

### Node Operations

1. List all nodes:
   ```powershell
   ./cluster-cli nodes
   ```

2. Add a new node:
   ```powershell
   ./cluster-cli add-node --cpus <number_of_cpus>
   ```
   Example:
   ```powershell
   ./cluster-cli add-node --cpus 4
   ```

3. Delete a node:
   ```powershell
   ./cluster-cli delete-node --node-id <node_id>
   ```
   Example:
   ```powershell
   ./cluster-cli delete-node --node-id "node_container_ce27d8ec-5cf7-43ad-80c4-0aabd089d608"
   ```

4. Restart a node:
   ```powershell
   ./cluster-cli restart-node --node-id <node_id>
   ```
   Example:
   ```powershell
   ./cluster-cli restart-node --node-id "node_container_9c134f04-f5b3-475b-a6ac-7d53861652b3"
   ```

### Pod Operations

1. Add a new pod:
   ```powershell
   ./cluster-cli add-pod --cpus <number_of_cpus> --algorithm <scheduling_algorithm>
   ```
   Available scheduling algorithms:
   - `best_fit`: Places the pod on the node with the smallest available capacity that can accommodate it
   - `worst_fit`: Places the pod on the node with the largest available capacity
   - `first_fit`: Places the pod on the first node that can accommodate it (default)

   Examples:
   ```powershell
   ./cluster-cli add-pod --cpus 2
   ./cluster-cli add-pod --cpus 2 --algorithm best_fit
   ./cluster-cli add-pod --cpus 2 --algorithm worst_fit
   ```

## 5. Monitoring and Troubleshooting

1. Check node health:
   - The system automatically monitors node health every 10 seconds
   - You can view node status using:
     ```powershell
     ./cluster-cli nodes
     ```

2. View pod status:
   - Pod status is included in the node information
   - Use the nodes command to see pod assignments:
     ```powershell
     ./cluster-cli nodes
     ```

3. Common issues and solutions:
   - If Docker is not running:
     ```powershell
     # Start Docker Desktop
     # Wait for Docker to be ready
     ```
   
   - If the API server is not responding:
     ```powershell
     # Check if the server is running
     # Restart the server if needed
     go run main.go
     ```
   
   - If a node fails:
     ```powershell
     # The system will automatically attempt to restart it
     # You can manually restart it using:
     ./cluster-cli restart-node --node-id <node_id>
     ```

## 6. Cleanup

1. Stop the API server:
   - Press Ctrl+C in the terminal running the server

2. Remove Docker containers (if needed):
   ```powershell
   docker ps -a
   docker rm <container_id>
   ```



