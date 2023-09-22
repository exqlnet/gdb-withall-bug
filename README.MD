
Follow steps below to reproduce `WithAll()` bug.

1. Run `hack/init.sql` in local/test MySQL server.
2. Configure MySQL connection in `hack/config.yaml`.
3. `go run hack/withall.go`