package dbservice

import (
	"github.com/gocql/gocql"
	"time"
)

type Voicenote struct {
	Id         gocql.UUID
	S3_key     string
	Created_at time.Time
}
