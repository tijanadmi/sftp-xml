package data

import (
	"encoding/xml"
)

// EnergyAccountReport was generated 2023-05-24 10:16:35 by https://xml-to-go.github.io/ in Ukraine.
type EnergyAccountReport struct {
	XMLName                xml.Name `xml:"EnergyAccountReport"`
	Text                   string   `xml:",chardata"`
	DtdVersion             string   `xml:"DtdVersion,attr"`
	DtdRelease             string   `xml:"DtdRelease,attr"`
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
	AccountTimeSeries []struct {
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

type EnergyAccountReport_100H struct {
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
