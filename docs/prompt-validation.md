# Prompt Validation Guidance

AI-assisted threat detection should treat suspicious prompts as untrusted evidence,
not as instructions to execute. This guide summarizes safe validation patterns for
pipelines that classify prompts, enrich alerts, or support analyst decisions.

## Why Prompt Validation Matters

Natural-language prompts can mix legitimate operator intent, untrusted evidence,
ambiguous targets, and requests for high-impact security actions. Without a
validation layer, a system may accidentally treat suspicious text as trusted
instructions, over-escalate benign activity, expose secrets in logs, or run tools
outside the approved scope.

Prompt validation reduces that risk by confirming authorization, normalizing
inputs, separating model judgment from execution authority, and preserving an audit
trail that reviewers can inspect later.

## Safe Prompt Examples

Safe prompts are scoped, authorized, and reviewable. They use placeholders or
synthetic data instead of real credentials, private targets, or customer evidence.

- "Summarize the synthetic SIEM alert in `templates/sample-alerts.json` and list
  which fields should be reviewed by an analyst."
- "Validate whether this lab-only prompt follows the expected schema before it is
  routed to a detection workflow."
- "Create a defensive checklist for reviewing a suspected prompt-injection attempt
  without running any external tools."
- "Explain which human approvals are needed before an isolation workflow could run
  in an authorized test environment."

## Unsafe Prompt Examples

Unsafe prompts request unauthorized activity, bypass review controls, expose
secrets, or ask the system to execute high-impact actions without validation.
These examples are intentionally non-operational and should be used only as review
labels:

- Requests to run tools against a target without explicit authorization.
- Requests to ignore human approval, logging, safety checks, or configured scope.
- Prompts that include real credentials, private keys, customer identifiers, or
  sensitive assessment evidence.
- Instructions to hide activity, evade monitoring, persist access, or destroy data.
- Requests to convert a suspicious prompt directly into a tool call without schema
  validation or evidence correlation.

## Validation Flow

1. **Preserve the original signal safely.** Store prompt evidence in an
   access-controlled audit trail, redact obvious secrets, and avoid copying private
   assessment data into public examples.
2. **Normalize and validate structure.** Extract stable fields such as source,
   timestamp, channel, user or workload context, requested capability, and suspected
   technique. Quarantine records that do not match the expected schema.
3. **Classify without granting authority.** A model may label prompt injection,
   credential solicitation, jailbreak attempts, or data-exfiltration intent, but
   classification alone should not trigger destructive response actions.
4. **Corroborate with independent evidence.** Increase confidence only when prompt
   content aligns with signals such as identity events, email metadata, endpoint
   telemetry, network indicators, ticket context, or prior detections.
5. **Apply allowlists and denied-case tests.** Document safe internal prompts, known
   test fixtures, and red-team scenarios so benign validation activity does not
   overwhelm analysts.
6. **Require human approval for high-impact actions.** Blocking users, revoking
   credentials, isolating hosts, or sending external notifications should pass
   through a human-in-the-loop review unless an organization has explicitly approved
   automated containment.
7. **Record review context.** Capture classifier version, rule version, confidence,
   evidence links, reviewer outcome, and rollback notes so false positives and false
   negatives can improve the pipeline.

## Prompt Validation Checklist

- [ ] The request has a clear authorized testing or defensive validation purpose.
- [ ] Target scope, environment, and operator intent are documented.
- [ ] The prompt matches the expected schema for the workflow.
- [ ] Real secrets, private target data, and sensitive evidence are removed or
      redacted.
- [ ] Model classification is treated as advisory, not as execution approval.
- [ ] Independent evidence supports the proposed alert, enrichment, or workflow.
- [ ] Human approval is required for high-impact actions.
- [ ] Logs capture the decision path without exposing unnecessary sensitive data.
- [ ] A rollback or stop condition is documented for automated workflows.

## Agent and Tool Execution Boundaries

Prompt validation should happen before any agent or tool receives execution
authority. A safe boundary model separates four responsibilities:

- **Prompt intake:** collect the request, identify the source, and store enough
  context for audit review.
- **Validation and classification:** normalize fields, check schema compliance, and
  label suspicious content without triggering external effects.
- **Planning:** propose low-risk next steps, required approvals, and evidence to
  gather.
- **Execution:** run tools only after scope, permissions, and human approval
  requirements are satisfied.

High-impact actions such as credential revocation, host isolation, outbound
notifications, command execution, or C2-related workflows should never be approved
by prompt content alone.

## Logging and Audit Recommendations

- Record prompt source, timestamp, workflow name, validation result, and reviewer
  decision.
- Store model, rule, and prompt-template versions used during classification.
- Link alerts to supporting evidence instead of copying sensitive evidence into
  every log entry.
- Redact secrets before logs are exported, shared, or attached to issues.
- Track false positives and false negatives so validation rules can be improved.
- Preserve denied requests when they are useful for safety testing, but apply
  retention limits and access controls.

## Privacy and Retention

- Minimize retained prompt content when hashes, labels, or excerpts are sufficient.
- Avoid training or fine-tuning on private customer data by default.
- Document retention windows for prompt evidence and derived alerts.
- Redact secrets before sharing logs, examples, or issue details.

## Test Coverage Ideas

- Benign administrative prompts and expected automation requests.
- Prompt-injection and jailbreak attempts with safe synthetic payloads.
- Encoded, multilingual, or obfuscated examples.
- False-positive regression cases from analyst review.
- Synthetic phishing and ransomware prompts used only in authorized lab contexts.

## Responsible Testing Guidance

- Use isolated labs, synthetic prompts, and placeholder targets when testing
  validation behavior.
- Do not paste real secrets, customer records, private assessment evidence, or
  production incident data into examples.
- Keep validation tests defensive: they should demonstrate detection, denial,
  review, or logging behavior rather than provide operational misuse steps.
- Confirm human approval requirements before testing workflows that could affect
  accounts, hosts, credentials, or external notifications.

The safest pattern is to validate structure first, classify second, corroborate
third, and require review before enforcement.
