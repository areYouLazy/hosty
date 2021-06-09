package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/areYouLazy/libhosty"
	"github.com/spf13/cobra"
)

func PrintOutput(action string, details bool, hfl []*libhosty.HostsFileLine) {
	// if quiet return
	if quietOutput {
		return
	}

	// if json and details print detailed json
	if jsonOutput && details {
		fmt.Printf("%s", DetailedJSONOutput(action, hfl))
		return
	}

	// if json and show print json
	if jsonOutput && action == "show" {
		fmt.Printf("%s", JSONOutput("show", hfl))
		return
	}

	// if details print detailed
	if details {
		fmt.Printf("%s", DetailedOutput(action, hfl))
		return
	}

	// if json print json
	if jsonOutput {
		fmt.Printf("%s", JSONDoneOutput())
		return
	}

	if action == "show" {
		fmt.Printf("%s", ShowOutput(hfl))
		return
	}

	// print something for humans
	fmt.Println("done")
}

type DetailedJSONOutputStruct struct {
	Action          string   `json:"action"`
	LineNumber      int      `json:"number"`
	LineType        string   `json:"type"`
	LineAddress     string   `json:"address"`
	LineHostnames   []string `json:"hostnames"`
	LineComment     string   `json:"comment"`
	LineIsCommented bool     `json:"is_commented"`
	LineRaw         string   `json:"raw"`
}

type JSONOutputStruct struct {
	LineRaw string `json:"raw"`
}

type JSONDone struct {
	Done bool `json:"done"`
}

func DetailedJSONOutput(action string, hfl []*libhosty.HostsFileLine) []byte {
	var buffer [][]byte

	for _, v := range hfl {
		var stringType string

		switch v.Type {
		case 0:
			stringType = "unknown"
		case 10:
			stringType = "empty"
		case 20:
			stringType = "comment"
		case 30:
			stringType = "address"
		}

		o := DetailedJSONOutputStruct{
			Action:          action,
			LineNumber:      v.Number,
			LineType:        stringType,
			LineAddress:     v.Address.String(),
			LineHostnames:   v.Hostnames,
			LineComment:     v.Comment,
			LineIsCommented: v.IsCommented,
			LineRaw:         strings.ReplaceAll(v.Raw, "\t", ""),
		}

		b, err := json.Marshal(&o)
		if err != nil {
			cobra.CheckErr(err)
		}

		buffer = append(buffer, b)
	}

	return bytes.Join(buffer, []byte("\n"))
}

func DetailedOutput(action string, hfl []*libhosty.HostsFileLine) string {
	var buffer []string

	for _, v := range hfl {
		var stringType string
		var buf []string

		switch v.Type {
		case 0:
			stringType = "unknown"
		case 10:
			stringType = "empty"
		case 20:
			stringType = "comment"
		case 30:
			stringType = "address"
		}

		buf = append(
			buf,
			fmt.Sprintf("action: %s", action),
			fmt.Sprintf("number: %d", v.Number),
			fmt.Sprintf("type: %s", stringType),
			fmt.Sprintf("address: %s", v.Address),
			fmt.Sprintf("hostnames: %s", strings.Join(v.Hostnames, ", ")),
			fmt.Sprintf("comment: %s", v.Comment),
			fmt.Sprintf("is_commented: %t", v.IsCommented),
			fmt.Sprintf("raw: %s", strings.ReplaceAll(v.Raw, "\t", "")),
		)

		buffer = append(buffer, strings.Join(buf, "\n"))
	}

	return strings.Join(buffer, "\n\n")
}

func JSONOutput(action string, hfl []*libhosty.HostsFileLine) []byte {
	var buffer [][]byte

	for _, v := range hfl {
		o := JSONOutputStruct{
			LineRaw: strings.ReplaceAll(v.Raw, "\t", ""),
		}

		b, err := json.Marshal(&o)
		if err != nil {
			cobra.CheckErr(err)
		}

		buffer = append(buffer, b)
	}

	return bytes.Join(buffer, []byte("\n"))
}

func JSONDoneOutput() []byte {
	o := JSONDone{
		Done: true,
	}

	b, err := json.Marshal(&o)
	if err != nil {
		cobra.CheckErr(err)
	}

	return b
}

func ShowOutput(hfl []*libhosty.HostsFileLine) string {
	var buffer []string

	for idx := range hfl {
		buffer = append(buffer, strings.ReplaceAll(hfl[idx].Raw, "\t", ""))
	}

	return strings.Join(buffer, "\n")
}
