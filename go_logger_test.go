package go_logger

import (
	// "log"
	// "os"
	// "reflect"
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
		{"debug", args{"debug"}, 0},
		{"info", args{"info"}, 1},
		{"warning", args{"warning"}, 2},
		{"error", args{"error"}, 3},
		{"critical", args{"critical"}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertLevelToInt(tt.args.s); got != tt.want {
				t.Errorf("convertLevelToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

// func TestGoLogger_Debug(t *testing.T) {
// 	type fields struct {
// 		logger *log.Logger
// 		level  string
// 	}
// 	type args struct {
// 		v interface{}
// 	}
// 	tests := []struct {
// 		name   string
// 		fields fields
// 		args   args
// 	}{
// 		{},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			l := &GoLogger{
// 				logger: tt.fields.logger,
// 				level:  tt.fields.level,
// 			}
// 			l.Debug(tt.args.v)
// 		})
// 	}
// }

// func TestGoLogger_Info(t *testing.T) {
// 	type fields struct {
// 		logger *log.Logger
// 		level  string
// 	}
// 	type args struct {
// 		v interface{}
// 	}
// 	tests := []struct {
// 		name   string
// 		fields fields
// 		args   args
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			l := &GoLogger{
// 				logger: tt.fields.logger,
// 				level:  tt.fields.level,
// 			}
// 			l.Info(tt.args.v)
// 		})
// 	}
// }

// func TestGoLogger_Warning(t *testing.T) {
// 	type fields struct {
// 		logger *log.Logger
// 		level  string
// 	}
// 	type args struct {
// 		v interface{}
// 	}
// 	tests := []struct {
// 		name   string
// 		fields fields
// 		args   args
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			l := &GoLogger{
// 				logger: tt.fields.logger,
// 				level:  tt.fields.level,
// 			}
// 			l.Warning(tt.args.v)
// 		})
// 	}
// }

// func TestGoLogger_Error(t *testing.T) {
// 	type fields struct {
// 		logger *log.Logger
// 		level  string
// 	}
// 	type args struct {
// 		v interface{}
// 	}
// 	tests := []struct {
// 		name   string
// 		fields fields
// 		args   args
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			l := &GoLogger{
// 				logger: tt.fields.logger,
// 				level:  tt.fields.level,
// 			}
// 			l.Error(tt.args.v)
// 		})
// 	}
// }

// func TestGoLogger_Critical(t *testing.T) {
// 	type fields struct {
// 		logger *log.Logger
// 		level  string
// 	}
// 	type args struct {
// 		v interface{}
// 	}
// 	tests := []struct {
// 		name   string
// 		fields fields
// 		args   args
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			l := &GoLogger{
// 				logger: tt.fields.logger,
// 				level:  tt.fields.level,
// 			}
// 			l.Critical(tt.args.v)
// 		})
// 	}
// }

