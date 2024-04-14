package documents

import (
	"context"
	"doc-search-app-backend/internal/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r *Repository) GetYears(ctx context.Context) ([]entity.DocumentsInYearCount, error) {
	pipeline := bson.A{
		bson.D{
			{"$project",
				bson.D{
					{"year",
						bson.D{
							{"$year",
								bson.D{
									{"$cond",
										bson.D{
											{"if",
												bson.D{
													{"$eq",
														bson.A{
															"$article.published",
															"",
														},
													},
												},
											},
											{"then", primitive.Null{}},
											{"else", bson.D{{"$toDate", "$article.published"}}},
										},
									},
								},
							},
						},
					},
				},
			},
		},
		bson.D{
			{"$group",
				bson.D{
					{"_id", "$year"},
					{"count", bson.D{{"$count", bson.D{}}}},
				},
			},
		},
		bson.D{
			{"$project",
				bson.D{
					{"_id", 0},
					{"year", "$_id"},
					{"count", 1},
				},
			},
		},
		bson.D{{"$sort", bson.D{{"year", -1}}}},
	}

	cursor, err := r.docs.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}

	var years []entity.DocumentsInYearCount
	if err = cursor.All(ctx, &years); err != nil {
		return nil, err
	}

	return years, nil
}
