package cmd

import (
	"errors"
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
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) > 1 {
				return errors.New("invalid number of arguments")
			}

			return nil
		},
		ValidArgs: []string{"ip/fqdn"},
		Run: func(cmd *cobra.Command, args []string) {
			var hfl []*libhosty.HostsFileLine
			var arg string

			// get flags
			details, _ := cmd.Flags().GetBool("details")

			//check if there's an argument
			if len(args) > 0 {
				// get arg
				arg = string(args[0])

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
			} else {
				//no argument, so just return the whole file
				// hfl = hf.GetHostsFileLines()
			}

			PrintOutput("show", details, hfl)
		},
	}
)
