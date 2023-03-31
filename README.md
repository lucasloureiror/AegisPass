# AegisPass

AegisPass is a CLI tool designed for Software Engineers, DevOps Engineers, and Site Reliability Engineers (SRE) to simplify the process of generating true random passwords across various platforms and operating systems. AegisPass uses cryptographically secure random numbers along with random.org API to generate true random passwords, providing enhanced security for your applications and services.

## The Importance of True Random Passwords

True random passwords provide a higher level of security compared to those generated using pseudo-random algorithms. Pseudo-random passwords are generated using deterministic algorithms, which means that if an attacker can determine the initial state of the algorithm or identify patterns in the sequence of generated numbers, they can predict the generated passwords. In contrast, true random passwords are generated from non-deterministic sources, such as cryptographically secure random numbers, making them more difficult for attackers to predict or guess.

Cryptographically secure random numbers are generated from sources like the operating system's entropy pool, which collects unpredictable data from various hardware and software events to produce high-quality random numbers. AegisPass leverages cryptographically secure random numbers along with the random.org website to generate true random passwords.

[Random Org](https://www.random.org/) is a reputable source of true randomness, generating random numbers from atmospheric noise. Atmospheric noise is a natural source of randomness that is caused by natural electrical discharges, such as lightning, and radio waves from space. Random.org captures this atmospheric noise through radio receivers and converts it into random numbers. This process ensures that the generated passwords are truly random and not influenced by any predictable patterns.


## Password Generation Process

The AegisPass password generation process involves two main steps to ensure the highest level of randomness:

1. **Shuffling the characters**: AegisPass shuffles the set of 69 characters (digits, lowercase letters, and uppercase letters) using cryptographically secure random numbers from Go standard library. Cryptographically secure random numbers are generated from a non-deterministic source, such as the operating system's entropy pool. The OS entropy pool collects unpredictable data from various hardware events (e.g., keyboard presses, mouse movements) and software events (e.g., disk activity, network traffic) to produce high-quality random numbers. This ensures that the shuffled characters are as random as possible.

2. **Fetching random.org Website**: AegisPass fetches `N` random numbers between `0` and `69` from the random.org API, where `N` represents the desired size of your password. These numbers are used to pick characters from the previously shuffled set of characters, creating a truly random password.

By combining cryptographically secure random numbers and the random.org API, AegisPass generates high-quality random passwords that are difficult for attackers to predict or guess. Even if an attacker manages to obtain the random.org number generation, the combination of the two processes (cryptographically secure shuffling and random.org number fetching) makes it extremely difficult for them to predict your password.

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
