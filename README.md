# Bexio V3 Go Client

Public Go client for the Bexio V3 API. Provides a small HTTP wrapper with sensible defaults (base URL, bearer auth, JSON headers) to build higher-level services.

## Getting started
- Initialize in your project: `go get github.com/meisterwerk/bexio-client`
- Construct a client with your API token and optional overrides:

```go
c := bexio.NewClient("<token>", bexio.WithUserAgent("my-app/1.0"))
req, err := c.NewRequest(ctx, http.MethodGet, "/contacts", nil)
if err != nil { /* handle */ }
res, err := c.Do(req)
```

## Notes
- Defaults to `https://api.bexio.com/3.0`.
- Keeps dependencies to the standard library so you can layer your own models and pagination helpers.
- Extend by adding typed service methods on top of the low-level request helper.
