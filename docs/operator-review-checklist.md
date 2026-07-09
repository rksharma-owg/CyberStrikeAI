# Operator Review Checklist

Use this checklist before running CyberStrikeAI workflows in an authorized lab or
defensive validation environment. It is intended to help operators confirm scope,
inputs, approvals, and audit readiness before high-impact security actions occur.

## Before Starting

- Confirm written authorization and the approved target scope.
- Use a lab, staging environment, or production-safe validation window.
- Verify that credentials, API keys, and tokens are stored outside the repository.
- Confirm the operator understands which tools may run and which actions require
  human approval.
- Review planned prompts for ambiguity, unsafe targets, or hidden instructions.

## Prompt and Workflow Review

- Check that prompt intent is defensive, authorized, and specific.
- Validate target names, IP ranges, URLs, and account identifiers against the
  approved scope.
- Remove real secrets, customer records, and private assessment evidence from
  examples or shared logs.
- Treat AI-generated plans as proposals until reviewed by an operator.
- Require explicit approval before running tools that can modify systems, revoke
  credentials, isolate hosts, or trigger external notifications.

## Evidence and Audit Readiness

- Enable logging for prompt intake, validation decisions, tool requests, and
  reviewer approvals.
- Record correlation identifiers for alerts, prompts, tool calls, and findings.
- Redact sensitive output before exporting reports or attaching logs to issues.
- Keep rollback or stop conditions visible to the operator.
- Review denied or blocked requests for lessons that improve validation rules.

## Closeout

- Confirm no credentials, private targets, or sensitive evidence were committed.
- Summarize findings with enough context for reviewers to reproduce safely.
- Track false positives and false negatives separately from confirmed findings.
- Update documentation when a workflow requires new approvals, warnings, or
  validation checks.
