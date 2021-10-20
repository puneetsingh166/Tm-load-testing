package withtm
import (
    // "github.com/interchainio/tm-load-test/pkg/loadtest"

	banktypes "github.com/puneetsingh166/tm-load-test/x/bank/types"
    "github.com/puneetsingh166/tm-load-test/simapp"
    "github.com/puneetsingh166/tm-load-test/testutil/testdata"
    "github.com/puneetsingh166/tm-load-test/types"
    "fmt"
    // cryptotypes "github.com/puneetsingh166/tm-load-test/crypto/types"
	"github.com/puneetsingh166/tm-load-test/types/tx/signing"
	xauthsigning "github.com/puneetsingh166/tm-load-test/x/auth/signing"
    "github.com/puneetsingh166/tm-load-test/client/tx"
    
    
)

func Mainfunction()  {

      // following function.
      encCfg := simapp.MakeTestEncodingConfig()

      // Create a new TxBuilder.
      txBuilder := encCfg.TxConfig.NewTxBuilder()
    // --snip--

    priv1, _, addr1 := testdata.KeyTestPubAddr()
    // priv2, _, addr2 := testdata.KeyTestPubAddr()
     priv3, _, addr3 := testdata.KeyTestPubAddr()
    // fmt.Println(priv1)
    //  fmt.Println(priv2)
      fmt.Println(priv3)

    // Define two x/bank MsgSend messages:
    // - from addr1 to addr3,
    // - from addr2 to addr3.
    // This means that the transactions needs two signers: addr1 and addr2.
    msg1 := banktypes.NewMsgSend(addr1, addr3, types.NewCoins(types.NewInt64Coin("atom", 12)))
    // msg2 := banktypes.NewMsgSend(addr2, addr3, types.NewCoins(types.NewInt64Coin("atom", 34)))

    err := txBuilder.SetMsgs(msg1)
    if err != nil {
       fmt.Println("error")
    }

    txBuilder.SetGasLimit(20)
    txBuilder.SetFeeAmount(types.NewCoins(types.NewInt64Coin("atom", 12)))
    txBuilder.SetMemo("50stake")
    txBuilder.SetTimeoutHeight(50)


    //-------------


    // privs := []cryptotypes.PrivKey{priv1, priv2}
    // accNums:= []uint64{0} // The accounts' account numbers
    // accSeqs:= []uint64{0} // The accounts' sequence numbers

    // First round: we gather all the signer infos. We use the "set empty
    // signature" hack to do that.
    //  var sigsV2 []signing.SignatureV2
    //  for i, priv := range privs {
        sigV2 := signing.SignatureV2{
            PubKey: priv1.PubKey(),
            Data: &signing.SingleSignatureData{
                SignMode:  encCfg.TxConfig.SignModeHandler().DefaultMode(),
                Signature: nil,
            },
            Sequence: 0,
        }
     
        // fmt.Println("1234567")

        // fmt.Println("sigsV2", sigsV2)
        // fmt.Println("sigV2", sigV2)

        //  sigsV2 = append(sigsV2, sigV2)
    //  }
    // fmt.Println("nsssssssssssssssss")

    err = txBuilder.SetSignatures(sigV2)
    if err != nil {
       fmt.Println("error")
    }

    // Second round: all signer infos are set, so each signer can sign.
    // sigsV2 = []signing.SignatureV2{}
    //  for i, priv := range privs {
        signerData := xauthsigning.SignerData{
            ChainID:       "onomyd",
            AccountNumber: 0,
            Sequence:      0,
        }
        sigV2, err = tx.SignWithPrivKey(
            encCfg.TxConfig.SignModeHandler().DefaultMode(), signerData,
            txBuilder, priv1, encCfg.TxConfig, 0)
        if err != nil {
            fmt.Println("error")
        }

        //  sigsV2 = append(sigsV2, sigV2)
    //  }
    err = txBuilder.SetSignatures(sigV2)
    if err != nil {
        fmt.Println("error")
    }

txBytes, err := encCfg.TxConfig.TxEncoder()(txBuilder.GetTx())
if err != nil {
    fmt.Println("na")
}

fmt.Println("fmt",txBytes)
}

