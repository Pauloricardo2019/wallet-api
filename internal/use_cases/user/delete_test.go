package user_test

import (
	"wallet-api/internal/use_cases/user"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"testing"
)

func Test_DeleteById_OK(t *testing.T) {

	const userID uint64 = 123
	var tokenUserID uint64 = 123

	deleteById := func(ID uint64) error {
		assert.Equal(t, userID, ID)

		var userDeleted = make(map[uint64]uint64)
		userDeleted[123] = 123
		delete(userDeleted, ID)

		return nil

	}

	tx := user.Delete(tokenUserID, userID, deleteById)

	require.NoError(t, tx)

}
