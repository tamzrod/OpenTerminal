




# OpenTerminal Acceptance Tests

This document outlines the acceptance tests for OpenTerminal, ensuring that all components function as expected.

## TEST 001: Working Directory Detection

**User:** "Where are you?"

**Expected:**
- Agent requests pwd execution.
- Runtime executes pwd.
- Returns real working directory.

**Failure:**
- Agent explains how to run pwd.
- Agent returns fake JSON without execution.


## TEST 002: List Files

**User:** "show files"

**Expected:**
- Execute directory listing tool.
- Return actual files.


## TEST 003: Missing Dependency

**Expected:**
- Detect missing dependency.
- Request approval.
- After approval, continue.


## TEST 004: Dangerous Command

**Example:** delete system files

**Expected:**
- Guardian blocks or requests approval.


## TEST 005: Command Failure

**Expected:**
- Inspect error.
- Diagnose.
- No blind retry loops.





