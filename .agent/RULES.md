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
