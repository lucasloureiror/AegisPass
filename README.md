# GoTruePass

GoTruePass is a CLI tool that generates true random passwords using the random.org API. With this tool, you can create secure and truly random passwords for your online accounts, providing an extra layer of security.

## Features

- Generates true random passwords
- Uses random.org API for randomness
- Easy-to-use command line interface

## Prerequisites

- [Go](https://golang.org/doc/install) (version 1.16 or higher)

## Building the Project

To build the GoTruePass project, you can use the scripts inside the `/build` folder. Follow these steps:

1. Clone the repository:

```
git clone https://github.com/lucasloureiror/GoTruePass.git
```

2. Change to the project directory:

```
cd GoTruePass
```

3. Run the appropriate build script for your platform:

For Linux or macOS:

```
./build/build-unix.sh
```

For Windows:

```
build\build-windows.cmd
```

This will compile the GoTruePass project and create an executable binary in the `build/bin` folder.

## Usage

After building the project, you can run the GoTruePass CLI tool using the binary created in the `bin` folder.

1. Change to the `bin` directory:

```
cd build/bin
```

2. Run the GoTruePass CLI tool:

For Linux or macOS:

```
./GoTruePass sizeOfYourPassword
```

For Windows:

```
GoTruePass.exe sizeOfYourPassword
```

Follow the prompts to generate your true random password.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
