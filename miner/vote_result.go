package miner

import (
	"github.com/ethereum/go-ethereum/core/types"
	"fmt"
	"github.com/ethereum/go-ethereum/core"
)

//type VoteContent struct {
//
//	//beatTime common.Hash
//	Address common.Address
//	Num int64
//}

//每n块获取计票,第n块整除块获取统计
func GetWitNessResult(n int, currentBlock *types.Block, currentChain *core.BlockChain) (uint64) {
	//当前区块高度与n的同余

	hm := currentBlock.Number().Uint64()
	fmt.Println(hm)
	////totalVote := []VoteContent{}
	////记票
	voten := 0  //总交易数
	totalNum := make(map[int]int)
	totalNum[0] = 0
	if (currentBlock.Number().Uint64()+1) % uint64(n) == 0  && hm != 0{
		for i:=0;i<n ;i++  {
			//得到块
			block := currentChain.GetBlockByNumber(hm)
			//得到块交易数据
			transactions := block.Body().Transactions
			if transactions != nil {
				//遍历交易切片取值，根据交易属性witness
				for _,v := range transactions{
					//if v.GetTranType() == ethapi.TransactionType0 {
					//取交易随机数结果
				//	index,err := strconv.Atoi(string(v.Data()))
					//if err != nil {
					//	fmt.Println("votetotal error")
					//} else {
					//	//得到所有投票
					//	//totalVote = append(totalVote,VoteContent{v.GetAddress(),int64(index)})
					//	totalNum[index] = totalNum[index]+1
					//	//}
					//}
					voten++
					fmt.Println(v)

				}
			}
			if hm == 0 {
				break
			}
			//遍历递减
			hm--
		}
	}
	//唱票数, 得到中奖数
	// keyluck  := 0
	//var valueluck = 0
	//for k,v := range totalNum{
	//	if v > valueluck {
	//		valueluck = v
	//		keyluck = k
	//	}
	//}
	//////得到中奖人地址
	////luckAddress := []common.Address{}
	////for _,vote := range totalVote{
	////	if keyluck == int(vote.Num) {
	////		luckAddress = append(luckAddress,vote.Address)
	////	}
	////}
	return uint64(voten)
}
