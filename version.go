package main

import "github.com/hashicorp/go-version"

const Version = "0.1.0"

var GitCommit string
var SemVersion = version.Must(version.NewVersion(Version))
