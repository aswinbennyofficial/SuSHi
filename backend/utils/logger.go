package utils

import (
	"os"

	"github.com/aswinbennyofficial/SuSHi/models"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var (
    AppLogFile *os.File
    AppLogger  zerolog.Logger
)

type Config models.Config



// LoadLogger is a function that loads the logger
func LoadLogger(config models.Config) {
    // Open log file for writing
    var err error
    AppLogFile, err = os.OpenFile(config.LogPath+"/activity-metrics-backend.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        log.Panic().Err(err).Msg("Failed to open app log file")
    }

    // Create multi-level writer to write to both file and stdout
    multiLevelWriter := zerolog.MultiLevelWriter(AppLogFile, os.Stdout)

    // Create logger instance for all log levels
    AppLogger = zerolog.New(multiLevelWriter).With().Timestamp().Logger()

    // Set global logger
    log.Logger = AppLogger
}

func CloseLogFiles() {
    // Close log file
    if AppLogFile != nil {
        if err := AppLogFile.Close(); err != nil {
            log.Error().Err(err).Msg("Failed to close app log file")
        }
    }
}