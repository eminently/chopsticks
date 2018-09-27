<img src="https://preview.ibb.co/dVzECU/chopsticks_name.png" width="417x863">

If you are afraid of Bitcoin Cash [forks](https://en.wikipedia.org/wiki/List_of_bitcoin_forks), use [chopsticks](https://api.chopsticks.cash) !


## Context

Sadly, the probability that the BitcoinCash (BCH/XBC) community will have to deal with a contentious fork on November 15th, 2018 is pretty high at this point.

Here is the list of the contentious forks that have been announced in the past weeks:
- **Bitcoin ABC** (BCH/XBC) maintained by Bitcoin ABC & bcoin-org/bcash which is basically the original and historical Bitcoin Cash chain that is going to be changed with new concensus rules [[more info](https://bitcoinabc.org)];
- **Bitcoin SV** (BCH/XBS) maintained by nChain which is basically the contentious fork who wants to push different consensus rule changes [[more info](https://github.com/bitcoin-sv)];
- **Bitcoin NayBC** (BCH/XBN) maintained by Tom Harding which is basically the chain of the naysayers to both of the proposed forks and who want to stay on the current chain [[more info](https://github.com/dgenr8/bitcoin-abc)].

Additional implementations:
- **Bitcoin Unlimited** is not yet mentioned as strictly compatible to XBC nor XBS nor XBN [[more info](https://www.bitcoincash.org)]. They are committed to avoiding a split, hence following the dominant chain. Also, they allowed each consensus rule change to be assessed independently and activated independently.
- The above applies to **Bitcoin XT** also maintained by Tom Harding.

Final note:
This list may evolve and will be updated if any new announcement is made by the community.
The owners of each implementations can process to a pull-request on chopsticks Github, if they want to improve the above definition.


## Project Goals

No one knows how things will go during this contentious forking period, and which Bitcoin Cash forks are going to survive and to be supported by the community.

Early adopters of Bitcoin Cash and application developers like us at [eminent.ly](https://eminent.ly) (a business social network that records proofs of referral and interest on-chain) cannot take the risk of losing any of their transactions during this period, and absolutely need to continue operating during the contentious fork.

The solution we propose here with [chopsticks.cash](https://api.chopsticks.cash) is to provide an API that will record transactions on each Bitcoin Cash post-fork chains (XBC, XBS, XBN, etc.). 

That way, Bitcoin Cash application developers will not be taken hostage in the conflicts opposing miners and/or protocol developers, and will be able to follow each chain incuring no additional costs.

In fact, the process of forking the chain will duplicate the funds you own on each chain, by definition. Thus, processing your transaction will not cost you more coins than if you were recording it on a unique and non-forked Bitcoin Cash chain. 

Bitcoin Cash application developers will be able to operate their business normally during the conflicting period. 

And finally, we will see how things will go and which chain(s) the market will decide to support, and which one to continue operating on.


## Technical architecture

### Chopsticks API

Here are the main features of [chopsticks.cash](https://api.chopsticks.cash) API:
1. take your pre-signed transaction (POST request) and execute your transaction on the 3+ chains;
2. cast your vote about your chain preference for example (XBC, XBN, XBS) by descending order of preference;
3. give statistics to the community.

API URL: https://api.chopsticks.cash


#### 1. Send a raw transaction to all chains' node

We decided to provide only one way for you to send us your transaction: a raw signed transaction (see:
```sendrawtransactions``` command of bitcoin-cli). 

In fact, we don't want you to pass your private key to our API, even if it is encrypted, for the security of your funds. 

So in order to use our API, you will have to sign your transaction on the client-side (your side) prior to calling our API and passing the hexadecimal representation of the signed transaction.

We intend to provide a Golang library to help you sign your transaction on the client-side. Contributions will be welcome for other programming languages.

Also, you will have to pass the chains you want us to execute your transaction on.

Optionally, you can cast a vote for your most preferred forks.


##### Request
 
```
POST /api/transactions
```

```json
{ 
  tx_hex: "aaaa...bbb", 
  blockchains: ["XBC", "XBS", "XBN"],
  voting: false
}
```

Note that the order of the blockchains array will count as vote if you pass ```voting:true``` within your JSON request.

So if you want to rank XBC 1st, XBN 2nd, XBS 3rd, you will need to write:

```json
{ 
  ...
  blockchains: ["XBC", "XBN", "XBS"],
  voting: true
}
```


##### Response

The API will send you a response containing your original hexadecimal representation of the signed transaction, the hash of your transaction as well as some info related to the chain it was excecuted on, the casted vote data and its signature.
Note that block_height is the last block mined not the block that your transaction will be added to. See [model/Transaction](https://github.com/eminently/chopsticks/blob/master/model/Transaction.go) and [model/Vote](https://github.com/eminently/chopsticks/blob/master/model/Vote.go) to the full list of JSON attributes.

```json
{ 
  tx_hex: "aaaa...bbb", 
  transactions: [
    { blockchain_type:"XBC", hash: "aaa...bbb", blockchain_version:"v0.18.2.0-unk", block_height:555555, ... },
    { blockchain_type:"XBS", hash: "aaa...bbb", blockchain_version:"0.1.0.0-beta-200015661", block_height:555555, ... },
    { blockchain_type:"XBN", hash: "aaa...bbb", blockchain_version:"v0.17.2.0-5210f8f46", block_height:555555, ... }
  ],
  vote: {
    uuid: "1234-...-accd",
    preferredChains: ["XBC", "XBN", "XBS"],
    created: 123456789,
    ...
  },
  vote_signature: "bbbb...ccc",
  errors: []
}
```


### Chopsticks Infrastructure

We will run the 3+ different nodes and chopsticks API on AWS.

[chopsticks.cash](https://api.chopsticks.cash) API will exclusively connect to these 3+ nodes that we will maintain.


## Contribute To And Support The Project

Join chopsticks.cash [Telegram group](https://t.me/joinchat/FmkGFhJBwEvLb00XQ1ztIA) and share your feedback and ideas with the team! 

In order to help us develop the API and run the AWS nodes, you can donate some [$BCH](https://coinmarketcap.com/currencies/bitcoin-cash/), please send it to our wallet address: [bitcoincash:qzuw0upmzk83lupq86lc6l62znpqs3k6svtf292dql](https://www.blocktrail.com/BCC/address/1HrhBfFRFovHv8EMxsuB9EcZgamtuH3fMc). You will be part of the journey, a founding member!



## License

The MIT License

Copyright (c) 2018 Eminently, LLC. https://www.eminent.ly

Permission is hereby granted, free of charge, to any person obtaining a copy
of the software provided within this source code repository and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

<sub>
THIS SOFTWARE AND SERVICE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS “AS IS” AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
</sup>
