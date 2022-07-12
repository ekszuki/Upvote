package main

import (
	"context"

	"klever.io/interview/pkg/infra/startup"
)

func main() {
	ctx := context.Background()

	st := startup.NewStartup(ctx)
	st.Initialize()
}
