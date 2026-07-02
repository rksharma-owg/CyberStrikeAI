# Threat Hunt Plan

Copy this file for an authorized hunt. Replace bracketed prompts and remove any
section that is not applicable. Do not include credentials or unredacted sensitive
data.

## Hunt Metadata

- **Hunt ID:** `[stable identifier]`
- **Owner:** `[team or role]`
- **Approver:** `[authorized role]`
- **Planned window:** `[start and end]`
- **Review date:** `[date]`

## Hypothesis

> If `[behavior]` is present in `[approved scope]`, then `[data sources]` should
> contain `[observable events]` during `[time range]`.

## Scope and Limits

- **Included systems/identities:** `[scope]`
- **Excluded systems/identities:** `[explicit exclusions]`
- **Maximum query window/cost:** `[limit]`
- **Allowed actions:** `[read-only queries by default]`
- **Stop conditions:** `[conditions requiring pause or escalation]`

## Data Readiness

| Data source | Owner | Coverage | Key fields | Known limitations |
| --- | --- | --- | --- | --- |
| `[source]` | `[owner]` | `[window]` | `[fields]` | `[gaps]` |

## Query Log

| Time | Query or reference | Reason for change | Result count | Analyst |
| --- | --- | --- | --- | --- |
| `[UTC time]` | `[sanitized query]` | `[initial/refinement]` | `[count]` | `[role]` |

## Findings

| Finding | Classification | Evidence reference | Rationale | Owner |
| --- | --- | --- | --- | --- |
| `[summary]` | `[expected/benign/suspicious/confirmed/inconclusive]` | `[case ID]` | `[reason]` | `[role]` |

## Closeout

- **Hypothesis outcome:** `[supported/rejected/inconclusive]`
- **Coverage limitations:** `[gaps]`
- **Detection changes:** `[linked work]`
- **Telemetry changes:** `[linked work]`
- **Incident or escalation:** `[reference or none]`
- **Next hunt date:** `[date or not planned]`
- **Evidence retention location:** `[approved system and policy]`
