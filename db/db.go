package db

import (
	"context"
	"courseProject/models"
	"courseProject/utils"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	jsoniter "github.com/json-iterator/go"
	"github.com/sirupsen/logrus"
)

type PostgresDriver struct {
	conn *pgxpool.Pool
}

func InitNewDriver() (*PostgresDriver, error) {
	drv := PostgresDriver{}
	pool, err := pgxpool.New(context.Background(), "postgres://bashkatov_nikita:Dpj1eEp0omMk-3tjSPKf@176.113.81.99:5432/filmsCatalog?sslmode=disable") // pgxpool.New(context.Background(), "postgres://postgres:postgres@localhost:5432/complex?sslmode=disable")
	if err != nil {
		logrus.Errorf("cannot connect to db")
		return nil, err
	}

	drv.conn = pool
	return &drv, nil
}

func (d *PostgresDriver) ExecuteSP(sp string, model any, params any) error {
	paramsBytes, err := jsoniter.Marshal(params)
	if err != nil {
		return err
	}

	var execQuery string
	paramLength := utils.CountParamLength(params)
	if paramLength == 0 {
		execQuery = fmt.Sprintf("SELECT %s() LIMIT 5;", sp)
		paramsBytes = nil
	} else {
		execQuery = fmt.Sprintf("SELECT %s($1) LIMIT 5;", sp)
	}

	dbBytes, err := d.ExecuteQuery(execQuery, paramsBytes)
	if err != nil {
		return fmt.Errorf("db error: %w", err)
	}

	if len(dbBytes) == 0 && model == nil {
		return nil
	}

	err = jsoniter.Unmarshal(dbBytes, model)
	if err != nil {
		return fmt.Errorf("cannot unmarshal data: %w", err)
	}

	dbError := models.DataBaseError{}
	err = jsoniter.Unmarshal(dbBytes, &dbError)
	if err != nil {
		return fmt.Errorf("cannot unmarshal data: %w", err)
	}
	if dbError.Error != "" {
		return fmt.Errorf("cannot get hero from db: %s", dbError.Error)
	}

	return nil
}

func (d *PostgresDriver) ExecuteQuery(query string, params any) ([]byte, error) {
	paramLength := utils.CountParamLengthUInt(params)
	if paramLength == 0 {
		rows, err := d.conn.Query(context.Background(), query)
		if err != nil {
			return nil, err
		}

		var result []byte
		for rows.Next() {
			if err = rows.Scan(&result); err != nil {
				return nil, err
			}
		}
		rows.Close()
		if err = rows.Err(); err != nil {
			return nil, err
		}
		return result, nil
	}

	rows, err := d.conn.Query(context.Background(), query, params)
	if err != nil {
		return nil, err
	}

	var result []byte
	for rows.Next() {
		if err = rows.Scan(&result); err != nil {
			return nil, err
		}
	}
	rows.Close()
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return result, nil
}
