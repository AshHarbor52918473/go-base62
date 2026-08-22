# Base62 encode/decode

```
base62.go
```
Check the Go test suite next to the implementation for concrete usage examples.

We use this to compress large integer IDs into short Base62 strings for URL parameters and OTP tracking. It helps keep our SMS payloads under the character limit when passing state. You decode them back without losing precision.

The implementation relies entirely on the Go standard library. You do not need to install external packages or run a separate service to get it working.