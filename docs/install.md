---
title: Install the Flow CLI
description: How to install the Flow command-line interface (CLI)
---

The Flow CLI can be installed on macOS, Windows (7 or greater) and most Linux systems.

## macOS

### Homebrew

```sh
brew install flow-cli
```

### From a pre-built binary

_This installation method only works on x86-64._

This script downloads and installs the appropriate binary for your system:

```sh
sh -ci "$(curl -fsSL https://storage.googleapis.com/flow-cli/install.sh)"
```

To update, simply re-run the installation command above.

It is currently not possible to install earlier versions of the Flow CLI with Homebrew.
## Linux

### From a pre-built binary

_This installation method only works on x86-64._

This script downloads and installs the appropriate binary for your system:

```sh
sh -ci "$(curl -fsSL https://storage.googleapis.com/flow-cli/install.sh)"
```

To update, simply re-run the installation command above.

### Install a specific version

To install a specific version of Flow CLI, append the version tag to the command (e.g. the command below installs CLI version v0.33.3).

```sh
sh -ci "$(curl -fsSL https://storage.googleapis.com/flow-cli/install.sh)" -- v0.33.3
```

## Windows

### From a pre-built binary

_This installation method only works on Windows 10, 8.1, or 7 (SP1, with [PowerShell 3.0](https://www.microsoft.com/en-ca/download/details.aspx?id=34595)), on x86-64._

1. Open PowerShell ([Instructions](https://docs.microsoft.com/en-us/powershell/scripting/install/installing-windows-powershell?view=powershell-7#finding-powershell-in-windows-10-81-80-and-7))
2. In PowerShell, run:

    ```powershell
    iex "& { $(irm 'https://storage.googleapis.com/flow-cli/install.ps1') }"
    ```

To update, simply re-run the installation command above.

# Upgrade the Flow CLI

## macOS

### Homebrew

```sh
brew upgrade flow-cli
```

### From a pre-built binary

_This update method only works on x86-64._

This script downloads and updates the appropriate binary for your system:

```sh
sh -ci "$(curl -fsSL https://storage.googleapis.com/flow-cli/install.sh)"
```

## Linux

### From a pre-built binary

_This update method only works on x86-64._

This script downloads and updates the appropriate binary for your system:

```sh
sh -ci "$(curl -fsSL https://storage.googleapis.com/flow-cli/install.sh)"
```

## Windows

### From a pre-built binary

_This update method only works on Windows 10, 8.1, or 7 (SP1, with [PowerShell 3.0](https://www.microsoft.com/en-ca/download/details.aspx?id=34595)), on x86-64._

1. Open PowerShell ([Instructions](https://docs.microsoft.com/en-us/powershell/scripting/install/installing-windows-powershell?view=powershell-7#finding-powershell-in-windows-10-81-80-and-7))
2. In PowerShell, run:

    ```powershell
    iex "& { $(irm 'https://storage.googleapis.com/flow-cli/install.ps1') }"
    ```
