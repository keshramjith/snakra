package dbservice

import (
	"time"

	"github.com/gofrs/uuid/v5"
)

type Voicenote struct {
	Id         uuid.UUID
	S3_key     string
	Created_at time.Time
}
