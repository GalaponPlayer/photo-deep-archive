package lib

import (
	"encoding/json"
	"runtime"

	"log/slog"
)

func LogInfo(message string, v interface{}) {
	b, err := json.Marshal(v)
	if err != nil {
		slog.Error("lib.LogInfo jsonMarshal", err)
	}
	_, file, line, _ := runtime.Caller(1)
	slog.Info(message, "file: ", file, "line: ", line, "data: ", string(b))
}
