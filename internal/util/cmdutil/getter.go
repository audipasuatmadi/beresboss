package cmdutil

import (
	"os"
	"strings"
)

func GetUserCommandsAndDirectories(args []string) (usrCommand string, usrCommandArgs []string, directories []string) {
	allCommands := strings.Split(os.Args[1], " ")

	usrCommand = allCommands[0]
	usrCommandArgs = allCommands[1:]
	directories = os.Args[2:]
	return
}
