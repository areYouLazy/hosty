package cmd

import (
	"fmt"
	"net"

	"github.com/areYouLazy/libhosty"
	"github.com/spf13/cobra"
)

var (
	uncommentCmd = &cobra.Command{
		Use:     "uncomment [ip/fqdn]",
		Aliases: []string{"u", "un", "unc", "unco", "uncom", "uncomm", "uncomme", "uncommen"},
		Short:   "Uncomment file data",
		Long: `Uncomment hosts data on file.
Query can be done by IP or FQDN`,
		Args:      cobra.ExactArgs(1),
		ValidArgs: []string{"ip/fqdn"},
		Run: func(cmd *cobra.Command, args []string) {
			var hfl []*libhosty.HostsFileLine
			arg := args[0]

			// get flags
			details, _ := cmd.Flags().GetBool("details")

			// try to parse arg as ip
			ip := net.ParseIP(arg)
			switch ip {
			case nil:
				hf.UncommentHostsFileLinesByHostnameAsRegexp(arg)

				if backupFile {
					err := Backup()
					cobra.CheckErr(err)
				}
				err := hf.SaveHostsFile()
				cobra.CheckErr(err)

				hfl = hf.GetHostsFileLinesByHostnameAsRegexp(arg)
				if len(hfl) <= 0 {
					fmt.Printf("nothing found for fqdn %s\n", arg)
					return
				}
			default:
				hf.UncommentHostsFileLinesByIP(ip)

				if backupFile {
					err := Backup()
					cobra.CheckErr(err)
				}
				err := hf.SaveHostsFile()
				cobra.CheckErr(err)

				hfl = hf.GetHostsFileLinesByIP(ip)
				if len(hfl) <= 0 {
					fmt.Printf("nothing found for ip %s\n", arg)
					return
				}
			}

			PrintOutput("uncomment", details, hfl)
		},
	}
)
