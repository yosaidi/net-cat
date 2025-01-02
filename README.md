# Netcat Server
A simple chat server written in Go.

## Features
* Handles multiple clients

* Broadcasts messages to all connected clients
* Logs messages and user actions (join, leave, rename)
* Supports renaming of users
* Limits the number of connected clients to 10
* Supports flags for renaming and changing color
## Usage
To run the server, simply execute the following command:

```sh
# clone the project
git clone https://github.com/yosaidi/net-cat.git 
# build application
go build -o TCPChat
# launch server
./TCPChat [Port number]
# in a new terminal window, launch the client
nc localhost [Port number]
```
Replace `[Port number]` with the desired port number.

## Code Structure
The code is organized into several packages:

* `server`: contains the main server logic, including the Server and Client structs, as well as functions for handling connections, broadcasting messages, and logging.

* `main`: contains the `main` function, which creates a new `Server` instance and starts it.
## Functions
`AsciiArt()`
Returns a string containing an ASCII art logo.

`Broadcast(client, message, messagetype)`
Broadcasts a message to all connected clients. The `messagetype` parameter determines the type of message (join, leave, or regular message).

`HandleConnection(client)`
Handles incoming messages from a client and broadcasts them to all other clients.

`User(conn)`
Checks if a username is already taken and prompts the user to enter a new name if necessary.

`HandleError(err)`
Handles errors by printing an error message and returning.

`Run()`
Starts the server and listens for incoming connections.

## Structs
`Server`
Represents the server instance, containing a slice of Client structs and a mutex for synchronization.

`Client`
Represents a connected client, containing a `net.Conn` object, a username, and a slice of messages.

`Historic`
Represents a log entry, containing a timestamp, username, and message.

## Constants
`IP`
The IP address to listen on (localhost).

`PORT`
The port number to listen on (8081).

# Flags
The following flags are available:

* `rename`: Renames the user to a new name.
* `color`: Changes the`color of the user's name.
# Commands
The following commands are available:

* `/rename <newname>`: Renames the user to  `<newname>`.

* `/color <color>`: Changes the color of the user's name to `<color>`. 

 #### Available colors are:

* yellow
* red
* blue
* magenta
* cyan
* green
* white

 