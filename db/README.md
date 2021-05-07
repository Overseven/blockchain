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
- Id
- number of transactions
- transactions
- difficulty
- miner pubKey
- hash
- Nonce

### Transaction
#### Airdrop
- type
- receiver
- pay
- fee
- message
- timestamp
- node pubKey
- sign

#### Transfer
- type
- sender
- receiver
- pay
- fee
- message
- timestamp
- node pubKey
- sign

#### Voting transaction
- type
- sender
- voting id
- vote 
- fee
- timestamp
- node pubKey
- sign

### Balance
- amount of tokens

### Snapshot
#### Balance
- id of last block that was counted
- amount of tokens

#### Parameter
- id of last block that was counted
- value
