// ISC prints personalized ISC style license
package main

import (
	"flag"
	"log"
	"os"
	"os/user"
	"text/template"
	"time"
)

// ISC style license
const license = `{{define "license" -}}
{{if .Short}}{{template "banner" .}}{{else}}{{template "isc" .}}{{end}}
{{end}}

{{define "isc" -}}
Copyright (c) {{.Year}} {{.Name}}{{with .Mail}} <{{.}}>{{end}}

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
{{- end}}

{{define "banner" -}}
// Copyright (c) {{.Year}} {{.Name}}. All rights reserved.
// Use of this source code is governed by ISC-style license
// that can be found in the LICENSE file.
{{- end}}`

type page struct {
	Name  string
	Mail  string
	Year  int
	Short bool
}

func main() {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	host, err := os.Hostname()
	if err != nil {
		log.Fatal(err)
	}

	var (
		name  = flag.String("name", usr.Name, "full name")
		mail  = flag.String("mail", usr.Username+"@"+host, "mail address")
		year  = flag.Int("year", time.Now().Year(), "copyright year")
		short = flag.Bool("banner", false, "print banner")
	)
	flag.Parse()

	args := page{Name: *name, Mail: *mail, Year: *year, Short: *short}

	tmpl, err := template.New("").Parse(license)
	if err != nil {
		log.Fatal(err)
	}

	if err = tmpl.ExecuteTemplate(os.Stdout, "license", args); err != nil {
		log.Fatal(err)
	}
}
