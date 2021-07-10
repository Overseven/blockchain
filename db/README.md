# Database structure

## Алгоритм

### Исходные данные

- Локальная версия блокчейна (__ЛБ__), последний блок с номером 12

### Ситуация 1
Из удаленного блокчейна (__УБ__) поступает блок с меньшим номером, 
чем есть в __ЛБ__, например, 8

Действия:

- возвращаем отправителю сообщение с информацией о номере последнего блока в __ЛБ__ (12)

### Ситуация 2
Из __УБ__ поступает блок с номером, совпадающим с номером последнего блока в __ЛБ__, то есть, 12

Действия:

- возвращаем отправителю сообщение о равенстве номеров последних блоков

### Ситуация 3
Из __УБ__ поступает блок с номером превышающим номер последнего блока в __ЛБ__, например, 20

Действия:
- запрашиваем недостающие блоки [13; 20] порцией по N блоков (например, по 4 блока), начиная с 13
  
- если поле "хэш предыдущего блока" пришедшего блока 13 идентично хэшу блока 12 из __ЛБ__:
  - нет расхождения, данные о транзакциях и событиях из новых блоков записываем в бд
  - перед добавлением каждого нового блока проверяем его корректность
  - если один из новых блоков не корректен (например, 17):
    - останавливаем запись новых блоков (теперь последний блок - 16)
    - возвращаем отправителю сообщение об ошибке в блоке
  - в противном случае завершаем запись, запрашиваем следующую порцию из N блоков
    
- в противном случае
  - запрашиваем по N блоков (только его хэш и хэш предыдущего блока) в обратном порядке, ищем блок, с которого начинается расхождение
  - например, последний одинаковый блок, после которого начинается расхождение, имеет id 9
  - ищем снапшот с ближайшим номером последнего блока, чтобы блок расхождением имел более высокий id (например, снапшот с последним блоком 5)
  - добавляем во временную область бд (temp) балансы, голосования из блоков __ЛБ__ с 5 < id <= 9 
  - запрашиваем и обрабатываем блоки пачками по __N__ штук, например, по 4:
    - проверяем каждый новый блок на корректность относительно данных из temp и снапшота 
    - если новый блок не корректен:
        - если id <= 12, очищаем temp, возвращаем отправителю сообщение об ошибке в блоке
        - в противном случае останавливаем запись в temp, заменяем блоки из бд блоками из temp


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
| transaction info | `tiXXXXXXXX` | `XXX` - transaction hash | [transaction](#transaction) |
| transaction by counter | `taXXXXYYY` | `XXX` - address, `YYY` - count | [transaction ref](#transaction-by-counter) |
| voting info | `viXXXXXXX` | `XXX` - voting id | [voting](#voting) |
| latest balance | `lbXXXXXXX` | `XXX` - wallet pubKey | [balance](#balance) |
| latest param | `lpXXX` | `XXX` - param id | [parameter](#parameter) |
| latest fee distribution ratio | `lfXXXXXXXX` | `XXX` - node address | [ratio](#fee-distribution-ratio) |
| number of snapshots | `cs` | | __uint8__|
| last block | `cb` | | __uint64__|

### Snapshot dependent

| Name | Key | Description | Value |
| --- | --- | --- | --- |
| last block | `sYn` | `Y` - snapshot | __uint64__ |
| balance |`sYbXXXXXXX` | `Y` - snapshot, `XXX` - wallet pubKey | [balance](#balance) |
| param | `sYpXXX` | `Y` - snapshot, `XXX` - param | [parameter](#parameter) |
| fee distribution ratio | `sYfXXXXXXXX` |  `Y` - snapshot, `XXX` - node address | [ratio](#fee-distribution-ratio) |

### Temporary structures
This section describes temporary data structures that used 
when trying to replace local blockchain with external. 

Snapshot with the closest **last block** id is head of temporary transactions data. 
If external blockchain has no errors inside, 
part of local blockchain with **block** id > snapshot **last block** id
will be removed and replaced by temporary data


| Name | Key | Description | Value |
| --- | --- | --- | --- |
| block | `0bXXXXXXXX` |`XXX` - block id | [block](#block) |
| transaction info | `0tiXXXXXXXX` | `XXX` - transaction hash | [transaction](#transaction) |
| transaction by counter | `0taXXXXYYY` | `XXX` - address, `YYY` - count | [transaction ref](#transaction-by-counter) |
| voting info | `0viXXXXXXX` | `XXX` - voting id | [voting](#voting) |
| last block | `0Yn` | `Y` - snapshot | __uint64__ |
| balance |`0YbXXXXXXX` | `Y` - snapshot, `XXX` - wallet pubKey | [balance](#balance) |
| param | `0YpXXX` | `Y` - snapshot, `XXX` - param | [parameter](#parameter) |
| fee distribution ratio | `0YfXXXXXXXX` |  `Y` - snapshot, `XXX` - node address | [ratio](#fee-distribution-ratio) |

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
sn_id = get_most_closely()  # 0, id of snapshot with max last block number and lower then delete_starting
sn_last = get_snapshot_last_block(sn_id)  # 10, last block id from snapshot

# deleting blocks in descending order
for (i = last_block; i >= delete_starting; i-=1){
    trans = get_transactions_from_block(i)  # list of transactions in current block
    for t in trans.reverse() {  # deleting trans in descending order
        if (t.type == Voting) {
            delete_vote_from_voting(t.voting_id, get_hash(t)) 
        } else if (t.type == VotingInit) {
            delete_voting(t.voting_id)
        }
        delete_transaction(t)
    }
    delete_block(i)
}

# updating cache from snapshot
delete_all_with_prefix('l')  # deleting latest block number, balances, params, fee ratios 
copy_snapshot_to_latest(sn_id)  # copy last block number, balances, param, fee ratios from snapshot ('sY...' -> 'l...')

# insert new blocks
for b in new_block {
    for t in b.trans() {
        if (t.type == Airdrop) {
            add_to_balance(t.receiver, t.pay)
        } else if (t.type == Transfer) {
            add_to_balance(t.sender, -1 * t.pay)
            add_to_balance(t.receiver, t.pay)
        } else if (t.type == VotingInit) {
            new_voting(t.voting_id, t). # write into db new voting
        } if (t.type == Voting) {
            add_vote_to_voting(t.voting_id, get_hash(t)) 
        } else if (t.type == FinishVoting) {
            if (voting_can_be_finished(t.voting_id, b.id)) {  # voting_can_be_finished() return true, 
                                                              #  if current block id > voting.endBlock && voting.finished == false 
                set_voting_finished(t.voting_id, finished=true)
                reward = calc_reward(t.voting_id, b.id)
                add_to_balance(t.sender, reward)  # reward for FinishVoting transaction sender
            }
        }
        
        # distribution of the reward between Node and Miner
        add_to_balance(t.sender, -1 * t.fee)
        fee_ratio = get_distr_fee_ratio(t.node)
        add_to_balance(t.node, fee_ratio * t.fee + get_emission()/2)
        add_to_balance(t.miner, (1.0 - fee_ratio) * t.fee + get_emission()/2)
        
        add_transaction(t)  # write into db new transaction
    }
}

# update number of last block
last_block = delete_ending
```

### Transaction
#### Transaction by counter
Every address has an individual counter of sent transactions. 
This counter insert in every transaction from the address.

This structure is using to quick find user transaction with the defined transaction counter value.

| №   | Field | Size |
| --- | --- | --- |
| 1 | transaction hash | __uint8__ \[32\] |

Example of usage:

```python
address = 0x02ca6a856ad061102fa1fecdb566fd6c34df15bc42815015cf7e31974fa6a3fbd6
trans_counter = 12

trans_hash = get_from_db(key='ta' + address + trans_counter)
if trans_hash == null:
  print('transaction with counter = ', trans_counter, " not exist)
  return

transaction = get_from_db(key='ta' + trans_hash)
```


#### Airdrop
| №   | Field | Size |
| --- | --- | --- |
| 1 | block number | __uint64__ |
| 2 | type | __uint8__ |
| 3 | transaction counter | __uint32__ |
| 4 | receiver | __uint8__ \[32\] |
| 5 | pay | __float64__ |
| 6 | fee | __float64__ |
| 7 | message len | __uint32__ |
| 8 | message | __uint8__ \[N\] |
| 9 | timestamp len | __uint8__ |
| 10| timestamp | __uint8__ \[N\] |
| 11| node pubKey | __uint8__ \[32\] |
| 12| sign | __uint8__ \[32\] |

#### Transfer
| №   | Field | Size |
| --- | --- | --- |
| 1 | block number | __uint64__ |
| 2 | type | __uint8__ |
| 3 | sender | __uint8__ \[32\] |
| 4 | transaction counter | __uint32__ |
| 5 | receiver | __uint8__ \[32\] |
| 6 | pay | __float64__ |
| 7 | fee | __float64__ |
| 8 | message len | __uint32__ |
| 9 | message | __uint8__ \[N\] |
| 10| timestamp len | __uint8__ |
| 11| timestamp | __uint8__ \[N\] |
| 12| node pubKey | __uint8__ \[32\] |
| 13| sign | __uint8__ \[32\] |

#### Voting init transaction
| №   | Field | Size |
| --- | --- | --- |
| 1 | block number | __uint64__ |
| 2 | type | __uint8__ |
| 3 | sender | __uint8__ \[32\] |
| 4 | transaction counter | __uint32__ |
| 5 | voting id | __uint64__ |
| 6 | parameter | __uint16__ |
| 7 | value len | __uint32__ |
| 8 | value | __uint8__ \[N\] |
| 9 | fee | __float64__ |
| 10| timestamp len | __uint8__ |
| 11| timestamp | __uint8__ \[N\] |
| 12| node pubKey | __uint8__ \[32\] |
| 13| sign | __uint8__ \[32\] |

#### Voting transaction
| №   | Field | Size |
| --- | --- | --- |
| 1 | block number | __uint64__ |
| 2 | type |  __uint8__ |
| 3 | sender | __uint8__ \[32\] |
| 4 | transaction counter | __uint32__ |
| 5 | voting id | __uint64__ |
| 6 | vote len | __uint32__ |
| 7 | vote | __uint8__ \[N\] |
| 8 | fee | __float64__ |
| 9 | timestamp len | __uint8__ |
| 10| timestamp | __uint8__ \[N\] |
| 11| node pubKey | __uint8__ \[32\] |
| 12|sign | __uint8__ \[32\] |

#### Voting finish transaction
| №   | Field | Size |
| --- | --- | --- |
| 1 | block number | __uint64__ |
| 2 | type | __uint8__ |
| 3 | sender | __uint8__ \[32\] |
| 4 | transaction counter | __uint32__ |
| 5 | voting id | __uint64__ |
| 6 | fee | __float64__ |
| 7 | timestamp len | __uint8__ |
| 8 | timestamp | __uint8__ \[N\] |
| 9 | node pubKey | __uint8__ \[32\] |
| 10 | sign | __uint8__ \[32\] |

### Voting
| №   | Field | Size |
| --- | --- | --- |
| 1 | id | __uint64__ |
| 2 | start on block | __uint64__ |
| 3 | end on block | __uint64__ |
| 4 | finished | __bool__ |
| 5 | vote trans. hash | __uint8__ \[N*32\]


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
