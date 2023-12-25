package wrappers_test

import (
	"context"
	"fmt"
	"testing"

	_ "github.com/golang-migrate/migrate/v4/database/mongodb"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/komandakycto/integro/wrappers"
)

func TestNew_MongoDB(t *testing.T) {
	image := "mongo:7"

	// Create a new test container with MongoDB image
	container, err := wrappers.New(image)
	require.NoError(t, err)

	// Apply migration from file
	err = container.Migrate(fmt.Sprintf("file://%s", "./../internal/test_data/mongodb/migrations/"))
	require.NoError(t, err)

	// Connect to the database
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(container.Conn()))
	require.NoError(t, err)

	// Execute a query to the database
	collection := client.Database("public").Collection("users")
	cursor, err := collection.Find(context.Background(), bson.M{})
	require.NoError(t, err)
	var results []bson.M
	err = cursor.All(context.Background(), &results)
	require.NoError(t, err)
	require.Len(t, results, 3)

	// Close the connection
	err = client.Disconnect(context.Background())
	require.NoError(t, err)

	// Stop the container
	err = container.Stop()
	require.NoError(t, err)
}
