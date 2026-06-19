# OpenTerminal Execution Rules

OpenTerminal is the mandatory execution layer.

All system commands must use:

/usr/local/bin/openterminal run "<command>"

Do not execute shell commands directly.

Forbidden direct execution:
- git
- rm
- cp
- mv
- bash
- sh

Examples:

Wrong:
git status

Correct:
/usr/local/bin/openterminal run "git status"

Wrong:
rm file.txt

Correct:
/usr/local/bin/openterminal run "rm file.txt"

Reason:
OpenTerminal provides:
- audit trail
- loop protection
- controlled execution
