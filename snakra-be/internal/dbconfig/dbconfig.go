package dbconfig

import (
	"fmt"
	"os"

	"github.com/gocql/gocql"
	"github.com/scylladb/gocqlx/v2"
)

func NewDbConn() *gocqlx.Session {
	cluster := gocql.NewCluster("localhost:9042")
	session, err := gocqlx.WrapSession(cluster.CreateSession())
	if err != nil {
		fmt.Printf("Error connecting to dbconfig: %s", err)
		os.Exit(1)
	}
	return &session
}
