# How to use Aegis?

After installation you can run the following command to get help and updated instructions:

```bash
aegis --help
```

## Modes

> Flags should come before the length of the password.

1.**Rapid Fire:** Generate a quick password with random mode with default length of 15 characters.


```bash
$ aegis
E#lBVszMJ7
```
2.**Random:** Is the default mode that generates a random password with the length the user provides, which means that every possible character will be considered to generate the password and it's mathematically possible that some group of character won't appear in the password like special characters.


```bash
$ aegis 10
a5uEiG9HNK
```

3.**Bulk**: Generate a N number of passwords with the length the user provides. The default number of passwords is 1.

```bash
$ aegis --bulk 2 10
-%>(2IT!A(
Uec9Z)@zJf
```

4.**Standard**: Generate password with one upper case, one number and one special character at least. Useful for online services that require a password with these characteristics.

```bash
$ aegis --standard 15
lc8lErAp!0LzIa6
```

5.**Numeric**: Generate a password with only numbers.
```bash
$ aegis --numeric 13
3899182977605
```

6.**Online**: Generate a password without using random.org, using only the OS entropy pool to generate the password as random as possible.
```bash
 $ aegis --online 20
 O8@f50*larH!LlDb6BbB
```
