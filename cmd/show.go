package cmd

import (
	"fmt"
	"net"

	"github.com/areYouLazy/libhosty"
	"github.com/spf13/cobra"
)

var (
	showCmd = &cobra.Command{
		Use:     "show [ip/fqdn]",
		Aliases: []string{"s", "sh", "sho"},
		Short:   "Show file data",
		Long: `Query hosts file for data.
Query can be done by IP or FQDN`,
		Args:      cobra.ExactArgs(1),
		ValidArgs: []string{"ip/fqdn"},
		Run: func(cmd *cobra.Command, args []string) {
			var hfl []*libhosty.HostsFileLine

			// get flags
			details, _ := cmd.Flags().GetBool("details")

			// get args
			arg := string(args[0])

			// try to parse arg as ip
			ip := net.ParseIP(arg)
			switch ip {
			case nil:
				// search by FQDN
				hfl = hf.GetHostsFileLinesByHostnameAsRegexp(arg)
			default:
				// arg is a valid IP address
				hfl = hf.GetHostsFileLinesByIP(ip)
				if len(hfl) <= 0 {
					fmt.Printf("nothing found for ip %s\n", ip)
					return
				}
			}

			PrintOutput("show", details, hfl)
		},
	}
)
