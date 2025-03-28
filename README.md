# az-secret2env

**az-secret2env** is a command-line tool designed to streamline the process of executing applications with environment
variables temporarily populated by secrets stored in Azure Key Vault. This tool allows you to reference secrets in your
environment variables and seamlessly replace those references with the actual secret values when starting a new process.
The original environment variables remain unchanged, ensuring that the secret values are only exposed during the
execution of the desired process.

## Key Features

- **Secure Execution**: Inject secrets into your environment variables at runtime, minimizing exposure.
- **Azure Integration**: Directly fetch secrets from Azure Key Vault using existing environment variable references.
- **Process Isolation**: Ensure that secret values are only available for the duration of the executed process.
- **Simple and Efficient**: A straightforward command-line interface that integrates easily into your existing
  workflows.

## Use Cases

- **Securely launching applications**: Run applications that require sensitive configuration without permanently
  altering your environment variables.
- **Temporary secret access**: Provide short-lived access to secrets, ideal for CI/CD pipelines or secure script
  execution.
- **Environment-specific configurations**: Dynamically inject environment-specific secrets at runtime.

## Installation with Homebrew on macOS

```bash
# First, install the wim-vdw tap, a repository of all my Homebrew packages.
brew tap wim-vdw/tap

# Then, install az-secret2env.
brew install wim-vdw/tap/az-secret2env

# To update to the latest version of az-secret2env, first update Homebrew.
brew uppdate

# Then, upgrade az-secret2env.
brew upgrade wim-vdw/tap/az-secret2env
```

## Installation on Linux/Windows

Download one of the pre-built releases from the [releases page](https://github.com/wim-vdw/az-secret2env/releases).
