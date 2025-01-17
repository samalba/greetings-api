package tasks

import (
	"context"
	"fmt"

	"go.dagger.io/dagger/sdk/go/dagger"
)

func Test(ctx context.Context) error {
	client, err := dagger.Connect(ctx)
	if err != nil {
		return err
	}
	defer client.Close()

	builder, err := goBuilder(
		client,
		ctx,
		[]string{"go", "test"},
	)
	if err != nil {
		return err
	}

	// Get Command Output
	out, err := builder.Stdout().Contents(ctx)
	if err != nil {
		return err
	}

	fmt.Println(out)

	return nil
}
