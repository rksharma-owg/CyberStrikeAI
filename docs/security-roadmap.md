# Security Architecture and Roadmap

This document describes CyberStrikeAI's high-level security architecture and the
project's current direction. It is guidance rather than a promise that every roadmap
item will ship on a particular date.

## Architecture Overview

CyberStrikeAI separates user interaction, orchestration, tool execution, and
persistence so security-sensitive actions can be reviewed and audited.

```text
Web UI / chatbots / MCP clients
              |
       Server entry points
              |
 Authentication, handlers, and HITL approval
              |
  Agent and multi-agent orchestration
              |
     MCP tools and tool definitions
              |
 Security executors and external utilities
              |
 Databases, knowledge stores, audit logs, and results
```

The main responsibilities are:

- **Entry points and handlers:** `cmd/` and `internal/handler/` accept web, API, and
  MCP traffic and connect requests to application services.
- **Identity and approval controls:** authentication configuration and
  `internal/hitl/` protect access and support human review of sensitive tool calls.
- **Orchestration:** `internal/agent/`, `internal/agents/`, and
  `internal/multiagent/` coordinate model interactions, roles, plans, and tasks.
- **Tool boundary:** `internal/mcp/`, `tools/`, and `skills/` describe available
  capabilities and mediate calls to local or federated tools.
- **Security-sensitive execution:** `internal/security/` and `internal/c2/` contain
  high-impact functionality that should only be used in an explicitly authorized
  scope.
- **Evidence and state:** `internal/database/`, `internal/audit/`,
  `internal/knowledge/`, and `internal/attackchain/` preserve configuration, events,
  findings, and relationships for review.
- **Operator experience:** `web/`, monitoring, and project views expose status and
  results while keeping approval and audit information visible.

## Security Principles

- Require explicit authorization and a defined target scope.
- Prefer human approval for destructive or high-impact operations.
- Keep credentials out of source control, logs, examples, and exported evidence.
- Record enough context to trace tool calls, decisions, and findings.
- Minimize enabled capabilities and installed tools for each engagement.
- Treat model output as untrusted input that requires validation at execution
  boundaries.

## Contribution Workflow

1. Start from the latest `main` and create a branch named for one focused change.
2. Explain the security impact before implementation, including affected trust
   boundaries, permissions, stored data, and failure modes.
3. Keep commits reviewable: separate documentation, behavior, tests, and generated
   files when that improves clarity.
4. Run relevant tests and inspect the final diff for secrets, unsafe defaults, and
   unrelated changes.
5. Open a pull request that describes the motivation, implementation, validation,
   operational impact, and rollback approach.
6. Address review findings with follow-up commits so the discussion remains
   traceable. Squash only when merging if a compact `main` history is preferred.

Changes to authentication, HITL policy, command execution, MCP exposure, C2
functionality, secret handling, or audit retention should receive explicit security
review. Include documentation and tests for both allowed and denied behavior where
applicable.

## Future Roadmap

Roadmap work is grouped by outcome and should be delivered through reviewed,
testable pull requests.

### Short-Term Goals

- Keep README, usage, deployment, and contribution documentation aligned with
  current safe-testing expectations.
- Maintain synthetic examples for alerts, prompts, and API payloads so contributors
  do not need real assessment data.
- Improve validation guidance for prompts, tool calls, and security-sensitive
  workflows before they reach execution boundaries.
- Make documentation changes small, reviewable, and easy to roll back.

### Medium-Term Goals

- Expand dashboard planning for approval queues, sensitive action history, and
  security posture summaries.
- Add stronger validation around API request examples, exported evidence, and
  audit log retention settings.
- Improve integration guidance for MCP clients and external security tools in
  isolated lab environments.
- Document operational recovery paths for misconfiguration, failed tool calls, and
  interrupted workflows.

### Roadmap Review Cadence

Roadmap items should be reviewed when major security workflows, MCP integrations,
or execution boundaries change. During each review, maintainers should confirm:

- The item still supports authorized defensive validation or safe lab use.
- New examples use synthetic data and placeholder credentials.
- Validation, logging, and rollback expectations are clear enough for reviewers.
- The work can be delivered through small pull requests instead of large,
  difficult-to-audit changes.

### Planned Security Modules

- Prompt validation and safe execution review.
- API key and secret handling checks.
- Audit-log integrity and retention review.
- Security Dashboard UI planning.
- Synthetic SIEM alert examples for defensive detection testing.

### Safe Testing Scope

CyberStrikeAI roadmap work should assume explicit authorization, controlled lab
targets, synthetic examples, and least-privilege credentials. Roadmap items should
avoid instructions that enable unauthorized access, persistence, evasion, data
theft, or destructive activity.

### Future AI-Assisted Security Features

- Context-aware prompt validation that separates user intent, evidence, and tool
  execution authority.
- Analyst review summaries that explain why a workflow is safe, risky, or blocked.
- Suggested detection logic based on synthetic examples and documented telemetry.
- False-positive review loops that improve rules without exposing private data.

### Logging, Validation, and Audit Improvements

- Define standard validation outcomes such as allowed, denied, needs review, and
  out of scope.
- Record reviewer decisions and approval context for high-impact workflows.
- Add redaction checks for prompt text, command output, API payloads, and exported
  evidence.
- Preserve correlation identifiers across prompts, agent plans, tool calls, and
  resulting findings.
- Document retention expectations for audit events and generated artifacts.

## Roadmap Timeline

This timeline is planning guidance, not a delivery commitment. Dates should be
added only when maintainers have agreed on scope, ownership, and validation
criteria.

| Stage | Focus | Example outcomes |
| --- | --- | --- |
| Now | Documentation, safe examples, and review workflows | Clear contribution guidance, safe attack-scenario examples, synthetic alert templates, and secure deployment notes |
| Next | Dashboard planning and API validation | Security Dashboard UI requirements, tested API examples, schema validation, and improved prompt-safety review patterns |
| Later | Evidence integrity and operational hardening | Export integrity checks, audit retention controls, stronger secret redaction, and expanded tool health indicators |
| Ongoing | Community and security hygiene | Responsible disclosure handling, dependency updates, issue triage, and reviewer-friendly pull requests |

### Safer Execution

- Expand policy controls for target scope, command risk, and network boundaries.
- Add clearer approval summaries for commands that modify remote systems.
- Improve secret redaction across logs, task output, and exported artifacts.

### Security Dashboard UI

- Present approval queues, active scopes, and recent sensitive actions together.
- Add filters for severity, project, tool, and execution status.
- Surface audit coverage and configuration warnings without exposing credentials.
- Provide accessible empty, loading, and error states for operational clarity.

### Observability and Evidence

- Standardize correlation identifiers across conversations, tasks, and tool calls.
- Make audit retention easier to configure and define redaction plus integrity
  verification requirements for exported evidence.
- Add health indicators for external tools, MCP connections, and background tasks.

### Documentation and Integration

- Maintain tested API request and response examples for common workflows.
- Use the [prompt validation guide](prompt-validation.md) when designing
  suspicious prompt handling for AI-assisted detection pipelines.
- Document secure deployment patterns and backup or recovery procedures.
- Expand contributor checks for documentation links and security-sensitive changes.

## Responsible Disclosure

Do not publish suspected vulnerabilities, credentials, exploit details, or sensitive
assessment data in a public issue.

1. Use the repository's **Report a vulnerability** option when private vulnerability
   reporting is available.
2. Include the affected version or commit, impact, reproduction steps, and suggested
   mitigations. Remove real secrets and use a minimal test environment.
3. Allow maintainers reasonable time to confirm, remediate, and coordinate a release
   before public discussion.
4. If private reporting is unavailable, request a private contact channel from the
   maintainers without posting the vulnerability details.

For ordinary hardening ideas that do not reveal an exploitable weakness, open a
regular issue with a clear scope and expected security benefit.
