## Installing and running

### PATH

```bash
export PATH=$PATH:$(go list -f '{{.Target}}')
go install
hello
```

### GOBIN

```bash
go env -w GOBIN=$HOME/bin
go install
hello
```
