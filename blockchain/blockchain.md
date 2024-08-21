# 什么是区块链？什么是智能合约？是什么gas?

这是一篇介绍性的文章，目的是给出一些我们将在这里所做事情的背景信息。区块链有什么特别之处？为什么我们要使用它？这篇文章还将定义一些基本的术语，像智能合约和gas。如果你感觉迷茫，这是一个开始的好地方。

## Web2

The term Web 2.0 is used to describe the traditional Internet as you know it. This Internet is operated primarily by large corporations like Google. The **formal reason** for the existance of corporations is to generate profit for their shareholders. This means that the global good or the benefits of their users are a byproduct **and not the end goal.**

Web 2.0这个术语被用来描述你所知道的传统的互联网。这种互联网主要由谷歌这样的大公司运营。大公司存在的主要原因是去为他们的股东创造利润。这意味着全球的利益或用户的利益是一个副产品，而不是最终目的。

Let's take a Web2 service like Gmail. You, as a user of this service, are **inherently** not equal to the creator of the service, Google. We call this **property centralization**. Centralized services don't provide users with true ownership. If Google **decides** that you've breached its terms of service, for example, **it is allowed to** take your email access away from you. Centralized services are also permissioned, meaning that to use the service and send an email you must ask for permission. If Google decides that your email is spam, it is not required to send it.

让我们拿Gmail这样的Web2服务举例。你作为这个服务的用户，本质上不等同于这个服务的创建者Google。我们把这叫做所有物中心化。中心化服务不会给用户提供真正所有权。例如，如果谷歌认为你违反了他的服务条款，则会拿走你电子邮件的访问权限。中心化服务也是需要获得许可的，这意味着当你使用这个服务和发送一个邮件时，你必须请求许可。如果谷歌确定你的邮件是垃圾邮件，那么它不允许被发送。

Centralization is based on trust. Users allow Google to maintain a position of power because they trust Google with this power.

中心化是基于信任的。用户允许谷歌区维持权力地位，**因为他们相信拥有这种权力的谷歌**。

## Web3 

Many of us view the Internet as a common good. We see it as a tool that turns the world into a global village. A tool that allows users to communicate and form communities. As such, we would like to see a shift of power from corporations to users. Web3 is the implementation of this ideal, what we would like to see as the next stage of evolution of the Internet.

我们很多人把互联网看作是一个共同利益。我们把它当作一个把世界变成一个地球村的工具。**一种允许用户交流然后形成社区的工具**。因此，我们将**希望看到**权力从公司转移到用户手里。**Web3就是这一理念的实现**，我们**希望看到**互联网的演变的下一个阶段。

Under this ideal, you, as a user of a Web3 service, must be inherently equal to the creator of the service. We **call this property decentralization**. Decentralized services provide users with true ownership. This is not only true for data, but also for assets. Decentralized assets like your Bitcoins or your TON coins are yours, and nobody can take them away from you. Decentralized services are also permissionless, meaning that to transfer your TON coin to somebody else requires nobody's permission but your own. Nobody can stop this from happening or censor you.

在这种理念下，你作为Web3服务的用户，**本质上一定与这个服务的创建者平等**。我们这称之为所有物去中心化（财产分散化）。去中心化的服务给用户提供真正的所有权。**这不仅适用于数据，也适用于资产**。**像Bitcoins或者TON coins等去中心化的资产是你的**，没人可以从你手里拿走它们。去中心化的服务也是无需许可的，**这意味着去转移你的TON coin给其他人不需要任何人的许可， 除了你自己的许可。**没人可以阻止这种事的发生或者审查你。

Decentralization allows systems to be trustless. Since there are no positions of authority, authority cannot be abused to hurt users.

去中心化使系统变得无需信任。由于没有权威职位，**不会滥用权力去伤害用户**。

## The blockchain

The ideal of Web3 sounds great on paper, but is it practical? As developers, how can we implement services where we are inherently equal to our users? Implementing a service normally requires the developer to write a backend. This backend runs on some server. Who owns the server? the developer. The developer can change the server without asking or even take it down. This relationship is inherently not equal. Backend servers are centralized.

Web3的理念在理论上听起来很美好，但是它真的实用吗？作为开发人员，**我们如何才能实现与用户本质上平等的服务？**实现一个服务通常需要开发者去写一个后端。这个后端运行在一些服务器上。谁拥有这个服务器呢？开发商。开发人员无需询问改变服务器甚至关闭服务器。这种关系本质上是不平等的。后端服务是中心化的。

区块链技术被发明去解决这个问题并且使开发人员去创建一个去中心化的后端。谁运行这个后端？是用户去做。由于关系是平等的，任何希望去参与运行后端的用户都被允许去这么做。区块链以在用户之间协作的方式运行。

Collaboration is governed by consensus. For an execution result to hold true, multiple users, a majority to be exact, must all vote to confirm the result. This makes blockchains very inefficient since every calculation must be executed by many users. This also makes blockchains expensive to run compared to a traditional centralized server.

合作以共识为基础。**为了使执行结果成立，多个用户，确切的说是大多数用户，必须全部投票去确认结果。**由于每个计算必须被很多用户执行，这使得区块链很低效。**与一个传统的中心化服务相比，这也使得区块链运行成本更昂贵**。

## The token 

We mentioned that execution results require voting. How does it work? Is it - one user, one vote - like in democratic elections? It turns out that this doesn't work on the Internet due to something known as the Sybil attack. It is very easy to create fake users on the Internet. Since Web3 is decentralized, we can't have a centralized source of authority that decides who's fake and who's real.

我们提到执行结果需要投票。它是如何运作的呢？是像民主选举那样一人一票吗？事实证明，**由于所谓的Sybil攻击，**这在互联网上不适用。在互联网上创建一个假的用户是很容易的。由于Web3是去中心化的，我们没有一个中心化的权利来源去决定谁是假的谁是真的。


一个主流的去中心化的解决这个问题的方法是把投票的权利围绕代币。如果你拥有10代表，你拥有10票。代币不能被造假，辨别代币的真假很容易。TON区块链围绕TON coin。以太坊区块链围绕以太币。这意味着每一个区块链都是一个经济体。代币作为一种激励工具可以确保去中心化的社区向着相同的方向发展。

## Network validators 

All blockchains are networks because they are operated by a group of users. Users that do the heavy lifting of operating the network and actively participate in the consensus process are called validators. The voting weight of every validator is proportional to the amount of tokens they hold. To keep validators honest, they are normally required to put their tokens at stake. If the consensus deems that a validator is dishonest, their tokens will be taken away as punishment. This governance process is called proof-of-stake.

所有的区块链都是互联网，因为他们由一个用户群体运营。**那些承担网络运营重任并积极参与共识过程的用户被称为做验证者**。每一个验证者的投票比重与他们所持有的代币数量是成正比的。为了保持验证者的诚实，他们通常需要把他们的代币置于危险中。如果共识认为一个验证者是不诚实的，他们的代币将被没收作为惩罚。这种管理方法被叫做**权益证明。**

Being a network validator is usually hard work. You need to run the blockchain node code on a server that you own and stake it with a lot of tokens. **Smaller users** that want to participate but don't have enough tokens to **warrant going to all this effort** can often delegate their tokens to one of the **larger** validators. These participants are called **nominators**.

作为一个网络验证者通常是艰难的工作。你需要在你自己的服务器上运行区块链节点代码并用大量代币抵押它。**那些想要去参加但是没有足够多代币来保证进行所有这些努力较小的用户通常可以把他们的代币委托给较大的验证者之一**。这种参与者被称为提名人。

## Gas fees 

We said earlier that blockchains are economies. The equipment for network validators is not free, so they must be paid for their efforts. Payment naturally takes place with the token of the blockchain. On the TON blockchain, users pay fees using the TON coin. TON network validators earn TON coin for performing the validation process and executing all the apps that are running on the blockchain.

我们之前说区块链是经济体。网络验证者的设备不是免费的，**因此他们的努力必须得到报酬**。**支付自然用区块链的代币进行的**。在TON区块链中，用户使用TON币支付费用。TON网络验证者从执行验证过程和执行在区块链上运行的所有应用赚取TON币。

When a user is performing some action on the blockchain, they must send a transaction. The transaction includes a fee payment called gas. The **analogy** comes from cars. Just like a car needs gas to run, so does a blockchain transaction. Users must sign transactions using their blockchain wallets. This signature guarantees that only the owner of the wallet can authorize the payment of gas and sending the transaction.

当一个用户在区块链上执行某些操作时，他们必须发送一个交易。这个交易包含的支付费用被称为gas。这个类比来自汽车。就像汽车需要汽油才能行驶。区块链交易也是一样。用户必须用他们的区块链钱包给交易签名。签名可以保证只有钱包的拥有者可以授权gas的费用支付并发送交易。

## Dapps 

We said earlier that the purpose of blockchains is to run decentralized backends. A simpler name for these services that run on the blockchain network **is apps - decentralized apps to be exact** or dapps for short. Developers create dapps and have network validators execute them.**Users interact with dapps** by sending them transactions. The developer of a dapp is equal to the dapp's users. **The developer should have no special privileges since the app is decentralized.**

我们之前说过区块链的目的是去运行去中心化的后端。这些运行在区块链网络的服务有一个更简单的**名字是应用，确切的说是去中心化应用**，或者简称为dapp。开发人员创建dapps并让网络验证者去执行他们。用户通过发送他们的交易**来和dapp交互**。一个dapp的开发人员与这个dappd的用户是平等的。**由于应用是去中心化的，开发人员没有特殊特权**。

Let's **reiterate** over the last point with an example. Let's take a Web2 service like Google Search. The developer of the service, Google, ranks search results for the benefit of users. Google enjoys this position of power and is allowed to promote its own products in search results. For example, when searching for "storage", Google can promote the result "Google Drive" over a competitor like "Dropbox". In a Web3 version of Google Search, the developer of the service will not be allowed to promote their own products in search.

让我们用一个例子重申最后一点。让我们拿谷歌搜索这种Web2服务举例。这个服务的开发人员，谷歌，为了用户的利益排名搜索结果。谷歌享有权利地位并且被允许在搜索结果中推广谷歌的产品。例如，当搜索“storage”时，谷歌可以优先推广“Google Drive”的结果而不是像“Dropbox”这样的竞争对手。在Web3版本的谷歌搜索中，服务的开发人员将不被允许去在搜索中推广他们自己的产品。

## Smart contracts

Every Web2 service like Google Search has terms of service. If a user feels that they were wronged under these terms, they can sue and ask a judge to rule on the dispute. Web3 is decentralized and **cannot have centralized sources of authority like judges.** On blockchain, code is law. The code of the dapp is the only agreement between its users. Unlike traditional legal agreements, this agreement is not open to interpretation. Code always executes **in the same way**.

每一个像谷歌搜索这样的Web2服务都有服务条款。如果用户感觉他们被这些条款不公正的对待了，他们可以提起诉讼并且要求法官对争议做出裁决。Web3是去中心化并且没有像法官这样拥有中心化的权利来源。在区块链中，代码就是法律。dapp的代码是在用户之间唯一的协议。不像传统合法的协议，这个协议不开放解释。代码总是以相同的方式执行。

Blockchain replaces legal contracts with code. The code of the dapp is therefore called a smart contract. Before users decide to participate in a dapp, by sending it transactions for example, they are expected to review the dapp's source code to understand its terms. Just like you wouldn't sign a lease contract for your apartment without reading the contract, you should not sign a transaction without reviewing the smart contract. Since not all users are tech savvy enough to do this, communities often rely on each other for this purpose.

区块链用代码替换了法律合同。dapp的代码因此被叫做智能合约。在用户决定去参与一个dapp之前，例如通过发送交易，他们需要去审查dapp的源代码去理解它的条款。就像你不会没有阅读合约就为你的公寓签署租赁合同。在没有仔细研究智能合约前，你不会签署交易。由于不是所有用户都足够精通技术来理解这么做，社区经常为了这个目的相互依赖。

After the developer finished writing the smart contract, the act of publishing this contract to the blockchain is called deployment. The contract code is deployed on to the chain where everybody can find it by its contract address.

在开发人员完成编写智能合约后，把智能合约发布到区块链的行为被称为部署。这个合约的代码被部署链上。每个人都能通过它的合约地址找到它。

Block  and explorers 

We said earlier that network validators must vote on the execution result of every transaction. To streamline this process, groups of transactions are batched together in blocks. Every block of transactions gets its own block number and undergoes the consensus process where a majority of network validators is required to approve it. When you order all the blocks one after the other you get a chain of blocks - this is the source of the word blockchain.

我们先前提到网络验证者必须要给每一笔交易的执行结果投票。为了提高这个过程的效率，在blocks中交易组会一起被成批处理。每一个交易块得到它自己的块号码并且经历共识过程。大多数的网络验证者都被要求去赞同它。当你一个一个排列所有区块后，你就得到了一连串的块-——这是区块链这个单词的来源。

After sending a transaction, a user must wait until this transaction is included in a block. A new block on TON blockchain is created every 5 seconds on average. **Users can inspect transactions, check if they succeeded or not and see which block they were added to by using a tool called a block explorer, or an explorer for short.**

在发送一个交易后，一个用户必须等待，直到这个交易被包含在一个块中。在TON区块链中一个新的区块平均每5s被创建。用户可以通过一个叫区块链浏览器，简称浏览器的工具简称交易是否成功并查看它们被添加到了哪个区块。

## So what is blockchain good for? 

在这篇文章中，我们主要介绍了什么是区块链以及定义了很多涉及到术语。实现像去中心化和去信任这种抽象的好处听起来是一件可怕的麻烦事。我们能给出更多实际的例子来说明区块链能改善你的生活？


根据谷歌翻译，翻译自[blockchain](https://blog.ton.org/what-is-blockchain "sleep")