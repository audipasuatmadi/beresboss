package usecase

import (
	"log"
	"os"
	"os/exec"

	"github.com/audipasuatmadi/beresbos/internal/dto"
	"github.com/audipasuatmadi/beresbos/internal/util/dirutil"
)

func RunCommands(args dto.RunCommandsArgs) {
	var absDir string
	var err error

	initialDir, err := os.Getwd()
	if err != nil {
		log.Fatalf("cannot get initial directory: %+v", err)
	}

	for _, directory := range args.Directories {
		absDir, err = dirutil.GetAbsDirectory(directory)
		if err != nil {
			log.Fatalf("cannot get abs directory: %+v", err)
		}

		err = os.Chdir(absDir)
		if err != nil {
			log.Fatalf("error checking out to %s, err :%+v", absDir, err)
		}

		cmd := prepareCommand(args)

		err := cmd.Run()
		if err != nil {
			log.Fatalf("command failed: %v", err)
		}

		os.Chdir(initialDir)
	}
}

func prepareCommand(args dto.RunCommandsArgs) *exec.Cmd {
	cmd := exec.Command(args.UserCommand, args.UserArgs...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd
}
