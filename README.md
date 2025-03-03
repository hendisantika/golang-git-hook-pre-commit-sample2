# GoLang Git Hook Pre Commit Sample

In Go (Golang) projects, using a pre-commit Git hook can help ensure that code quality standards are maintained before
any changes are committed to the repository. A pre-commit hook is a script that runs automatically before each commit,
allowing you to perform checks like linting, formatting, testing, etc.

Here’s how you can set up a pre-commit Git hook for a Go project:

### Steps to Set Up a Pre-Commit Hook in Go Projects:

Steps to Use pre-commit Framework:

1. Install pre-commit :

   Install the pre-commit tool globally on your system.

```shell
pip install pre-commit
```

2. Create a .pre-commit-config.yaml File :

   In the root of your Go project, create a .pre-commit-config.yaml file to define the hooks you want to run.

```yaml
repos:
   - repo: https://github.com/pre-commit/pre-commit-hooks
     rev: v4.3.0
     hooks:
        - id: trailing-whitespace
        - id: end-of-file-fixer
        - id: check-yaml
        - id: check-added-large-files

   - repo: https://github.com/dnephin/pre-commit-golang
     rev: v0.5.0
     hooks:
        - id: go-fmt
        - id: go-imports
        - id: no-go-testing
        - id: golangci-lint
        - id: go-unit-tests

   - repo: https://github.com/alessandrojcm/commitlint-pre-commit-hook
     rev: v8.0.0
     hooks:
        - id: commitlint
          stages: [ commit-msg ]
          additional_dependencies: [ '@commitlint/config-conventional' ]

```

3. Install the Hooks :
   After creating the .pre-commit-config.yaml file, install the hooks by running:

```shell
pre-commit install
```

or using brew

```shell
brew pre-commit install
```

4. Run the Hooks Manually (Optional) :
   You can manually run the hooks on all files using:

```shell
pre-commit run --all-files
```

### Conclusion

Using a pre-commit Git hook in your Go project ensures that basic checks like formatting, linting and testing are
performed before any code is committed. You can either write a custom shell script or use the pre-commit framework for
a more structured approach. Both methods help maintain code quality and prevent issues from being introduced into the
codebase.
