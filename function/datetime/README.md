# Datetime Functions
This function package exposes common datetime functions.

## formatUnixTime()
Converts datetime to unix timestamp.
### Input Args

| Arg      | Type     | Description                  |
|:---------|:---------|:-----------------------------|
| datetime | datetime | The datetime object          |

### Output

| Arg       | Type   | Description                                     |
|:----------|:-------|:------------------------------------------------|
| timestamp | string | Unix timestamp of the given datetime object     |

## Dependabot

If dependabot finds any vulnerabilities related to a go package, add this line to the go.mod file:

```
replace {vulnerable-go-pkg} => {vulnerable-go-pkg} v0.0.0
// The version should be the patched version for the found vulnerability
```

And then run:
```
go mod tidy
``` 