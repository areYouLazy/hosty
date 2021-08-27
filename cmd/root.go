package cmd

import (
	"github.com/areYouLazy/libhosty"
	"github.com/spf13/cobra"
)

var (
	// expose HostsFile and HostsFileConfig
	hc *libhosty.HostsFileConfig
	hf *libhosty.HostsFile

	// Used for flags
	customFile  string
	jsonOutput  bool
	quietOutput bool
	backupFile  bool

	rootCmd = &cobra.Command{
		Use:   "hosty",
		Short: "A command-line interface to the /etc/hosts file",
		Long: `Hosty is a command-line tool to interact with the /etc/hosts file.
It allows for fast inspect and edit of the file. Main goals of this tool are to be fast, reliable and scriptable.
Hosty uses libhosty to manipulate the file. You can find more about libhosty at https://github.com/areYouLazy/libhosty`,
		Version: "1.1.1",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			// init libhosty here because only here we're sure to have
			// all flags parsed by cobra
			var err error

			if customFile != "" {
				hc, err = libhosty.NewHostsFileConfig(customFile)
				cobra.CheckErr(err)

				hf, err = libhosty.InitWithConfig(hc)
				cobra.CheckErr(err)
			} else {

				hf, err = libhosty.Init()
				cobra.CheckErr(err)
			}
		},
	}
)

// Execute executes the root command
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	// parse flags
	rootCmd.PersistentFlags().BoolVarP(&jsonOutput, "json", "j", false, "print output in json format for easy parsing")
	rootCmd.PersistentFlags().BoolVarP(&quietOutput, "quiet", "q", false, "suppress every output except for errors")
	rootCmd.PersistentFlags().StringVarP(&customFile, "file", "f", "", "parse a custom /etc/hosts-like file instead of the system default one")
	rootCmd.PersistentFlags().BoolVarP(&backupFile, "backup", "b", false, "backup file before editing. Backup is hidden and named with the format: .YYYYMMDDHHmmss-hosts.bck")

	// hosty show
	rootCmd.AddCommand(showCmd)
	showCmd.Flags().BoolP("details", "d", false, "show lines details")

	// hosty delete
	rootCmd.AddCommand(deleteCmd)
	deleteCmd.Flags().BoolP("details", "d", false, "show operation details")

	// hosty add
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().BoolP("details", "d", false, "show operation details")

	// hosty comment
	rootCmd.AddCommand(commentCmd)
	commentCmd.Flags().BoolP("details", "d", false, "show operation details")

	// hosty uncomment
	rootCmd.AddCommand(uncommentCmd)
	uncommentCmd.Flags().BoolP("details", "d", false, "show operation details")

	// hosty export
	rootCmd.AddCommand(exportCmd)

	// hosty restore
	rootCmd.AddCommand(restoreCmd)
	restoreCmd.Flags().BoolP("details", "d", false, "show operation details")
}
