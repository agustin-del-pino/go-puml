package client

import (
	"bytes"
	"compress/flate"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
)

type PlantUMLRender string

const (
	SVG   PlantUMLRender = "svg"
	PNG   PlantUMLRender = "png"
	ASCII PlantUMLRender = "ascii"
)

const b64 = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz-_"

type PlantUMLClient interface {
	CompressDiagram([]byte) (string, error)

	Render(PlantUMLRender, []byte) (string, []byte, error)
	RenderFile(PlantUMLRender, string) (string, []byte, error)

	BytesToSVG([]byte) (string, []byte, error)
	StringToSVG(string) (string, []byte, error)
	FileToSVG(string) (string, []byte, error)

	BytesToPNG([]byte) (string, []byte, error)
	StringToPNG(string) (string, []byte, error)
	FileToPNG(string) (string, []byte, error)

	BytesToASCII([]byte) (string, []byte, error)
	StringToASCII(string) (string, []byte, error)
	FileToASCII(string) (string, []byte, error)
}

type client struct {
	url    string
	get    func(string) ([]byte, error)
	zip    func([]byte) ([]byte, error)
	encode func([]byte) string
}

func zipDiagram(b []byte) ([]byte, error) {
	w, err := flate.NewWriter(nil, 7)

	if err != nil {
		return nil, err
	}

	buf := new(bytes.Buffer)

	w.Reset(buf)

	if _, err := w.Write(b); err != nil {
		return nil, err
	}
	if e := w.Close(); e != nil {
		return nil, e
	}

	return buf.Bytes(), nil
}

func encodeDiagram(b []byte) string {
	return base64.NewEncoding(b64).EncodeToString(b)
}

func request(u string) ([]byte, error) {
	if r, err := http.Get(u); err != nil {
		return nil, err
	} else {
		defer r.Body.Close()

		if b, err := io.ReadAll(r.Body); err != nil {
			return nil, err
		} else {
			return b, nil
		}
	}
}

func (c *client) CompressDiagram(b []byte) (string, error) {
	if z, err := c.zip(b); err != nil {
		return "", err
	} else {
		return c.encode(z), nil
	}
}

func (c *client) Render(r PlantUMLRender, b []byte) (string, []byte, error) {
	if d, err := c.CompressDiagram(b); err != nil {
		return "", nil, err
	} else {
		u := fmt.Sprintf("%s/%s/%s", c.url, r, d)
		_b, err := c.get(u)
		return u, _b, err
	}
}

func (c *client) RenderFile(r PlantUMLRender, p string) (string, []byte, error) {
	if b, err := os.ReadFile(p); err != nil {
		return "", nil, err
	} else {
		return c.Render(r, b)
	}
}

func (c *client) BytesToSVG(b []byte) (string, []byte, error) {
	return c.Render(SVG, b)
}

func (c *client) StringToSVG(s string) (string, []byte, error) {
	return c.Render(SVG, []byte(s))
}

func (c *client) FileToSVG(p string) (string, []byte, error) {
	return c.RenderFile(SVG, p)
}

func (c *client) BytesToPNG(b []byte) (string, []byte, error) {
	return c.Render(PNG, b)
}

func (c *client) StringToPNG(s string) (string, []byte, error) {
	return c.Render(PNG, []byte(s))
}

func (c *client) FileToPNG(p string) (string, []byte, error) {
	return c.RenderFile(PNG, p)
}

func (c *client) BytesToASCII(b []byte) (string, []byte, error) {
	return c.Render(ASCII, b)
}

func (c *client) StringToASCII(s string) (string, []byte, error) {
	return c.Render(ASCII, []byte(s))
}

func (c *client) FileToASCII(p string) (string, []byte, error) {
	return c.RenderFile(ASCII, p)
}

func NewPlantUMLClient(u string) (PlantUMLClient, error) {
	if _u, err := url.Parse(u); err != nil {
		return nil, err
	} else {
		return &client{
			url:    _u.String(),
			zip:    zipDiagram,
			encode: encodeDiagram,
			get:    request,
		}, nil
	}
}
