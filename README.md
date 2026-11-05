# Base62 encode/decode

```
base62.go
```
When generating short IDs for OTP payloads or tracking tokens, predictable lengths matter. Run the Go Base62 test next to the implementation to see how it handles boundary conditions.

Encode integers to short Base62 ids and back. We keep this strictly dependency-free. Adding third-party packages to a core routing module just invites supply chain risk.

The Go Base62 logic relies entirely on the standard library. There are no external services to provision or manage.

## FAQ

**Do I need anything besides `INFRAI_API_KEY`?**  
No, just `go run .` and the key. `base62.go` handles the network layer by wrapping the call in a standard HTTPS request. You avoid SDK version drift entirely. For a base62 example, that covers the full dependency tree.