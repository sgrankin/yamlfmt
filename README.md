# yamlfmt

Barebones comment-preserving YAML formatting via https://github.com/go-yaml/yaml.

## Install:
```
; go install sgrankin.dev/yamlfmt@latest
```

## Use:
```
; yamlfmt  # Read stdin.
; yamlfmt *.yaml  # Read files and output them as separate docs.
; yamlfmt -indent=3  # Set the indent.
```
