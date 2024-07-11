package mongo

import (
	"context"

	"github.com/charmingruby/brandwl/internal/domain/search/search_entity"
	"github.com/charmingruby/brandwl/internal/infra/database/mongo/mongo_mapper"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	RESEARCH_COLLETION = "search_result"
)

func NewMongoSearchRepository(db *mongo.Database) *MongoSearchRepository {
	return &MongoSearchRepository{
		db: db,
	}
}

type MongoSearchRepository struct {
	db *mongo.Database
}

func (r *MongoSearchRepository) Store(dr *search_entity.DomainResearch) error {
	mdr := mongo_mapper.DomainResearchToMongo(dr)

	collection := r.db.Collection(RESEARCH_COLLETION)

	if _, err := collection.InsertOne(context.Background(), mdr); err != nil {
		return err
	}

	return nil
}
