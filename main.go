package main

import (
	"flag"
	"log"
	"os"
	"os/user"
	"text/template"
	"time"
)

const ISC = `Copyright (c) {{.Year}} {{.Name}}{{with .Mail}} <{{.}}>{{end}}

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

func main() {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	host, err := os.Hostname()
	if err != nil {
		log.Fatal(err)
	}

	name := flag.String("name", usr.Name, "Full name")
	mail := flag.String("mail", usr.Username+"@"+host, "Mail address")
	year := flag.Int("year", time.Now().Year(), "Copyright year")
	flag.Parse()

	l := struct {
		Name, Mail string
		Year       int
	}{
		Name: *name,
		Mail: *mail,
		Year: *year,
	}

	template.Must(template.New("ISC").Parse(ISC)).Execute(os.Stdout, l)
}
