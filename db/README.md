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
| voting info | `viXXXXXXX` | `XXX` - voting id | [voting](#voting) |
| latest balance | `lbXXXXXXX` | `XXX` - wallet pubKey | [balance](#balance) |
| latest param | `lpXXX` | `XXX` - param id | [parameter](#parameter) |
| latest fee distribution ratio | `lfXXXXXXXX` | `XXX` - node address | [ratio](#fee-distribution-ratio) |
| number of snapshots | `cs` | | __uint8__|

### Snapshot dependent

| Name | Key | Description | Value |
| --- | --- | --- | --- |
| last block | `sYn` | `Y` - snapshot | __uint64__ |
| balance |`sYbXXXXXXX` | `Y` - snapshot, `XXX` - wallet pubKey | [balance](#balance) |
| param | `sYpXXX` | `Y` - snapshot, `XXX` - param | [parameter](#parameter) |
| fee distribution ratio | `sYfXXXXXXXX` |  `Y` - snapshot, `XXX` - node address | [ratio](#fee-distribution-ratio) |


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

Block keys are always takes 8 byte, the least significant bytes places at the end of the key for correct sort order:

|    Key  | Number of block |
|   ---   |  --- |
| 0000000000000001 |    1 |
| 0000000000000010 |   16 |
| 0000000000000A00 | 2560 |

Pseudocode with algo when some blocks need to be replaced:

``` python
last_block = 18  # number of last block in db 
delete_starting = 11  # id of the block that will be deleted with all next (delete [11; 18])
delete_ending = 20  # id of new last block
snapshots = get_number_of_shapshots()
sn_id = get_most_closely()  # id of snapshot with max last block number and lower then delete_starting
sn_last = get_snapshot_last_block(sn_id)  # last block id from snapshot

# deleting blocks in descending order
for (i = last_block; i >= delete_starting; i-=1){
    trans = get_transactions_from_block(i)  # list of transactions in current block
    for t in trans {
        if (t.type = Voting) {
            delete_vote_from_voting(t.voting_id, get_hash(t)) 
        } else if (t.type = VotingInit) {
            delete_voting(t.voting_id)
        }
        delete_transaction(t)
    }
    delete_block(i)
}



last_block = delete_ending
```

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
| 5 | finished | __bool__ |

### Balance
| №   | Field | Size |
| --- | --- | --- |
| 1 | amount of tokens | __float64__ |


### Parameter
| №   | Field | Size |
| --- | ---   | ---  |
|  1  | len   | __uint32__ |
|  2  | value | __uint8__ \[N\] |


### Fee distribution ratio
| №   | Field | Size |
| --- | ---   | ---  |
|  1  | ratio |  __float64__ |
