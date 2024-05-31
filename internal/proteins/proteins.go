package proteins

type Protein struct {
	ID            string  `json:"id"`
	ImageInactive string  `json:"imageInactive"`
	ImageActive   string  `json:"imageActive"`
	Name          string  `json:"name"`
	Description   string  `json:"description"`
	Price         float64 `json:"price"`
}

// List of available proteins
var proteins = []Protein{
	{
		ID:            "1",
		ImageInactive: "https://tech.redventures.com.br/icons/pork/inactive.svg",
		ImageActive:   "https://tech.redventures.com.br/icons/pork/active.svg",
		Name:          "Chasu",
		Description:   "A sliced flavourful pork meat with a selection of season vegetables.",
		Price:         10,
	},
	// Add more proteins here
}

// List returns a list of available proteins
func List() []Protein {
	return proteins
}
