package go_logger

import (
	"log"
	"os"
	"reflect"
	"testing"
)

func Test_convertLevelToInt(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertLevelToInt(tt.args.s); got != tt.want {
				t.Errorf("convertLevelToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLogger_Debug(t *testing.T) {
	type fields struct {
		logger *log.Logger
		level  string
	}
	type args struct {
		o interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &Logger{
				logger: tt.fields.logger,
				level:  tt.fields.level,
			}
			l.Debug(tt.args.o)
		})
	}
}

func TestLogger_Info(t *testing.T) {
	type fields struct {
		logger *log.Logger
		level  string
	}
	type args struct {
		o interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &Logger{
				logger: tt.fields.logger,
				level:  tt.fields.level,
			}
			l.Info(tt.args.o)
		})
	}
}

func TestLogger_Warning(t *testing.T) {
	type fields struct {
		logger *log.Logger
		level  string
	}
	type args struct {
		o interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &Logger{
				logger: tt.fields.logger,
				level:  tt.fields.level,
			}
			l.Warning(tt.args.o)
		})
	}
}

func TestLogger_Error(t *testing.T) {
	type fields struct {
		logger *log.Logger
		level  string
	}
	type args struct {
		o interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &Logger{
				logger: tt.fields.logger,
				level:  tt.fields.level,
			}
			l.Error(tt.args.o)
		})
	}
}

func TestLogger_Critical(t *testing.T) {
	type fields struct {
		logger *log.Logger
		level  string
	}
	type args struct {
		o interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &Logger{
				logger: tt.fields.logger,
				level:  tt.fields.level,
			}
			l.Critical(tt.args.o)
		})
	}
}

func TestCreateLogger(t *testing.T) {
	type args struct {
		f     *os.File
		level string
	}
	tests := []struct {
		name string
		args args
		want *Logger
	}{

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CreateLogger(tt.args.f, tt.args.level); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateLogger() = %v, want %v", got, tt.want)
			}
		})
	}
}
