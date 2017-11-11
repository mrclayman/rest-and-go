# CLIENT

## Generic description

The client component has been implemented to facilitate testing of the server and allow for handling of WebSocket connections that are more persistent in nature than those that only use the server's REST API calls.

As such, it has been written rather quickly and not following meticulous analysis and design. It shares no functionality with the server, which would be more preferable in a real development environment. It also does not support OS signal processing (like Ctrl+C interrupting the execution etc., which would otherwise log the player out of the server before terminating).

## Architecture

The client comprises of two packages:
 * The *client* package that implements the client's respective parts of the functionality.
 * The *main* package that ties the workflow of the client together.

## Usage

Upon start, the client application asks the user to pick an identity to log into the server as. Then, the player is offered several actions in the main menu and in subsequent menus depending on the type of action taken.  Menu options are usually selected using numbers, but in some cases more elaborate input is required.

As the actions are picked, the client communicates synchronously with the server and transforms into readable form the server's responses.
