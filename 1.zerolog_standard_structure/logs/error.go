package logs

import (
	// "fmt"
	"os"
	"runtime/debug"
	"time"

	"github.com/rs/zerolog"
	// "github.com/rs/zerolog/log"
)

var buildInfo *debug.BuildInfo
var logger zerolog.Logger

// TRACE (-1): for tracing the code execution path.
// DEBUG (0): messages useful for troubleshooting the program.
// INFO (1): messages describing the normal operation of an application.
// WARNING (2): for logging events that need may need to be checked later.
// ERROR (3): error messages for a specific operation.
// FATAL (4): severe errors where the application cannot recover. os.Exit(1) is called after the message is logged.
// PANIC (5): similar to FATAL, but panic() is called instead.

func init() {

	buildInfo, _ = debug.ReadBuildInfo()

	logger = zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339}).
		Level(zerolog.TraceLevel).
		With().
		Timestamp().
		Caller().
		Int("pid", os.Getpid()).
		Str("go_version", buildInfo.GoVersion).
		Logger()
}

func Error(s string, err error) {

	if err != nil {

		logger.Error().Msg("❌🔥 Error message :" + s + err.Error())
	}
}

func Info(s string) {
	if len(s) != 0 {

		logger.Info().Msg("📢 Info message :" + s)
	}
}

func Debug(s string) {
	if len(s) != 0 {

		logger.Debug().Msg("🐞 Debug message :" + s)
	}
}

func Warn(s string) {
	if len(s) != 0 {

		logger.Warn().Msg("⚠️ Warn message :" + s)
	}
}

func Fatal(s string) {
	if len(s) != 0 {

		logger.WithLevel(zerolog.FatalLevel).Msg("💀 Fatal message :" + s)
	}
}

func Panic(s string) {
	if len(s) != 0 {
		logger.WithLevel(zerolog.PanicLevel).Msg("🔥 Panic message :" + s)

	}
}

func Trace(s string) {
	if len(s) != 0 {
		logger.Trace().Msg("🔎 Trace message :" + s)
	}
}



// output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}
// output.FormatLevel = func(i interface{}) string {
//     return strings.ToUpper(fmt.Sprintf("| %-6s|", i))
// }
// output.FormatMessage = func(i interface{}) string {
//     return fmt.Sprintf("***%s****", i)
// }
// output.FormatFieldName = func(i interface{}) string {
//     return fmt.Sprintf("%s:", i)
// }
// output.FormatFieldValue = func(i interface{}) string {
//     return strings.ToUpper(fmt.Sprintf("%s", i))
// }

// log := zerolog.New(output).With().Timestamp().Logger()

// log.Info().Str("foo", "bar").Msg("Hello World")

// Output: 2006-01-02T15:04:05Z07:00 | INFO  | ***Hello World**** foo:BAR