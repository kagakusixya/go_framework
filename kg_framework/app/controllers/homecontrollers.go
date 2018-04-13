package controllers

/*
type T struct {
	Url       string
	Html_data map[string]string
}
*/
func (t *T) Homehellowrld() {
	t.Html_data = map[string]string{
		"key1": "hello",
		//"key2": name,
		//"ket3": say,
	}
	t.Url = "keiziban.html.tpl"
}

func (t *T) Homehell() {
	t.Url = "tesut.html.tpl"
}
