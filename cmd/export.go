package cmd

import (
	"errors"

	"github.com/spf13/cobra"
)

var (
	exportCmd = &cobra.Command{
		Use:     "export [location]",
		Aliases: []string{"e", "ex", "exp", "expo", "expor"},
		Short:   "Export file to a custom location",
		Long:    `Export hosts file to a custom location`,
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 || len(args) > 2 {
				return errors.New("invalid number of arguments")
			}

			return nil
		},
		ValidArgs: []string{"location"},
		Run: func(cmd *cobra.Command, args []string) {
			arg := string(args[0])

			if backupFile {
				err := Backup()
				cobra.CheckErr(err)
			}
			err := hf.SaveHostsFileAs(arg)
			cobra.CheckErr(err)
		},
	}
)
