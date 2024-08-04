# Eve is Listening

You are a collegue of Alice and Bob. You know that they are working on a new wonderful project together, but without you... However, you need to be involved in this project to have a change to get promoted in the future.

Fortunatly, you were able to intercept a message between them, by listening their channel of communication. In this chat, all messages are 128-bytes long, written in English (ASCII charset).

Because you have some hacking abilities, you even get information about the company's chat communication protocol, based on the *Diffie-Hellman key exchange*. The protocol works as follow:
- The company uses the generator $g$ and the number $p$. These values are public.
- Alice picks up a secret number $a$ and Bob picks up a secret number $b$ on his side.
- Alice computes $\mathcal{A}=g^a\ mod\ p$ then sends it to Bob.
- Bob computes $\mathcal{B}=g^b\ mod\ p$ then sends it to Alice.

Now, each can get the secret shared key $k$ as follow:
- On Alice's side, $k=\mathcal{B}^a\ mod\ p$.
- On Bob's side, $k=\mathcal{A}^b\ mod\ p$.

Finally, a sent message $m$ is ciphered as $m^{\prime}$ such as $m^{\prime}=m \oplus k$, where $m$ is zero padded first, if its length is lesser than 128 bytes, then converted to a little-endian integer.

You are given:
- $g=$ 20713387470590905255573036258836721934728492752390521374064507177089990210962343331928418955536103194720235532725338679759288301525140653512087629238835871170633543949795607865540123535064826464925528784090341658306794839708039239341640351981655246845864770402104880102447526054707525127448850020961623527.
- $p=$ 106528951761249025729412125479197260910308638225544451426813760411773819654979331756107858688322178730446171344806416830002019734743798381012666677175332885430568316533798811252472855340838402509111994536576627148671845860618445807934056330241652934528282514178025398366887626499360801730469435657805629799361.
- And the ciphered message you just intercepted in the input file.

Are you able to find the original message?
