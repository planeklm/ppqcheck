# ppqcheck

a cli tool for checking if a provisioning profile (.mobileprovision) has PPQ enabled.

[![Go Reference](https://pkg.go.dev/badge/github.com/planeklm/ppqcheck.svg)](https://pkg.go.dev/github.com/planeklm/ppqcheck)

# installation
install [go](https://go.dev)

```
go install github.com/planeklm/ppqcheck@latest
```

# usage
you can get usage info with ppqcheck -h.

```
ppqcheck -h
Usage of ppqcheck:
  -path string
        Path to the .mobileprovision
```

# example
```
ppqcheck /path/to/file.mobileprovision
ppqcheck /path/to/file1.mobileprovision /path/to/file2.mobileprovision
ppqcheck -path /path/to/file.mobileprovision
```
