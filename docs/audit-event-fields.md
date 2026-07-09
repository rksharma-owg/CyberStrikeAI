# Audit Event Field Guide

CyberStrikeAI workflows should produce audit events that are useful for review
without exposing unnecessary secrets, private targets, or sensitive assessment
evidence. This guide describes common fields that documentation, examples, and
future implementation work can reference consistently.

## Core Fields

| Field | Purpose | Example |
| --- | --- | --- |
| `event_id` | Stable identifier for the audit record | `evt_20260709_001` |
| `timestamp` | UTC event time | `2026-07-09T23:30:00Z` |
| `actor` | User, service, or workflow that initiated the action | `analyst@example.org` |
| `workflow` | Human-readable workflow name | `prompt_validation_review` |
| `correlation_id` | Shared identifier across prompt, plan, tool call, and finding | `corr_lab_1234` |
| `action` | Requested or completed operation | `validate_prompt` |
| `decision` | Review outcome | `allowed`, `denied`, `needs_review` |
| `scope_ref` | Reference to approved target scope, not the full sensitive scope text | `lab-scope-2026-q3` |

## Security Review Fields

- `risk_level`: low, medium, high, or critical based on possible operational impact.
- `approval_required`: whether human-in-the-loop approval is required before
  execution.
- `approval_id`: identifier for the approval record when one exists.
- `denial_reason`: short reason when a workflow is blocked or returned for review.
- `redaction_status`: whether secrets, target details, or private evidence were
  removed before storage or export.

## Evidence References

Audit events should link to evidence rather than copying sensitive evidence into
every event. Prefer references such as:

- Alert identifiers from synthetic or internal SIEM data.
- Ticket IDs from approved case-management systems.
- Hashes or object IDs for exported artifacts.
- Links to access-controlled logs or reports.

Avoid embedding raw credentials, private keys, customer records, screenshots with
secrets, or full command output when a reference is sufficient.

## Example Event

```json
{
  "event_id": "evt_20260709_001",
  "timestamp": "2026-07-09T23:30:00Z",
  "actor": "analyst@example.org",
  "workflow": "prompt_validation_review",
  "correlation_id": "corr_lab_1234",
  "action": "validate_prompt",
  "decision": "needs_review",
  "scope_ref": "lab-scope-2026-q3",
  "risk_level": "medium",
  "approval_required": true,
  "redaction_status": "redacted"
}
```

Use synthetic values in public examples and keep implementation-specific schemas
small enough for reviewers to understand and test.
