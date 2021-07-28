package log

import (
	"os"

	"github.com/huprince/operate-mysql-tool/config"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Logger zap 日志
var Logger *zap.Logger

func init() {
	Logger = InitLogger()
}

// InitLogger 自定义日志记录器
func InitLogger() *zap.Logger {
	//appName := config.GetEnv().AppName
	debug := config.GetEnv().Debug
	hook := lumberjack.Logger{
		Filename: config.GetEnv().LogPath,
		MaxSize: config.GetEnv().LogMaxSize,
		MaxAge: config.GetEnv().LogMaxAge,
		MaxBackups: config.GetEnv().LogMaxBackups,
		Compress: config.GetEnv().LogIsCompress,
	}
	encodeConfig := zapcore.EncoderConfig{
		MessageKey: "msg",
		LevelKey: "level",
		TimeKey: "time",
		NameKey: "logger",
		CallerKey: "file",
		StacktraceKey: "stacktrace",
		LineEnding: zapcore.DefaultLineEnding,
		EncodeLevel: zapcore.LowercaseLevelEncoder,
		EncodeTime: zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller: zapcore.ShortCallerEncoder,
		EncodeName: zapcore.FullNameEncoder,
	}
	atomicLevel := zap.NewAtomicLevel()
	atomicLevel.SetLevel(zap.InfoLevel)
	writes := []zapcore.WriteSyncer{zapcore.AddSync(&hook)}
	if debug {
		writes = append(writes, zapcore.AddSync(os.Stdout))
	}
	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(encodeConfig),
		zapcore.NewMultiWriteSyncer(writes...),
		atomicLevel,
	)
	caller := zap.AddCaller()
	development := zap.Development()
	//field := zap.Fields(zap.String("appName", appName))
	return zap.New(core, caller, development)
}