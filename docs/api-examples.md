# API Request and Response Examples

These examples use the API routes and schemas defined by the current application.
Run them only against a CyberStrikeAI instance you control.

## Conventions

The snippets assume the server was started locally with `./run.sh --http`:

```bash
export BASE_URL="http://127.0.0.1:8080"
```

For HTTPS, set `BASE_URL` to the configured HTTPS origin and use a trusted
certificate or pass the appropriate CA certificate to `curl` with `--cacert`.
Do not put real passwords or bearer tokens in committed scripts, screenshots, or
shell history.

Example values such as `local-password`, `TOKEN`, and UUIDs below are placeholders.

## Log In

Request:

```http
POST /api/auth/login HTTP/1.1
Host: 127.0.0.1:8080
Content-Type: application/json

{
  "password": "local-password"
}
```

Equivalent `curl` command:

```bash
curl --fail-with-body \
  --request POST \
  --header "Content-Type: application/json" \
  --data '{"password":"local-password"}' \
  "${BASE_URL}/api/auth/login"
```

Example response:

```json
{
  "token": "TOKEN",
  "expires_at": "2030-01-01T12:00:00Z",
  "session_duration_hr": 24
}
```

Store the returned token in a temporary shell variable for subsequent examples:

```bash
export CYBERSTRIKE_TOKEN="TOKEN"
```

## Validate a Token

```bash
curl --fail-with-body \
  --header "Authorization: Bearer ${CYBERSTRIKE_TOKEN}" \
  "${BASE_URL}/api/auth/validate"
```

Example response:

```json
{
  "token": "TOKEN",
  "expires_at": "2030-01-01T12:00:00Z"
}
```

## Create a Conversation

Use a descriptive title that does not contain credentials or sensitive target data:

```bash
curl --fail-with-body \
  --request POST \
  --header "Authorization: Bearer ${CYBERSTRIKE_TOKEN}" \
  --header "Content-Type: application/json" \
  --data '{"title":"Authorized lab validation"}' \
  "${BASE_URL}/api/conversations"
```

Example response:

```json
{
  "id": "550e8400-e29b-41d4-a716-446655440000",
  "title": "Authorized lab validation",
  "createdAt": "2030-01-01T10:00:00Z",
  "updatedAt": "2030-01-01T10:00:00Z",
  "projectId": ""
}
```

An optional `projectId` may be supplied when the conversation should use a project's
shared facts:

```json
{
  "title": "Authorized lab validation",
  "projectId": "PROJECT_ID"
}
```

## List Conversations

```bash
curl --fail-with-body \
  --get \
  --header "Authorization: Bearer ${CYBERSTRIKE_TOKEN}" \
  --data-urlencode "limit=20" \
  --data-urlencode "offset=0" \
  "${BASE_URL}/api/conversations"
```

Example response shape:

```json
[
  {
    "id": "550e8400-e29b-41d4-a716-446655440000",
    "title": "Authorized lab validation",
    "createdAt": "2030-01-01T10:00:00Z",
    "updatedAt": "2030-01-01T10:00:00Z"
  }
]
```

## Download the OpenAPI Specification

The authenticated specification is the source of truth for additional routes and
schemas:

```bash
curl --fail-with-body \
  --header "Authorization: Bearer ${CYBERSTRIKE_TOKEN}" \
  "${BASE_URL}/api/openapi/spec"
```

The interactive documentation page is available at `${BASE_URL}/api-docs`.

## Handle Errors Safely

An unsuccessful request may return an error object:

```json
{
  "error": "unauthorized"
}
```

Treat response bodies as potentially sensitive. Avoid verbose tracing in shared
terminals, redact tokens before attaching output to an issue, and unset temporary
credentials when finished:

```bash
unset CYBERSTRIKE_TOKEN
```
