package repository

import (
	"database/sql"
	"sftp_xml/pkg/data"
)

type DatabaseRepo interface {
	Connection() *sql.DB
	GetSovaIdentify() ([]*data.SovaIdentify, error)
}
