package types

type ReviewDTO struct{
	Info  string `json:"info"`
	Stars string `json:"stars"`
}

type ReviewsDTO struct {
	Reviews []ReviewDTO	`json:"reviews"`
}
