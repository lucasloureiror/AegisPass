# AegisPass

AegisPass is a CLI tool designed for Software Engineers, DevOps Engineers, and Site Reliability Engineers (SRE) to simplify the process of generating true random passwords across various platforms and operating systems. Using the random.org API, AegisPass enhances the security of your applications and services by creating truly random passwords.

## The Importance of True Random Passwords

True random passwords provide a higher level of security compared to those generated using pseudo-random algorithms. This is because true random passwords are generated from a non-deterministic source, making them more difficult for attackers to predict or guess.

[random.org](https://www.random.org/) is a reputable source of true randomness, generating random numbers from atmospheric noise. This process ensures that the generated passwords are truly random and not influenced by any predictable patterns.

## Features

- Tailored for Software Engineers, DevOps Engineers, and SREs
- Generates true random passwords
- Uses random.org API for randomness
- Supports various platforms and operating systems
- Easy-to-use command line interface

## Prerequisites

- [Go](https://golang.org/doc/install) (version 1.16 or higher)

## Building the Project

To build the AegisPass project, you can use the scripts inside the `/build` folder. Follow these steps:

1. Clone the repository:

```
git clone https://github.com/lucasloureiror/AegisPass.git
```

2. Change to the project directory:

```
cd AegisPass
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

This will compile the AegisPass project and create an executable binary in the `build/bin` folder.

## Usage

After building the project, you can run the AegisPass CLI tool using the binary created in the `bin` folder.

1. Change to the `bin` directory:

```
cd build/bin
```

2. Run the AegisPass CLI tool:

For Linux or macOS:

```
./aegis sizeOfYourPassword
```

For Windows:

```
aegis.exe sizeOfYourPassword
```

Follow the prompts to generate your true random password.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
