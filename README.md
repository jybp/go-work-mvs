- Without `go.work`:

```
$ cd svc-a
$ GOWORK=off go list -m all | grep "github.com/jackc/pgx"
github.com/jackc/pgx/v5 v5.0.1
````

```
$ cd svc-b
$ GOWORK=off go list -m all | grep "github.com/jackc/pgx"
github.com/jackc/pgx/v5 v5.1.0
```

- With `go.work`:

```
$ cd svc-a
$ GOWORK=on go list -m all | grep "github.com/jackc/pgx"
github.com/jackc/pgx/v5 v5.1.0
```

```
$ cd svc-b
$ GOWORK=on go list -m all | grep "github.com/jackc/pgx"
github.com/jackc/pgx/v5 v5.1.0
````