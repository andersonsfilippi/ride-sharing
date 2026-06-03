package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RideFareModel struct {
	ID                primitive.ObjectID // for mongoDB
	UserID            string
	PackageSlug       string  // ex: van, luxury, sedan
	TotalPriceInCents float64 // calculated price of price
	ExpiresAt         time.Time
}
