# gRPC Client Demo with Round-Robin Load Balancing

This repository contains a demo for a gRPC client configured to use the **round-robin** load balancing policy in a Kubernetes environment running on Docker Desktop.

## Prerequisites

Before you begin, ensure you have the following installed:

- [Docker Desktop](https://www.docker.com/products/docker-desktop/) with Kubernetes enabled
- [kubectl](https://kubernetes.io/docs/tasks/tools/) command-line tool
- Bash-compatible shell (for running `setting.sh`)

## Setup Instructions

Follow these steps to set up and run the demo:

1. **Enable Kubernetes on Docker Desktop**
   - Open Docker Desktop.
   - Navigate to **Settings > Kubernetes**.
   - Check the box for **Enable Kubernetes** and click **Apply & Restart**.
   - Wait for Kubernetes to start (this may take a few minutes).

2. **Install kubectl**
   - Follow the official [kubectl installation guide](https://kubernetes.io/docs/tasks/tools/#kubectl) for your operating system.
   - Verify installation by running:
     ```bash
     kubectl version --client
     ```

3. **Set Kubernetes Context**
   - Configure `kubectl` to use the Docker Desktop context:
     ```bash
     kubectl config use-context docker-desktop
     ```
   - Verify the context:
     ```bash
     kubectl config current-context
     ```

4. **Run the Setup Script**
   - Execute the provided `setting.sh` script to configure the environment:
     ```bash
     ./setting.sh
     ```
   - Ensure the script has executable permissions. If needed, run:
     ```bash
     chmod +x setting.sh
     ```

## Running the Demo

1. Ensure your Kubernetes cluster is running:
   ```bash
   kubectl get nodes