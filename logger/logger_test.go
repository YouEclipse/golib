package logger

import (
	"testing"
)

func init() {
	InitLogger()
}

func TestDebugf(t *testing.T) {
	type args struct {
		format string
		params []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Debugf Test",
			args: args{
				format: "%s:我能吞下玻璃而不伤身体,The quick brown fox jumps over the lazy dog.",
				params: []interface{}{"这是一条debug日志"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Debugf(tt.args.format, tt.args.params...)
		})
	}
}

func TestInfof(t *testing.T) {
	type args struct {
		format string
		params []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Infof Test",
			args: args{
				format: "%s:我能吞下玻璃而不伤身体,The quick brown fox jumps over the lazy dog.",
				params: []interface{}{"这是一条info日志"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Infof(tt.args.format, tt.args.params...)
		})
	}
}

func TestWarnf(t *testing.T) {
	type args struct {
		format string
		params []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Warnf Test",
			args: args{
				format: "%s:我能吞下玻璃而不伤身体,The quick brown fox jumps over the lazy dog.",
				params: []interface{}{"这是一条warn日志"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Warnf(tt.args.format, tt.args.params...)
		})
	}
}

func TestErrorf(t *testing.T) {
	type args struct {
		format string
		params []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Errorf Test",
			args: args{
				format: "%s:我能吞下玻璃而不伤身体,The quick brown fox jumps over the lazy dog.",
				params: []interface{}{"这是一条err日志"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Errorf(tt.args.format, tt.args.params...)
		})
	}
}

func TestFatalf(t *testing.T) {
	type args struct {
		format string
		params []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Fatalf Test",
			args: args{
				format: "%s:我能吞下玻璃而不伤身体,The quick brown fox jumps over the lazy dog.",
				params: []interface{}{"这是一条fatal日志"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Fatalf(tt.args.format, tt.args.params...)
		})
	}
}

func TestDebug(t *testing.T) {
	type args struct {
		v []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Debug Test",
			args: args{
				v: []interface{}{"这是一条debug日志:", "我能吞下玻璃而不伤身体", "The quick brown fox jumps over the lazy dog."},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Debug(tt.args.v)
		})
	}
}

func TestInfo(t *testing.T) {
	type args struct {
		v []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Info Test",
			args: args{
				v: []interface{}{"这是一条info日志:", "我能吞下玻璃而不伤身体", "The quick brown fox jumps over the lazy dog."},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Info(tt.args.v)
		})
	}
}

func TestWarn(t *testing.T) {
	type args struct {
		v []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Warn Test",
			args: args{
				v: []interface{}{"这是一条warn日志:", "我能吞下玻璃而不伤身体", "The quick brown fox jumps over the lazy dog."},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Warn(tt.args.v)
		})
	}
}

func TestError(t *testing.T) {
	type args struct {
		v []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Error Test",
			args: args{
				v: []interface{}{"这是一条error日志:", "我能吞下玻璃而不伤身体", "The quick brown fox jumps over the lazy dog."},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Error(tt.args.v)
		})
	}
}

func TestFatal(t *testing.T) {
	type args struct {
		v []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Fatal Test",
			args: args{
				v: []interface{}{"这是一条fatal日志:", "我能吞下玻璃而不伤身体", "The quick brown fox jumps over the lazy dog."},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Fatal(tt.args.v)
		})
	}
}
