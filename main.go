// CocosGoSdkDemo project main.go
package main

import (
	sdk "CocosSDK"
	"fmt"
)

func main() {
	//初始化sdk
	sdk.InitSDK("123.56.98.47", 80, false)
	//导入私钥
	//sdk.Wallet.AddAccountByPrivateKey("5KSUYbLbDFaEeFEKRjHSW1ZZDu6a4QXf6nTwMvFEzCVeXuJGnWh", "123456")
	//sdk.Wallet.AddAccountByPrivateKey("5KSUYbLbDFaEeFEKRjHSW1ZZDu6a4QXf6nTwMvFEzCVeXuJGnWh", "123456")
	sdk.Wallet.ImportAccount("gggg2", "12345678")
	return
	//钱包另存为
	sdk.Wallet.SaveAs("./wallets.dat")
	//钱包导入
	sdk.Wallet.LoadWallet("./wallets.dat")
	//设置默认账户
	sdk.Wallet.SetDefaultAccount("gggg2", "12345678")
	//升级终身账户
	sdk.Wallet.UpgradeAccount("gggg2")
	//注册开发者
	sdk.Wallet.RegisterNhAssetCreator("gggg2")
	//创建新账户
	sdk.Wallet.CreateAccount("gggg10", "12345678")
	//创建代币
	sdk.CreateToken("CCS", 10000, 5)
	//更新代币
	sdk.UpdateToken("CCS", 10000, 6)
	//发行代币
	sdk.IssueToken("CCS", "gggg2", 1)
	//转账代币
	sdk.Wallet.Transfer("cocos1025", "C0C0S", 1)
	//销毁代币
	sdk.ReserveToken("C0C0S", 1)
	//注资手续费池
	sdk.TokenFundFeePool("C0C0S", 100)

	//创建世界观
	sdk.CreateWorldView("MyWorld")
	//提议关联世界观
	sdk.RelateWorldView("block_chain")
	//创建NH资产
	sdk.CreateNhAsset("C0C0S", "MyWorld", "cocos1024", `{"name":"航空母舰","version":"船新版本"}`)
	//NH资产转账
	sdk.TransferNhAsset("gggg3", "4.2.77730")
	//NH资产删除
	sdk.DeleteNhAsset("4.2.77730")
	//卖出NH资产
	sdk.SellNhAsset("gggg3", "4.2.77731", "物美价廉快来买!", "1.3.0", "1.3.0", 5, 1000)
	//取消订单
	sdk.CancelNhAssetOrder("4.3.934")
	//吃单,买入
	sdk.FillNhAsset("4.3.924")
	//部署合约
	sdk.CreateContractByFile("contract.cocos1024", sdk.Wallet.CreateKey().GetPublicKey().ToBase58String(), "./test.lua")
	//调用合约
	sdk.InvokeContract("contract.cocos1024", "hello", 1, 2)
	//质押gas
	err := sdk.Pledgegas("gggg2", 100)
	fmt.Println(err)
	//赎回
	err = sdk.Pledgegas("gggg2", 0)
	fmt.Println(err)
	//投票
	err = sdk.Vote("1.5.6", 100)
	fmt.Println(err)
	//投票赎回
	err = sdk.Vote("1.5.6", 0)
	fmt.Println(err)
	//查询可领取的冻结资产
	sdk.GetVestingBalances()
	//领取冻结资产
	err = sdk.WithdrawVestingBalance("1.13.30")
	fmt.Println(err)
}
