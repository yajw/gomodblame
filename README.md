# gomodblame

Find the chains which import a specific module.

# Install

```bash
go install github.com/yajw/gomodblame
```

`gomodblame` requires go version `>=1.18`

# Usage

```bash
gomodblame github.com/json-iterator/go@v1.1.12
```

or 
```bash
gomodblame github.com/json-iterator/go
```

an example output:
```
github.com/json-iterator/go@v1.1.12 <- a@v1 <- b@v2
github.com/json-iterator/go@v1.1.12 <- c@v3
```

This output means `github.com/json-iterator/go@v1.1.12` is imported originally by `b@v2` and `c@v3`, and by the separate chain of `b@v2 -> a@v1` and `c@v3`.
