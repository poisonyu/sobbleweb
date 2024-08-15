# 区块链怎样在一个无人居住的小岛上帮助你

现今有很多关于区块链的炒作，但是如果你想问区块链实际上有什么用——很难得到一个明确的答案。有人常说区块链围绕信任展开，但是信任是一个很抽象的概念。我们能用一个具体的例子去证明它吗？这里有一个简单的东西将帮助我们理解为什么去中心化的共识（区块链）是有利于商业的。

## 那么，为什么去中心化的共识（区块链）对商业是有好处的？

想想电视剧迷失的开场——一架载满乘客的飞机坠毁在南太平洋的某处的一个荒凉小岛上。当幸存者探索这个岛屿并且弄明白救援无望的时候，他们开始过上一种新的生活——一个微型的文明开始出现。

Hugo是一个擅长种植的人，他喜欢去种植蔬菜。Sawyer在残骸种发现了斧子，开始劈柴。Kate是一个优秀的猎人，她擅长抓捕野猪。Jack打捞了一些药，当你病了并需要一些抗生素，他是你的帮手。

一个以物换物的经济发展起来了。当Sawyer饿了，他用一捆木柴与Hugo交换了4个西红柿。Kate用一头野猪交换了20个西红柿。Jack用一粒抗生素交换4捆木柴，但也会接受一头野猪。

一天早上Jack带着一个好主意醒来。让我们的文明迈入一个新的等级，引入金钱这个抽象的概念。他把这个想法扔给其他人。我们将创建一个IslandCoin并在我们的经济中使用它，而不是物物交换——这样即使Sawyer不需要任何抗生素，他仍然能够从Sawyer那里得到木柴。这个提议很简单：一个西红柿价值2个硬币，一头野猪价值40硬币，诸如此类。Kate问我们每个人开始有多少硬币？好，让我们公平一点——我们每个人开始将有100个硬币。

Since the survivors don’t have access to metal on the island, they obviously can’t mint actual coins. Jack suggests to keep track of how many coins each of them has. He comes up with a simple system. He’ll write on a piece of paper **a balance of 100** next to each of their names. Every time you want to transfer **somebody** some coins, just let him know the amount and he’ll update the **numbers** on his piece of paper.

由于幸存者在岛上不能获得金属，他们明显不能铸造真的硬币。Jack建议去记录他们每个人拥有的硬币。他想出一个简单的系统。他将在一张纸上写下100旁边写下他们每个人的名字。每次你想要转移某人的一些硬币时，只需要让他知道金额，他会更新他的纸上数字。

Sawyer doesn’t trust Jack very much, **especially with the whole alpha male tension around Kate**. Why should he hold the piece of paper? he asks. What prevents Jack from messing around with the numbers behind his back? Sawyer pitches his own idea — replace IslandCoin with his version called PacificToken! When you want to make a transfer, just let him know and he’ll keep track. Well, it’s becoming clear that Jack and Sawyer won’t play nice together. Kate tries to resolve the conflict by offering to track the numbers herself. It’s great because Jack and Sawyer will do anything she says anyways. But Hugo isn’t too happy about that though.

Sawyer非常不信任Jack,尤其是整个最有权势的雄性矛盾围绕着Kate。为什么他持有这张纸？他问道。什么阻止Jack在背后玩弄这些数字？Swayer抛出一个他自己的想法——用他的版本被叫做PacificToken来替换IslandCoin!当你想要去转移金额时，只要让他知道，他会跟踪记录。好吧，很明显Jack和Sawyer不会玩的好。Kate尝试去通过提议自己去记录这些数字来解决这个冲突。这很好，因为Jack和Sawyer无论如何都会做她说的任何事。但是Hugo对此不太满意。

The gang can’t decide on one individual that will keep track of the balance sheet — we have a deadlock. The idea falls through. Fast forward one year. The barter economy lives on. It sucks and they suffer, **but that’s the only thing they manage to get working**.

这群人不能决定一个人去跟踪记录资产负债表——我们陷入了一个僵局。这个想法落空了。一年很快过去了。物物交换的经济持续着。这很糟糕，他们很痛苦，但是这是唯一能设法工作的事情。



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