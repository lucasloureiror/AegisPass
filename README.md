# AegisPass

AegisPass is a CLI tool designed for Software Engineers and DevOps Engineers/Site Reliability Engineers (SRE) to simplify the process of generating high quality random passwords across various platforms and operating systems. AegisPass uses cryptographically secure random numbers along with random.org API to generate random passwords, providing enhanced security for your applications and services with to generate Randomness Mixing.

AegisPass is written in Go and is available for Linux, Windows, and macOS and it does not use anything out of the standard library and random.org, in the principle of no-trust.

## The Importance of True Random Passwords

High quality random passwords provide a higher level of security compared to those generated using ordinary pseudo-random algorithms. Pseudo-random passwords are generated using deterministic algorithms, which means that if an attacker can determine the initial state of the algorithm or identify patterns in the sequence of generated numbers, they can predict the generated passwords. In contrast, high quality random passwords are generated from sources that are very hard to predict - if not impossible with current technology, such as cryptographically secure random numbers, making them more difficult for attackers to predict or guess based on the origin of the password.

Cryptographically secure random numbers are generated from sources like the operating system's entropy pool, which collects unpredictable data from various hardware and software events to produce high-quality random numbers. AegisPass leverages cryptographically secure random numbers along with the random.org website to generate random passwords.

[Random Org](https://www.random.org/) is a reputable source of randomness, generating random numbers from atmospheric noise. Atmospheric noise is a natural source of randomness that is caused by natural electrical discharges, such as lightning, and radio waves from space. 

Random.org captures this atmospheric noise through radio receivers and converts it into random numbers. This process ensures that the generated numbers are as random as possible and not influenced by any easily predictable patterns. By incorporating atmospheric noise-generated random numbers into AegisPass with Cryptographically secure random numbers, the software ensures a higher degree of unpredictability in the resulting passwords, making them more secure and robust against potential attacks.

The process of randomness mixing, as employed by AegisPass, is also used by other systems like [Cloudflare LavaRand](https://blog.cloudflare.com/lavarand-in-production-the-nitty-gritty-technical-details/), which combines the chaotic patterns of lava lamps with other sources of randomness, such as cryptographically secure random numbers, to generate true random numbers. This further enhances the security and unpredictability of the generated passwords by incorporating multiple layers of randomness.

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

- [Go](https://golang.org/doc/install) (version 1.20 or higher)

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

3. Run the appropriate Makefile:

For Linux, Windows and macOS:

```
make build
```

This will compile the AegisPass project and create an executable binary in the `build/bin` folder called aegis.exe.

## Usage

After building the project, you can run the AegisPass CLI tool using the binary created in the `bin` folder.

1. Change to the `bin` directory:

```
cd build/bin
```

2. Run the AegisPass CLI tool:

For Linux or macOS just run the executable with the size of your password as an argument:

```
./aegis.exe 7
```

For Windows:

```
aegis.exe 7
```

In this example a password with size of 7 characters will be generated.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
