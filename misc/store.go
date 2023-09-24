package misc

import (
	"gopkg.in/ini.v1"
)

var storePath = STORE_FILE

var storeFile *ini.File

func init() {
	f, err := ini.Load(storePath)
	if err != nil {
		f, err = ini.Load([]byte("[Store]"))
		if err != nil {
			panic(err)
		}
		storeFile = f
		return
	}
	storeFile = f
}

func SetItem(key string, value string) {
	storeFile.Section("Store").NewKey(key, value)
	persistence()
}

func GetItem(key, fallback string) string {
	k, err := storeFile.Section("Store").GetKey(key)
	if err != nil {
		return fallback
	}
	return k.String()
}

func persistence() {
	storeFile.SaveTo(storePath)
}
