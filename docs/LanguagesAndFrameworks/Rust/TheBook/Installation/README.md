# Installation

Rust can be downloaded through `rustup`, which is a command line tool for managing Rust versions and associated tools.

## Installing rustup on Linux or macOS

On Linux or macOS, open a terminal and enter the following command:

```bash
$ curl --proto '=https' --tlsv1.2 https://sh.rustup.rs -sSf | sh
```

The command downloads a script and starts the installation of the `rustup` tool, which installs the latest stable version of Rust. 

## Updating and Uninstalling

After installing Rust via `rustup`, updating to the latest version is easy.

```bash
$ rustup update
```

And to uninstall Rust and `rustup`, run the following uninstall script from you shell:

```bash
$ rustup self uninstall
```