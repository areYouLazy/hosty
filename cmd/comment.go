package cmd

import (
	"fmt"
	"net"

	"github.com/areYouLazy/libhosty"
	"github.com/spf13/cobra"
)

var (
	commentCmd = &cobra.Command{
		Use:     "comment [ip/fqdn]",
		Aliases: []string{"c", "co", "com", "comm", "comme", "commen"},
		Short:   "Comment file data",
		Long: `Comment hosts data on file.
Query can be done by IP or FQDN`,
		Args:      cobra.ExactArgs(1),
		ValidArgs: []string{"ip/fqdn"},
		Run: func(cmd *cobra.Command, args []string) {
			var hfl []*libhosty.HostsFileLine
			arg := string(args[0])

			// get flags
			details, _ := cmd.Flags().GetBool("details")

			// try to parse arg as ip
			ip := net.ParseIP(arg)
			switch ip {
			case nil:
				hf.CommentHostsFileLinesByHostnameAsRegexp(arg)

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
				hf.CommentHostsFileLinesByIP(ip)

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

			PrintOutput("comment", details, hfl)
		},
	}
)
