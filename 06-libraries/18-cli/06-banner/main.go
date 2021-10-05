package main

import (
	"github.com/dimiro1/banner"
	"github.com/mattn/go-colorable"
)

func init() {
	template := `{{ .Title "Here We Go" "" 4 }}
   {{ .AnsiColor.BrightCyan }}The title will be ascii and indented 4 spaces{{ .AnsiColor.Default }}
   GoVersion: {{ .GoVersion }}
   GOOS: {{ .GOOS }}
   GOARCH: {{ .GOARCH }}
   NumCPU: {{ .NumCPU }}
   GOPATH: {{ .GOPATH }}
   GOROOT: {{ .GOROOT }}
   Compiler: {{ .Compiler }}
   ENV: {{ .Env "GOPATH" }}
   Now: {{ .Now "Monday, 2 Jan 2006" }}
   {{ .AnsiColor.BrightGreen }}This text will appear in Green
   {{ .AnsiColor.BrightRed }}This text will appear in Red{{ .AnsiColor.Default }}`

	banner.InitString(colorable.NewColorableStdout(), true, true, template)
}

func main() {

}