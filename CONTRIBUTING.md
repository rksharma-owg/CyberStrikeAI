# Contributing to CyberStrikeAI

Thank you for helping improve CyberStrikeAI. Contributions should make authorized
security work safer, clearer, or more reliable.

## Project Overview

CyberStrikeAI is an AI-assisted security testing and orchestration project for
explicitly authorized environments. The repository combines web interfaces, MCP
integrations, security tooling, audit workflows, and documentation that help
operators plan, execute, and review defensive validation work.

Useful contributions usually improve one of these areas:

- Safer defaults for authentication, approval, command execution, and audit trails.
- Clearer documentation for authorized testing workflows and local setup.
- Better examples that use synthetic data and placeholder credentials.
- Tests or templates that make security-sensitive behavior easier to review.

## Welcome Contribution Types

Maintainers especially welcome small, focused changes that are easy to review and
safe to merge:

- Documentation fixes, setup clarifications, and broken-link repairs.
- Synthetic examples for API payloads, alert templates, and safe lab workflows.
- Tests for validation, authorization, audit, and denied-case behavior.
- Improvements to contributor tooling, issue templates, or review checklists.
- Accessibility and usability improvements for dashboards or documentation.

Large feature proposals, security-sensitive behavior changes, and new tool
integrations should start with an issue or discussion before implementation.

## Before You Start

- Search existing issues and pull requests for related work.
- Use a focused issue or discussion for changes that affect architecture, security
  boundaries, or user-facing behavior.
- Report exploitable vulnerabilities privately according to [SECURITY.md](SECURITY.md).
- Keep real credentials, customer data, and assessment evidence out of the repository.

## Local Setup

Use an isolated development environment and avoid pointing local tests at systems
outside your authorized scope.

1. Clone the repository and switch to a feature branch.
2. Install the Go version declared in `go.mod`.
3. Install any optional security tools only when they are needed for your change.
4. Copy local configuration from documented examples instead of committing secrets.
5. Run relevant tests, linters, or documentation checks before opening a pull
   request.

For setup details and safe example commands, see the [usage guide](docs/usage.md).

## Reporting Issues

Open a GitHub issue for bugs, documentation gaps, safe enhancement ideas, or
reproducible unexpected behavior. A useful issue includes:

- The affected area, such as API examples, MCP integration, dashboard behavior, or
  documentation.
- Expected behavior and actual behavior.
- Safe reproduction steps using placeholder hosts, synthetic data, or a lab target.
- Environment details such as operating system, Go version, and relevant tool
  versions.
- Security impact, if any, without posting secrets or exploit details.

Use the private reporting process in [SECURITY.md](SECURITY.md) for suspected
vulnerabilities.

## Development Workflow

1. Fork the repository and update your fork from the latest `main`.
2. Create a descriptive branch such as `docs/api-timeouts` or
   `fix/scope-validation`.
3. Make one cohesive change with tests or documentation where applicable.
4. Run relevant checks and inspect the full diff for unrelated files and secrets.
5. Open a pull request using a clear title and complete description.
6. Address review feedback with traceable commits. Do not force-push a shared branch
   without coordinating with reviewers.

See the [usage guide](docs/usage.md) for local setup and the
[security roadmap](docs/security-roadmap.md) for architecture and trust boundaries.

## Pull Request Process

Before opening a pull request:

1. Rebase or merge the latest `main` into your branch.
2. Keep the diff focused on one reviewable change.
3. Run relevant checks and note the exact validation steps in the PR description.
4. Inspect the diff for secrets, private targets, generated files, and unrelated
   formatting changes.
5. Link related issues with `Closes #123` only when the PR fully resolves them.

After opening a pull request, respond to review feedback with clear follow-up
commits. Avoid rewriting shared history unless maintainers explicitly ask for it.

## Branch Naming Conventions

Use short, lowercase, hyphenated branch names that describe the change and its
scope. Prefer one of these prefixes:

- `docs/<topic>` for documentation-only improvements.
- `feature/<scope>` for user-visible enhancements or new examples.
- `fix/<scope>` for bug fixes or corrections to existing behavior.
- `security/<scope>` for changes that affect trust boundaries, permissions,
  secrets, disclosure workflows, or defensive controls.
- `chore/<scope>` for repository maintenance that does not change behavior.

Keep branches focused enough that reviewers can understand the risk quickly. Avoid
including real target names, secrets, customer identifiers, or private assessment
details in branch names.

## Pull Request Checklist

- [ ] The change has a clear user or maintainer benefit.
- [ ] Security impact and affected trust boundaries are described.
- [ ] Tests or manual validation steps are included.
- [ ] Documentation and examples match current behavior.
- [ ] No secrets, private targets, or sensitive evidence are present.
- [ ] Generated files and dependency changes are intentional.
- [ ] The PR is focused enough to review and roll back safely.

## Pull Request Workflow Examples

Use the PR description to make the review path obvious. These examples are good
starting points for common contribution types:

### Documentation-only update

- Summary: explain which guide, template, or README section changed.
- Validation: run Markdown formatting checks if available and review links locally.
- Risk: note whether examples are illustrative or tied to runtime behavior.
- Rollback: revert the documentation commit if the guidance becomes inaccurate.

### Security-sensitive change

- Summary: describe the affected trust boundary, permission, or secret-handling path.
- Validation: include tests, manual review steps, and any denied-case behavior.
- Risk: list false-positive, false-negative, privacy, and operational impacts.
- Rollback: identify the prior safe behavior and any configuration changes needed.

### API or sample-data update

- Summary: describe the endpoint, payload, or example being improved.
- Validation: confirm JSON/YAML parses and commands use placeholder values.
- Risk: state that examples are synthetic and safe for authorized lab use.
- Rollback: revert the example or template without affecting application logic.

## Documentation and Examples

Use placeholders for tokens, credentials, hostnames, and identifiers. Label portable
templates that are not consumed by the application at runtime. Commands should be
bounded, minimally privileged, and safe for an authorized lab by default.

## Responsible Use Statement

CyberStrikeAI should only be used where the operator has explicit permission to
test, monitor, or validate the target environment. Contributions must not include
instructions that encourage unauthorized access, credential theft, persistence,
evasion, or destructive activity. Prefer defensive framing, clear authorization
checks, and human review for high-impact workflows.

## Security-Sensitive Changes

Changes involving authentication, HITL approval, command execution, MCP exposure,
C2 functionality, secret handling, audit retention, or data export require explicit
security review. Document both allowed and denied behavior, failure modes, and a
rollback approach.

## Commit Attribution

Use co-author trailers only for people who materially contributed and agreed to be
credited. The GitHub-recognized format is:

```text
Co-authored-by: Full Name <verified-email@example.com>
```

Never add a placeholder or unrelated account to a commit.
