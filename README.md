# OpenTerminal

## Overview
OpenTerminal is a local-first AI computer operator designed for reliable control over terminal, files, scripts, services, containers, and networking. The system prioritizes reliability over complexity, ensuring secure and predictable interactions with system components.

## Status
- **Implemented Components**: CLI interface, Ollama local model support, agent runtime loop, shell execution tool, filesystem tools, permission checking, audit logging
- **Documented Architecture**: Full architecture documentation available in [architecture.md](docs/architecture.md)

## Architecture
OpenTerminal's architecture consists of:
1. **Daemon**: Central control unit managing agent lifecycles
2. **CLI**: Command-line interface for user interaction
3. **Agent Runtime**: Executes tasks with permission constraints
4. **Tool Router**: Directs requests to appropriate system tools
5. **Permission Engine**: Enforces security policies (SAFE/APPROVAL/BLOCKED)
6. **Audit Logs**: Comprehensive action tracking and monitoring

The system follows a validated flow from user intent to execution, with all operations requiring permission checks and logging.

## Development
- Run CLI: `./cmd/openterminal`
- Build daemon: `go build -o openterminal`
- View architecture: `markdown docs/architecture.md`

## License
[License](LICENSE)
