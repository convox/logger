# logger

Easy logging in Go

## Examples

```go
var log = logger.New("ns=project")

// date="2015-01-01T00:00:00.000000" ns=project foo=bar num=5 pct=68.9 arg="test"
log.Log("foo=bar num=%d pct=%0.1f arg=%q", 5, 68.99, "test")

// date="2015-01-01T00:00:00.000000" ns=project sub=worker state=success foo=bar
log.Namespace("sub=worker").Success("foo=bar")

// date="2015-01-01T00:00:00.000000" ns=project at=buyProduct state=error error="invalid token"
log.At("buyProduct").Error(fmt.Errorf("invalid token"))

// date="2015-01-01T00:00:00.000000" ns=project foo=bar elapsed=2.398
l := log.Start()
l.Log("foo=bar")
```

## License

Apache 2.0 &copy; 2015 David Dollar
