package struct_block

import (
	"fmt"
	"github.com/szulx0/block/struct_block/block"
)

func main() {
	flag := 0
	index := 0
	data := ""
	bc := block.CreateBlockChain()
	for {
		fmt.Println("请输入命令: 1:添加块 2:查找块  3:输出区块链 ")
		fmt.Scanf("%d", &flag)
		switch flag {
		case 1:
			fmt.Println("请输入进行的操作(如发生了转账):")
			fmt.Scanf("%s", &data)
			bc.AddBlock(data)
			fmt.Println("检查索引、当前块哈希是否计算正确、前后块哈希是否一致")
			err := bc.Check()
			if err != nil {
				break
			}
			fmt.Println("检查完成，没有出现错误")
			bc.PrintBlocks()
			flag = 0
		case 2:
			fmt.Println("请输入查找块的索引: ")
			fmt.Scanf("%d", &index)
			bc.Find(index).PrintBlock()
			flag = 0
		case 3:
			bc.PrintBlocks()
			flag = 0
		default:
			fmt.Println("请输入1-3之间的数字")
			flag = 0
		}
	}
}
