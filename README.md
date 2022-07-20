 # Monty Hall problem using Go


## Objective
Write a programme which will print all the pairs of prime numbers whose sum equals the number entered by the user.


## What does this program doe?
Take a number by user's input and discover all prime number from 2 to it. After that, print all the pairs of prime numbers whose sum equals the number entered by the user.

`.Eg 10 = 5 + 5, 7 + 3; 12 = 11 + 1, 5 + 7`

For perfomance reasons, it's done a cache of all prime numbers found.
---

## How to test by yourself
```sh
$ git clone https://github.com/PabloEmidio/primes-sum
$ cd primes-sum
$ go build -o prime.out .
```

---

## Some examples
```bash
[example@example ~]$ ./prime.out
> Type a number: 50
47 + 3 = 50
43 + 7 = 50
37 + 13 = 50
31 + 19 = 50
```

```bash
[example@example ~]$ ./prime.out
> Type a number: 100
97 + 3 = 100
97 + 3 = 100
97 + 3 = 100
89 + 11 = 100
83 + 17 = 100
71 + 29 = 100
59 + 41 = 100
53 + 47 = 100
```

```bash
[example@example ~]$ ./prime.out
> Type a number: 0
Program ended!
```
