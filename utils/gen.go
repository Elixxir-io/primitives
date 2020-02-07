////////////////////////////////////////////////////////////////////////////////
// Copyright © 2018 Privategrity Corporation                                   /
//                                                                             /
// All rights reserved.                                                        /
////////////////////////////////////////////////////////////////////////////////

// The following directive is necessary to make the package coherent:

// This program generates cmd/version_vars.go. It can be invoked by running
// go generate
package utils

import (
	"bufio"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
	"text/template"
	"time"
)

// Version file generation consumed by higher-level repos
func GenerateVersionFile() {
	gitversion := GenerateGitVersion()
	deps := ReadGoMod()

	f, err := os.Create("version_vars.go")
	die(err)
	defer f.Close()

	packageTemplate.Execute(f, struct {
		Timestamp    time.Time
		GITVER       string
		DEPENDENCIES string
	}{
		Timestamp:    time.Now(),
		GITVER:       gitversion,
		DEPENDENCIES: deps,
	})
}

// Determine current Git version information
func GenerateGitVersion() string {
	cmd := exec.Command("git", "show", "--oneline")
	stdoutStderr, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(strings.NewReader(string(stdoutStderr)))
	for scanner.Scan() {
		return scanner.Text()
	}
	return "UNKNOWNVERSION"
}

// Read in go modules file
func ReadGoMod() string {
	r, _ := ioutil.ReadFile("../go.mod")
	return string(r)
}

// Exit the program
func die(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// Template for version_vars.go
var packageTemplate = template.Must(template.New("").Parse(
	"// Code generated by go generate; DO NOT EDIT.\n" +
		"// This file was generated by robots at\n" +
		"// {{ .Timestamp }}\n" +
		"package cmd\n\n" +
		"const GITVERSION = `{{ .GITVER }}`\n" +
		"const SEMVER = \"1.0.0\"\n" +
		"const DEPENDENCIES = `{{ .DEPENDENCIES }}`\n"))