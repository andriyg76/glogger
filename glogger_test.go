package glogger

import (
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
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

	defer func() {
		r := recover()
		assert.NotNil(t, r)

		assert.Equal(t, r, "Panic panic")
	}()

	_logger.Panic("Panic %s", "panic")
}

func TestLevels(t *testing.T) {
	str := "trace"
	checkLog(t, TRACE, str, TRACE, str, "")
	checkLog(t, TRACE, str, DEBUG, str, "")
	checkLog(t, TRACE, str, WARN, "", str)

	str = "debug"
	checkLog(t, DEBUG, str, TRACE, "", "")
	checkLog(t, DEBUG, str, DEBUG, str, "")
	checkLog(t, DEBUG, str, WARN, "", str)

	str = "info"
	checkLog(t, INFO, str, DEBUG, "", "")
	checkLog(t, INFO, str, INFO, str, "")
	checkLog(t, INFO, str, WARN, "", str)

	str = "warn"
	checkLog(t, WARN, str, INFO, "", "")
	checkLog(t, WARN, str, WARN, "", str)
	checkLog(t, WARN, str, ERROR, "", str)

	str = "error"
	checkLog(t, ERROR, str, INFO, "", "")
	checkLog(t, ERROR, str, WARN, "", "")
	checkLog(t, ERROR, str, ERROR, "", str)
}

func checkLog(t *testing.T, ll LogLevel, str string, pl LogLevel, std_out string, stderr_out string) {
	_logger := Create(ll)

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
