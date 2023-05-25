package main

import (
	"fmt"
	"os"
	"sftp_xml/pkg/data"
	"strings"

	"github.com/pkg/sftp"
)

func (app *application) getFiles(sc sftp.Client, remoteDir string) ([]*data.SftpFile, error) {
	files, err := sc.ReadDir(remoteDir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to list remote dir: %v\n", err)
		return nil, err
	}

	var fs []*data.SftpFile

	for _, f := range files {
		var name, modTime, size string
		var fl data.SftpFile
		name = f.Name()
		modTime = f.ModTime().Format("2006-01-02 15:04:05")
		size = fmt.Sprintf("%12d", f.Size())

		if f.IsDir() {
			name = name + "/"
			modTime = ""
			size = "PRE"
		} else {
			s := strings.Split(name, "_")
			fl.FileDate = s[0]
			fl.FileSender = s[2]
			fl.FileArea = s[3]
			fl.FileVersion = s[4]
		}
		// Output each file name and size in bytes
		//fmt.Fprintf(os.Stdout, "%19s %12s %s\n", modTime, size, name)
		fl.Name = name
		fl.ModTime = modTime
		fl.Size = size

		fs = append(fs, &fl)
	}

	return fs, nil
}

/*func listFiles(sc sftp.Client, remoteDir string) (err error) {
	fmt.Fprintf(os.Stdout, "Listing [%s] ...\n\n", remoteDir)

	files, err := sc.ReadDir(remoteDir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to list remote dir: %v\n", err)
		return
	}

	for _, f := range files {
		var name, modTime, size string

		name = f.Name()
		modTime = f.ModTime().Format("2006-01-02 15:04:05")
		size = fmt.Sprintf("%12d", f.Size())

		if f.IsDir() {
			name = name + "/"
			modTime = ""
			size = "PRE"
		}
		// Output each file name and size in bytes
		fmt.Fprintf(os.Stdout, "%19s %12s %s\n", modTime, size, name)
	}

	return
}*/
