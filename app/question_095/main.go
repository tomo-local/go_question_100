package main

type Datastore interface {
	GetUser(id int) string
}

type App struct {
	store Datastore
}

func (a *App) ShowUser(id int) string {
	return a.store.GetUser(id)
}
