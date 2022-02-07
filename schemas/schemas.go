package schemas

type AboutMeSchema struct {
	Alias string `form:"alias"`
}

type Resume struct {
	Section string `form:"section"`
}
