package _map

import (
	"context"

	"github.com/hazelcast/hazelcast-commandline-client/internal/output"
	"github.com/hazelcast/hazelcast-commandline-client/internal/plug"
	"github.com/hazelcast/hazelcast-commandline-client/internal/proto/codec"

	. "github.com/hazelcast/hazelcast-commandline-client/internal/check"
)

type MapEntrySetCommand struct{}

func (mc *MapEntrySetCommand) Init(cc plug.InitContext) error {
	help := "Get all entries of an IMap"
	cc.SetCommandHelp(help, help)
	return nil
}

func (mc *MapEntrySetCommand) Exec(ec plug.ExecContext) error {
	ctx := context.TODO()
	mapName := ec.Props().GetString(mapFlagName)
	showType := ec.Props().GetBool(mapFlagShowType)
	ci, err := ec.ClientInternal(ctx)
	if err != nil {
		return err
	}
	req := codec.EncodeMapEntrySetRequest(mapName)
	resp, err := ci.InvokeOnRandomTarget(ctx, req, nil)
	if err != nil {
		return err
	}
	pairs := codec.DecodeMapEntrySetResponse(resp)
	rows := output.DecodePairs(ci, pairs, showType)
	ec.AddOutputRows(rows...)
	return nil
}

func init() {
	Must(plug.Registry.RegisterCommand("map:entry-set", &MapEntrySetCommand{}))
}