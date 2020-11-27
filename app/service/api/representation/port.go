package representation

type UploadPortsRequest map[string]*UploadPortRequest

type UploadPortRequest struct {
	Name        string    `json:"name"`
	Coordinates []float32 `json:"coordinates"`
	City        string    `json:"city"`
	Province    string    `json:"province"`
	Country     string    `json:"country"`
	Alias       []string  `json:"alias"`
	Regions     []string  `json:"regions"`
	Timezone    string    `json:"timezone"`
	Unlocs      []string  `json:"unlocs"`
	Code        string    `json:"code"`
}

type GetPortsResponse struct {
	Ports []*GetPortResponse `json:"ports"`
}

type GetPortResponse struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Coordinates []float32 `json:"coordinates,omitempty"`
	City        string    `json:"city"`
	Province    string    `json:"province"`
	Country     string    `json:"country"`
	Alias       []string  `json:"alias,omitempty"`
	Regions     []string  `json:"regions,omitempty"`
	Timezone    string    `json:"timezone"`
	Unlocs      []string  `json:"unlocs"`
	Code        string    `json:"code"`
}
