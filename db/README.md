# Database structure


## Elements:
1) Block
2) Transaction
3) Snapshot
4) Balance
5) Parameter
6) Voting


## Address structure:
- `bXXXXXXXX` - block, `XXX` - block id
- `tXXXXXXXX` - transaction, `XXX` - transaction hash
- `vtXXXZZZZ` - voting transaction, `XXX` - voting id, `ZZZ` - vote id
- `viXXXXXXX` - voting info, `XXX` - voting id

- `s0bXXXXXXX` - snapshot `0`, `XXX` - wallet pubKey
- `s1bXXXXXXX` - snapshot `1`, `XXX` - wallet pubKey
- `s2bXXXXXXX` - snapshot `2`, `XXX` - wallet pubKey

- `s0p1`  - snapshot `0`, param `1`
- `s0p2`  - snapshot `0`, param `2`
- `s1p0`  - snapshot `1`, param `0`

- `lbXXXXXXX` - latest balance, `XXX` - wallet pubKey
- `lp0` - latest param `0`
- `lp1` - latest param `1`

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
