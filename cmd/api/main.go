package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"io/ioutil"

	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

type EnergyAccountReport struct {
	XMLName                xml.Name `xml:"EnergyAccountReport"`
	Text                   string   `xml:",chardata"`
	DtdRelease             string   `xml:"DtdRelease,attr"`
	DtdVersion             string   `xml:"DtdVersion,attr"`
	DocumentIdentification struct {
		Text string `xml:",chardata"`
		V    string `xml:"v,attr"`
	} `xml:"DocumentIdentification"`
	DocumentVersion struct {
		Text string `xml:",chardata"`
		V    string `xml:"v,attr"`
	} `xml:"DocumentVersion"`
	DocumentType struct {
		Text string `xml:",chardata"`
		V    string `xml:"v,attr"`
	} `xml:"DocumentType"`
	DocumentStatus struct {
		Text string `xml:",chardata"`
		V    string `xml:"v,attr"`
	} `xml:"DocumentStatus"`
	ProcessType struct {
		Text string `xml:",chardata"`
		V    string `xml:"v,attr"`
	} `xml:"ProcessType"`
	ClassificationType struct {
		Text string `xml:",chardata"`
		V    string `xml:"v,attr"`
	} `xml:"ClassificationType"`
	SenderIdentification struct {
		Text         string `xml:",chardata"`
		V            string `xml:"v,attr"`
		CodingScheme string `xml:"codingScheme,attr"`
	} `xml:"SenderIdentification"`
	SenderRole struct {
		Text string `xml:",chardata"`
		V    string `xml:"v,attr"`
	} `xml:"SenderRole"`
	ReceiverIdentification struct {
		Text         string `xml:",chardata"`
		V            string `xml:"v,attr"`
		CodingScheme string `xml:"codingScheme,attr"`
	} `xml:"ReceiverIdentification"`
	ReceiverRole struct {
		Text string `xml:",chardata"`
		V    string `xml:"v,attr"`
	} `xml:"ReceiverRole"`
	DocumentDateTime struct {
		Text string `xml:",chardata"`
		V    string `xml:"v,attr"`
	} `xml:"DocumentDateTime"`
	AccountingPeriod struct {
		Text string `xml:",chardata"`
		V    string `xml:"v,attr"`
	} `xml:"AccountingPeriod"`
	Domain struct {
		Text         string `xml:",chardata"`
		V            string `xml:"v,attr"`
		CodingScheme string `xml:"codingScheme,attr"`
	} `xml:"Domain"`
	AccountTimeSeries struct {
		Text                            string `xml:",chardata"`
		SendersTimeSeriesIdentification struct {
			Text string `xml:",chardata"`
			V    string `xml:"v,attr"`
		} `xml:"SendersTimeSeriesIdentification"`
		BusinessType struct {
			Text string `xml:",chardata"`
			V    string `xml:"v,attr"`
		} `xml:"BusinessType"`
		Product struct {
			Text string `xml:",chardata"`
			V    string `xml:"v,attr"`
		} `xml:"Product"`
		ObjectAggregation struct {
			Text string `xml:",chardata"`
			V    string `xml:"v,attr"`
		} `xml:"ObjectAggregation"`
		Area struct {
			Text         string `xml:",chardata"`
			V            string `xml:"v,attr"`
			CodingScheme string `xml:"codingScheme,attr"`
		} `xml:"Area"`
		MeasurementUnit struct {
			Text string `xml:",chardata"`
			V    string `xml:"v,attr"`
		} `xml:"MeasurementUnit"`
		AccountingPoint struct {
			Text         string `xml:",chardata"`
			V            string `xml:"v,attr"`
			CodingScheme string `xml:"codingScheme,attr"`
		} `xml:"AccountingPoint"`
		Period struct {
			Text         string `xml:",chardata"`
			TimeInterval struct {
				Text string `xml:",chardata"`
				V    string `xml:"v,attr"`
			} `xml:"TimeInterval"`
			Resolution struct {
				Text string `xml:",chardata"`
				V    string `xml:"v,attr"`
			} `xml:"Resolution"`
			AccountInterval []struct {
				Text string `xml:",chardata"`
				Pos  struct {
					Text string `xml:",chardata"`
					V    string `xml:"v,attr"`
				} `xml:"Pos"`
				InQty struct {
					Text string `xml:",chardata"`
					V    string `xml:"v,attr"`
				} `xml:"InQty"`
				OutQty struct {
					Text string `xml:",chardata"`
					V    string `xml:"v,attr"`
				} `xml:"OutQty"`
			} `xml:"AccountInterval"`
		} `xml:"Period"`
	} `xml:"AccountTimeSeries"`
}

func main() {

	serverPort := 22
	var serverAddr string
	var username string
	var password string

	flag.StringVar(&serverAddr, "servername", "", "SFTP server adres")
	flag.StringVar(&username, "username", "", "SFTP user")
	flag.StringVar(&password, "password", "", "SFTP user password")

	flag.Parse()

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
	// Open the remote file and read its contents
	remoteFile, err := sftpClient.Open("/cges/20230322_SOVA_10YCS-CG-TSO---S_10Y1001C--00100H_001.xml")
	if err != nil {
		panic(err)
	}
	defer remoteFile.Close()

	var xmlData EnergyAccountReport
	fileContents, err := ioutil.ReadAll(remoteFile)
	if err != nil {
		panic(err)
	}

	err = xml.Unmarshal(fileContents, &xmlData)
	if err != nil {
		panic(err)
	}

	// Print the file contents to the console
	//fmt.Println(string(fileContents))
	//fmt.Println(xmlData)
	for i := range xmlData.AccountTimeSeries.Period.AccountInterval {
		fmt.Printf("Za %s interval %s %s\n", xmlData.AccountTimeSeries.Period.AccountInterval[i].Pos.V, xmlData.AccountTimeSeries.Period.AccountInterval[i].InQty.V, xmlData.AccountTimeSeries.Period.AccountInterval[i].OutQty.V)
	}
}
