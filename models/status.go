package models

type Status struct{
	Stats struct{
		Water int `json:"water"`
		Wind int `json:"wind"`
	} `json:"stats"`
}