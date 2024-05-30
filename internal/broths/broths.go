package broths

type Broth struct {
	ID            string  `json:"id"`
	ImageInactive string  `json:"imageInactive"`
	ImageActive   string  `json:"imageActive"`
	Name          string  `json:"name"`
	Description   string  `json:"description"`
	Price         float64 `json:"price"`
}

// List of available broths
var broths = []Broth{
	{
		ID:            "1",
		ImageInactive: "https://tech.redventures.com.br/icons/salt/inactive.svg",
		ImageActive:   "https://tech.redventures.com.br/icons/salt/active.svg",
		Name:          "Salt",
		Description:   "Simple like the seawater, nothing more",
		Price:         10,
	},
	// ToDO: Add more broths here
}

// Get all available broths
func List() []Broth {
	return broths
}
