# XOR is The Key

The logical XOR (^) operation is a base function for all symmetric-key algorithms. In these algorithms the same key is used for both the encryption of plaintext and the decryption of ciphertext. Let's see if you know how XOR works...

Given these XOR operations between several keys:

```
k1 = e080e59551eca330dc2ba63345c420acd369e74fd586d638a0d3f990286c10fc
k2 ^ k1 = 4a841f0dec61ba634c09d1ae8d588a616efa1028f03d4bb2c169079fb51db549
k2 ^ k3 = d39c13bb20780ca035b3be1b1d9be1c1f91f83aa505293fe5fb1ccd3916b87e7
flag ^ k1 ^ k3 ^ k2 = 05fe12bdc7658615c055a2378d5e77da74d1b1d3657cee82aac56bc38182d036
```

What is the value of `flag`?

*Before you XOR these keys, be sure to decode from hex to bytes.*
