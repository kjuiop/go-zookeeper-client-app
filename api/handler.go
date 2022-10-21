package api

type API struct {
}

func NewHandler() (*API, error) {
	a := new(API)
	return a, nil
}

func (a *API) Close() {
}
