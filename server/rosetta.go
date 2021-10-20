package server

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/puneetsingh166/tm-load-test/codec"
	codectypes "github.com/puneetsingh166/tm-load-test/codec/types"
	"github.com/puneetsingh166/tm-load-test/server/rosetta"
)

// RosettaCommand builds the rosetta root command given
// a protocol buffers serializer/deserializer
func RosettaCommand(ir codectypes.InterfaceRegistry, cdc codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "rosetta",
		Short: "spin up a rosetta server",
		RunE: func(cmd *cobra.Command, args []string) error {
			conf, err := rosetta.FromFlags(cmd.Flags())
			if err != nil {
				return err
			}

			protoCodec, ok := cdc.(*codec.ProtoCodec)
			if !ok {
				return fmt.Errorf("exoected *codec.ProtoMarshaler, got: %T", cdc)
			}
			conf.WithCodec(ir, protoCodec)

			rosettaSrv, err := rosetta.ServerFromConfig(conf)
			if err != nil {
				return err
			}
			return rosettaSrv.Start()
		},
	}
	rosetta.SetFlags(cmd.Flags())

	return cmd
}
