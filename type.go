package iniopt

import "github.com/go-ini/ini"

type INIComp struct {
	Original *ini.File
	Current  *ini.File
}
