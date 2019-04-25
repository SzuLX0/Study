package block

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"strconv"
	"time"
)

var (
	index []int
	timestamp []int64
	data [][]byte
	preBlockHash [][]byte
	hash [][]byte
)

//计算哈希
func GetHash(Index int, Timestamp int64, Data []byte, PreBlockHash []byte) []byte {
	index := []byte(strconv.Itoa(Index))
	timestamp := []byte(strconv.FormatInt(Timestamp, 10))
	headers := bytes.Join([][]byte{PreBlockHash, Data, timestamp}, index)
	hash := sha256.Sum256(headers)
	return hash[:]
}


func PrintBlock(Index int){
	time := time.Unix(timestamp[Index], 0).Format("2006-01-02 15:04:05")
	fmt.Printf("index = %d, TimeStamp = %s\n", Index, time)
	fmt.Printf("Hash = %x\n", hash[Index])
	fmt.Printf("PreBlockHash = %x\n", preBlockHash[Index])
	fmt.Println()
}

func AddBlock(Data string) {
	fmt.Println("创建区块...")
	length := len(index)
	index = append(index, length)
	data = append(data, []byte(Data))
	timestamp = append(timestamp, time.Now().Unix())
	preBlockHash = append(preBlockHash, hash[length-1])
	hash = append(hash, GetHash(index[length], timestamp[length], data[length], preBlockHash[length]))
	fmt.Println("添加完成")
}

func Check() error{
	length := len(index)
	for i := range index{
		//1.检查索引是否正确
		if i != index[i]{
			fmt.Printf("索引错误，错误位置为%d\n", i)
			return fmt.Errorf("索引错误")
		}
		//2.检查hash是否正确
		if !bytes.Equal(hash[i], GetHash(i, timestamp[i], data[i], preBlockHash[i])){
			fmt.Printf("hash值错误，错误位置为%d\n", i)
			return fmt.Errorf("hash值错误")
		}
		//3.检查前后hash是否一致
		if i+1<length && !bytes.Equal(hash[i], preBlockHash[i+1]){
			fmt.Println("前后hash值不一致,错误位置为%d\n", index)
			return fmt.Errorf("前后hash值不一致")
		}
	}
	return nil
}

func  Find(Index int)  {
	if 0<=Index && Index<len(index){
		PrintBlock(Index)
	} else{
		fmt.Printf("不存在索引为%d 的区块\n", Index)
	}
}

func PrintBlocks(){
	for i:= range index{
		PrintBlock(i)
	}
}


//初始化区块链，创建第一个区块
func CreateBlockChain() {
	fmt.Println("初始化：建立第一个区块")
	index = append(index, 0)
	timestamp = append(timestamp, time.Now().Unix())
	data = append(data, []byte("The blockChain init"))
	preBlockHash = append(preBlockHash, []byte(""))
	hash = append(hash,GetHash(index[0], timestamp[0], data[0], preBlockHash[0]))
	fmt.Println("初始化完成")

}