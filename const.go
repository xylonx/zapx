package zapx

import (
	"os"
	"runtime"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	goVersion       string
	hostName        string
	buildAppVersion string
	buildUser       string
	buildHost       string
	buildTime       string
)

var levelMap = map[string]zapcore.Level{
	"debug":  zapcore.DebugLevel,
	"info":   zapcore.InfoLevel,
	"warn":   zapcore.WarnLevel,
	"error":  zapcore.ErrorLevel,
	"dpanic": zapcore.DPanicLevel,
	"panic":  zapcore.PanicLevel,
	"fatal":  zapcore.FatalLevel,
}

func init() {
	name, err := os.Hostname()
	if err != nil {
		name = "unknown"
	}
	hostName = name
	goVersion = runtime.Version()
}

func getLoggerLevel(lvl string) zapcore.Level {
	if level, ok := levelMap[lvl]; ok {
		return level
	}
	return zapcore.InfoLevel
}

func defaultFields() []zap.Field {
	var fields []zap.Field
	projectName := os.Getenv("ELASTIC_APM_SERVICE_NAME")
	if len(projectName) == 0 {
		projectName = os.Args[0]
		projectName = strings.Trim(projectName, "./")
	}
	if len(projectName) == 0 {
		projectName = "empty-project-name"
	}
	fields = append(fields, zap.String("application.name", projectName))
	if buildAppVersion != "" {
		fields = append(fields, zap.String("version", buildAppVersion))
	}
	if buildTime != "" {
		fields = append(fields, zap.String("build.time", buildTime))
	}
	if buildHost != "" {
		fields = append(fields, zap.String("build.host", buildHost))
	}
	if buildUser != "" {
		fields = append(fields, zap.String("build.user", buildUser))
	}
	if hostName != "" {
		fields = append(fields, zap.String("hostname", hostName))
	}
	if goVersion != "" {
		fields = append(fields, zap.String("go.version", goVersion))
	}
	return fields
}
