# ppqcheck

a cli tool for checking if a provisioning profile (.mobileprovision) has PPQ enabled.

[![Go Reference](https://pkg.go.dev/badge/github.com/planeklm/ppqcheck.svg)](https://pkg.go.dev/github.com/planeklm/ppqcheck)

# installation
install [go](https://go.dev)

```bash
$ go get github.com/planeklm/ppqcheck
```

# usage
you can get usage info with `ppqcheck -h`.

```
$ ppqcheck -h
Usage of ppqcheck:
  -path string
        Path of the .mobileprovision (required)
```

# example
```
$ ppqcheck -path /path/to/file.mobileprovision
```
