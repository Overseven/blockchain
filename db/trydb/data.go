package trydb

type Block struct {
    Id uint64
    NumOfTrans uint8
    TransHashes [][]byte
    Difficulty uint64
    MinerPubKey []byte
    Hash []byte
    Nonce uint64
}

type Voting struct {
    Id uint64
    StartBlockId uint64
    EndBlockId uint64
    Param uint16
    Value []byte
    Finished bool
}

type Balance struct {
    Address []byte
    Amount float64
    LastBlock uint64
}

type Parameter struct {
    Id uint16
    PrevValue []byte
    CurrValue []byte
}

type FeeDistRatio struct {
    Address []byte
    PrevRatio float64
    CurrRatio float64
}