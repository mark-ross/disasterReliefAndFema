package fema

type MetaData struct {
	Skip   int    `json:"skip"`
	Top    int    `json:"top"`
	Count  int    `json:"count"`
	Filter string `json:"filter"`
	Format string `json:"format"`
	//OrderBy // FIXME
	Select     []string `json:"select"`
	EntityName string   `json:"entityname"`
	Version    string   `json:"v2"`
	URL        string   `json:"url"`
	RunDate    string   `json:"rundate"`
}
