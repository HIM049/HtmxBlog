package template

type App struct {
	PageTitle  string
	Navigation []NavigationItem
}

type NavigationItem struct {
	Name string
	Url  string
}
