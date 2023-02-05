package utils

import (
	"errors"
	"fmt"
	"strings"

	"github.com/jackc/pgconn"
)

// Функция форматирования запроса
// Декоративная история, чтобы запросы смотрелись понятнее
func FormatQuery(q string) string {
	return strings.ReplaceAll(strings.ReplaceAll(q, "\t", ""), "\n", " ")
}

// Функция преобразования ошибки
func FormatErr(err error) error {
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		pgErr = err.(*pgconn.PgError)
		err = fmt.Errorf(
			fmt.Sprintf(
				"SQL Error: %s, Detail: %s, Where: %s, Code: %s, SQLState: %s",
				pgErr.Message,
				pgErr.Detail,
				pgErr.Where,
				pgErr.Code,
				pgErr.SQLState(),
			))
	}
	return err
}
