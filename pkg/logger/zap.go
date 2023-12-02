package logger

import (
	"github.com/natefinch/lumberjack"
	"go-project-frame/server/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"time"
)

// getZapCore 创建一个zapcore.Core，负责处理日志记录器的核心功能
func getZapCore(cfgLog config.LogOptions) (zapcore.Core, error) {
	// 获取日志级别
	level, err := getLevel(cfgLog)
	if err != nil {
		return nil, err
	}
	// 指定日志将写到哪里去
	writeSyncer := getLogWriter(cfgLog)
	if err != nil {
		return nil, err
	}
	// 获取日志编码
	encoder := getEncoder()

	// debug级别，需要额外输出一份到控制台
	if cfgLog.Level == "debug" {
		// 获取一个控制台的编码器
		consoleEncoder := getConsoleEncoder()
		// 使用NewTee将多个core合并到core
		return zapcore.NewTee(
			// 日志正常记录到文件中去
			zapcore.NewCore(encoder, writeSyncer, level),
			// 将debug级别的日志，也输出到控制台一份
			zapcore.NewCore(consoleEncoder, zapcore.Lock(os.Stdout), zapcore.DebugLevel),
		), nil
	}
	return zapcore.NewCore(encoder, writeSyncer, level), nil
}

// getLevel 指定记录的日志级别
func getLevel(cfgLog config.LogOptions) (zapcore.Level, error) {
	level := new(zapcore.Level)
	if err := level.UnmarshalText([]byte(cfgLog.Level)); err != nil {
		return 0, err
	}
	return *level, nil
}

// getEncoder 编码器，指示如何写入日志
func getEncoder() zapcore.Encoder {
	/*
	 * zap.NewProductionEncoderConfig()
	 * - 返回一个预定义好的Encoder配置，输出格式示例：{"level":"debug","ts":1572160754.994731,"msg":"Trying to hit GET request for www.sogo.com"}
	 * - 可以看出，存在两个问题：1）时间"ts":1572160754.994731，不是人类可读方式；2）没有函数调用方法+行号等信息
	 */
	encoderConfig := zap.NewProductionEncoderConfig()
	// 修改时间编码器，并指定time的key，不使用默认的ts
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.TimeKey = "time"
	// 在日志文件中使用大写字母记录日志级别
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	// 配置日志中的时间持续时间，以秒为单位进行编码
	encoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder
	// 对调用者信息进行简短的编码
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	// 创建一个 JSON 格式的日志编码器
	return zapcore.NewJSONEncoder(encoderConfig)
}

// getConsoleEncoder 创建一个基于控制台输出的日志编码器，指示如何写入日志
func getConsoleEncoder() zapcore.Encoder {
	// zap.NewDevelopmentEncoderConfig() 创建一个适合开发环境使用的日志编码器配置，会输出一些额外的信息，例如调用者信息、堆栈跟踪等
	encoderConfig := zap.NewDevelopmentEncoderConfig()
	// 使用自定义的时间编码器
	encoderConfig.EncodeTime = getCustomEncodeTime
	// 在日志文件中使用 小写字母+带颜色 记录日志级别
	encoderConfig.EncodeLevel = zapcore.LowercaseColorLevelEncoder
	// 创建一个基于控制台输出的日志编码器
	return zapcore.NewConsoleEncoder(encoderConfig)
}

// getLogWriter 指定日志将写到哪里去
func getLogWriter(cfgLog config.LogOptions) zapcore.WriteSyncer {
	// Zap本身不支持切割归档日志文件，可以使用 Lumberjack 进行日志切割归档，目前只支持按文件大小切割，原因是按时间切割效率低且不能保证日志数据不被破坏
	// 想按日期切割可以使用github.com/lestrrat-go/file-rotatelogs这个库，虽然目前不维护了，但也够用了。
	lumberjackLogger := &lumberjack.Logger{
		// 日志输出文件路径
		Filename: cfgLog.Filename,
		// 日志文件最大大小(MB)
		MaxSize: cfgLog.MaxSize,
		// 保留旧日志文件的最大天数
		MaxBackups: cfgLog.MaxBackups,
		// 最大保留日志个数
		MaxAge: cfgLog.MaxAge,
	}
	// zapcore.AddSync，将添加到 Zap 日志库的核心同步器中
	return zapcore.AddSync(lumberjackLogger)
}

// getCustomEncodeTime 自定义时间编码器，该函数的参数部分，符合 Zap 日志库中的时间编码器接口要求
func getCustomEncodeTime(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	// 自定义输出格式
	enc.AppendString("[kubemanage] " + t.Format("2006/01/02 - 15:04:05.000"))
}
