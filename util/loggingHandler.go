package util

import (
	"errors"
	"io"
	"os"
	"path"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"gopkg.in/natefinch/lumberjack.v2"
)

var Log zerolog.Logger

func init() {

	if err := logDirectory(); err != nil {
		log.Error().Err(err).Msg("Directory creation failed with error: " + err.Error())
		os.Exit(1)
		return
	}

	if err := logSubDirectory(); err != nil {
		log.Error().Err(err).Msg("sub Directory creation failed with error: " + err.Error())
		os.Exit(1)
		return
	}

	var writers []io.Writer
	writers = append(writers, zerolog.ConsoleWriter{Out: os.Stderr})
	writers = append(writers, newRollingFile())
	mw := io.MultiWriter(writers...)
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	Log = zerolog.New(mw).With().Timestamp().Logger()

}

func newRollingFile() io.Writer {

	DirName := "./log/" + time.Now().Format("01-02-2006") + "/"

	logFileName := time.Now().Format("01-02-2006") + ".log"

	return &lumberjack.Logger{
		Filename:   path.Join(DirName, logFileName),
		MaxSize:    100,  // Size in MB before file gets rotated
		MaxBackups: 5,    // Max number of files kept before being overwritten
		MaxAge:     30,   // Max number of days to keep the files
		Compress:   true, // Whether to compress log files using gzip
	}
}

func logDirectory() error {

	dirName := "Log"

	err := os.Mkdir(dirName, 0700)
	if err == nil {
		return nil
	}
	if os.IsExist(err) {
		info, err := os.Stat(dirName)
		if err != nil {
			return err
		}
		if !info.IsDir() {
			return errors.New("path exists but is not a directory")
		}
		return nil
	}
	return err
}

func logSubDirectory() error {

	subDirName := "./Log/" + time.Now().Format("01-02-2006")

	err := os.Mkdir(subDirName, 0700)
	if err == nil {
		return nil
	}
	if os.IsExist(err) {
		info, err := os.Stat(subDirName)
		if err != nil {
			return err
		}
		if !info.IsDir() {
			return errors.New("path exists but is not a directory")
		}
		return nil
	}
	return err
}
