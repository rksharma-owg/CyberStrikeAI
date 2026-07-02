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

## Pull Request Checklist

- [ ] The change has a clear user or maintainer benefit.
- [ ] Security impact and affected trust boundaries are described.
- [ ] Tests or manual validation steps are included.
- [ ] Documentation and examples match current behavior.
- [ ] No secrets, private targets, or sensitive evidence are present.
- [ ] Generated files and dependency changes are intentional.
- [ ] The PR is focused enough to review and roll back safely.

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
