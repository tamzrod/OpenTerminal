


# OpenTerminal Security Model

OpenTerminal follows a strict security model to ensure that AI agents operate within defined permissions and do not compromise the system. The permission model categorizes operations into three levels:

## SAFE
SAFE operations are read-only and do not modify any system state. Examples include reading files, listing directories, and viewing logs.

## APPROVAL
APPROVAL operations require human approval before execution. These operations involve making changes to the system that could affect its stability or security. Examples include installing software, modifying configuration files, and restarting services.

## BLOCKED
BLOCKED operations are explicitly denied and cannot be executed by any agent. These operations pose a significant risk to the system's integrity and safety. Examples include deleting critical system files, altering network settings, and executing arbitrary code.

## Human Approval Flow

1. **Request**: An AI agent requests an APPROVAL operation.
2. **Review**: The human operator reviews the request and assesses its potential impact on the system.
3. **Approval/Rejection**: The human operator either approves or rejects the request based on the review.
4. **Execution**: If approved, the operation is executed; if rejected, the agent receives a denial message.

This approval flow ensures that all potentially risky operations are carefully considered and authorized by a human operator before execution.


