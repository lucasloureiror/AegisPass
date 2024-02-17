# How to install AegisPass


## Homebrew

[Homebrew](https://brew.sh) is one of the quickest ways to install AegisPass on macOS and Linux and you can stay up to date with the latest version.

```bash
brew install lucasloureiror/tools/aegis
```

## Go builtin install command

If you have Go installed, you can use the following command to install AegisPass:

```bash
go install github.com/lucasloureiror/AegisPass/cmd/aegis@latest
```

> If you don't have $GOPATH/bin in your PATH for your shell, don't forget to add it so you can use AegisPass globally.

This command will show you the path to your GOPATH so you can add it to your PATH in your shell:

```bash
go env GOPATH
```

## Download the binary from the releases page

You can also download the binary from the [releases page](https://www.github.com/lucasloureiror/AegisPass/releases) and add it to your PATH.

## Building the Project

If you want to build AegisPass yourself you can use the makefile with the following steps:

1. Clone the repository:

```bash
git clone https://github.com/lucasloureiror/AegisPass.git
```

2. Change to the project directory:

```bash
cd AegisPass
```

3. Run the appropriate Makefile:

For Linux, Windows and macOS:

```bash
make build
```

This will compile the AegisPass project and create an executable binary in the `build/bin` folder called aegis.exe. Add to your path to use it globally.
