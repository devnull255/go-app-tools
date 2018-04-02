package cli

import (
	"errors"
	"fmt"
)

var version = "0.1"

const (
	Help      = "help"
	Apps      = "apps"
	Projects  = "projects"
	Files     = "files"
	Packages  = "packages"
	Dirs      = "dirs"
	Configure = "configure"
)

type CommandHelp struct {
	Name string
	ShortText string
	LongText string
}

var CommandHelpMap map[string]CommandHelp

func init() {
	CommandHelpMap = map[string]CommandHelp{}
	CommandHelpMap[Help] = CommandHelp{Name: Help, ShortText: "Lists help", LongText: ""}
	CommandHelpMap[Apps] = CommandHelp{Name: Apps, ShortText: "Manages apps", LongText: ""}
	CommandHelpMap[Projects] = CommandHelp{Name: Projects, ShortText: "Manages projects", LongText: ""}
	CommandHelpMap[Files] = CommandHelp{Name: Files, ShortText: "Manages files"}
	CommandHelpMap[Dirs] = CommandHelp{Name: Dirs, ShortText: "Manages directories"}
	CommandHelpMap[Packages] = CommandHelp{Name: Packages, ShortText: "Manages packages"}
	CommandHelpMap[Configure] = CommandHelp{Name: Configure, ShortText: "Used to configure gat tool"}
}

// ShowVersion displays the version of the gat CLI.
func ShowVersion() {
	fmt.Println("gat Version:", version)
}

// Description returns a string describing the gat tool's basic use.
func Description() string {
	return `Description: The gat command-line tool provides the quick generation of Go Application project shells utilizing built-in as well as user-provided templates.`
}

// Usage returns a string on basic usage of the gat command.
func Usage() string {
	return `Usage: gat <command> [cmd args]
Valid commands are:
`
}

// ShowHelp displays general and sub-command level help.
func CmdHelp(args []string) error {
	fmt.Println(Description())
	fmt.Println(Usage())
	for _,ch := range CommandHelpMap {
		fmt.Printf("%-20s: %s\n",ch.Name, ch.ShortText)
	}
	return nil
}

// CmdProjects is the subcommand entry point for project management tasks.
// Among the supported operations are:
//     list - lists current projects created and managed by gat
//     create - creates a gat-managed project
//     update - updates a gat-managed project
//     delete - deletes a gat-managed project
//     configure - configures project parameters
func CmdProjects(args []string) error {
	fmt.Println("In CmdProjects: TODO")
	return nil
}

// CmdFiles provides file-level CRUD commands for a given project.
func CmdFiles(args []string) error {
	fmt.Println("In CmdFiles: TODO")
	return nil
}

// CmdApps provides the subcommand entry point for application management tasks.
func CmdApps(args []string) error {
	fmt.Println("In CmdApps: TODO")
	return nil
}

// CmdTemplates is the subcommand entry point for project, application and file templates.
func CmdTemplates(args []string) error {
	fmt.Println("In CmdTemplates: TODO")
	return nil
}

// CmdDirs is the subcommand entry point for directory management operations.
func CmdDirs(args []string) error {
	fmt.Println("In CmdDirs: TODO")
	return nil
}

// CmdPackages is the subcommand entry point for package management operations.
func CmdPackages(args []string) error {
	fmt.Println("In CmdPackages: TODO")
	return nil
}

// CmdConfigure is the subcommand entry point for gat configuration items, like
// gat project root location, server application launching, debug parameters,
// logging, etc.
func CmdConfigure(args []string) error {
	fmt.Println("In CmdConfigure: TODO")
	return nil
}

func getCommands() map[string]func([]string) error {
	var cmdMap = map[string]func([]string)error{}
	cmdMap[Projects] = CmdProjects
	cmdMap[Apps] = CmdApps
	cmdMap[Files] = CmdFiles
	cmdMap[Dirs] = CmdDirs
	cmdMap[Configure] = CmdConfigure
	cmdMap[Help] = CmdHelp
	return cmdMap
}

// RunCommand executes the command named by cmdName.
func RunCommand(cmdName string, args []string) error {
	cmds := getCommands()
	f, found := cmds[cmdName]
	if !found {
		err := errors.New(cmdName + " is not a valid command.")
		return err
	}
	return f(args)
}
