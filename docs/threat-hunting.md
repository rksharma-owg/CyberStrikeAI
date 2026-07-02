# Threat Hunting Workflow

This workflow helps defensive teams use CyberStrikeAI to organize authorized,
hypothesis-driven threat hunts. It favors repeatable evidence and bounded queries
over broad, unsupervised scanning.

## 1. Define the Hunt

Write a falsifiable hypothesis before selecting tools:

> If the behavior is present, then the approved data sources should contain these
> observable events during this time range.

Record:

- business reason and expected defensive value;
- systems, identities, networks, and dates that are in scope;
- data owners and the person authorized to approve the hunt;
- expected observables and nearby benign behavior;
- time, query-cost, and data-retention limits;
- escalation and stop conditions.

Do not expand scope simply because a query returns an interesting adjacent result.
Request approval and record the change first.

## 2. Check Data Readiness

Before querying, confirm:

- clocks and time zones are understood;
- required data sources are available for the full hunt window;
- field names, normalization, and sampling behavior are documented;
- access is read-only unless a response action is separately approved;
- service accounts and API credentials have the minimum necessary permissions;
- sensitive fields have an approved handling and redaction policy.

A missing signal is not evidence that the behavior did not occur when collection is
incomplete.

## 3. Run a Bounded Search

1. Start with the narrowest query that can test the hypothesis.
2. Save the query, parameters, data-source version, and execution time.
3. Review a small result sample before expanding the time range or scope.
4. Compare results with a known-benign baseline.
5. Record refinements and explain why each one improves precision or recall.
6. Stop when the exit criteria are met or evidence quality becomes insufficient.

High-impact actions such as endpoint isolation, credential revocation, or blocking
must follow the organization's response process and should not be embedded in a
hunting query.

## 4. Evaluate Findings

Classify each result as:

- **Expected:** understood activity covered by a documented baseline.
- **Benign anomaly:** unusual but authorized activity worth tuning or monitoring.
- **Suspicious:** requires additional evidence or owner confirmation.
- **Confirmed:** meets the organization's incident criteria.
- **Inconclusive:** data quality or scope prevents a reliable decision.

Preserve both supporting and contradictory evidence. Avoid upgrading severity based
only on an AI-generated summary.

## 5. Record Evidence

For every retained finding, capture:

- hunt identifier and hypothesis revision;
- query and sanitized parameters;
- data source, time range, and collection limitations;
- relevant event identifiers and timestamps;
- analyst decision and supporting rationale;
- follow-up owner, due date, and linked incident or tuning task.

Store evidence in the approved case system. Redact credentials, personal data, and
unrelated customer information before adding excerpts to project documentation.

## 6. Close the Hunt

A hunt is complete when the hypothesis has a documented outcome, findings have
owners, and limitations are explicit. The closeout should identify:

- what was learned;
- detections or telemetry that should change;
- false positives that need safe tuning;
- remaining blind spots;
- whether and when the hunt should be repeated.

Do not treat “no results” as “no risk” without documenting data coverage and query
limitations.
