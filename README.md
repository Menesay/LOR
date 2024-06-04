# Lightweight Onion Router

![demo](/LOR/demo.png "Demo")

## Overview

Lightweight Onion Router (LOR) is a minimalist proxy chain structure that routes your HTTP traffic through multiple proxy servers. This project was built using Go and provides a simple setup to get you started.

## Setup

### Step 1: Install Go

Ensure you have Go installed on your machine. You can download it from the official [Go website](https://golang.org/dl/).

### Step 2: Clone the Repository

Clone this repository to your local machine:

```bash
git clone https://github.com/Menesay/LOR.git
cd LOR
```

### Step 3: Configure Server IPs and Ports

Open the source code and configure the IP addresses and ports for the LOR servers. The default configuration is as follows:

| Server | IP Address    | Host Port | Proxy Port |
|--------|---------------|-----------|------------|
| lor-0  | 127.0.0.10    | 1400      | 1401       |
| lor-1  | 127.0.0.20    | 1400      | 1401       |
| lor-2  | 127.0.0.30    | 1400      | 1401       |

### Step 4: Compile the Code

Compile the source code for each server using the following commands:

```bash
go build lor-0.go
go build lor-1.go
go build lor-2.go
```

## Running the Servers

### Step 1: Start the Servers in Sequence

Due to the network structure, the servers must be started in reverse order:

1. First, run `lor-2`:
    ```bash
    ./lor-2
    ```

2. Then, run `lor-1`:
    ```bash
    ./lor-1
    ```

3. Finally, run `lor-0`:
    ```bash
    ./lor-0
    ```

### Step 2: Connect to LOR

To connect to the LOR network, follow these steps:

1. Open a terminal and use `nc` (netcat) to connect to `lor-0`:
    ```bash
    nc 127.0.0.10 1400
    ```

2. Send the hostname:
    ```bash
    menesay.duckdns.org
    ```

3. Send the port number:
    ```bash
    80
    ```

Wait for a moment, and the HTTP proxy will be running at `127.0.0.10:1401`.

## Additional Information

- Make sure your firewall settings allow connections on the specified IP addresses and ports.
- The order of operations is crucial for the proxychain to function correctly.
- For any issues or contributions, feel free to open a pull request or raise an issue on the GitHub repository.

Enjoy with Lightweight Onion Router!

---

Feel free to reach out if you have any questions or need further assistance.

Happy browsing! 🚀
