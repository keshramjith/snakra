package dbservice

import (
	"context"
	"fmt"
	"os"

	pgxuuid "github.com/jackc/pgx-gofrs-uuid"
	"github.com/jackc/pgx/v5"
)

type DbService struct {
	conn *pgx.Conn
}

func NewDbConn() *DbService {
	db_url := os.Getenv("DATABASE_URL")
	conn, err := pgx.Connect(context.Background(), db_url)
	if err != nil {
		fmt.Printf("Error connecting to dbconfig: %s", err)
		os.Exit(1)
	}
	pgxuuid.Register(conn.TypeMap())

	return &DbService{conn: conn}
}

func (dbs *DbService) InsertVoicenote(vn *Voicenote) error {
	args := pgx.NamedArgs{
		"id":             vn.Id.String(),
		"s3_key":         vn.S3_key,
		"created_at":     vn.Created_at,
		"url_short_form": vn.UrlShortForm,
	}
	_, err := dbs.conn.Exec(context.Background(), `INSERT INTO voicenotes (id, s3_key, created_at, url_short_form) VALUES (@id, @s3_key, @created_at, @url_short_form)`, args)
	if err != nil {
		return err
	}
	return nil
}

func (dbs *DbService) GetVoicenote(vn *Voicenote) error {
	err := dbs.conn.QueryRow(context.Background(), `SELECT * FROM voicenotes WHERE url_short_form = @url_short_form`, pgx.NamedArgs{"url_short_form": vn.UrlShortForm}).Scan(&vn.Id, &vn.S3_key, &vn.Created_at, &vn.UrlShortForm)
	if err != nil {
		return err
	}
	return nil
}

func (db *DbService) CloseDb() {
	fmt.Println("Closing db connection...")
	db.conn.Close(context.Background())
}
