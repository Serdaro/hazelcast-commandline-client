package plug

import (
	"context"
	"io"

	"github.com/hazelcast/hazelcast-go-client"

	"github.com/hazelcast/hazelcast-commandline-client/internal/log"
	"github.com/hazelcast/hazelcast-commandline-client/internal/output"
)

type CommandContext interface {
	AddStringFlag(long, short, value string, required bool, help string)
	AddBoolFlag(long, short string, value bool, required bool, help string)
	Interactive() bool
}

type ExecContext interface {
	Logger() log.Logger
	Stdout() io.Writer
	Stderr() io.Writer
	Args() []string
	Props() ReadOnlyProperties
	Client(ctx context.Context) (*hazelcast.Client, error)
	Interactive() bool
	AddOutputRows(row ...output.Row)
}