package misc

import (
	"fmt"
	"os"
)

const (
	VERSION = "1.0"
)

var Env string

func IsDev() bool {
	return Env != "prod"
}

var APP_STORE_PATH = func() string {
	path := os.TempDir()
	if IsDev() {
		path, _ = os.Getwd()
	}
	return path + string(os.PathSeparator)
}()

var LOG_PATH = APP_STORE_PATH + "mydata.log"
var STORE_FILE = fmt.Sprintf("%smydata_store.ini", APP_STORE_PATH)

var AppLogger = NewCustomLogger(LOG_PATH)
