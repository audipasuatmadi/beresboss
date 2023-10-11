package main

import (
	"log"
	"os"

	"github.com/audipasuatmadi/beresbos/internal/dto"
	"github.com/audipasuatmadi/beresbos/internal/usecase"
	"github.com/audipasuatmadi/beresbos/internal/util/cmdutil"
)

func main() {
	// contributing? add your names here!
	log.Println("[main]: starting tool, created by Putu Audi Pasuatmadi")

	if len(os.Args) < 2 {
		log.Fatalf("args are [command] [...paths], e.g: beresbos \"mvn clean install -T2C --settings \\\"some-settings\\\"\" ./project1 ./module/project-2")
	}

	userCommand, userArgs, directories := cmdutil.GetUserCommandsAndDirectories(os.Args)
	usecase.RunCommands(dto.RunCommandsArgs{
		UserCommand: userCommand,
		UserArgs:    userArgs,
		Directories: directories,
	})

	log.Println("[main]: stopping application")
}
