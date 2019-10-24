// CocosGoSdkDemo project main.go
package main

import (
	sdk "cocos-go-sdk"
)

func main() {
	//初始化sdk
	sdk.InitSDK("47.93.62.96", 8049, false)
	//导入私钥
	sdk.Wallet.AddAccountByPrivateKey("5KSUYbLbDFaEeFEKRjHSW1ZZDu6a4QXf6nTwMvFEzCVeXuJGnWh", "123456")
	sdk.Wallet.AddAccountByPrivateKey("5KSUYbLbDFaEeFEKRjHSW1ZZDu6a4QXf6nTwMvFEzCVeXuJGnWh", "123456")
	//钱包另存为
	sdk.Wallet.SaveAs("./wallets.dat")
	//钱包导入
	sdk.Wallet.LoadWallet("./wallets.dat")
	//设置默认账户
	sdk.Wallet.SetDefaultAccount("cocos1024", "123456")
	//升级终身账户
	sdk.Wallet.UpgradeAccount("cocos1024")
	//注册开发者
	sdk.Wallet.RegisterNhAssetCreator("cocos1024")
	//创建新账户
	sdk.Wallet.CreateAccount("cocos1025", "123456")
	//创建代币
	sdk.CreateToken("C1C1S", "1.3.0", "1.3.1", 10000, 1, 5, 1)
	//更新代币
	sdk.UpdateToken("C1C1S", "1.3.60", "1.3.0", 10000, 2, 50, 1)
	//发行代币
	sdk.IssueToken("C0C0S", "cocos1024", 1)
	//转账代币
	sdk.Wallet.Transfer("cocos1025", "C0C0S", "0000", 1)
	//销毁代币
	sdk.ReserveToken("C0C0S", 1)
	//注资手续费池
	sdk.TokenFundFeePool("C0C0S", 100)
	//创建Vesting Balance
	sdk.CreateVestingBalance("C0C0S", 100)
	//创建世界观
	sdk.CreateWorldView("MyWorld")
	//提议关联世界观
	sdk.RelateWorldView("block_chain")
	//创建NH资产
	sdk.CreateNhAsset("C0C0S", "MyWorld", "cocos1024", `{"name":"航空母舰","version":"船新版本"}`)
	//NH资产转账
	sdk.TransferNhAsset("cocos1025", "4.2.77730")
	//NH资产删除
	sdk.DeleteNhAsset("4.2.77730")
	//卖出NH资产
	sdk.SellNhAsset("cocos1025", "4.2.77731", "物美价廉快来买!", "1.3.0", "1.3.0", 5, 1000)
	//取消订单
	sdk.CancelNhAssetOrder("4.3.934")
	//吃单,买入
	sdk.FillNhAsset("4.3.924")
	//部署合约
	sdk.CreateContractByFile("contract.cocos1024", sdk.Wallet.CreateKey().GetPublicKey().ToBase58String(), "./test.lua")
	//调用合约
	sdk.InvokeContract("contract.cocos1024", "hello", 1, 2)
}
