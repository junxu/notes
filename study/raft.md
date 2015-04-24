# raft论文阅读笔记
## 1. Motivation
Consensus algorithms typically arise in the context of replicated state machines. In this approach, state machines on a collection of servers compute identical copies of the same state and can continue operating even if some of the servers are down. Replicated state machines are used to solve a variety of fault tolerance problems in distributed systems.

在实际系统中Consensus algorithms应该具有以下属性:

* safety (never returning an incorrect result): under all non-Byzantine conditions, including network delays,  partitions, and packet loss, duplication, and reordering.
* available: as long as any majority of the servers are operational and can communicate with each other and with clients.
* do not depend on timing to ensure the consistency of the logs
* a command can complete as soon as a majority of the cluster has responded to a single round of remote procedure calls

## 2. Basic Raft algorithm
raft协议的设计主要是为了understandable.
### 2.1 Designing for understandable
采用了以下技术：
* problem decomposition: wherever possible, we divided problems into separate pieces that could be solved, explained, and understood relatively independently. Raft算法分成三个部分：leader selection, log replication and safety.
* Simplify the state space by reducing the number of states to consider, making the system more coherent and eliminating nondeterminism where possible. 如在raft算法中，logs是不允许有holes.

### 2.2 Raft overview

### 2.3 Raft basics
### 2.4 Leader election
### 2.5 Log replication
### 2.6 Safty
#### 2.6.1 Election restriction
#### 2.6.2 Committing enties from previous terms
#### 2.6.3 Safety argument
### 2.7 Follower and candidate crashed
### 2.8 Persisted state and server restart
### 2.9 Timeing and availability
### 2.10 Leadship transfer extension

## 3. Cluster membership changes

## 4. Log compaction

