package domain


type Modulos struct {
	Name   string
	Videos []Videos
}
type Videos struct {
	Name string
	ID   int
}