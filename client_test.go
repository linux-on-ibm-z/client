package client

import (
	"os"
	"path/filepath"
	"testing"

	sourcepath "github.com/GeertJohan/go-sourcepath"
	"github.com/rai-project/config"
	"github.com/stretchr/testify/assert"
)

// TestClient ...
func DISABLED_TestClient(t *testing.T) {
	clt, err := New(
		Directory(filepath.Join(sourcepath.MustAbsoluteDir(), "_fixtures")),
		BuildFileBaseName("rai_build"),
		DisableRatelimit(),
	)
	if !assert.NoError(t, err) {
		return
	}
	assert.NotNil(t, clt)

	err = clt.Validate()
	assert.NoError(t, err)

	err = clt.Upload()
	if !assert.NoError(t, err) {
		return
	}
	assert.NotEmpty(t, clt.uploadKey, "upload key must be set after upload")

	err = clt.Publish()
	if !assert.NoError(t, err) {
		return
	}
	err = clt.Subscribe()
	if !assert.NoError(t, err) {
		return
	}

	err = clt.Connect()
	if !assert.NoError(t, err) {
		return
	}

	clt.Wait()

	defer clt.Disconnect()
}

// TestMain ...
func TestMain(m *testing.M) {
	config.Init(
		config.VerboseMode(true),
		config.DebugMode(true),
		config.ColorMode(true),
	)
	os.Exit(m.Run())
}
