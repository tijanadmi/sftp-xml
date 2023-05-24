package data

import (
	"database/sql"
)

type OracleDBRepo struct {
	DB *sql.DB
}

type Models struct {
	DB OracleDBRepo
}

func NewModels(db *sql.DB) Models {
	return Models{
		DB: OracleDBRepo{DB: db},
	}
}

type SovaIdentify struct {
	ID          int    `json:"id"`
	SftpFolder  string `json:"sftp_folder"`
	FileName    string `json:"file_name"`
	FileSender  string `json:"file_sender"`
	FileArea    string `json:"file_area"`
	FileVersion string `json:"file_version"`
}

type SovaDay struct {
	ID                              int    `json:"id"`
	FileSender                      string `json:"file_sender"`
	FileArea                        string `json:"file_area"`
	FileVersion                     string `json:"file_version"`
	SenderIdentification            string `json:"sender_identification"`
	ReceiverIdentification          string `json:"receiver_identification"`
	DocumentDateTime                string `json:"document_datatime"`
	AccountingPeriod                string `json:"accounting_period"`
	SendersTimeSeriesIdentification string `json:"senders_timeseries_identification"`
	Area                            string `json:"area"`
	AccountingPoint                 string `json:"accounting_point"`
	ResolutionOfPeriod              string `json:"resolution"`
}

type SovaAccountInterval struct {
	ID        int    `json:"id"`
	IdSovaDay string `json:"id_sova_day"`
	Pos       string `json:"pos"`
	InQty     string `json:"in_qty"`
	OutQty    string `json:"out_qty"`
}
