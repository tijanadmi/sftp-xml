package dbrepo

import (
	"context"
	"database/sql"
	"fmt"
	"sftp_xml/pkg/data"
	"time"
)

type OracleDBRepo struct {
	DB *sql.DB
}

func (m *OracleDBRepo) Connection() *sql.DB {
	return m.DB
}

func (m *OracleDBRepo) GetSovaIdentify() ([]*data.SovaIdentify, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `select id, sftp_folder, file_name,file_sender, file_area, file_version from vet.sova_idetify
	`

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		fmt.Println("Pogresan upit ili nema rezultata upita")
		return nil, err
	}
	defer rows.Close()

	var is []*data.SovaIdentify

	for rows.Next() {
		var i data.SovaIdentify
		err := rows.Scan(
			&i.ID,
			&i.SftpFolder,
			&i.FileName,
			&i.FileSender,
			&i.FileArea,
			&i.FileVersion,
		)

		if err != nil {
			return nil, err
		}

		is = append(is, &i)
	}

	return is, nil
}
