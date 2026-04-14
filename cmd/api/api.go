package main

type Api struct {
	DataStore *Store
}

func NewApi(store *Store) *Api {
	return &Api{
		DataStore: store,
	}
}
func (api *Api) Start() {

}
