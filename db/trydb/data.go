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
