# YesWeCode

The YesWeCode repository is a set of coding and CTF challenges made for my internal company competitions. For each challenge, the goal is to find the result, called the *flag*, given a problem to solve.

While most of them are original, some were inspired by these famous and great websites: [CodinGame](https://www.codingame.com), [CryptoHack](https://cryptohack.org/) and [RootMe](https://www.root-me.org/). 

Challenges are dispatched into directories, representing a level of difficulty: the higher the number, the harder the challenge. The level is figured out following the understanding of the statement and the time it takes to solve. It's quite subjective, because it depends a lot on the knownledge and experience of people. Feel free to do it and use it ;).

All flags are encoded using Base64, in order not to spoil you, in case you want to do it, instead of just using it. To test your solution aginst a flag, just use the following command:

```bash
cat flag.b64 | base64 -d | diff your_solution.txt -
```

## Make it Work for CTFd

The [CTFd](https://ctfd.io/) platform was used to host and manage challenges. But, even if they are written in the Markdown language, some of them contain Latex formulas that can't be handled by the CTFd platform. So, in order to make them work, you can use the tool [Pandoc](https://pandoc.org/), to convert Latex formulas into a compatible format. For example, if your Markdown file is `challenge.md`, you can convert using the command:

```bash
pandoc --toc --standalone --webtex -f markdown -t markdown challenge.md -o challenge_final.md
```

then copy/paste the text of the final file directly into the challenge description.
