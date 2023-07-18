# String Functions
This function package exposes common string functions.

## replaceAllMultiple()
Returns a copy of the input string with all replacements performed from a list of old, new string pairs.
### Input Args

| Arg      | Type     | Description                  |
|:---------|:---------|:-----------------------------|
| str           | string    | The original string to perform replace on         |
| replaceList   | array     | The old/new string pairs that will replace str    |

### Output

| Arg       | Type   | Description                                     |
|:----------|:-------|:------------------------------------------------|
| newStr    | string | Copy of str with all replacements performed     |

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