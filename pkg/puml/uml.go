package puml

import (
	"fmt"
	"strings"
)

type ArrowDirection int

const (
	ArrowRight ArrowDirection = iota
	ArrowLeft
)

type UML interface {
	Title(string) UML
	Skinparam(string, string) UML
	Participant(string, string) UML
	ArrowR(string, string, string) UML
	ArrowL(string, string, string) UML
	String() string
}

type Participant struct {
	Name  string
	Alias string
}

func (p Participant) String() string {
	return fmt.Sprintf(`participant "%s" as %s`, p.Name, p.Alias)
}

type Arrow struct {
	From      string
	To        string
	Message   string
	Direction ArrowDirection
}

func (a Arrow) String() string {
	b := strings.Builder{}

	b.WriteString(a.From)

	if a.Direction == ArrowRight {
		b.WriteString("->")
	} else {
		b.WriteString("<-")
	}

	b.WriteString(a.To)

	if a.Message != "" {
		b.WriteString(fmt.Sprintf(`: %s`, a.Message))
	}

	return b.String()
}

type Skinparams map[string]string

func (s Skinparams) String() string {
	b := strings.Builder{}

	for k, v := range s {
		b.WriteString(fmt.Sprintf("skinparam %s %s\n", k, v))
	}

	return b.String()
}

type uml struct {
	title        string
	skinparams   Skinparams
	participants []Participant
	arrows       []Arrow
}

func (u *uml) Title(t string) UML {
	u.title = t
	return u
}

func (u *uml) Skinparam(p string, v string) UML {
	u.skinparams[p] = v
	return u
}

func (u *uml) Participant(n string, a string) UML {
	u.participants = append(u.participants, Participant{
		Name:  n,
		Alias: a,
	})
	return u
}

func (u *uml) ArrowR(f string, t string, m string) UML {
	u.arrows = append(u.arrows, Arrow{
		From:      f,
		To:        t,
		Message:   m,
		Direction: ArrowRight,
	})
	return u
}

func (u *uml) ArrowL(f string, t string, m string) UML {
	u.arrows = append(u.arrows, Arrow{
		From:      f,
		To:        t,
		Message:   m,
		Direction: ArrowLeft,
	})
	return u
}

func (u *uml) String() string {
	b := strings.Builder{}

	b.WriteString("@startuml\n")

	if u.title != "" {
		b.WriteString(fmt.Sprintf(`title "%s"%s`, u.title, "\n"))
	}

	if len(u.skinparams) > 0 {
		b.WriteString(fmt.Sprintf("%s\n", u.skinparams))
	}

	for _, p := range u.participants {
		b.WriteString(fmt.Sprintf("%s\n", p))
	}

	for _, a := range u.arrows {
		b.WriteString(fmt.Sprintf("%s\n", a))
	}

	b.WriteString("@enduml")

	return b.String()
}
