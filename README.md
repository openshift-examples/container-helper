# Container helper

## Init

```bash

go get -u github.com/spf13/cobra/cobra
cobra init . --pkg-name container-helper
cobra add serve

```

## Build

```bash
go install container-helper
# cannot find package "container-helper" in any of:
#        /usr/local/Cellar/go/1.15.5/libexec/src/container-helper (from $GOROOT)
#        /Volumes/Development/Go/src/container-helper (from $GOPATH)

go mod init container-helper

go install container-helper
```