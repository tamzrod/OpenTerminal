
# OpenTerminal Architecture

OpenTerminal is a local-first AI computer operator engine designed to give AI agents controlled access to various system components while ensuring security and maintainability. The architecture is divided into several key components:

## Daemon
The daemon acts as the central control unit, managing the lifecycle of agents and handling communication between them.

## CLI
The Command Line Interface (CLI) provides a command-line interface for users to interact with OpenTerminal and manage their agents.

## Agent Runtime
The agent runtime is responsible for executing tasks requested by agents. It ensures that each agent runs within its designated permissions and resources.

## Tool Router
The tool router directs requests from agents to the appropriate tools based on the type of operation required (e.g., filesystem access, terminal commands).

## Permission Engine
The permission engine enforces security policies, ensuring that only authorized operations are performed. It categorizes operations into SAFE, APPROVAL, and BLOCKED levels.

## Audit Logs
Audit logs record all actions taken by agents for auditing and monitoring purposes.

## Model Adapters
Model adapters facilitate communication between the OpenTerminal system and different AI models (e.g., Ollama, local models).


