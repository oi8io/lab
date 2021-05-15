package log

import (
	"testing"
)

func TestLog(t *testing.T) {
	SetLevel(LevelDebug)
	Debugf("oksdfasdf [%s]", "sfasdfsafas")
	Infof("oksdfasdf [%s]", "sfasdfsafas")
	Warnf("oksdfasdf [%s]", "sfasdfsafas")
	Errorf("oksdfasdf [%s]", "sfasdfsafas")
}
