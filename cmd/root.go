package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/briandowns/spinner"
	"github.com/kekda-py/create-nextcord-app/utils"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "create-nextcord-app",
	Short: "Set up a nextcord app by running a single command.",

	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println(utils.Colorize("blue", "Creating a new nextcord app...\n"))

		var (
			Name  string
			Dir   string
			Token string

			err error
		)

		utils.Ask(utils.Question{
			Message: "What is the bot's name?",
			Validate: func(value string) error {
				return nil
			},
			Default: "the greatest bot",
		}, &Name)

		utils.Ask(utils.Question{
			Message: "What should be the directory's name?",
			Validate: func(value string) error {
				if value == "." || value == "./" {
					return nil
				}
				if _, err := os.Stat("./" + value); err != nil {
					return nil
				}
				return fmt.Errorf("directory `%s` already exists", value)
			},
			Default: "./",
		}, &Dir)

		utils.Ask(utils.Question{
			Message: "Bots token?",
			Validate: func(value string) error {
				return nil
			},
			Default: "",
		}, &Token)

		ans := struct {
			Name  string
			Dir   string
			Token string
		}{}

		ans.Name = Name
		ans.Dir = Dir
		ans.Token = Token

		fmt.Println()

		s := spinner.New(spinner.CharSets[34], 100*time.Millisecond)

		s.Suffix = utils.Colorize("blue", " Creating the directory...")
		s.Start()

		if err = utils.CreateDir(ans.Dir); err != nil {
			fmt.Println(utils.Colorize("red", "Error: "+err.Error()))
			os.Exit(1)
		}

		time.Sleep(1 * time.Second)

		s.Stop()

		s = spinner.New(spinner.CharSets[34], 100*time.Millisecond)

		s.Suffix = utils.Colorize("blue", " Creating the files...")
		s.Start()

		token := ans.Token
		if token == "" {
			token = "TOKEN"
		}

		if err = utils.CreateFiles(ans.Dir, token); err != nil {
			fmt.Println(utils.Colorize("red", "Error: "+err.Error()))
			os.Exit(1)
		}

		time.Sleep(1 * time.Second)
		s.Stop()

		s = spinner.New(spinner.CharSets[34], 100*time.Millisecond)

		s.Suffix = utils.Colorize("blue", " Initializing a github repository...")

		s.Start()

		if err = utils.InitializeGit(ans.Dir); err != nil {
			fmt.Println(utils.Colorize("red", "Error: "+err.Error()))
			os.Exit(1)
		}

		time.Sleep(1 * time.Second)
		s.Stop()

		s = spinner.New(spinner.CharSets[34], 100*time.Millisecond)

		s.Suffix = utils.Colorize("blue", " Creating a virtual environment and installing packages...")

		s.Start()

		if err = utils.InitializeVenv(ans.Dir); err != nil {
			fmt.Println(utils.Colorize("red", "Error: "+err.Error()))
			os.Exit(1)
		}

		time.Sleep(1 * time.Second)
		s.Stop()

		fmt.Println(utils.Colorize("green", "Successfully created a nextcord app!\n"))
		fmt.Println("To run the app do -\n ")
		fmt.Println(utils.Colorize("blue", "\tcd "+ans.Dir+"\n"+"\tenv/Scripts/activate\n"+"\tpython main.py\n"))

	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
