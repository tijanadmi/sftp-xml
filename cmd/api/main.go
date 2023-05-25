package main

import (
	"context"
	"database/sql"
	"encoding/xml"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	"sftp_xml/pkg/data"
	"sftp_xml/pkg/repository"
	"sftp_xml/pkg/repository/dbrepo"

	_ "github.com/godror/godror"
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

type application struct {
	db struct {
		dsn string
	}
	logger *log.Logger
	DB     repository.DatabaseRepo
}

func main() {
	var app application
	serverPort := 22
	var serverAddr string
	var username string
	var password string

	flag.StringVar(&serverAddr, "servername", "", "SFTP server adres")
	flag.StringVar(&username, "username", "", "SFTP user")
	flag.StringVar(&password, "password", "", "SFTP user password")

	flag.StringVar(&app.db.dsn, "dsn", "", "Oracle connection string")

	flag.Parse()

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	db, err := openDB(app)
	if err != nil {
		logger.Fatal(err)
	}
	defer db.Close()

	logger.Println("Connected to database")

	app = application{
		logger: logger,
		DB:     &dbrepo.OracleDBRepo{DB: db},
	}

	// Create an SSH client config with the given username and password
	fmt.Println(serverAddr)
	fmt.Println(username)
	fmt.Println(password)
	sshConfig := &ssh.ClientConfig{
		User: username,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	// Connect to the SSH server using the client config

	sshConn, err := ssh.Dial("tcp", fmt.Sprintf("%s:%d", serverAddr, serverPort), sshConfig)
	if err != nil {
		panic(err)
	}
	fmt.Println("Connect to SFTP server!")
	sftpClient, err := sftp.NewClient(sshConn)
	if err != nil {
		panic(err)
	}
	defer sftpClient.Close()
	fmt.Println("Create a client!")

	//*
	//* List working directory files
	//*
	//listFiles(*sftpClient, "cges")

	listOfFiles, err := app.getFiles(*sftpClient, "cges")

	if err != nil {
		panic(err)
	}
	for _, f := range listOfFiles {
		fmt.Println(f.Name, f.FileDate, f.FileSender, f.FileArea, f.FileVersion)
		if f.FileArea == "10Y1001C--00100H" {
			filename := fmt.Sprintf("/cges/%s", f.Name)
			remoteFile, err := sftpClient.Open(filename)
			if err != nil {
				panic(err)
			}
			defer remoteFile.Close()
			var xmlData data.EnergyAccountReport_100H
			fileContents, err := ioutil.ReadAll(remoteFile)
			if err != nil {
				panic(err)
			}

			err = xml.Unmarshal(fileContents, &xmlData)
			if err != nil {
				panic(err)
			}

			id, err := app.DB.InsertSovaDayAndReturnId(xmlData, f.FileDate, f.FileSender, f.FileArea, f.FileVersion)
			if err != nil {
				panic(err)
			}
			for i := range xmlData.AccountTimeSeries.Period.AccountInterval {
				fmt.Printf("Za %s interval %s %s\n", xmlData.AccountTimeSeries.Period.AccountInterval[i].Pos.V, xmlData.AccountTimeSeries.Period.AccountInterval[i].InQty.V, xmlData.AccountTimeSeries.Period.AccountInterval[i].OutQty.V)
				err := app.DB.InsertSovaAccountInterval(id, xmlData.AccountTimeSeries.Period.AccountInterval[i].Pos.V, xmlData.AccountTimeSeries.Period.AccountInterval[i].InQty.V, xmlData.AccountTimeSeries.Period.AccountInterval[i].OutQty.V)
				if err != nil {
					panic(err)
				}
			}
		}
	}

	//*

}

func openDB(app application) (*sql.DB, error) {
	db, err := sql.Open("godror", app.db.dsn)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	return db, nil
}
