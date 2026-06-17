

# OpenTerminal Tool System

OpenTerminal uses a modular tool system to provide various functionalities. Each tool has a defined schema, permission level, timeout, and logging mechanism.

## Shell Tool
The shell tool allows agents to execute terminal commands. It supports basic command execution and redirection.

- **Schema**: Defines the command structure.
- **Permission Level**: Determines the required permissions for executing the command.
- **Timeout**: Sets a maximum time limit for command execution.
- **Logging**: Records all executed commands and their outputs.

## Filesystem Tool
The filesystem tool provides access to the local filesystem. It supports file operations such as reading, writing, and deleting files.

- **Schema**: Defines the file operation structure.
- **Permission Level**: Determines the required permissions for accessing the filesystem.
- **Timeout**: Sets a maximum time limit for file operations.
- **Logging**: Records all filesystem activities.

## Docker Tool
The Docker tool allows agents to manage Docker containers. It supports container creation, deletion, and execution.

- **Schema**: Defines the Docker operation structure.
- **Permission Level**: Determines the required permissions for managing Docker containers.
- **Timeout**: Sets a maximum time limit for Docker operations.
- **Logging**: Records all Docker activities.

## Network Tool
The network tool provides access to network functionalities. It supports operations such as making HTTP requests and managing network interfaces.

- **Schema**: Defines the network operation structure.
- **Permission Level**: Determines the required permissions for accessing network resources.
- **Timeout**: Sets a maximum time limit for network operations.
- **Logging**: Records all network activities.

## Service/Systemd Tool
The service/systemd tool allows agents to manage system services. It supports starting, stopping, and restarting services.

- **Schema**: Defines the service operation structure.
- **Permission Level**: Determines the required permissions for managing system services.
- **Timeout**: Sets a maximum time limit for service operations.
- **Logging**: Records all service activities.



