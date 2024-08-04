# XOR is The Key

The logical XOR (^) operation is a base function for all symmetric-key algorithms. In these algorithms the same key is used for both the encryption of plaintext and the decryption of ciphertext. Let's see if you know how XOR works...

Given these XOR operations between several keys:

```
k1 = {{k1value}}
k2 ^ k1 = {{k2k1value}}
k2 ^ k3 = {{k2k3value}}
flag ^ k1 ^ k3 ^ k2 = {{flagk1k3k2}}
```

What is the value of `flag`?

*Before you XOR these keys, be sure to decode from hex to bytes.*
