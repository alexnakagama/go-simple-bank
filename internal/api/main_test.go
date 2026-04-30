package api

import (
	"os"
	"testing"
	"time"

	"github.com/alexnakagama/go-simple-bank/config"
	db "github.com/alexnakagama/go-simple-bank/internal/db/sqlc"
	"github.com/alexnakagama/go-simple-bank/util"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

func newTestServer(t *testing.T, store db.Store) *Server {
	config := config.Config{
		TokenSymmetricKey:   util.RandomString(32),
		AccessTokenDuration: time.Minute,
	}

	server, err := NewServer(config, store)
	require.NoError(t, err)

	return server
}

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	os.Exit(m.Run())
}
