package cmd

import (
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

func Backup() error {
	var backupFilePath string

	// get time as string
	timeString := time.Now().Format("20060102150405")

	// format filename
	backupFileName := fmt.Sprintf(".%s-hosts.bak", timeString)

	// get path of current file
	tmpPath := strings.Split(hf.Config.FilePath, "/")

	// ensure path is absolute or make it relative
	if len(tmpPath) > 1 && !strings.HasPrefix(tmpPath[0], "/") {
		backupFilePath = strings.Join(tmpPath[:len(tmpPath)-1], "/")
	} else {
		backupFilePath = "./"
	}

	// format full file path + name
	backupFile := fmt.Sprintf("%s/%s", backupFilePath, backupFileName)

	// backup file
	_, err := copy(hf.Config.FilePath, backupFile)
	return err
}

func copy(src, dst string) (int64, error) {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return 0, err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return 0, fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return 0, err
	}
	defer destination.Close()
	nBytes, err := io.Copy(destination, source)
	return nBytes, err
}
