package repository

import (
	"database/sql"
	"sftp_xml/pkg/data"
)

type DatabaseRepo interface {
	Connection() *sql.DB
	GetSovaIdentify() ([]*data.SovaIdentify, error)
	InsertSovaDayAndReturnId(e data.EnergyAccountReport_100H, FileDate string, FileSender string, FileArea string, FileVersion string) (int, error)
	InsertSovaAccountInterval(id int, pos string, in string, out string) error
}
