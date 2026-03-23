package main

import "github.com/wim-vdw/az-secret2env/cmd"

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func main() {
	cmd.SetVersion(version)
	cmd.SetCommit(commit)
	cmd.SetBuildTime(date)
	cmd.Execute()
}
