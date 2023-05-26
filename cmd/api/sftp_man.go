package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
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

func (app *application) insertDataFromFile00100H(folderName string, sftpClient sftp.Client, f data.SftpFile) error {
	filename := fmt.Sprintf("/%s/%s", folderName, f.Name)
	remoteFile, err := sftpClient.Open(filename)
	if err != nil {
		return err
	}
	defer remoteFile.Close()
	var xmlData data.EnergyAccountReport_100H
	fileContents, err := ioutil.ReadAll(remoteFile)
	if err != nil {
		return err
	}

	err = xml.Unmarshal(fileContents, &xmlData)
	if err != nil {
		panic(err)
	}

	id, err := app.DB.InsertSovaDayAndReturnId(xmlData, f.FileDate, f.FileSender, f.FileArea, f.FileVersion)
	if err != nil {
		return err
	}
	for i := range xmlData.AccountTimeSeries.Period.AccountInterval {
		//fmt.Printf("Za %s interval %s %s\n", xmlData.AccountTimeSeries.Period.AccountInterval[i].Pos.V, xmlData.AccountTimeSeries.Period.AccountInterval[i].InQty.V, xmlData.AccountTimeSeries.Period.AccountInterval[i].OutQty.V)
		err := app.DB.InsertSovaAccountInterval(id, xmlData.AccountTimeSeries.Period.AccountInterval[i].Pos.V, xmlData.AccountTimeSeries.Period.AccountInterval[i].InQty.V, xmlData.AccountTimeSeries.Period.AccountInterval[i].OutQty.V)
		if err != nil {
			return err
		}
	}
	return nil
}

func (app *application) insertDataFromFile(folderName string, sftpClient sftp.Client, f data.SftpFile) error {
	filename := fmt.Sprintf("/%s/%s", folderName, f.Name)
	remoteFile, err := sftpClient.Open(filename)
	if err != nil {
		return err
	}
	defer remoteFile.Close()
	var xmlDataA data.EnergyAccountReport
	fileContents, err := ioutil.ReadAll(remoteFile)
	if err != nil {
		return err
	}

	err = xml.Unmarshal(fileContents, &xmlDataA)
	if err != nil {
		return err
	}
	for j := range xmlDataA.AccountTimeSeries {
		id, err := app.DB.InsertSovaDayIndexAndReturnId(xmlDataA, j, f.FileDate, f.FileSender, f.FileArea, f.FileVersion)
		if err != nil {
			return err
		}
		for i := range xmlDataA.AccountTimeSeries[j].Period.AccountInterval {
			//fmt.Printf("Za %s interval %s %s\n", xmlDataA.AccountTimeSeries[j].Period.AccountInterval[i].Pos.V, xmlDataA.AccountTimeSeries[j].Period.AccountInterval[i].InQty.V, xmlDataA.AccountTimeSeries[j].Period.AccountInterval[i].OutQty.V)
			err := app.DB.InsertSovaAccountInterval(id, xmlDataA.AccountTimeSeries[j].Period.AccountInterval[i].Pos.V, xmlDataA.AccountTimeSeries[j].Period.AccountInterval[i].InQty.V, xmlDataA.AccountTimeSeries[j].Period.AccountInterval[i].OutQty.V)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (app *application) readFromFolder(folderName string, sftpClient sftp.Client) error {

	listOfFiles, err := app.getFiles(sftpClient, "cges")

	if err != nil {
		return err
	}
	for _, f := range listOfFiles {
		fmt.Println(f.Name, f.FileDate, f.FileSender, f.FileArea, f.FileVersion)
		if f.FileArea == "10Y1001C--00100H" {
			err = app.insertDataFromFile00100H("cges", sftpClient, *f)
			if err != nil {
				return err
			}
		} else {
			err = app.insertDataFromFile("cges", sftpClient, *f)
			if err != nil {
				return err
			}

		}
	}
	return nil
}
