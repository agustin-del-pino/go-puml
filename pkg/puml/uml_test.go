package puml

import (
	"github.com/stretchr/testify/assert"
	"testing"
)
func TestSkinparams_String(t *testing.T) {
	s := Skinparams{
		"test": "on",
	}
	assert.Equal(t, "skinparam test on\n", s.String())
}

func TestParticipant_String(t *testing.T) {
	p := Participant{
		Name:  "test test",
		Alias: "tst",
	}

	assert.Equal(t, `participant "test test" as tst`, p.String())
}

func TestArrow_String(t *testing.T) {
	t.Run("left arrow", func(t *testing.T) {
		a := Arrow{
			From:      "test",
			To:        "test",
			Message:   "test",
			Direction: ArrowLeft,
		}

		assert.Equal(t, `test<-test: test`, a.String())
	})
	t.Run("right arrow", func(t *testing.T) {
		a := Arrow{
			From:      "test",
			To:        "test",
			Message:   "test",
			Direction: ArrowRight,
		}

		assert.Equal(t, `test->test: test`, a.String())
	})
	t.Run("without message", func(t *testing.T) {
		a := Arrow{
			From:      "test",
			To:        "test",
			Direction: ArrowLeft,
		}

		assert.Equal(t, `test<-test`, a.String())
	})
}

func TestUml_Title(t *testing.T) {
	t.Run("with title", func(t *testing.T) {
		u := uml{}
		u.Title("test")
		assert.Equal(t, "test", u.title)
	})
	t.Run("without title", func(t *testing.T) {
		u := uml{}
		assert.Equal(t, "", u.title)
	})
}

func TestUml_Participant(t *testing.T) {
	u := uml{}
	u.Participant("test", "tst")
	assert.Len(t, u.participants, 1)
	assert.Equal(t, "test", u.participants[0].Name)
	assert.Equal(t, "tst", u.participants[0].Alias)
}

func TestUml_Arrow(t *testing.T) {
	t.Run("left arrow", func(t *testing.T) {
		u := uml{}
		u.ArrowL("test", "test", "test")
		assert.Len(t, u.arrows, 1)
		assert.Equal(t, "test", u.arrows[0].From)
		assert.Equal(t, "test", u.arrows[0].To)
		assert.Equal(t, "test", u.arrows[0].Message)
		assert.Equal(t, ArrowLeft, u.arrows[0].Direction)
	})

	t.Run("right arrow", func(t *testing.T) {
		u := uml{}
		u.ArrowR("test", "test", "test")
		assert.Len(t, u.arrows, 1)
		assert.Equal(t, "test", u.arrows[0].From)
		assert.Equal(t, "test", u.arrows[0].To)
		assert.Equal(t, "test", u.arrows[0].Message)
		assert.Equal(t, ArrowRight, u.arrows[0].Direction)
	})

	t.Run("without comment", func(t *testing.T) {
		u := uml{}
		u.ArrowL("test", "test", "")
		assert.Len(t, u.arrows, 1)
		assert.Equal(t, "test", u.arrows[0].From)
		assert.Equal(t, "test", u.arrows[0].To)
		assert.Equal(t, "", u.arrows[0].Message)
	})
}

func TestUml_String(t *testing.T) {
	u := uml{}
	u.Title("test").
		Participant("test", "tst").
		ArrowR("tst", "test1", "testA").
		ArrowL("test1", "tst", "testB")

	assert.Equal(t, testUml, u.String())
}

const testUml string = `@startuml
title "test"
participant "test" as tst
tst->test1: testA
test1<-tst: testB
@enduml`
