package utils

import (
	"go.mongodb.org/mongo-driver/mongo"

	"errors"
)

// Return the MongoDB collection matching the desired selection
func ImageCollectionSelection (mongoClient *mongo.Client, selection string) (*mongo.Collection, error) {
	switch selection {
	case "flickr":
		return mongoClient.Database(DotEnvVariable("SCRAPPER_DB")).Collection(DotEnvVariable("FLICKR_COLLECTION")), nil
	case "unsplash":
		return mongoClient.Database(DotEnvVariable("SCRAPPER_DB")).Collection(DotEnvVariable("UNSPLASH_COLLECTION")), nil
	case "pexels":
		return mongoClient.Database(DotEnvVariable("SCRAPPER_DB")).Collection(DotEnvVariable("PEXELS_COLLECTION")), nil
	default:
		return nil, errors.New("Collection does not exist")
	}
}

// Returns a map of all collections with images
func ImageCollections (mongoClient *mongo.Client) map[string]*mongo.Collection {
	collections := make(map[string]*mongo.Collection)
	collections["flickr"] = mongoClient.Database(DotEnvVariable("SCRAPPER_DB")).Collection(DotEnvVariable("FLICKR_COLLECTION"))
	collections["unsplash"] = mongoClient.Database(DotEnvVariable("SCRAPPER_DB")).Collection(DotEnvVariable("UNSPLASH_COLLECTION"))
	collections["pexels"] = mongoClient.Database(DotEnvVariable("SCRAPPER_DB")).Collection(DotEnvVariable("PEXELS_COLLECTION"))
	return collections
}

func ImageOrigins () []string {
	return []string{"flickr", "unsplash", "pexels"}
}
