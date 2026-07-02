# Detection Rule Authoring Guide

CyberStrikeAI can help validate defensive detections during authorized security
testing. This guide defines a portable documentation format for describing those
detections alongside test evidence.

The example format is not loaded by the application at runtime. Adapt it to the
schema used by your SIEM, EDR, or analytics platform.

## Rule Design

A useful detection rule should explain:

- **Intent:** the observable behavior and defensive outcome.
- **Scope:** supported data sources, platforms, and environments.
- **Logic:** the minimum conditions that produce a match.
- **Exclusions:** known benign behavior and narrowly justified exceptions.
- **Severity:** the operational impact if the behavior is confirmed.
- **Evidence:** fields an analyst needs to validate the signal.
- **Tests:** positive and negative cases with safe reproduction steps.
- **Ownership:** the team responsible for tuning and response.

Avoid rules based only on a tool name, a single brittle string, or an indicator that
cannot be refreshed. Prefer behavior and context that remain meaningful when tooling
changes.

## Portable Template

```yaml
id: unique-rule-id
title: Short analyst-facing title
status: experimental
description: Behavior detected and why it matters
references:
  - https://example.invalid/advisory
log_sources:
  - product: example
    service: audit
detection:
  selection:
    event.action: example-action
  condition: selection
false_positives:
  - Documented administrative workflow
level: medium
tags:
  - attack.tactic
test_cases:
  positive:
    - Safe event fixture that should match
  negative:
    - Benign event fixture that should not match
```

Use stable identifiers and version control so findings can refer to the exact rule
revision that generated them.

## Validation Workflow

1. Confirm the data source is enabled and fields are normalized as expected.
2. Replay a sanitized positive fixture in a non-production environment.
3. Confirm the rule produces the intended signal and preserves useful evidence.
4. Replay at least one nearby benign event and confirm it does not match.
5. Record query latency, result count, and any exclusions introduced during tuning.
6. Have a second reviewer check that exclusions do not create an obvious blind spot.
7. Promote the rule only after its owner and response path are documented.

Testing must stay within an explicitly authorized scope. Use synthetic or sanitized
events whenever possible; do not copy customer data or credentials into fixtures.

## Review Checklist

- [ ] Title and description are understandable without reading the query.
- [ ] Data sources and required fields are listed.
- [ ] Severity reflects impact rather than test urgency.
- [ ] Positive and negative fixtures are included.
- [ ] False positives are specific and reviewable.
- [ ] References use durable, trustworthy sources.
- [ ] Secrets, real identities, and sensitive infrastructure details are absent.
- [ ] Rule owner, review date, and rollback path are known.

## Recording Results

When a rule is validated through CyberStrikeAI, link the resulting conversation or
task identifier to the rule revision and test case. Preserve only the evidence
required for review, apply the project's retention policy, and redact credentials
before exporting results.
