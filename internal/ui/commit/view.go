package commit

type View struct {
	headers []string
	content string
}

func NewView() *View {
	return &View{
		content: "",
	}
}

func (v *View) SelectHeader() error {

	return nil
}

func (v *View) Render() string {
	return v.content
}
