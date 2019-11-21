// Code generated by goa v3.0.7, DO NOT EDIT.
//
// enduro HTTP client CLI support package
//
// Command:
// $ goa gen github.com/artefactual-labs/enduro/internal/api/design -o
// internal/api

package cli

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	collectionc "github.com/artefactual-labs/enduro/internal/api/gen/http/collection/client"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// UsageCommands returns the set of commands and sub-commands using the format
//
//    command (subcommand1|subcommand2|...)
//
func UsageCommands() string {
	return `collection (list|show|delete|cancel|retry|workflow)
`
}

// UsageExamples produces an example of a valid invocation of the CLI tool.
func UsageExamples() string {
	return os.Args[0] + ` collection list --original-id "Est recusandae dolorum porro nam." --transfer-id "Voluptatum qui nam et molestiae." --aip-id "Maiores minus debitis et accusantium." --pipeline-id "Eveniet dolores omnis omnis." --query "Repellendus voluptatem optio et asperiores aut." --cursor "Alias ex vel nulla."` + "\n" +
		""
}

// ParseEndpoint returns the endpoint and payload as specified on the command
// line.
func ParseEndpoint(
	scheme, host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restore bool,
) (goa.Endpoint, interface{}, error) {
	var (
		collectionFlags = flag.NewFlagSet("collection", flag.ContinueOnError)

		collectionListFlags          = flag.NewFlagSet("list", flag.ExitOnError)
		collectionListOriginalIDFlag = collectionListFlags.String("original-id", "", "")
		collectionListTransferIDFlag = collectionListFlags.String("transfer-id", "", "")
		collectionListAipIDFlag      = collectionListFlags.String("aip-id", "", "")
		collectionListPipelineIDFlag = collectionListFlags.String("pipeline-id", "", "")
		collectionListQueryFlag      = collectionListFlags.String("query", "", "")
		collectionListCursorFlag     = collectionListFlags.String("cursor", "", "")

		collectionShowFlags  = flag.NewFlagSet("show", flag.ExitOnError)
		collectionShowIDFlag = collectionShowFlags.String("id", "REQUIRED", "Identifier of collection to show")

		collectionDeleteFlags  = flag.NewFlagSet("delete", flag.ExitOnError)
		collectionDeleteIDFlag = collectionDeleteFlags.String("id", "REQUIRED", "Identifier of collection to delete")

		collectionCancelFlags  = flag.NewFlagSet("cancel", flag.ExitOnError)
		collectionCancelIDFlag = collectionCancelFlags.String("id", "REQUIRED", "Identifier of collection to remove")

		collectionRetryFlags  = flag.NewFlagSet("retry", flag.ExitOnError)
		collectionRetryIDFlag = collectionRetryFlags.String("id", "REQUIRED", "Identifier of collection to retry")

		collectionWorkflowFlags  = flag.NewFlagSet("workflow", flag.ExitOnError)
		collectionWorkflowIDFlag = collectionWorkflowFlags.String("id", "REQUIRED", "Identifier of collection to look up")
	)
	collectionFlags.Usage = collectionUsage
	collectionListFlags.Usage = collectionListUsage
	collectionShowFlags.Usage = collectionShowUsage
	collectionDeleteFlags.Usage = collectionDeleteUsage
	collectionCancelFlags.Usage = collectionCancelUsage
	collectionRetryFlags.Usage = collectionRetryUsage
	collectionWorkflowFlags.Usage = collectionWorkflowUsage

	if err := flag.CommandLine.Parse(os.Args[1:]); err != nil {
		return nil, nil, err
	}

	if flag.NArg() < 2 { // two non flag args are required: SERVICE and ENDPOINT (aka COMMAND)
		return nil, nil, fmt.Errorf("not enough arguments")
	}

	var (
		svcn string
		svcf *flag.FlagSet
	)
	{
		svcn = flag.Arg(0)
		switch svcn {
		case "collection":
			svcf = collectionFlags
		default:
			return nil, nil, fmt.Errorf("unknown service %q", svcn)
		}
	}
	if err := svcf.Parse(flag.Args()[1:]); err != nil {
		return nil, nil, err
	}

	var (
		epn string
		epf *flag.FlagSet
	)
	{
		epn = svcf.Arg(0)
		switch svcn {
		case "collection":
			switch epn {
			case "list":
				epf = collectionListFlags

			case "show":
				epf = collectionShowFlags

			case "delete":
				epf = collectionDeleteFlags

			case "cancel":
				epf = collectionCancelFlags

			case "retry":
				epf = collectionRetryFlags

			case "workflow":
				epf = collectionWorkflowFlags

			}

		}
	}
	if epf == nil {
		return nil, nil, fmt.Errorf("unknown %q endpoint %q", svcn, epn)
	}

	// Parse endpoint flags if any
	if svcf.NArg() > 1 {
		if err := epf.Parse(svcf.Args()[1:]); err != nil {
			return nil, nil, err
		}
	}

	var (
		data     interface{}
		endpoint goa.Endpoint
		err      error
	)
	{
		switch svcn {
		case "collection":
			c := collectionc.NewClient(scheme, host, doer, enc, dec, restore)
			switch epn {
			case "list":
				endpoint = c.List()
				data, err = collectionc.BuildListPayload(*collectionListOriginalIDFlag, *collectionListTransferIDFlag, *collectionListAipIDFlag, *collectionListPipelineIDFlag, *collectionListQueryFlag, *collectionListCursorFlag)
			case "show":
				endpoint = c.Show()
				data, err = collectionc.BuildShowPayload(*collectionShowIDFlag)
			case "delete":
				endpoint = c.Delete()
				data, err = collectionc.BuildDeletePayload(*collectionDeleteIDFlag)
			case "cancel":
				endpoint = c.Cancel()
				data, err = collectionc.BuildCancelPayload(*collectionCancelIDFlag)
			case "retry":
				endpoint = c.Retry()
				data, err = collectionc.BuildRetryPayload(*collectionRetryIDFlag)
			case "workflow":
				endpoint = c.Workflow()
				data, err = collectionc.BuildWorkflowPayload(*collectionWorkflowIDFlag)
			}
		}
	}
	if err != nil {
		return nil, nil, err
	}

	return endpoint, data, nil
}

// collectionUsage displays the usage of the collection command and its
// subcommands.
func collectionUsage() {
	fmt.Fprintf(os.Stderr, `The collection service manages packages being transferred to Archivematica.
Usage:
    %s [globalflags] collection COMMAND [flags]

COMMAND:
    list: List all stored collections
    show: Show collection by ID
    delete: Delete collection by ID
    cancel: Cancel collection processing by ID
    retry: Retry collection processing by ID
    workflow: Retrieve workflow status by ID

Additional help:
    %s collection COMMAND --help
`, os.Args[0], os.Args[0])
}
func collectionListUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] collection list -original-id STRING -transfer-id STRING -aip-id STRING -pipeline-id STRING -query STRING -cursor STRING

List all stored collections
    -original-id STRING: 
    -transfer-id STRING: 
    -aip-id STRING: 
    -pipeline-id STRING: 
    -query STRING: 
    -cursor STRING: 

Example:
    `+os.Args[0]+` collection list --original-id "Est recusandae dolorum porro nam." --transfer-id "Voluptatum qui nam et molestiae." --aip-id "Maiores minus debitis et accusantium." --pipeline-id "Eveniet dolores omnis omnis." --query "Repellendus voluptatem optio et asperiores aut." --cursor "Alias ex vel nulla."
`, os.Args[0])
}

func collectionShowUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] collection show -id UINT

Show collection by ID
    -id UINT: Identifier of collection to show

Example:
    `+os.Args[0]+` collection show --id 7262811908532712622
`, os.Args[0])
}

func collectionDeleteUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] collection delete -id UINT

Delete collection by ID
    -id UINT: Identifier of collection to delete

Example:
    `+os.Args[0]+` collection delete --id 7686560774918400757
`, os.Args[0])
}

func collectionCancelUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] collection cancel -id UINT

Cancel collection processing by ID
    -id UINT: Identifier of collection to remove

Example:
    `+os.Args[0]+` collection cancel --id 11533904348193230542
`, os.Args[0])
}

func collectionRetryUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] collection retry -id UINT

Retry collection processing by ID
    -id UINT: Identifier of collection to retry

Example:
    `+os.Args[0]+` collection retry --id 2998639704067115621
`, os.Args[0])
}

func collectionWorkflowUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] collection workflow -id UINT

Retrieve workflow status by ID
    -id UINT: Identifier of collection to look up

Example:
    `+os.Args[0]+` collection workflow --id 17765174268271371906
`, os.Args[0])
}
