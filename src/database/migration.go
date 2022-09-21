package database

import (
	"fmt"
	"os"
	"pismo-transactions/src/errs"
	"pismo-transactions/src/utils"
)

func ExecuteMigration() error {
	Conn, err := Conn()
	if err != nil {
		return fmt.Errorf("%s %w", errs.ERR_DB_CONN, err)
	}

	utils.LoadEnv()
	path := os.Getenv("SQL_SCRIPTS_DIR")

	files, err := os.ReadDir(path)
	if err != nil {
		return fmt.Errorf("%s %w", errs.ERR_READ_FILES, err)
	}

	for _, file := range files {
		f, err := os.ReadFile(path + file.Name())
		if err != nil {
			return fmt.Errorf("%s %w", errs.ERR_READ_FILES, err)
		}

		sql := string(f)
		_, err = Conn.Exec(sql)
		if err != nil {
			return fmt.Errorf("%s %w", errs.ERR_EXEC_QUERY, err)
		}
	}
	return nil
}
