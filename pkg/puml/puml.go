package puml

type PlantUML interface {
	StartUML() UML
}

type plantuml struct{}

func (p *plantuml) StartUML() UML {
	return &uml{
		skinparams: make(Skinparams),
	}
}

func NewPlantUML() PlantUML {
	return &plantuml{}
}
