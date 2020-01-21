package iniopt

import (
	"fmt"
	"log"
	"path/filepath"
	"strings"

	"github.com/go-ini/ini"
	"golang.org/x/text/encoding/charmap"
)

type INIComp struct {
	Name        string
	Original    *ini.File
	Current     *ini.File
	Differences map[string]map[string]string // map - section -> keys -> value of differences
}

func CompareINI(original, current string, makeAll *bool) (b []byte, err error) {

	// read the INI files and create the compare object
	ic, err := readFiles(original, current)
	if err != nil {
		return
	}

	ic.Name = filepath.Base(strings.TrimSuffix(original, filepath.Ext(original)))

	for _, section := range ic.Current.Sections() {
		// Default section is not used by LOC
		// error out if current file contains un sectioned keys
		if section.Name() == "DEFAULT" {
			if len(section.Keys()) > 0 {
				return b, fmt.Errorf("current file has entries with no section: %v", section.Keys())
			}
			continue
		}

		oSection := ic.Original.Section(section.Name())

		// loop over keys and compare values
		log.Println("Section Found: ", section.Name())

		for _, key := range section.Keys() {
			// log.Println(key.Name(), key.String())
			oKey := oSection.Key(key.Name())

			// section/key/value matches between old and new, do nothing
			if key.String() == oKey.String() {
				log.Printf("values match for key %s in section %s\n", key.Name(), section.Name())
				// if makeall is false then continue, otherwise add the key
				if !*makeAll {
					continue
				}
			}

			ic.addDifference(section.Name(), key.Name(), key.String())
		}
	}

	b, err = ic.makeBytes()

	return
}

func (ic *INIComp) makeBytes() (b []byte, err error) {
	b = ic.writeDifferences()
	b = append(b, []byte(fmt.Sprintf("@dbHot(%s.INI,SET,%s.INI[*);\r\n", ic.Name, ic.Name))...)

	return encode(b)
}

func (ic *INIComp) writeDifferences() (b []byte) {
	for section, keys := range ic.Differences {
		for key, value := range keys {
			b = append(b, []byte(fmt.Sprintf("@WIZSET(%s.INI[%s]%s=%s);\r\n", ic.Name, section, key, value))...)
		}

	}
	return
}

func (ic *INIComp) addDifference(section, key, value string) {
	if ic.Differences[section] == nil {
		ic.Differences[section] = make(map[string]string)
	}
	ic.Differences[section][key] = value
}

// ReadFiles reads two INI's returning a type containing them both
func readFiles(original, current string) (*INIComp, error) {
	var ic INIComp
	var err error

	ic.Current, err = ini.Load(current)
	if err != nil {
		return nil, err
	}

	ic.Original, err = ini.Load(original)
	if err != nil {
		return nil, err

	}

	ic.Differences = make(map[string]map[string]string)

	return &ic, nil
}

// encode takes the bytes and returns it windows 1252 encoded
func encode(b []byte) ([]byte, error) {
	return charmap.Windows1252.NewEncoder().Bytes(b)
}
