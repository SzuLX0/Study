package block

import (
	"bytes"
	"fmt"
	"time"
)

type BlockChain struct {
	blocks []*Block
}

func CreateBlockChain() *BlockChain{
	fmt.Println("初始化：建立第一个区块")
	blockChain := BlockChain{
		blocks:[]*Block{},
	}
	blockChain.blocks = append(blockChain.blocks, BlockInit("The blockChain init"))
	fmt.Println("初始化完成")
	return &blockChain
}

func (bc *BlockChain)AddBlock(data string) {
	fmt.Println("创建区块...")
	index := len(bc.blocks)
	newBlock := NewBlock(data, bc.blocks[index-1].Hash, index-1)
	bc.blocks = append(bc.blocks, newBlock)
	fmt.Println("添加完成")
}

func (bc *BlockChain)Check() error{
	length := len(bc.blocks)
	for index, b := range bc.blocks{
		//1.检查索引是否正确
		if b.Index != index{
			fmt.Printf("索引错误，错误位置为%d\n", index)
			return fmt.Errorf("索引错误")
		}
		//2.检查hash是否正确
		if !bytes.Equal(b.Hash, b.GetHash()){
			fmt.Printf("hash值错误，错误位置为%d\n", index)
			return fmt.Errorf("hash值错误")
		}
		//3.检查前后hash是否一致
		if index+1<length && !bytes.Equal(b.Hash, bc.blocks[index+1].PreBlockHash){
			fmt.Println("前后hash值不一致,错误位置为%d\n", index)
			return fmt.Errorf("前后hash值不一致")
		}
	}
	return nil
}

func (bc *BlockChain) Find(index int) Block {
	if 0<=index && index<len(bc.blocks){
		return *bc.blocks[index]
	} else{
		return Block{}
	}
}

func (bc *BlockChain) PrintBlocks(){
	for _,b := range bc.blocks{
		time := time.Unix(b.Timestamp, 0).Format("2006-01-02 15:04:05")
		fmt.Printf("index = %d, TimeStamp = %s\n", b.Index, time)
		fmt.Printf("Hash = %x\n", b.Hash)
		fmt.Printf("PreBlockHash = %x\n", b.PreBlockHash)
		fmt.Println()
	}
}
