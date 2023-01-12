package character

type TemplateClass struct {
	FactionGroup uint8
	ClassID      uint8
}

type Template struct {
	TemplateSetID uint32
	Classes       []TemplateClass
	Name          string
	Description   string
	Level         uint8
}
