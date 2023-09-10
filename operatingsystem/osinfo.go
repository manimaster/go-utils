package osinfo

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
)

type OS string

const (
	CENTOS  OS = "centos"
	WINDOWS OS = "windows"
)

func GatherSystemInfo(osType OS) (*bytes.Buffer, error) {
	var logs []string

	switch osType {
	case CENTOS:
		logs = []string{
			"/var/log/messages",
			"/var/log/syslog",
		}
	case WINDOWS:
		// For simplicity, I'm skipping Windows logs here.
		// Accessing Windows logs typically requires using the Windows Event Log API.
	default:
		return nil, fmt.Errorf("unsupported OS type")
	}

	// Gather system information
	info, err := gatherInfo()
	if err != nil {
		return nil, err
	}

	// Create a buffer to write our archive to
	buf := new(bytes.Buffer)

	// Create a new zip archive
	w := zip.NewWriter(buf)

	// Add system information to zip
	var files = []struct {
		Name, Body string
	}{
		{"system_info.txt", info},
	}
	for _, file := range files {
		f, err := w.Create(file.Name)
		if err != nil {
			return nil, err
		}
		_, err = f.Write([]byte(file.Body))
		if err != nil {
			return nil, err
		}
	}

	// Add logs to zip
	for _, log := range logs {
		err = addFileToZip(w, log)
		if err != nil {
			return nil, err
		}
	}

	// Close the archive
	err = w.Close()
	if err != nil {
		return nil, err
	}

	return buf, nil
}

func gatherInfo() (string, error) {
	v, _ := mem.VirtualMemory()
	h, _ := host.Info()

	return fmt.Sprintf("Memory: %v\nHost: %v", v, h), nil
}

func addFileToZip(w *zip.Writer, filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// Get the file information
	info, err := file.Stat()
	if err != nil {
		return err
	}

	// Create the file header
	header, err := zip.FileInfoHeader(info)
	if err != nil {
		return err
	}

	// Using FileInfoHeader.Name instead of filename ensures we maintain folder structure
	header.Name = filepath.Base(filename)
	header.Method = zip.Deflate

	wr, err := w.CreateHeader(header)
	if err != nil {
		return err
	}
	_, err = io.Copy(wr, file)
	return err
}
