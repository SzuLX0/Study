package block

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"strconv"
	"time"
)

//块结构的信息：索引，时间戳，数据，哈希值，前置哈希值
//块的生成：注意首块和其他块的区别
//块的存储结构（账本）：块链表
//检查块：前后索引一致？前后hash值一致？当前hash值正确？

type Block struct{
	Index int
	Timestamp	int64
	Data []byte
	PreBlockHash []byte
	Hash []byte
}

//计算哈希
func (b *Block)GetHash() []byte {
	index := []byte(strconv.Itoa(b.Index))
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	headers := bytes.Join([][]byte{b.PreBlockHash, b.Data, timestamp}, index)
	hash := sha256.Sum256(headers)
	return hash[:]
}

//设置哈希值
func (b *Block) setHash(){
	hash := b.GetHash()
	b.Hash = hash
}

//创建区块
func NewBlock(data string, preBlockHash []byte, index int) *Block{
	Index := index+1
	block := &Block{Index, time.Now().Unix(),[]byte(data), preBlockHash, []byte{}}

	block.setHash()
	return block
}

//创建第一个区块
func BlockInit(data string) *Block{
	block := &Block{0, time.Now().Unix(), []byte(data),
		[]byte{}, []byte{}}

	block.setHash()
	return block
}


func (b Block) PrintBlock(){
	time := time.Unix(b.Timestamp, 0).Format("2006-01-02 15:04:05")
	fmt.Printf("index = %d, TimeStamp = %s\n", b.Index, time)
	fmt.Printf("Hash = %x\n", b.Hash)
	fmt.Printf("PreBlockHash = %x\n", b.PreBlockHash)
	fmt.Println()
}