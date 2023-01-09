package database

import (
	"aozorarodoku-service/base/uuid"
	"context"

	"github.com/jmoiron/sqlx"
)

type Content struct {
	Id             uuid.UUID `db:"id"`
	TitleRuby      string    `db:"title_ruby"`
	Title          string    `db:"title"`
	AuthorRuby     string    `db:"author_ruby"`
	Author         string    `db:"author"`
	SpeakerRuby    string    `db:"speaker_ruby"`
	Speaker        string    `db:"speaker"`
	FileName       string    `db:"file_name"`
	NewArrivalDate string    `db:"new_arrival_date"`
	Time           string    `db:"time"`
}

func FindContents(ctx context.Context, db *sqlx.DB) ([]Content, error) {
	query := `SELECT * FROM contents`
	rows, err := db.Queryx(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]Content, 0)
	for rows.Next() {
		var c Content
		err := rows.StructScan(&c)
		if err != nil {
			return nil, err
		}
		results = append(results, c)
	}
	return results, nil
}

func InsertContent(ctx context.Context, tx *sqlx.Tx, c ...Content) error {
	query := `INSERT INTO contents (id, title_ruby, title, author_ruby, author, speaker_ruby, speaker, file_name, new_arrival_date, time) VALUES (:id, :title_ruby, :title, :author_ruby, :author, :speaker_ruby, :speaker, :file_name, :new_arrival_date, :time);`

	if err := ExecQuery(ctx, tx, query, c); err != nil {
		return err
	}
	return nil
}

func TruncateContents(ctx context.Context, tx *sqlx.Tx) error {
	query := `TRUNCATE TABLE contents;`
	if err := ExecQuery(ctx, tx, query, nil); err != nil {
		return err
	}
	return nil
}
