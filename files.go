package iniopt

import "github.com/go-ini/ini"

// ReadFiles reads two INI's returning a type containing them both
func ReadFiles(original, current string) (ic *INIComp, err error) {

	ic.Current, err = ini.Load(current)
	if err != nil {
		return
	}

	ic.Original, err = ini.Load(original)
	if err != nil {
		return
	}

	return
}
