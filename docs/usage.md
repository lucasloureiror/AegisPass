# How to use Aegis?

After installation you can run the following command to get help and updated instructions: 

```bash
aegis help
```

## Modes

> Flags should come after the length of the password, support for flags before the length of the password is to be implemented.

1.**Rapid Fire:** Generate a quick password with random mode with default length of 10 characters.


```bash
$ aegis
E#lBVszMJ7
```
2.**Random:** Is the default mode that generates a random password with the length the user provides, which means that every possible character will be considered to generate the password and it's mathematically possible that some group of character won't appear in the password like special characters.


```bash
$ aegis 10
a5uEiG9HNK 
```

3.**Standard**: Generate password with one upper case, one number and one special character at least. Useful for online services that require a password with these characteristics.

```bash
$ aegis 15 --standard
lc8lErAp!0LzIa6 
```

4.**Numeric**: Generate a password with only numbers.
```bash
$ aegis 13 --numeric
3899182977605
```

5.**Offline**: Generate a password without using random.org, using only the OS entropy pool to generate the password as random as possible.
```bash
 $ aegis 20 --offline
 O8@f50*larH!LlDb6BbB
```




