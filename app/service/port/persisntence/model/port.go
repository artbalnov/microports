package model

// Port business model
type Port struct {
	ID          string    `bson:"_id"`
	Name        string    `bson:"name"`
	Code        string    `bson:"code"`
	Alias       []string  `bson:"alias"`
	Unlocs      []string  `bson:"unlocs"`
	Country     string    `bson:"country"`
	Regions     []string  `bson:"regions"`
	Province    string    `bson:"province"`
	City        string    `bson:"city"`
	Coordinates []float32 `bson:"coordinates"`
	Timezone    string    `bson:"timezone"`
}
