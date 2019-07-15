// ISC prints personalized ISC style license
package main

import (
	"flag"
	"io"
	"log"
	"net/mail"
	"os"
	"os/user"
	"text/template"
	"time"
)

// ISC style license
const license = `Copyright (c) {{.Year}} {{.Name}} <{{.Mail}}>

Permission to use, copy, modify, and distribute this software for any
purpose with or without fee is hereby granted, provided that the above
copyright notice and this permission notice appear in all copies.

THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES
WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF
MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR
ANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES
WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN
ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF
OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE.
`

type page struct {
	Name string
	Mail string
	Year int
}

func execute(w io.Writer, p page) error {
	tmpl, err := template.New("").Parse(license)
	if err != nil {
		return err
	}
	return tmpl.Execute(w, p)
}

func owner() (*mail.Address, error) {
	usr, err := user.Current()
	if err != nil {
		return nil, err
	}
	host, err := os.Hostname()
	if err != nil {
		return nil, err
	}
	return &mail.Address{
		Name:    usr.Name,
		Address: usr.Username + "@" + host,
	}, nil
}

func main() {
	usr, err := owner()
	if err != nil {
		log.Fatal(err)
	}
	var (
		name = flag.String("name", usr.Name, "full name")
		mail = flag.String("mail", usr.Address, "mail address")
		year = flag.Int("year", time.Now().Year(), "copyright year")
	)
	flag.Parse()
	args := page{Name: *name, Mail: *mail, Year: *year}
	if err = execute(os.Stdout, args); err != nil {
		log.Fatal(err)
	}

}
