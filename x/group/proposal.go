package group

import (
	"github.com/puneetsingh166/tm-load-test/codec/types"
	sdk "github.com/puneetsingh166/tm-load-test/types"
	"github.com/puneetsingh166/tm-load-test/types/tx"
)

func (p *Proposal) GetMsgs() []sdk.Msg {
	msgs, err := tx.GetMsgs(p.Msgs, "proposal")
	if err != nil {
		panic(err)
	}
	return msgs
}

func (p *Proposal) SetMsgs(msgs []sdk.Msg) error {
	anys, err := tx.SetMsgs(msgs)
	if err != nil {
		return err
	}
	p.Msgs = anys
	return nil
}

// UnpackInterfaces implements UnpackInterfacesMessage.UnpackInterfaces
func (p Proposal) UnpackInterfaces(unpacker types.AnyUnpacker) error {
	return tx.UnpackInterfaces(unpacker, p.Msgs)
}
