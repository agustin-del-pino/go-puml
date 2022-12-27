# Go PULM
Go API for Plant UML

# Constructors

## `NewPlantUML() PlantUML`
Instance a new PlantUML implementation.

````go
puml := NewPlantUML()
````

## `NewPlantUMLClient`
Instance a new PlantUMLClient implementation.
````go
cli := NewPlantUMLClient()
````

# API Reference

## PlantUML Interface

### `StartUml() UML`
Start a new UML diagram and returns its implementation.

````go
uml := puml.StartUml()
````

## UML Interface

### `Title(title string) UML`
Adds a title to the diagram.

````go
uml.Title("The UML Diagram")
````

### `Skinparam(param string, value string) UML`
Adds a Skinparam to the diagram.

````go
uml.Skinparam("responseMessageBelowArrow", "true")
````

### `Participant(displayName string, name string) UML`
Adds a new Participant to the diagram.

````go
uml.Participant("Service", "srv")
````

*It will translate in puml as: `participant Service as srv`*

### `ArrowR(from string, to string, msg string) UML`
Adds an arrow to the diagram.

````go
uml.ArrowR("srv1", "srv2", "msg")
````
*It will translate in puml as: `srv1->srv2: msg`*


### `ArrowL(from string, to string, msg string) UML`
Adds an arrow with style to the diagram

````go
uml.ArrowL("srv1", "srv2", "msg")
````

*It will translate in puml as: `srv1<-srv2: msg`*

### `String() string`
Returns the UML diagram as string representation. Understands this method as *Marshal*.

````go
diagram := uml.String()
````

## PlantUMLClient
Rest Client that render a PlantUML diagram.

### `Render(renderType PlantUMLRender, diagram []byte) (string, []byte, error)`

Render a diagram from bytes to specified Render Type. 

Retrieves the: Generated URL for render, Rendered Diagram, Error.


````go
url, render, err := cli.Render(PlantUMLRender, []byte)
````

### `RenderFile(PlantUMLRender, string) (string, []byte, error)`
Render a diagram from file to specified Render Type. 

Retrieves the: Generated URL for render, Rendered Diagram, Error.

````go
url, render, err := cli.RenderFile(PlantUMLRender, string)
````

### `BytesToSVG([]byte) (string, []byte, error)`
Render a diagram from bytes to SVG. 

Retrieves the: Generated URL for render, Rendered Diagram, Error.

````go
url, render, err := cli.BytesToSVG([]byte(`@startuml
actor1->actor2: Hello World
@enduml`))
````

### `StringToSVG(string) (string, []byte, error)`
Render a diagram from string to SVG. 

Retrieves the: Generated URL for render, Rendered Diagram, Error.

````go
url, render, err := cli.StringToSVG(`@startuml
actor1->actor2: Hello World
@enduml`)
````

### `FileToSVG(string) (string, []byte, error)`
Render a diagram from file to SVG. 

Retrieves the: Generated URL for render, Rendered Diagram, Error.

````go
url, render, err := cli.FileToSVG("./my-diagram.puml")
````

### `BytesToPNG([]byte) (string, []byte, error)`
Render a diagram from bytes to PNG. 

Retrieves the: Generated URL for render, Rendered Diagram, Error.

````go
url, render, err := cli.BytesToPNG([]byte(`@startuml
actor1->actor2: Hello World
@enduml`))
````

### `StringToPNG(string) (string, []byte, error)`
Render a diagram from string to PNG. 

Retrieves the: Generated URL for render, Rendered Diagram, Error.

````go
url, render, err := cli.StringToPNG(`@startuml
actor1->actor2: Hello World
@enduml`)
````

### `FileToPNG(string) (string, []byte, error)`
Render a diagram from file to PNG. 

Retrieves the: Generated URL for render, Rendered Diagram, Error.

````go
url, render, err := cli.FileToPNG("./my-diagram.puml")
````

### `BytesToASCII([]byte) (string, []byte, error)`
Render a diagram from bytes to ASCII. 

Retrieves the: Generated URL for render, Rendered Diagram, Error.

````go
url, render, err := cli.BytesToASCII([]byte(`@startuml
actor1->actor2: Hello World
@enduml`))
````

### `StringToASCII(string) (string, []byte, error)`
Render a diagram from string to ASCII. 

Retrieves the: Generated URL for render, Rendered Diagram, Error.

````go
url, render, err := cli.StringToASCII(`@startuml
actor1->actor2: Hello World
@enduml`)
````

### `FileToASCII(string) (string, []byte, error)`
Render a diagram from file to ASCII. 

Retrieves the: Generated URL for render, Rendered Diagram, Error.

````go
url, render, err := cli.FileToASCII("./my-diagram.puml")
````

