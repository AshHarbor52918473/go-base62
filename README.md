# Base62 encode/decode

```
base62.go
```
Short IDs matter when stuffing tracking tokens into SMS OTP payloads. This utility compresses integers into Base62 strings. Check the adjacent Go test files for the exact encoding logic.

It maps integers to Base62 identifiers and decodes them back. No external packages required. The implementation relies strictly on the Go standard library, meaning you avoid extra dependencies entirely.