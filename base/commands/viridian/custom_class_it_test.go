package viridian_test

import (
	"context"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/hazelcast/hazelcast-commandline-client/internal/check"
	"github.com/hazelcast/hazelcast-commandline-client/internal/it"
)

func customClass_NonInteractiveTest(t *testing.T) {
	viridianTester(t, func(ctx context.Context, tcx it.TestContext) {
		// setup
		f := "foo.zip"
		fd := "testdata/" + f
		c := createOrGetClusterWithState(ctx, tcx, "RUNNING")
		// test upload custom class
		tcx.WithReset(func() {
			tcx.CLCExecute(ctx, "viridian", "upload-custom-class", c.ID, fd)
			tcx.AssertStderrContains("OK")
			check.Must(waitCustomClassOperation(ctx, tcx, "Custom class uploaded successfully."))
		})
		id := ""
		// test list custom class
		tcx.WithReset(func() {
			tcx.CLCExecute(ctx, "viridian", "list-custom-classes", c.ID)
			tcx.AssertStderrContains("OK")
			id = customClassID(tcx.ExpectStdout.String())
			tcx.AssertStdoutContains(f)
		})
		// test download custom class
		tcx.WithReset(func() {
			tcx.CLCExecute(ctx, "viridian", "download-custom-class", c.ID, f)
			tcx.AssertStderrContains("OK")
			tcx.AssertStdoutContains("Custom class downloaded successfully.")
		})
		// test delete custom class
		tcx.WithReset(func() {
			check.Must(waitState(ctx, tcx, c.ID, "RUNNING"))
			tcx.CLCExecute(ctx, "viridian", "delete-custom-class", c.ID, id)
			check.Must(waitCustomClassOperation(ctx, tcx, "Custom class deleted successfully."))
			tcx.AssertStderrContains("OK")
		})
		// check the list output again to be sure that delete was really successful
		tcx.WithReset(func() {
			tcx.CLCExecute(ctx, "viridian", "list-custom-classes", c.ID)
			tcx.AssertStderrContains("OK")
			tcx.AssertStderrNotContains(f)
		})
	})
}

func waitCustomClassOperation(ctx context.Context, tcx it.TestContext, expected string) error {
	tryCount := 0
	for {
		if tryCount == 5 {
			return fmt.Errorf("custom class operation exceeded try limit")
		}
		if ctx.Err() != nil {
			return ctx.Err()
		}
		if strings.Contains(tcx.ExpectStdout.String(), expected) {
			return nil
		}
		tryCount++
		time.Sleep(5 * time.Second)
	}
}

func customClassID(l string) string {
	return strings.Split(l, "\t")[0]
}
