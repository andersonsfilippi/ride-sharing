package domain

import (
	"time"

	tripTypes "github.com/andersonsfilippi/ride-sharing/services/trip-service/pkg/types"
	pb "github.com/andersonsfilippi/ride-sharing/shared/proto/trip"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RideFareModel struct {
	ID                primitive.ObjectID // for mongoDB
	UserID            string
	PackageSlug       string  // ex: van, luxury, sedan
	TotalPriceInCents float64 // calculated price of price
	Route             *tripTypes.OsrmApiResponse
	ExpiresAt         time.Time
}

func (r *RideFareModel) ToProto() *pb.RideFare {
	return &pb.RideFare{
		Id:                r.ID.Hex(),
		UserID:            r.UserID,
		PackageSlug:       r.PackageSlug,
		TotalPriceInCents: r.TotalPriceInCents,
	}
}

func ToRideFaresProto(fares []*RideFareModel) []*pb.RideFare {
	var protoFares []*pb.RideFare

	for _, f := range fares {
		protoFares = append(protoFares, f.ToProto())
	}
	return protoFares
}
