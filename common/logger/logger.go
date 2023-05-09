package logger

import (
	"io"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Log *zap.SugaredLogger

// InitLogger 初始化Logger
func InitLogger(cfg *Config) error {
	writeSyncer, err := getWriter(cfg.Filename)
	if err != nil {
		return err
	}
	encoder := getEncoder()
	var l = new(zapcore.Level)
	err = l.UnmarshalText([]byte(cfg.Level))
	if err != nil {
		return err
	}

	w := zapcore.NewMultiWriteSyncer(zapcore.AddSync(writeSyncer))
	core := zapcore.NewCore(encoder, w, l)

	log := zap.New(core, zap.AddCaller())
	zap.ReplaceGlobals(log)
	Log = log.Sugar()
	return nil
}

func getWriter(filename string) (io.Writer, error) {
	hook, err := rotatelogs.New(
		filename+".%Y-%m-%d",
		rotatelogs.WithLinkName(filename),
		rotatelogs.WithMaxAge(time.Hour*24*7),
		rotatelogs.WithRotationTime(time.Hour*24),
	)

	if err != nil {

		return nil, err
	}
	return hook, nil
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05")
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	return zapcore.NewJSONEncoder(encoderConfig)
}
