package api

type Usecase interface {
	FetchQuery() (string, error)
	SetQuery(name string) error
	ChageQuery(name string) error
}
