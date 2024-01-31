# Why use AegisPass?

The truth is that SysAdmins, DevOps Engineers, SREs and Software Engineers in general are always looking for ways to improve security and make their lives easier. Sometimes we have to generate passwords using services that we can't guarantee are secure, and that's where AegisPass comes in.

AegisPass is a CLI tool designed for Software Engineers and DevOps Engineers/Site Reliability Engineers (SRE) to simplify the process of generating high-quality random passwords across various platforms and operating systems. AegisPass uses cryptographically secure random numbers along with random.org API to generate random passwords, providing enhanced security for your applications and services with to generate Randomness Mixing. 

So, now you can generate passwords with confidence and security, without leaving the terminal.

## The Importance of True Random Passwords

High quality random passwords provide a higher level of security compared to those generated using ordinary pseudo-random algorithms. Pseudo-random passwords are generated using deterministic algorithms, which means that if an attacker can determine the initial state of the algorithm or identify patterns in the sequence of generated numbers, they can predict the generated passwords. In contrast, high quality random passwords are generated from sources that are very hard to predict - if not impossible with current technology, such as cryptographically secure random numbers, making them more difficult for attackers to predict or guess based on the origin of the password.

Cryptographically secure random numbers are generated from sources like the operating system's entropy pool, which collects unpredictable data from various hardware and software events to produce high-quality random numbers. AegisPass leverages cryptographically secure random numbers along with the random.org website to generate random passwords, but also supports offline password generation.

The process of randomness mixing, as employed by AegisPass, is also used by other systems like [Cloudflare LavaRand](https://blog.cloudflare.com/lavarand-in-production-the-nitty-gritty-technical-details/), which combines the chaotic patterns of lava lamps with other sources of randomness, such as cryptographically secure random numbers, to generate true random numbers. This further enhances the security and unpredictability of the generated passwords by incorporating multiple layers of randomness.

If you want to learn more about random.org, that AegisPass also uses by default to achieve a higher degree of randomness, you can check their [website](https://www.random.org/) and decide for youself if you will utilize or not.


## Why should I trust Aegis?

Short answer: You probably shouldn't trust AegisPass, and that's okay.

Trust is gained, so you can take a look at AegisPass's source code and see that it does not use anything out of the standard library and random.org, in the principle of no-trust. You can also disable the use of random.org with the --offline flag.

- **Open Source:** AegisPass is open source under the MIT License, so you can use it for free and modify it as you wish.

- **Randomness Mixing**: AegisPass uses a technique called Random Mixing to generate passwords, that is a combination of events that current knowledge deems random, something like that is used by [Cloudflare's LavaRand](https://blog.cloudflare.com/randomness-101-lavarand-in-production).

- **No trust:** AegisPass does not use anything out of the standard library in the principle of no-trust and you can disable the use of random.org with the --offline flag.

- **Cryptographically secure:** AegisPass uses cryptographically secure random numbers to generate passwords, so the OS entropy is used to generate the passwords.

## How AegisPass works?

You can take a deeper look at the source code and see for yourself, but the basics are here:

1. **Shuffling the characters:** AegisPass shuffles the set of 69 characters (digits, lowercase letters, and uppercase letters) using cryptographically secure random numbers from Go standard library. Cryptographically secure random numbers are generated from a non-deterministic source, such as the operating system's entropy pool. 

    1.  OS entropy pool:  collects unpredictable data from various hardware events (e.g., keyboard presses, mouse movements) and software events (e.g., disk activity, network traffic) to produce high-quality random numbers. This ensures that the shuffled characters are as random as possible.

2. **Picking the characters**: The characters are picked from the shuffled set of characters to generate the password. There are ways for this to happen:

    1. **Random.org**: Aegis uses random.org to generate random numbers to pick the characters in a already shuffled set of characters. Random.org is a random number service that generates randomness from atmospheric noise. This is another step to ensure that the password is as random as possible. 

    > Note: Random.org is a third-party service so we can't guarantee that it's secure, but with only the numbers generated by random.org, it's impossible to generate the password with current computing power. The numbers are used to pick the characters from the shuffled set of characters.

    2. **Offline mode**: You can disable the use of random.org and Aegis will apply the Fisher-Yates shuffle algorithm to the shuffled set of characters to generate the password. Aegis use the OS entropy pool again to pick the characters and Fisher-Yates algorithm guarantee that all permutations are equally likely.

## Is AegisPass secure?

AegisPass tries to generate password with a thecnique called Random Mixing, that is a combination of events that current knowledge deems random. **However, we can't guarantee that in the future the passwords generated by AegisPass will be secure, but we can guarantee that the passwords generated by AegisPass are as random as possible with the current knowledge.**

