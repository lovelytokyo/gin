# Ready

```
mkdir -p $GOPATH/src/github.com/mnuma/
cd $GOPATH/src/github.com/mnuma/
git clone git@github.com:mnuma/goapp-example.git
```

```
brew install direnv
cd $GOPATH/src/github.com/mnuma/goapp-example
direnv allow
```

# Usage

```
go get github.com/Masterminds/glide
go install github.com/Masterminds/glide
```

```
Ex:
glide get github.com/k0kubun/pp
```

```
GO15VENDOREXPERIMENT=1 go run main.go
```

# Using Idea

Preference → Go → Go Libraries → Project Libraries

- add(+) `$GOPATH`
- add(+) `$GOPATH/src/github.com/mnuma/goapp-example/vendor`

Edit configurations... → Environment


| Name| Value|
|-----|------|
|GO15VENDOREXPERIMENT     |1      |
