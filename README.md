# Flogo Extensions

Flogo functions and activities for use in Tibco Cloud Integration Flogo.

# Adding New Functions or Activities

Do not use package names that are already being imported in flogo.json.  Only one will be loaded on Flogo startup for the same name.

Add your module to the go.work file, which may be done by the command:

```
go work use -r [directory]
```

Add the directory to the list of directory values in the dependabot.yml:

```
directory: "/activity/pushover,/function/unixtime,/function/hash,/function/string" # Location of package manifests
```


# Why this must be public

This repository must be public in order to be tracked by dependabot for dependency updates.

# To update dependencies

This repository is a workspace, identified by the go.work file.  This enables dependencies to be managed at the root level.
https://go.dev/ref/mod#workspaces

In the go.work file, add a line similar to:
```
replace example.com/bad/thing v1.4.5 => example.com/good/thing v1.4.5
```

Then run:

```
go work sync
```
To sync dependencies among all the workspace modules.