package main

import (
	"log"

	"github.com/Drafteame/taskrun/cmd/list"
	"github.com/Drafteame/taskrun/cmd/print"
	"github.com/Drafteame/taskrun/cmd/root"
	"github.com/Drafteame/taskrun/cmd/run"
)

func main() {
	cmd := root.GetCommand()

	cmd.AddCommand(
		list.GetCommand(root.GetStage(), root.GetJobsFile()),
	)
	cmd.AddCommand(
		print.GetCommand(root.GetStage(), root.GetJobsFile()),
	)
	cmd.AddCommand(
		run.GetCommand(root.GetStage(), root.GetJobsFile(), root.GetDebug()),
	)

	if err := root.GetCommand().Execute(); err != nil {
		log.Fatal(err)
	}
}
