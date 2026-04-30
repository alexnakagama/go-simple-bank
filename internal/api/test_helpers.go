package api

import (
	"testing"
	"time"

	"github.com/alexnakagama/go-simple-bank/config"
	db "github.com/alexnakagama/go-simple-bank/internal/db/sqlc"
	"github.com/alexnakagama/go-simple-bank/util"
	"github.com/stretchr/testify/require"
)

func NewTestServer(t *testing.T, store db.Store) *Server {
	config := config.Config{
		TokenSymmetricKey:   util.RandomString(32),
		AccessTokenDuration: time.Minute,
	}

	server, err := NewServer(config, store)
	require.NoError(t, err)

	return server
}
