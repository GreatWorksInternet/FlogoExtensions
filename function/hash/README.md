# Hash Functions
This function package exposes common hashing functions.

## b64Sha256()
Converts input to a Base 64 encoded Sha256 hash.
### Input Args

| Arg      | Type     | Description                  |
|:---------|:---------|:-----------------------------|
| body     | string   | The body being hashed        |

### Output

| Arg       | Type   | Description                                     |
|:----------|:-------|:------------------------------------------------|
| hash      | string | The body Sha256 hashed and base 64 encoded      |

## hmacB64Sha256()
Converts input to a Base 64 encoded Sha256 HMAC hash.
### Input Args

| Arg      | Type     | Description                     |
|:---------|:---------|:--------------------------------|
| body     | string   | The body being hashed           |
|:---------|:---------|:--------------------------------|
| key      | string   | The private key used in hashing |        |

### Output

| Arg       | Type   | Description                                                         |
|:----------|:-------|:--------------------------------------------------------------------|
| hash      | string | The body Sha256 HMAC hashed and base 64 encoded using the input key |

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