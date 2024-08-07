package repository

import (
	"regexp"
	"testing"

	mock_db "github.com/itzaddddd/ticket-reserve/util/mock/db"
	"github.com/stretchr/testify/assert"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestDecreseQuota(t *testing.T) {
	db, mock := mock_db.NewMockDb()

	numToDecrese := 1
	eventId := 1

	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(`UPDATE "events" SET "remain_quota"`)).
		WithArgs(numToDecrese, sqlmock.AnyArg(), eventId).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	repo := NewRepository(db)
	err := repo.DecreseQuota(uint(eventId), uint(numToDecrese))
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())

}
