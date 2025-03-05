# Broadcast Server

This project implements a simple broadcast server using WebSockets in Go. It allows multiple clients to connect to the server and send messages that are broadcasted to all connected clients. This is a great way to understand real-time communication in applications.

This project is one of the backend projects from [roadmap.sh](https://roadmap.sh/projects/broadcast-server).


## Project Structure

```
broadcast-server
├── cmd
│   ├── client
│   │   └── main.go        # Entry point for the client application
│   └── server
│       └── main.go        # Entry point for the server application
├── go.mod                  # Module definition and dependencies
└── README.md               # Project documentation
```

## Getting Started

### Prerequisites

- Go 1.16 or later installed on your machine.

### Running the Server

To start the server, navigate to the project directory and run the following command:

```
go run cmd/server/main.go start
```

This will start the server and listen for incoming client connections on the specified port.

### Connecting a Client

To connect a client to the server, open another terminal window and run:

```
go run cmd/client/main.go connect
```

Once connected, you can send messages to the server, which will be broadcasted to all connected clients.

### Testing the Application

You can open multiple terminal windows to run multiple clients and test the broadcasting feature by sending messages from different clients.

### Error Handling

The server is designed to handle client disconnections gracefully and will remove clients from the list of connected clients when they disconnect.

## License

This project is open-source and available under the MIT License.