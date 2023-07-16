package dbservice

import (
	"fmt"
	"os"
	"time"

	"github.com/gocql/gocql"
	"github.com/scylladb/gocqlx/v2"
	"github.com/scylladb/gocqlx/v2/qb"
	"github.com/scylladb/gocqlx/v2/table"
)

type Voicenote struct {
	Id         gocql.UUID
	S3_key     string
	Created_at time.Duration
}

type DbService struct {
	session *gocqlx.Session
}

var vnTable *table.Table

func NewDbConn() *DbService {
	cluster := gocql.NewCluster("localhost:9042")
	session, err := gocqlx.WrapSession(cluster.CreateSession())
	if err != nil {
		fmt.Printf("Error connecting to dbconfig: %s", err)
		os.Exit(1)
	}

	var voicenoteMetadata = table.Metadata{
		Name:    "voicenotes",
		Columns: []string{"id", "s3_key", "created_at"},
		PartKey: []string{"id"},
	}

	vnTable = table.New(voicenoteMetadata)

	return &DbService{session: &session}
}

func (dbs *DbService) InsertVoicenote(vn *Voicenote) {
	insertQ := qb.Insert(fmt.Sprintf(`%s.voicenotes`, "testdb")).Columns("id", "s3_key", "created_at").Query(*dbs.session)
	insertQ.BindStruct(vn)
	if err := insertQ.ExecRelease(); err != nil {
		fmt.Printf("Error inserting in db: %s\n", err)
		os.Exit(1)
	}
}

func (db *DbService) CloseDb() {
	db.session.Close()
}
