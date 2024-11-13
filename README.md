# go-logger

An example of how to implement a go structure logging using [log/slog](https://go.dev/blog/slog) standard
library and **heavily** inspired by [google/exposure-notifications-server](https://github.com/google/exposure-notifications-server/blob/main/pkg/logging/logger.go) applicaton.

## Example

Visit [main.go](./cmd/app/main.go) file for more details.

```go
func main() {
    logger := logging.NewLogger("dev", "debug")
    logger.Info("Hello from dev!")

    logger = logging.NewLogger("prod", "info")
    logger.Info("Hello from prod!")

    logger = logging.NewLoggerFromEnv()
    logger.Info("Hello from env mode!")

    logger = logging.NewLoggerFromEnv().With("app", "showcase")

    ctx := context.Background()
    ctx = logging.LoggerWithContext(ctx, logger)

    if err := start(ctx); err != nil {
        logger.Error(err.Error())
        os.Exit(2)
    }
}

func start(ctx context.Context) error {
    logger := logging.LoggerFromContext(ctx)
    logger.Info("calling from start")
    return nil
}
```

## License

[MIT License](./LICENSE)
