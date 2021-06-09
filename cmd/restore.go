package cmd

import (
	"runtime"

	"github.com/areYouLazy/libhosty"
	"github.com/spf13/cobra"
)

var (
	restoreCmd = &cobra.Command{
		Use:     "restore [os]",
		Aliases: []string{"r", "re", "res", "rest", "resto", "restor"},
		Short:   "Restore default hosts file",
		Long: `Restore the default hosts file.
If [os] is mitted, hosty will try to guess your OS and restore the appropriate file`,
		ValidArgs: []string{"os"},
		Run: func(cmd *cobra.Command, args []string) {
			os := ""
			if len(args) > 0 {
				os = string(args[0])
			} else {
				os = runtime.GOOS
			}

			// get flags
			details, _ := cmd.Flags().GetBool("details")

			switch os {
			case "windows":
				hf.RestoreDefaultWindowsHostsFile()
			case "win":
				hf.RestoreDefaultWindowsHostsFile()
			case "darwin":
				hf.RestoreDefaultDarwinHostsFile()
			default:
				hf.RestoreDefaultLinuxHostsFile()
			}

			if backupFile {
				err := Backup()
				cobra.CheckErr(err)
			}
			err := hf.SaveHostsFile()
			cobra.CheckErr(err)

			hfl := make([]*libhosty.HostsFileLine, len(hf.HostsFileLines))

			for idx := range hf.HostsFileLines {
				hfl = append(hfl, &hf.HostsFileLines[idx])
			}

			PrintOutput("restore", details, hfl)
		},
	}
)
