package logger

import (
	"fmt"
	"mindscribe-be/pkg/config"
	"os"
	"path/filepath"
	"sync"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var log *zap.Logger

// dateBasedRotatingWriter manages date-based log file rotation
type dateBasedRotatingWriter struct {
	mu            sync.Mutex
	logDir        string
	currentFile   *os.File
	currentDate   string
	filePrefix    string
	fileExtension string
}

// Write implements the io.Writer interface
func (w *dateBasedRotatingWriter) Write(p []byte) (n int, err error) {
	w.mu.Lock()
	defer w.mu.Unlock()

	// Get current date
	currentDate := time.Now().Format("2006-01-02")

	// Rotate if date has changed
	if currentDate != w.currentDate {
		if w.currentFile != nil {
			w.currentFile.Close()
		}

		// Create log filename
		filename := filepath.Join(w.logDir, fmt.Sprintf("%s_%s.%s", w.filePrefix, currentDate, w.fileExtension))

		// Ensure log directory exists
		if err := os.MkdirAll(w.logDir, 0755); err != nil {
			return 0, err
		}

		// Open new log file
		newFile, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
		if err != nil {
			return 0, err
		}

		w.currentFile = newFile
		w.currentDate = currentDate
	}

	// Write to current log file
	return w.currentFile.Write(p)
}

// Sync ensures all writes are completed
func (w *dateBasedRotatingWriter) Sync() error {
	w.mu.Lock()
	defer w.mu.Unlock()

	if w.currentFile != nil {
		return w.currentFile.Sync()
	}
	return nil
}

// Close closes the current log file
func (w *dateBasedRotatingWriter) Close() error {
	w.mu.Lock()
	defer w.mu.Unlock()

	if w.currentFile != nil {
		return w.currentFile.Close()
	}
	return nil
}

func getLogLevel(c *config.Config) zap.AtomicLevel {
	level := zap.NewAtomicLevel()
	level.UnmarshalText([]byte(c.Logger.LogLevel))
	return level
}

// Init initializes the logger with console and file logging
func Init(cfg *config.Config) error {
	// Create encoder config
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Format("2006-01-02T15:04:05.000Z07:00"))
	}
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder

	// Create console encoder (more human-readable)
	consoleEncoder := zapcore.NewConsoleEncoder(encoderConfig)

	// Create JSON encoder for file (more machine-friendly)
	fileEncoder := zapcore.NewJSONEncoder(encoderConfig)

	var core zapcore.Core

	// Create a console writer
	consoleWriter := zapcore.AddSync(os.Stdout)

	// Create a date-based file writer
	if cfg.Logger.LogToFile == "true" {
		fileWriter := zapcore.AddSync(&dateBasedRotatingWriter{
			logDir:        "./data/logs",
			filePrefix:    "application",
			fileExtension: "log",
		})

		// Create a multi-writer core that logs to both file and console
		core = zapcore.NewTee(
			// File logging with JSON encoder
			zapcore.NewCore(fileEncoder, fileWriter, getLogLevel(cfg)),

			// Console logging with console encoder
			zapcore.NewCore(consoleEncoder, consoleWriter, getLogLevel(cfg)),
		)
	} else {
		core = zapcore.NewTee(
			// Console logging with console encoder
			zapcore.NewCore(consoleEncoder, consoleWriter, getLogLevel(cfg)),
		)
	}

	// Create logger with additional options
	log = zap.New(core,
		zap.AddCaller(),
		zap.AddStacktrace(zapcore.ErrorLevel),
	)

	return nil
}

// Logger returns the global logger instance
func Logger() *zap.Logger {
	return log
}

// Sync flushes any buffered log entries
func Sync() error {
	return log.Sync()
}
