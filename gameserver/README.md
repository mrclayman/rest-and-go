# SERVER

As per the task requirements, the server component has been implemented in the Go language. The server has been implemented using only standard library features except the **WebSocket** library (https://github.com/gorilla/websocket), which is used to implement WebSocket functionality of the server's network interface.

The server both accepts requests and dispatches responses as JSON structures.

## Architecture
The component is split into 3 main packages:

 * The *core* package that implements general infrastructure of a game server, i.e. defines all the elementary data structures and functions/methods that manipulate them. For testing and demonstration purposes, the package also generates some dummy data that are stored within the internal structures to maintain some initial state.
 * The *handler* package that implements functionality directly related to handling network traffic. Apart from implementing all the endpoints defined in the assignment requirements, the handler package also implements the "/logout" endpoint to round out the use-case of player logging in, poking around, joining or creating matches, and then logging back out again. The WebSocket part of the package implements support for the required 3 message types plus the "match quit" message that ends the player's presence in a match.
 * The *main* package that only provides the entrypoint for launching the server.

The REST API parts of the server's network interface are handlers that simply accept a client's request, look up or generate the necessary information and send the response back to the client.

The WebSocket interface is a little more elaborate in that it persists the client's connection object and uses it to communicate with the client. Also, to allow for multiple clients to communicate with the server using the WebSocket interface, each time a client requests communication over WebSockets, a separate communication thread is spawned and runs until the client sends the "quit" message. At the moment, no limit is imposed on the maximum number of clients allowed.
