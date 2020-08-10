package types

type ReviewDTO struct{
	Info  string `json:"info"`
	Stars string `json:"stars"`
}

type Review struct{
	info 		string
	stars 		string
}

func NewReview(r ReviewDTO) Review{
	return Review{
		info: r.Info,
		stars: r.Stars,
	}
}