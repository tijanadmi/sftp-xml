package dbrepo

import (
	"context"
	"database/sql"
	"fmt"
	"log"
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

func (m *OracleDBRepo) InsertSovaDayAndReturnId(e data.EnergyAccountReport_100H, FileDate string, FileSender string, FileArea string, FileVersion string) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `insert into VET.SOVA_DAY (FILE_DATE, FILE_SENDER, FILE_AREA, FILE_VERSION, SENDER_IDENTIFICATION, RECEIVER_IDENTIFICATION, DOCUMENT_DATATIME, 
		ACCOUNTING_PERIOD, SENDERS_TIMESERIES_IDENT, AREA, ACCOUNTING_POINT, RESOLUTION_OF_PERIOD) 
		values (:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12) RETURNING id INTO :13
	`
	var id int
	stmt, err := m.DB.PrepareContext(ctx, query)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, FileDate, FileSender, FileArea, FileVersion, e.SenderIdentification.V, e.ReceiverIdentification.V, e.DocumentDateTime.V, e.AccountingPeriod.V, e.AccountTimeSeries.SendersTimeSeriesIdentification.V, e.AccountTimeSeries.Area.V, e.AccountTimeSeries.AccountingPoint.V, e.AccountTimeSeries.Period.Resolution.V, sql.Out{Dest: &id})
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (m *OracleDBRepo) InsertSovaAccountInterval(id int,pos string, in string, out string) error {

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `insert into VET.SOVA_ACCOUNT_INTERVAL (ID_SOVA_DAY, POS, IN_QTY, OUT_QTY) 
		values (:1, :2, :3, :4) `
	//var int status
	//var string message
	_, err := m.DB.ExecContext(ctx, query, id, pos, in, out)

	if err != nil {
		log.Println(err)
		return err
	} else {
		return nil
	}
}
