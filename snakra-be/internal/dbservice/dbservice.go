package dbservice

import (
	"fmt"
	"github.com/gocql/gocql"
	"github.com/scylladb/gocqlx/v2"
	"github.com/scylladb/gocqlx/v2/table"
	"os"
)

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
		Name:    "testdb.voicenotes",
		Columns: []string{"id", "s3_key", "created_at"},
		PartKey: []string{"id"},
	}

	vnTable = table.New(voicenoteMetadata)

	return &DbService{session: &session}
}

func (dbs *DbService) InsertVoicenote(vn *Voicenote) {
	insertQ := dbs.session.Query(vnTable.Insert()).BindStruct(vn)
	insertQ.BindStruct(vn)
	if err := insertQ.ExecRelease(); err != nil {
		fmt.Printf("Error inserting in db: %s\n", err)
		os.Exit(1)
	}
}

func (dbs *DbService) GetVoicenote(vn *Voicenote) {
	getQ := dbs.session.Query(vnTable.Get()).BindStruct(vn)
	if err := getQ.GetRelease(vn); err != nil {
		fmt.Printf("Error getting from db: %s\n", err)
		os.Exit(1)
	}
}

func (db *DbService) CloseDb() {
	db.session.Close()
}
