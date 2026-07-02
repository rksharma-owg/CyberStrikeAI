# Safe Attack Simulation Scenarios

These scenarios help defensive teams validate telemetry, detections, approvals, and
response playbooks in an isolated lab. They do not require real malware, credential
collection, persistence, or destructive actions.

## Ransomware Behavior Simulation

### Objective

Confirm that endpoint, file, identity, and network telemetry can identify a
ransomware-like sequence while preserving all original data.

### Prerequisites

- Written authorization and an isolated, disposable lab endpoint.
- A directory containing synthetic files with a tested backup or snapshot.
- Named observers for endpoint telemetry, SIEM alerts, and response decisions.
- A stop condition for unexpected resource use, scope changes, or production traffic.

### Safe Sequence

1. Create a small synthetic dataset in the approved lab directory.
2. Record a baseline of normal file creation, rename, and archive activity.
3. Copy—not encrypt—a limited subset to files with a clearly artificial extension
   such as `.simulation-locked`.
4. Write a harmless ransom-note marker that states it is a simulation and contains
   no payment instructions or external contact details.
5. Generate a synthetic alert or event fixture representing an attempted shadow-copy
   deletion; do not execute the destructive command.
6. Verify whether endpoint and SIEM controls correlate rapid file changes, the marker
   note, and the synthetic destructive-action event.
7. Restore the snapshot, remove simulation artifacts, and confirm file integrity.

### Expected Evidence

- process and file-event timestamps;
- affected synthetic paths and event counts;
- detection rule identifiers and severity;
- approval, escalation, and containment decisions;
- recovery duration and integrity verification.

### Success Criteria

- The scenario stays inside the approved directory and time window.
- Original files remain recoverable and unchanged.
- Analysts can distinguish the simulation from real malicious activity.
- Missed signals and noisy rules become assigned follow-up work.

## Phishing Detection Workflow

### Objective

Validate mail filtering, reporting, identity telemetry, and analyst triage without
collecting real credentials or sending messages outside an approved participant list.

### Safe Sequence

1. Obtain written approval for participants, timing, sender identity, and success
   metrics.
2. Use a reserved domain such as `example.com` in examples and a controlled internal
   domain for the authorized exercise.
3. Send a clearly bounded simulation message with a training landing page.
4. Make the landing page reject and discard all entered values; record only a
   synthetic event identifier needed for the exercise metric.
5. Include an approved reporting path and verify that reported messages retain the
   headers needed for triage.
6. Correlate mail-gateway events, user reports, identity alerts, and synthetic landing
   events by exercise identifier.
7. End the exercise at the approved time, remove landing content, and publish
   aggregate lessons without naming participants.

### Review Questions

- Did authentication and reputation controls label the message as expected?
- Could analysts distinguish the controlled domain and exercise identifier?
- Did the reporting workflow preserve headers and timestamps?
- Were any real secrets or personal values retained?
- Which controls need tuning, documentation, or additional telemetry?

Do not impersonate real vendors, request passwords or MFA codes, use emotional
pressure, or penalize participants. The goal is to improve systems and training, not
to test individuals covertly.
