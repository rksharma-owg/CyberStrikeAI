# Security Policy

CyberStrikeAI includes security-testing and command-execution capabilities. Use it
only in environments where you have explicit authorization.

## Supported Versions

Security fixes are prioritized for the latest release and the current `main` branch.
Older releases may not receive backported fixes. Confirm whether an issue still
affects the latest code before reporting it.

## Reporting a Vulnerability

Do not disclose suspected vulnerabilities, credentials, exploit details, or
sensitive assessment data in a public issue or discussion.

1. Use GitHub's **Report a vulnerability** option on the repository's Security tab
   when private vulnerability reporting is available.
2. If that option is unavailable, ask the maintainers for a private reporting
   channel without including vulnerability details.
3. Include the affected version or commit, impact, minimal reproduction steps,
   relevant configuration, and suggested mitigations.
4. Remove real secrets and use a controlled test environment.

Maintainers will evaluate the report, coordinate remediation where appropriate, and
communicate disclosure timing with the reporter. Response and release timing depend
on severity, reproducibility, and maintainer availability.

## Scope Guidance

Examples of relevant reports include:

- authentication or authorization bypass;
- unsafe command or tool execution across an approval boundary;
- secret exposure in logs, exports, or API responses;
- cross-project or cross-conversation data access;
- vulnerabilities in C2, WebShell, MCP, or external integration boundaries;
- stored or reflected injection affecting the web interface.

Configuration questions, hardening ideas without an exploitable weakness, and
ordinary bugs can use the public issue tracker.

## Coordinated Disclosure

Allow maintainers reasonable time to validate and remediate a report before public
discussion. Do not access data beyond what is required to demonstrate impact, and do
not disrupt systems or users while testing.
