# Flogo Extensions

Flogo functions and activities for use in Tibco Cloud Integration Flogo.

# Dependencies:
```
npm install go 
go get github.com/project-flogo/core/activity
go get github.com/project-flogo/core/data/coerce
go get github.com/project-flogo/core/data/metadata
go get github.com/joho/godotenv
```

# To build:
```
go build
```
# To Test:
```
go test
```

# Why this must be public

This repository must be public in order to be tracked by dependabot for dependency updates.

# To update dependencies

In the go.work file, add a line similar to:
```
replace example.com/bad/thing v1.4.5 => example.com/good/thing v1.4.5
```