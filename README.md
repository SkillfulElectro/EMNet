# EMNet: A Go TCP Server with Dynamic Port Selection and Peer Management

## Overview

EMNet is a Go server program that starts a TCP server based on command-line arguments or defaults, and it manages peer IPs using a global map. The server handles two types of JSON messages: `ADD_PEER` and `ASK_IP`.

## Features

1. **Dynamic Port Selection**:
    - If both IP and port are specified, the server starts on the given IP and port.
    - If only the IP is specified, the server starts on the first available port between 80 and 60000.
    - If no arguments are provided, the server starts on `0.0.0.0` and the first available port between 80 and 60000.

2. **Peer Management**:
    - `ADD_PEER`: Adds a peer IP to the global map with a generated random key.
    - `ASK_IP`: Retrieves and removes a peer IP from the global map based on a given key.

## Usage

### Command-Line Arguments

The server can be started with the following command-line arguments:

- `-ip IPV4` (Optional): IP address to bind the server.
- `-port PORT` (Optional): Port to bind the server.

### Examples

1. **Start server on a specified IP and port**:
    ```sh
    ./EMNet -ip 192.168.1.10 -port 8080
    ```

2. **Start server on a specified IP with dynamic port selection**:
    ```sh
    ./EMNet -ip 192.168.1.10
    ```

3. **Start server with default IP (`0.0.0.0`) and dynamic port selection**:
    ```sh
    ./EMNet
    ```

### JSON Message Handling

The server processes JSON messages sent by clients. The messages should be enclosed in `{}`. The following message types are supported:

1. **ADD_PEER**:
    - **Request**:
        ```json
        {
            "msg": "ADD_PEER",
            "IP_Peer": "192.168.1.20:63413"
        }
        ```
    - **Response (Success)**:
        ```json
        {
            "msg": "SUCCESS",
            "id": 123456
        }
        ```
    - **Response (Failure)**:
        ```json
        {
            "msg": "FAILED"
        }
        ```

2. **ASK_IP**:
    - **Request**:
        ```json
        {
            "msg": "ASK_IP",
            "id": 123456
        }
        ```
    - **Response (Success)**:
        ```json
        {
            "msg": "SUCCESS",
            "IP": "192.168.1.20:34721"
        }
        ```
    - **Response (Failure)**:
        ```json
        {
            "msg": "FAILED"
        }
        ```

## Building and Running EMNet

### Building

To build the server, navigate to the directory containing the code and run:

```sh
go build -o EMNet
```

### Running

To run the server, use one of the following commands based on your requirements:

1. **With specified IP and port**:
    ```sh
    ./EMNet -ip 192.168.1.10 -port 8080
    ```

2. **With specified IP and dynamic port**:
    ```sh
    ./EMNet -ip 192.168.1.10
    ```

3. **With default IP and dynamic port**:
    ```sh
    ./EMNet
    ```

## Contributing

Contributions are welcome! Please fork the repository and submit a pull request for any improvements or bug fixes.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
