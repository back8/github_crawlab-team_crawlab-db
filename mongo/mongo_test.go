package mongo

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func setupMongoTest() (err error) {
	return nil
}

func cleanupMongoTest() {
}

func TestMongoInitMongo(t *testing.T) {
	err := setupMongoTest()
	require.Nil(t, err)

	err = InitMongo()
	require.Nil(t, err)

	cleanupMongoTest()
}
