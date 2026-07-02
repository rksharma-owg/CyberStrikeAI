# Usage Guide

This guide covers a minimal local setup for an authorized lab environment. Use
CyberStrikeAI only against systems you own or have explicit permission to test.

## Prerequisites

- Go 1.21 or later
- Python 3.10 or later
- An OpenAI-compatible API endpoint and credential
- A disposable lab target or an explicitly approved testing scope

Install only the security tools needed for your planned workflow. Tool definitions
are available in `tools/`, and unavailable optional tools are skipped at runtime.

## Start the Application

From the repository root, run:

```bash
chmod +x run.sh
./run.sh
```

The launcher checks dependencies, creates the Python virtual environment, builds the
application, and starts it with HTTPS by default. Open
`https://127.0.0.1:8080/` and accept the local self-signed certificate warning.
The port may differ if `server.port` has been changed in `config.yaml`.

For local HTTP instead:

```bash
./run.sh --http
```

Then open `http://127.0.0.1:8080/`.

## Complete First-Run Setup

1. Copy the generated login password from the console, or set `auth.password` in
   `config.yaml` before launch.
2. Add the OpenAI-compatible endpoint, model, and API key under **Settings**.
3. Define the authorized target scope before starting a task.
4. Begin with a low-impact discovery command and review it before approval.

Do not commit real credentials or sensitive assessment results. If configuration is
edited locally, inspect `git diff` before creating a commit.

## Verify and Stop

Confirm that the dashboard loads, authentication succeeds, and the expected tools
appear in the tool list. Stop the foreground server with `Ctrl+C` when the lab
session is complete, then review logs and remove any temporary credentials or data
that should not persist.

## Troubleshooting

- If the browser reports a certificate warning, confirm you are using `https://`
  with the default launcher.
- If logs report that an HTTP request was sent to an HTTPS server, change the client
  URL to `https://` or restart with `./run.sh --http`.
- If a tool is unavailable, install it from its official source and confirm its
  executable is on `PATH`.
- If the configured port is busy, change `server.port` in `config.yaml` and restart.
