1. If we use context.WithTimeout, then we call ctx.Done(), it will return channel that will be closing in timeout defined.
    - So if we context.WithTimeout(context.Background(), 1*time.Second), then do <-ctx.Done(), then this will wait until 1 seconds, only then the execution continues.
    - ctx.Done() will return channel that will be closed when a job is done (job is to be cancelled)