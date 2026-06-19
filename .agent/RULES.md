LINEAR Execution State Machine
-e 
Steps:
VERIFY
INSPECT
PLAN
IMPLEMENT
TEST
COMPLETE
-e 
Git Workflow Rules

Branch authority:

main:
- protected human branch
- agents must never push directly
- agents must never merge into main

openhands:
- agent working branch
- agents may commit freely
- agents may push after successful TEST or COMPLETE step
- push is allowed as runtime checkpoint

Recovery rule:
- remote openhands branch is source of recovery after runtime loss
- unpushed work is considered unsafe

Merge rule:
- human reviews openhands branch
- human manually merges into main
- human deletes openhands branch after merge

Agent commands allowed:

Auto checkpoint:
git checkout -B openhands
git add .
git commit -m "<summary>"
git push -u origin openhands

Forbidden:
git push origin main
git merge openhands into main
-e 
Milestone Branch Rules

Purpose:
Milestones are stable recovery checkpoints.

Branch roles:

main:
- current validated project state
- receives only human approved merges
- not used for experiments

openhands:
- AI agent development branch
- auto commit allowed
- auto push allowed
- temporary working state

milestone/*:
- frozen known-good checkpoints
- created manually by humans only
- never used as active development branch
- never auto pushed by agents
- never rewritten

Milestone creation rule:

Create milestone only after:
- tests pass
- architecture is stable
- human approves current main state

Examples:

milestone/run-mode-stable
milestone/agent-governance
milestone/v0.1

Recovery rule:

If main becomes unstable:
- create recovery branch from milestone
- do not modify milestone directly

Authority:

Agent:
openhands only

Human:
main
milestone/*
