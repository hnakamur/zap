package zap_test

import (
	"io/ioutil"
	"os"
	"testing"

	"go.uber.org/zap"
)

func BenchmarkZapDevelopmentLog(b *testing.B) {
	tmpfile, err := ioutil.TempFile("", "benchmark")
	if err != nil {
		b.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())

	cfg := zap.NewDevelopmentConfig()
	cfg.OutputPaths = []string{tmpfile.Name()}
	logger, err := cfg.Build()
	if err != nil {
		b.Fatal(err)
	}
	for i := 0; i < b.N; i++ {
		logger.Info("sample log message", zap.String("key1", "value1"), zap.String("key2", "value2"))
	}
}
