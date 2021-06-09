package cmd

import (
	"fmt"
	"net"

	"github.com/areYouLazy/libhosty"
	"github.com/spf13/cobra"
)

var (
	deleteCmd = &cobra.Command{
		Use:     "delete [ip/fqdn]",
		Aliases: []string{"d", "de", "del", "dele", "delet", "rm", "rem"},
		Short:   "Delete file data",
		Long: `Delete hosts data on file.
Query can be done by IP or FQDN`,
		Args:      cobra.ExactArgs(1),
		ValidArgs: []string{"ip/fqdn"},
		Run: func(cmd *cobra.Command, args []string) {
			var hfl []*libhosty.HostsFileLine
			arg := string(args[0])

			// get flags
			details, _ := cmd.Flags().GetBool("details")
			fqdn, _ := cmd.Flags().GetString("fqdn")

			// try to parse arg as ip
			ip := net.ParseIP(arg)
			switch ip {
			case nil:
				// delete needs to get data before acting
				// otherwise we'll not have data for detailed output
				hfl = hf.GetHostsFileLinesByHostnameAsRegexp(arg)
				if len(hfl) <= 0 {
					fmt.Printf("nothing found for fqdn %s\n", arg)
					return
				}

				hf.RemoveHostsFileLinesByHostnameAsRegexp(fqdn)
			default:
				// delete needs to get data before acting
				// otherwise we'll not have data for detailed output
				hfl = hf.GetHostsFileLinesByIP(ip)
				if len(hfl) <= 0 {
					fmt.Printf("nothing found for ip %s\n", arg)
					return
				}

				hf.RemoveHostsFileLinesByIP(ip)
			}

			if backupFile {
				err := Backup()
				cobra.CheckErr(err)
			}
			err := hf.SaveHostsFile()
			cobra.CheckErr(err)

			PrintOutput("delete", details, hfl)
		},
	}
)
