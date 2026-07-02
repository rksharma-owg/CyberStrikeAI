# Secure Deployment Guide

CyberStrikeAI can invoke security tools and process sensitive assessment data. Deploy
it only for authorized users and targets, with isolation appropriate to that risk.

## Docker Security Notes

- Build from a reviewed commit and pin base images by immutable digest in controlled
  environments.
- Run as a non-root user and use a read-only root filesystem where supported.
- Drop Linux capabilities by default; add only capabilities required by an approved
  tool and document why.
- Do not mount the Docker socket, host root filesystem, SSH directories, or cloud
  credential directories into the container.
- Place the service on a dedicated network and restrict outbound traffic to approved
  model endpoints, registries, and authorized test targets.
- Publish only the required application port. Terminate TLS with a trusted
  certificate for shared deployments.
- Store databases, audit logs, and uploaded artifacts on access-controlled volumes
  with an explicit backup and retention policy.
- Set CPU, memory, process, and storage limits to contain runaway tools.
- Scan images and dependencies before promotion, then rebuild rather than patching a
  running container.
- Separate production, test, and offensive-security lab deployments.

## Runtime Checklist

- [ ] Authentication uses a strong, unique password.
- [ ] TLS and network boundaries match the deployment model.
- [ ] HITL approval is enabled for sensitive operations.
- [ ] Tool allowlists contain only reviewed capabilities.
- [ ] Logging avoids credentials and unnecessary target data.
- [ ] Backup restoration and incident shutdown are tested.
- [ ] Operators know how to revoke access and rotate secrets.

## API Key Protection

- Provision a dedicated key with the narrowest model and account permissions.
- Inject secrets at runtime from an approved secret manager or protected environment;
  never bake them into images or commit them to `config.yaml`.
- Restrict secret access to the service identity and authorized administrators.
- Prevent keys from appearing in command arguments, health endpoints, traces,
  screenshots, exported conversations, and support bundles.
- Rotate keys on a defined schedule and immediately after suspected exposure.
- Monitor provider usage for unexpected models, volume, regions, or source addresses.
- Test revocation and replacement without requiring an image rebuild.

Before publishing logs or examples, search for provider key prefixes, authorization
headers, session tokens, and private endpoint URLs. Treat redaction as a final safety
check rather than the primary secret-control mechanism.
