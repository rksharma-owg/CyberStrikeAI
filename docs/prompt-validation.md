# Prompt Validation Guidance

AI-assisted threat detection should treat suspicious prompts as untrusted evidence,
not as instructions to execute. This guide summarizes safe validation patterns for
pipelines that classify prompts, enrich alerts, or support analyst decisions.

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

The safest pattern is to validate structure first, classify second, corroborate
third, and require review before enforcement.
