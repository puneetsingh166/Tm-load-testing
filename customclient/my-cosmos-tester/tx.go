package main
import (
     banktypes "github.com/puneetsingh166/tm-load-test/x/bank/types"
     "github.com/puneetsingh166/tm-load-test/simapp"
     "github.com/puneetsingh166/tm-load-test/types"
     "fmt"
     "github.com/puneetsingh166/tm-load-test/types/tx/signing"
     xauthsigning "github.com/puneetsingh166/tm-load-test/x/auth/signing"
     "github.com/puneetsingh166/tm-load-test/client/tx"
     cryptotypes "github.com/puneetsingh166/tm-load-test/crypto/types"
     "github.com/puneetsingh166/tm-load-test/testutil/testdata"
)

func GenTX() ([]byte){
  encCfg := simapp.MakeTestEncodingConfig()
  
  // Create a new TxBuilder.
  
  txBuilder := encCfg.TxConfig.NewTxBuilder()
  priv1, _, addr1 := testdata.KeyTestPubAddr()
  _, _, addr2 := testdata.KeyTestPubAddr()
  msg1 := banktypes.NewMsgSend(addr1, addr2, types.NewCoins(types.NewInt64Coin("footoken", 1)))
  err := txBuilder.SetMsgs(msg1)
  if err != nil {
	  fmt.Println("error")
  }

  txBuilder.SetGasLimit(200000)
  txBuilder.SetFeeAmount(types.NewCoins(types.NewInt64Coin("footoken", 5)))
  txBuilder.SetMemo("footoken")
  txBuilder.SetTimeoutHeight(5)


  //------------------------------- ------Signing a Transaction---------------------------------//
  
  privs := []cryptotypes.PrivKey{priv1}
  var sigsV2 []signing.SignatureV2
   for _, priv := range privs {
	  sigV2 := signing.SignatureV2{
		  PubKey: priv.PubKey(),
		  Data: &signing.SingleSignatureData{
			  SignMode:  encCfg.TxConfig.SignModeHandler().DefaultMode(),
			  Signature: nil,
		  },
		  
		  Sequence: 0,
	  }
	   sigsV2 = append(sigsV2, sigV2)
  }
  err = txBuilder.SetSignatures(sigsV2...)
  if err != nil {
	  fmt.Println("error - ",err)
  }
  sigsV2 = []signing.SignatureV2{}

  for _, priv := range privs {
	  signerData := xauthsigning.SignerData{
		  ChainID:       "onomy",
		  AccountNumber: 0,
		  Sequence:      0,
	  }
	  sigV2, err := tx.SignWithPrivKey(
		  encCfg.TxConfig.SignModeHandler().DefaultMode(), signerData,
		  txBuilder, priv, encCfg.TxConfig, 0)
		  
	if err != nil {
		  fmt.Println("error - ",err)
	  }
	  sigsV2 = append(sigsV2, sigV2)
  }
  
  err = txBuilder.SetSignatures(sigsV2...)
  if err != nil {
	  fmt.Println("error - ",err)
  }

  txBytes, err := encCfg.TxConfig.TxEncoder()(txBuilder.GetTx())
  if err != nil {
    fmt.Println("error - ",err)
  }
  return txBytes
}
