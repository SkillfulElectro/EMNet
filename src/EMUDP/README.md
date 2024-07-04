- EMUDP is provided to help clients get their udp socket public_ip:port

### Usage:

1. **Start the Server**:
    - Run the server on the desired IP and port.
    ```
    EMUDP -ip 127.0.0.1 -port 12345
    ```
    - it supports all interfaces listening and dynamic port just like EMNet
2. **Expected Output**:
    - The client should print the response from the server, which will be its own IP and port in the format `ipv4:port`.
    ```
    127.0.0.1:client_port
    ```

Make sure the server is running before you run the client. The client will send a "ping" message to the server and should receive a response with its own IP and port. This will help verify that the server is functioning correctly.
