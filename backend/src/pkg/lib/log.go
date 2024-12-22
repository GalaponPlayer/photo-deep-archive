package lib

import (
	"encoding/json"
	"runtime"

	"log/slog"
)

func LogInfo(message string, v interface{}) {
	b, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		slog.Error("lib.LogInfo jsonMarshal", err)
		return
	}
	_, file, line, _ := runtime.Caller(1)
	slog.Info(message, "file: ", file, "line: ", line, "data: ", string(b))
}

func LogError(message string, v interface{}) {
	b, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		slog.Error("lib.LogError jsonMarshal", err)
		return
	}
	_, file, line, _ := runtime.Caller(1)
	slog.Error(message, "file: ", file, "line: ", line, "data: ", string(b))
}
