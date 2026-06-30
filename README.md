# Collider

A high-performance, standalone WebSocket-based signaling and static file server for WebRTC Peer-to-Peer (P2P) calls, written in Go.

Collider serves two purposes:
1.  **WebSocket Signaling Broker**: Relays SDP offers/answers and ICE candidates between two registered peers in a room.
2.  **Static HTTP Server**: Serves the built-in WebRTC client (`index.html`) on the root path `/` for quick P2P video call demos.

---

## Features

*   **100% Standalone**: No external room server (like Google App Engine) or database required.
*   **Built-in Web Client**: Serves a local HTML5/WebRTC video chat interface directly.
*   **Bypassed Origin Checks**: Allows local `file://` or cross-origin connections out of the box.
*   **Stateful Connection Management**: Cleans up resources when clients disconnect.

---

## Building

Since the Go files have been promoted to the root of the repository, you can build Collider directly:

1.  Clone or navigate to the repository root:
    ```bash
    cd apprtc
    ```
2.  Build the binary:
    ```bash
    go build -o bin/collider collidermain/main.go
    ```

The compiled executable will be located at `bin/collider`.

---

## Running

Start the server using the compiled binary from the root directory (so it can locate the root `index.html` file):

### A. Local HTTP / Unencrypted WebSocket (Recommended for testing)
To run on port `8089` without TLS:
```bash
./bin/collider -port=8089 -tls=false
```

### B. Production HTTPS / Secure WebSocket (WSS)
To run on port `443` with TLS enabled:
```bash
./bin/collider -port=443 -tls=true
```
*Note: Secure mode requires a valid certificate (`cert.pem`) and private key (`key.pem`) to be placed in the root of the repository.*

---

## Quick Start Demo

Once the server is running on port `8089` with `-tls=false`:

1.  Open your browser and navigate to **`http://localhost:8089/`** in two separate tabs.
2.  **Tab 1 (Peer A)**:
    *   Enter a Room ID (e.g., `test_room`).
    *   Click **"1. Start Call (Peer A)"**.
3.  **Tab 2 (Peer B)**:
    *   Enter the **same** Room ID (`test_room`).
    *   Click **"2. Join Call (Peer B)"**.
4.  The two tabs will negotiate and establish a direct P2P video call through the local Collider server.

---

## Protocol Specification

Collider communicates with clients using JSON frames over WebSocket at the `/ws` endpoint.

### 1. Registering in a Room (Client -> Server)
A client must register immediately after establishing the WebSocket connection:
```json
{
  "cmd": "register",
  "roomid": "room_name",
  "clientid": "unique_client_id"
}
```

### 2. Sending Signaling Messages (Client -> Server)
To send SDP or ICE candidates to the other peer in the room:
```json
{
  "cmd": "send",
  "msg": "<escaped_json_payload>"
}
```

### 3. Forwarding Messages (Server -> Client)
Collider wraps the payload and forwards it to the destination peer:
```json
{
  "msg": "<escaped_json_payload>"
}
```

---

## Testing

To run the Go test suite:
```bash
go test ./...
```
