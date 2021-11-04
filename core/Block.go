package core

import (
	"crypto/sha256"
	"encoding/hex"
	"time"
)

// 一个区块 包括 区块体和区块链头
type Block struct {
	// 区块头
	Index         int64  // 区块编号 代表区块位置
	Timestamp     int64  // 区块时间戳 代表区块创建的时间
	PrevBlockHash string // 上一个区块哈希值
	Hash          string // 当前区块哈希值

	// 区块体（数据）
	Data string // 区块数据
}

// 函数  计算哈希值
func calculateHash(b Block) string {
	blockData := string(b.Index) + string(b.Timestamp) + string(b.PrevBlockHash) + string(b.Data)
	hashInBytes := sha256.Sum256([]byte(blockData))
	return hex.EncodeToString(hashInBytes[:])
}

// 创建一个新区块
func GenerateNewBlock(preBlock Block, data string) Block {
	newBlock := Block{}
	newBlock.Index = preBlock.Index + 1
	newBlock.PrevBlockHash = preBlock.Hash
	newBlock.Timestamp = time.Now().Unix()
	newBlock.Data = data
	newBlock.Hash = calculateHash(newBlock)
	return newBlock
}

// 创始区块
func GenerateGenesisBlock() Block {
	preBlock := Block{}
	preBlock.Index = -1
	preBlock.Hash = ""
	return GenerateNewBlock(preBlock,"Genesis Block")
}
