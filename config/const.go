package config

const (
	ServiceName   = "Avalon"
	Version       = "0.1.0"
	RequestID     = "request_id"
	Authorization = "Authorization"

	envServerAddress = "SERVER_ADDRESS"
	envSecretKey     = "SECRET_KEY"

	// Headers
	HeaderRequestID = "X-Request-ID"

	// Database EnvVar
	envDatabaseHost        = "DATABASE_HOST"
	envDatabasePort        = "DATABASE_PORT"
	envDatabaseName        = "DATABASE_NAME"
	envDatabaseUsername    = "DATABASE_USERNAME"
	envDatabasePassword    = "DATABASE_PASSWORD"
	envDatabaseEnableLog   = "DATABASE_ENABLE_LOG"
	envDatabaseMaxOpenConn = "DATABASE_MAX_OPEN_CONN"
	envDatabaseMaxIdleConn = "DATABASE_MAX_IDLE_CONN"

	envLoggerFormat = "LOGGER_FORMAT"
	envLogLevel     = "LOG_LEVEL"
)
