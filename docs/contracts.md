



# OpenTerminal Internal Contracts

This document outlines the internal contracts used within OpenTerminal, ensuring that all components communicate and operate correctly.

## Agent Request Format
```json
{
  "agent_id": "string",
  "request_type": "string",
  "payload": {
    // request-specific data
  }
}
```

## Tool Call Format
```json
{
  "tool_name": "string",
  "parameters": {
    // tool-specific parameters
  }
}
```

## Tool Response Format
```json
{
  "status": "success" | "failure",
  "data": {
    // response data
  },
  "error_message": "string"
}
```

## Permission Decision Format
```json
{
  "decision": "allow" | "deny" | "approve",
  "reason": "string"
}
```

## Audit Event Format
```json
{
  "event_type": "string",
  "timestamp": "ISO8601 timestamp",
  "agent_id": "string",
  "action": "string",
  "result": "success" | "failure",
  "details": {
    // additional details
  }
}
```

## Error Handling Format
```json
{
  "error_code": "string",
  "message": "string",
  "details": {
    // additional error details
  }
}
```

### Flow

1. **User Intent**: The user expresses a desire to perform an action.
2. **Agent**: The Operator Agent interprets the user intent and formulates an agent request.
3. **Tool Request**: The Operator Agent sends a tool call request to the Tool Router.
4. **Permission Engine**: The Permission Engine evaluates the tool request and decides whether to allow, deny, or require approval.
5. **Runtime Executor**: If allowed, the Runtime Executor executes the requested operation.
6. **Result**: The result of the operation is returned to the Operator Agent.
7. **Agent**: The Operator Agent processes the result and communicates it back to the user.

### Rules

- Every operation must be traceable.
- Every operation must be logged.
- Every operation must have its permissions checked.
- Every operation must have a timeout controlled.



