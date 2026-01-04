package log

import (
	"fmt"
	"os"

	"github.com/Dreamacro/clash/common/observable"

	log "github.com/sirupsen/logrus"
)

var (
	logCh  = make(chan any)
	source = observable.NewObservable(logCh)
	level  = INFO
)

func init() {
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
}

type Event struct {
	LogLevel LogLevel
	Payload  string
}

func (e *Event) Type() string {
	return e.LogLevel.String()
}

func Infof(format string, v ...any) {
	event := newLog(INFO, fmt.Sprintf(format, v...))
	logCh <- event
	print(event)
}

func Warnf(format string, v ...any) {
	event := newLog(WARNING, fmt.Sprintf(format, v...))
	logCh <- event
	print(event)
}

func Errorf(format string, v ...any) {
	event := newLog(ERROR, fmt.Sprintf(format, v...))
	logCh <- event
	print(event)
}

func Debugf(format string, v ...any) {
	event := newLog(DEBUG, fmt.Sprintf(format, v...))
	logCh <- event
	print(event)
}

func Fatalf(format string, v ...any) {
	log.Fatalf(format, v...)
}

func Infoln(v ...any) {
	event := newLog(INFO, fmt.Sprint(v...))
	logCh <- event
	print(event)
}

func Warnln(v ...any) {
	event := newLog(WARNING, fmt.Sprint(v...))
	logCh <- event
	print(event)
}

func Errorln(v ...any) {
	event := newLog(ERROR, fmt.Sprint(v...))
	logCh <- event
	print(event)
}

func Debugln(v ...any) {
	event := newLog(DEBUG, fmt.Sprint(v...))
	logCh <- event
	print(event)
}

func Fatalln(v ...any) {
	log.Fatalln(v...)
}

func Subscribe() observable.Subscription {
	sub, _ := source.Subscribe()
	return sub
}

func UnSubscribe(sub observable.Subscription) {
	source.UnSubscribe(sub)
}

func Level() LogLevel {
	return level
}

func SetLevel(newLevel LogLevel) {
	level = newLevel
}

func print(data Event) {
	if data.LogLevel < level {
		return
	}

	switch data.LogLevel {
	case INFO:
		log.Infoln(data.Payload)
	case WARNING:
		log.Warnln(data.Payload)
	case ERROR:
		log.Errorln(data.Payload)
	case DEBUG:
		log.Debugln(data.Payload)
	case SILENT:
		return
	}
}

func newLog(logLevel LogLevel, payload string) Event {
	return Event{
		LogLevel: logLevel,
		Payload:  payload,
	}
}
