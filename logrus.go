package logrus

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
)

type PlainFormatter struct {
	TimestampFormat string
	LevelDesc       []string
}

func (f *PlainFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	return []byte(fmt.Sprintf("%s %s %s\n", f.LevelDesc[entry.Level], entry.Time.Format(f.TimestampFormat), entry.Message)), nil
}

var defaultFormatter = &PlainFormatter{
	TimestampFormat: "2006-01-02 15:04:05",
	LevelDesc:       []string{"PANC", "FATL", "ERRO", "WARN", "INFO", "DEBG", "TRAC"},
}

func init() {
	logrus.SetOutput(os.Stdout)
	logrus.SetFormatter(defaultFormatter)
}
