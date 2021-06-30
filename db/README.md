# Database structure


## Elements:
1) Block
2) Transaction
3) Snapshot
4) Balance
5) Parameter
6) Voting
7) Node/miner fee distribution ratio

## Address structure:
### Independent of snapshot 

| Name | Key | Description | Value |
| --- | --- | --- | --- |
| block | `bXXXXXXXX` |`XXX` - block id | [block](#block) |
| transaction | `tXXXXXXXX` | `XXX` - transaction hash | [transaction](#transaction) |
| voting transaction | `vtXXXZZZZ` | `XXX` - voting id, `ZZZ` - vote id | ??? |
| voting info | `viXXXXXXX` | `XXX` - voting id | [Voting](#voting) |
| latest balance | `lbXXXXXXX` | `XXX` - wallet pubKey | [balance](#balance) |
| latest param | `lpXXX` | `XXX` - param id | [parameter](#parameter) |

### Snapshot dependent

| Name | Key | Description | Value |
| --- | --- | --- | --- |
| balance |`sYbXXXXXXX` | `Y` - snapshot, `XXX` - wallet pubKey | [balance](#balance) |
| param | `sYpXXX` | `Y` - snapshot, `XXX` - param | [parameter](#parameter) |
| fee distribution ratio | `fYXXXXXXXX` |  `Y` - snapshot, `XXX` - node address | __TODO__ |


## Elements bytes-level structure:
### Block
| №   | Field | Size |
| --- | --- | --- |
| 1 | Id | __uint64__ |
| 2 | number of transactions | __uint8__ |
| 3 | transactions hashes | __uint8__ \[32 * N\] |
| 4 | difficulty | __uint64__ |
| 5 | miner pubKey | __uint8__ \[32\] |
| 6 | hash | __uint8__ \[32\] |
| 7 | Nonce | __uint64__ |

### Transaction
#### Airdrop
| №   | Field | Size |
| --- | --- | --- |
| 1 | block number | __uint64__ |
| 2 | type | __uint8__ |
| 3 | receiver | __uint8__ \[32\] |
| 4 | pay | __float64__ |
| 5 | fee | __float64__ |
| 6 | message len | __uint32__ |
| 7 | message | __uint8__ \[N\] |
| 8 | timestamp len | __uint8__ |
| 9 | timestamp | __uint8__ \[N\] |
| 10| node pubKey | __uint8__ \[32\] |
| 11| sign | __uint8__ \[32\] |

#### Transfer
| №   | Field | Size |
| --- | --- | --- |
| 1 | block number | __uint64__ |
| 2 | type | __uint8__ |
| 3 | sender | __uint8__ \[32\] |
| 4 | receiver | __uint8__ \[32\] |
| 5 | pay | __float64__ |
| 6 | fee | __float64__ |
| 7 | message len | __uint32__ |
| 8 | message | __uint8__ \[N\] |
| 9 | timestamp len | __uint8__ |
| 10| timestamp | __uint8__ \[N\] |
| 11| node pubKey | __uint8__ \[32\] |
| 12| sign | __uint8__ \[32\] |

#### Voting transaction
| №   | Field | Size |
| --- | --- | --- |
| 1 | block number | __uint64__ |
| 2 | type |  __uint8__ |
| 3 | sender | __uint8__ \[32\] |
| 4 | voting id | __uint64__ |
| 5 | vote len | __uint32__ |
| 6 | vote | __uint8__ \[N\] |
| 7 | fee | __float64__ |
| 8 | timestamp len | __uint8__ |
| 9 | timestamp | __uint8__ \[N\] |
| 10| node pubKey | __uint8__ \[32\] |
| 11|sign | __uint8__ \[32\] |

#### Voting init transaction
| №   | Field | Size |
| --- | --- | --- |
| 1 | block number | __uint64__ |
| 2 | type | __uint8__ |
| 3 | sender | __uint8__ \[32\] |
| 4 | voting id | __uint64__ |
| 5 | parameter | __uint16__ |
| 6 | value len | __uint32__ |
| 7 | value | __uint8__ \[N\] |
| 8 | fee | __float64__ |
| 9 | timestamp len | __uint8__ |
| 10| timestamp | __uint8__ \[N\] |
| 11| node pubKey | __uint8__ \[32\] |
| 12| sign | __uint8__ \[32\] |

### Voting
// TODO: check this
| №   | Field | Size |
| --- | --- | --- |
| 1 | id | __uint64__ |
| 2 | start on block | __uint64__ |
| 3 | end on block | __uint64__ |
| 4 | vote trans. hash | __uint8__ \[N*32\]

### Balance
| №   | Field | Size |
| --- | --- | --- |
| 1 | amount of tokens | __float64__ |


_А оно нам нужно?_
### Snapshot
#### Balance
| №   | Field | Size |
| --- | --- | --- |
| 1 | last transaction hash | __uint8__ \[32\] |
| 2 | amount of tokens | float64 |

#### Parameter
- id of last block that was counted
- value
