# logger

Easy logging

## Examples

```
var log = logger.Namespace("ns=project")

// date="2015-01-01T00:00:00.000000" foo=bar num=5 pct=68.9 arg="test"
log.Log("foo=bar num=%d pct=%0.1f arg=%q", 5, 68.99, "test")

// date="2015-01-01T00:00:00.000000" state=success foo=bar
log.Success("foo=bar")

// date="2015-01-01T00:00:00.000000" at=buyProduct state=error error="invalid token"
log.At("buyProduct").Error(fmt.Errorf("invalid token"))

// date="2015-01-01T00:00:00.000000" foo=bar elapsed=2.398
l := log.Start()
l.Log("foo=bar")
```

## License

Apache 2.0 &copy; 2015 David Dollar
