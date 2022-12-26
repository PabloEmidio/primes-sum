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
$ make compile
$ make exec
```

---

## Some examples
```bash
[example@example ~]$ make exec
> Type a number: 50
47 + 3 = 50
43 + 7 = 50
37 + 13 = 50
31 + 19 = 50

In 0h0min0s, it's was found 16 primes Number between 2 and 50
```

```bash
[example@example ~]$ make exec
> Type a number: 100000
...
80147 + 19853 = 100000
80111 + 19889 = 100000
80051 + 19949 = 100000
80039 + 19961 = 100000
80021 + 19979 = 100000

In 0h0min2s, it's was found 8918 primes Number between 2 and 100000
'
> Type a number: 0
Program ended!
```
