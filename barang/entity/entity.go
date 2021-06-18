package entity

type Barang struct {
	Name           string   `json:"name"`
	Category       []string `json:"category"`
	Price          float64  `json:"price"`
	Sold           int      `json:"sell"`
	StoreName      string   `json:"store"`
	CountAvailable int      `json:"count_available"`
	Description    string   `json:"description"`
	Review         []review `json:"review"`
}

type review struct {
	UserName string `json:"user_name"`
	Points   int    `json:"points"`
	Comments string `json:"comments"`
}
