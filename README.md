# YesWeCode

The YesWeCode repository is a set of coding and CTF challenges made for my internal company competitions. For each challenge, the goal is to find the result, called the *flag*, given a problem to solve.

While most of them are original, some were inspired by these famous and great websites: [CodinGame](https://www.codingame.com), [CryptoHack](https://cryptohack.org/) and [RootMe](https://www.root-me.org/). 

Challenges are dispatched into directories, representing a level of difficulty: the higher the number, the harder the challenge. The level is given following the understanding of the statement and the time it takes to solve. It's quite subjective, because it depends a lot on the knownledge and experience of people. Feel free to do it and use it ;).

All the flags are encoded using Base64, in order not to spoil you, in case you want to do it, instead of just using it. To test your solution aginst a flag, just use the following command:

```bash
cat flag.b64 | base64 -d | diff your_solution.txt -
```
