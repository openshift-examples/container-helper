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

## Resources

  * https://carolynvanslyck.com/blog/2020/08/sting-of-the-viper/
    https://github.com/carolynvs/stingoftheviper
  * https://blog.container-solutions.com/golang-configuration-in-12-factor-applications
  * https://github.com/openshift/origin/blob/master/examples/hello-openshift/hello_openshift.go