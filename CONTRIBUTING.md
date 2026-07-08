# Contributing to CyberStrikeAI

Thank you for helping improve CyberStrikeAI. Contributions should make authorized
security work safer, clearer, or more reliable.

## Before You Start

- Search existing issues and pull requests for related work.
- Use a focused issue or discussion for changes that affect architecture, security
  boundaries, or user-facing behavior.
- Report exploitable vulnerabilities privately according to [SECURITY.md](SECURITY.md).
- Keep real credentials, customer data, and assessment evidence out of the repository.

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
