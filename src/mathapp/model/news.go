package model

import (
	"database/sql"
)

type News struct {
	id           int64
	title        sql.NullString
	content      sql.NullString
	status       sql.NullString
	creationDate sql.NullString
	createdId    sql.NullString
	createdBy    sql.NullString
	changeDate   sql.NullString
	changedId    sql.NullString
	changedBy    sql.NullString
	_type        sql.NullString
}
