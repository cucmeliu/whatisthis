package models

type one_rec struct {
	name  string
	score float64
}

type Result_rec struct {
	Log_id  string
	Results []*one_rec
}
