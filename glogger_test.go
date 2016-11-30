package glogger

import (
	"testing"
	"log"
	"github.com/stretchr/testify/assert"
)

type ioWriterMock struct {
	string string
	log    *log.Logger
}

func (f *ioWriterMock) Write(buff []byte) (int, error) {
	f.string = string(buff)
	f.log.Print(f.string)
	return len(buff), nil
}

var so, se = ioWriterMock{"", stdout}, ioWriterMock{"", stderr}

func init() {
	stdout = log.New(&so, "", log.LstdFlags)
	stderr = log.New(&se, "", log.LstdFlags)
}

func TestPanic(t *testing.T) {
	_logger := Create(FATAL)

	defer func () {
		r := recover()
		assert.NotNil(t, r)

		assert.Equal(t, r, "Panic panic")
	} ()

	_logger.Panic("Panic %s", "panic")
}

func TestLevels(t *testing.T) {
	var _logger = Create(TRACE)

	str := "trace"
	checkLog(t, _logger, TRACE, str, TRACE, str, "")
	checkLog(t, _logger, TRACE, str, DEBUG, str, "")
	checkLog(t, _logger, TRACE, str, WARN, "", str)

	str = "debug"
	checkLog(t, _logger, DEBUG, str, TRACE, "", "")
	checkLog(t, _logger, DEBUG, str, DEBUG, str, "")
	checkLog(t, _logger, DEBUG, str, WARN, "", str)

	str = "info"
	checkLog(t, _logger, INFO, str, DEBUG, "", "")
	checkLog(t, _logger, INFO, str, INFO, str, "")
	checkLog(t, _logger, INFO, str, WARN, "", str)

	str = "warn"
	checkLog(t, _logger, WARN, str, INFO, "", "")
	checkLog(t, _logger, WARN, str, WARN, "", str)
	checkLog(t, _logger, WARN, str, ERROR, "", str)

	str = "error"
	checkLog(t, _logger, ERROR, str, INFO, "", "")
	checkLog(t, _logger, ERROR, str, WARN, "", "")
	checkLog(t, _logger, ERROR, str, ERROR, "", str)
}

func checkLog(t *testing.T, _logger Logger, ll LogLevel, str string, pl LogLevel, std_out string, stderr_out string) {
	_logger.SetLevel(ll)

	se.string = ""
	so.string = ""

	_logger.Log(pl, str)
	assertOut(t, pl, so.string, std_out, "stdout")
	assertOut(t, pl, se.string, stderr_out, "stderr")
}

func assertOut(t *testing.T, pl LogLevel, actual, expected, stream string) {
	if actual == "" && expected == "" {
		return
	}

	if actual != "" && expected != "" {
		return
	}

	assert.Fail(t, "", "Unexpected %s output expected %s: [%#v] actual: [%#v]", stream, pl, expected, actual)
}
