package internal

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/goexl/gox"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Zap struct {
	zap *zap.Logger
}

func New() (executor *Zap, err error) {
	config := zap.NewProductionConfig()
	config.EncoderConfig = zap.NewProductionEncoderConfig()
	config.EncoderConfig.TimeKey = "timestamp"
	config.EncoderConfig.EncodeTime = zapcore.RFC3339TimeEncoder
	config.Level = zap.NewAtomicLevelAt(zapcore.DebugLevel) // !确保在最低日志级别，由上层代码处理日志级别

	executor = new(Zap)
	executor.zap, err = config.Build(zap.WithCaller(false) /*不打印调用链路，由上层代码处理*/)

	return
}

func (e *Zap) Debug(msg string, required gox.Field[any], fields ...gox.Field[any]) {
	e.zap.Debug(msg, e.parse(required, fields...)...)
}

func (e *Zap) Info(msg string, required gox.Field[any], fields ...gox.Field[any]) {
	e.zap.Info(msg, e.parse(required, fields...)...)
}

func (e *Zap) Warn(msg string, required gox.Field[any], fields ...gox.Field[any]) {
	e.zap.Warn(msg, e.parse(required, fields...)...)
}

func (e *Zap) Error(msg string, required gox.Field[any], fields ...gox.Field[any]) {
	e.zap.Error(msg, e.parse(required, fields...)...)
}

func (e *Zap) Panic(msg string, required gox.Field[any], fields ...gox.Field[any]) {
	e.zap.Panic(msg, e.parse(required, fields...)...)
}

func (e *Zap) Fatal(msg string, required gox.Field[any], fields ...gox.Field[any]) {
	e.zap.Fatal(msg, e.parse(required, fields...)...)
}

func (e *Zap) Sync() error {
	return e.zap.Sync()
}

func (e *Zap) parse(required gox.Field[any], optionals ...gox.Field[any]) (parsed []zap.Field) {
	parsed = make([]zap.Field, 0, len(optionals)+1)
	for _, field := range append([]gox.Field[any]{required}, optionals...) {
		if "" == field.Key() || nil == field.Value() {
			continue
		}

		switch value := field.Value().(type) {
		case bool:
			parsed = append(parsed, zap.Bool(field.Key(), value))
		case *bool:
			parsed = append(parsed, zap.Boolp(field.Key(), value))
		case []bool:
			parsed = append(parsed, zap.Bools(field.Key(), value))
		case *[]bool:
			parsed = append(parsed, zap.Bools(field.Key(), *value))
		case int8:
			parsed = append(parsed, zap.Int8(field.Key(), value))
		case *int8:
			parsed = append(parsed, zap.Int8p(field.Key(), value))
		case int:
			parsed = append(parsed, zap.Int(field.Key(), value))
		case *int:
			parsed = append(parsed, zap.Intp(field.Key(), value))
		case []int:
			parsed = append(parsed, zap.Ints(field.Key(), value))
		case *[]int:
			parsed = append(parsed, zap.Ints(field.Key(), *value))
		case uint:
			parsed = append(parsed, zap.Uint(field.Key(), value))
		case *uint:
			parsed = append(parsed, zap.Uintp(field.Key(), value))
		case []uint:
			parsed = append(parsed, zap.Uints(field.Key(), value))
		case *[]uint:
			parsed = append(parsed, zap.Uints(field.Key(), *value))
		case time.Duration:
			parsed = append(parsed, zap.Duration(field.Key(), value))
		case *time.Duration:
			parsed = append(parsed, zap.Durationp(field.Key(), value))
		case int64:
			parsed = append(parsed, zap.Int64(field.Key(), value))
		case *int64:
			parsed = append(parsed, zap.Int64p(field.Key(), value))
		case []int64:
			parsed = append(parsed, zap.Int64s(field.Key(), value))
		case *[]int64:
			parsed = append(parsed, zap.Int64s(field.Key(), *value))
		case float32:
			parsed = append(parsed, zap.Float32(field.Key(), value))
		case *float32:
			parsed = append(parsed, zap.Float32p(field.Key(), value))
		case float64:
			parsed = append(parsed, zap.Float64(field.Key(), value))
		case *float64:
			parsed = append(parsed, zap.Float64p(field.Key(), value))
		case []float64:
			parsed = append(parsed, zap.Float64s(field.Key(), value))
		case *[]float64:
			parsed = append(parsed, zap.Float64s(field.Key(), *value))
		case *string:
			parsed = append(parsed, zap.Stringp(field.Key(), value))
		case []string:
			parsed = append(parsed, zap.Strings(field.Key(), value))
		case *[]string:
			parsed = append(parsed, zap.Strings(field.Key(), *value))
		case time.Time:
			parsed = append(parsed, zap.Time(field.Key(), value))
		case *time.Time:
			parsed = append(parsed, zap.Timep(field.Key(), value))
		case []time.Time:
			parsed = append(parsed, zap.Times(field.Key(), value))
		case []time.Duration:
			parsed = append(parsed, zap.Durations(field.Key(), value))
		case json.Marshaler, []json.Marshaler: // !不要调换顺序
			parsed = append(parsed, zap.Reflect(field.Key(), field.Value()))
		case fmt.Stringer:
			parsed = append(parsed, zap.Stringer(field.Key(), value))
		case []fmt.Stringer:
			parsed = append(parsed, zap.Stringers(field.Key(), value))
		case error:
			parsed = append(parsed, zap.Error(value))
		default:
			parsed = append(parsed, zap.Any(field.Key(), field.Value()))
		}
	}

	return
}
