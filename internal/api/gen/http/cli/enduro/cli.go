// Code generated by goa v3.5.5, DO NOT EDIT.
//
// enduro HTTP client CLI support package
//
// Command:
// $ goa gen github.com/penwern/enduro/internal/api/design -o internal/api

package cli

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	batchc "github.com/penwern/enduro/internal/api/gen/http/batch/client"
	collectionc "github.com/penwern/enduro/internal/api/gen/http/collection/client"
	pipelinec "github.com/penwern/enduro/internal/api/gen/http/pipeline/client"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// UsageCommands returns the set of commands and sub-commands using the format
//
//	command (subcommand1|subcommand2|...)
func UsageCommands() string {
	return `pipeline (list|show|processing)
batch (submit|status|hints)
collection (monitor|list|show|delete|cancel|retry|workflow|download|decide|bulk|bulk-status)
`
}

// UsageExamples produces an example of a valid invocation of the CLI tool.
func UsageExamples() string {
	return os.Args[0] + ` pipeline list --name "Iure nulla."` + "\n" +
		os.Args[0] + ` batch submit --body '{
      "completed_dir": "Laboriosam odit.",
      "path": "Sit nihil.",
      "pipeline": "Necessitatibus vel aut deleniti quia qui.",
      "processing_config": "Vel voluptatem.",
      "retention_period": "Sed perferendis illum illum omnis et officiis."
   }'` + "\n" +
		os.Args[0] + ` collection monitor` + "\n" +
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
	dialer goahttp.Dialer,
	collectionConfigurer *collectionc.ConnConfigurer,
) (goa.Endpoint, interface{}, error) {
	var (
		pipelineFlags = flag.NewFlagSet("pipeline", flag.ContinueOnError)

		pipelineListFlags    = flag.NewFlagSet("list", flag.ExitOnError)
		pipelineListNameFlag = pipelineListFlags.String("name", "", "")

		pipelineShowFlags  = flag.NewFlagSet("show", flag.ExitOnError)
		pipelineShowIDFlag = pipelineShowFlags.String("id", "REQUIRED", "Identifier of pipeline to show")

		pipelineProcessingFlags  = flag.NewFlagSet("processing", flag.ExitOnError)
		pipelineProcessingIDFlag = pipelineProcessingFlags.String("id", "REQUIRED", "Identifier of pipeline")

		batchFlags = flag.NewFlagSet("batch", flag.ContinueOnError)

		batchSubmitFlags    = flag.NewFlagSet("submit", flag.ExitOnError)
		batchSubmitBodyFlag = batchSubmitFlags.String("body", "REQUIRED", "")

		batchStatusFlags = flag.NewFlagSet("status", flag.ExitOnError)

		batchHintsFlags = flag.NewFlagSet("hints", flag.ExitOnError)

		collectionFlags = flag.NewFlagSet("collection", flag.ContinueOnError)

		collectionMonitorFlags = flag.NewFlagSet("monitor", flag.ExitOnError)

		collectionListFlags                   = flag.NewFlagSet("list", flag.ExitOnError)
		collectionListNameFlag                = collectionListFlags.String("name", "", "")
		collectionListOriginalIDFlag          = collectionListFlags.String("original-id", "", "")
		collectionListTransferIDFlag          = collectionListFlags.String("transfer-id", "", "")
		collectionListAipIDFlag               = collectionListFlags.String("aip-id", "", "")
		collectionListPipelineIDFlag          = collectionListFlags.String("pipeline-id", "", "")
		collectionListEarliestCreatedTimeFlag = collectionListFlags.String("earliest-created-time", "", "")
		collectionListLatestCreatedTimeFlag   = collectionListFlags.String("latest-created-time", "", "")
		collectionListStatusFlag              = collectionListFlags.String("status", "", "")
		collectionListCursorFlag              = collectionListFlags.String("cursor", "", "")

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

		collectionDownloadFlags  = flag.NewFlagSet("download", flag.ExitOnError)
		collectionDownloadIDFlag = collectionDownloadFlags.String("id", "REQUIRED", "Identifier of collection to look up")

		collectionDecideFlags    = flag.NewFlagSet("decide", flag.ExitOnError)
		collectionDecideBodyFlag = collectionDecideFlags.String("body", "REQUIRED", "")
		collectionDecideIDFlag   = collectionDecideFlags.String("id", "REQUIRED", "Identifier of collection to look up")

		collectionBulkFlags    = flag.NewFlagSet("bulk", flag.ExitOnError)
		collectionBulkBodyFlag = collectionBulkFlags.String("body", "REQUIRED", "")

		collectionBulkStatusFlags = flag.NewFlagSet("bulk-status", flag.ExitOnError)
	)
	pipelineFlags.Usage = pipelineUsage
	pipelineListFlags.Usage = pipelineListUsage
	pipelineShowFlags.Usage = pipelineShowUsage
	pipelineProcessingFlags.Usage = pipelineProcessingUsage

	batchFlags.Usage = batchUsage
	batchSubmitFlags.Usage = batchSubmitUsage
	batchStatusFlags.Usage = batchStatusUsage
	batchHintsFlags.Usage = batchHintsUsage

	collectionFlags.Usage = collectionUsage
	collectionMonitorFlags.Usage = collectionMonitorUsage
	collectionListFlags.Usage = collectionListUsage
	collectionShowFlags.Usage = collectionShowUsage
	collectionDeleteFlags.Usage = collectionDeleteUsage
	collectionCancelFlags.Usage = collectionCancelUsage
	collectionRetryFlags.Usage = collectionRetryUsage
	collectionWorkflowFlags.Usage = collectionWorkflowUsage
	collectionDownloadFlags.Usage = collectionDownloadUsage
	collectionDecideFlags.Usage = collectionDecideUsage
	collectionBulkFlags.Usage = collectionBulkUsage
	collectionBulkStatusFlags.Usage = collectionBulkStatusUsage

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
		case "pipeline":
			svcf = pipelineFlags
		case "batch":
			svcf = batchFlags
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
		case "pipeline":
			switch epn {
			case "list":
				epf = pipelineListFlags

			case "show":
				epf = pipelineShowFlags

			case "processing":
				epf = pipelineProcessingFlags

			}

		case "batch":
			switch epn {
			case "submit":
				epf = batchSubmitFlags

			case "status":
				epf = batchStatusFlags

			case "hints":
				epf = batchHintsFlags

			}

		case "collection":
			switch epn {
			case "monitor":
				epf = collectionMonitorFlags

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

			case "download":
				epf = collectionDownloadFlags

			case "decide":
				epf = collectionDecideFlags

			case "bulk":
				epf = collectionBulkFlags

			case "bulk-status":
				epf = collectionBulkStatusFlags

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
		case "pipeline":
			c := pipelinec.NewClient(scheme, host, doer, enc, dec, restore)
			switch epn {
			case "list":
				endpoint = c.List()
				data, err = pipelinec.BuildListPayload(*pipelineListNameFlag)
			case "show":
				endpoint = c.Show()
				data, err = pipelinec.BuildShowPayload(*pipelineShowIDFlag)
			case "processing":
				endpoint = c.Processing()
				data, err = pipelinec.BuildProcessingPayload(*pipelineProcessingIDFlag)
			}
		case "batch":
			c := batchc.NewClient(scheme, host, doer, enc, dec, restore)
			switch epn {
			case "submit":
				endpoint = c.Submit()
				data, err = batchc.BuildSubmitPayload(*batchSubmitBodyFlag)
			case "status":
				endpoint = c.Status()
				data = nil
			case "hints":
				endpoint = c.Hints()
				data = nil
			}
		case "collection":
			c := collectionc.NewClient(scheme, host, doer, enc, dec, restore, dialer, collectionConfigurer)
			switch epn {
			case "monitor":
				endpoint = c.Monitor()
				data = nil
			case "list":
				endpoint = c.List()
				data, err = collectionc.BuildListPayload(*collectionListNameFlag, *collectionListOriginalIDFlag, *collectionListTransferIDFlag, *collectionListAipIDFlag, *collectionListPipelineIDFlag, *collectionListEarliestCreatedTimeFlag, *collectionListLatestCreatedTimeFlag, *collectionListStatusFlag, *collectionListCursorFlag)
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
			case "download":
				endpoint = c.Download()
				data, err = collectionc.BuildDownloadPayload(*collectionDownloadIDFlag)
			case "decide":
				endpoint = c.Decide()
				data, err = collectionc.BuildDecidePayload(*collectionDecideBodyFlag, *collectionDecideIDFlag)
			case "bulk":
				endpoint = c.Bulk()
				data, err = collectionc.BuildBulkPayload(*collectionBulkBodyFlag)
			case "bulk-status":
				endpoint = c.BulkStatus()
				data = nil
			}
		}
	}
	if err != nil {
		return nil, nil, err
	}

	return endpoint, data, nil
}

// pipelineUsage displays the usage of the pipeline command and its subcommands.
func pipelineUsage() {
	fmt.Fprintf(os.Stderr, `The pipeline service manages Archivematica pipelines.
Usage:
    %[1]s [globalflags] pipeline COMMAND [flags]

COMMAND:
    list: List all known pipelines
    show: Show pipeline by ID
    processing: List all processing configurations of a pipeline given its ID

Additional help:
    %[1]s pipeline COMMAND --help
`, os.Args[0])
}
func pipelineListUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] pipeline list -name STRING

List all known pipelines
    -name STRING: 

Example:
    %[1]s pipeline list --name "Iure nulla."
`, os.Args[0])
}

func pipelineShowUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] pipeline show -id STRING

Show pipeline by ID
    -id STRING: Identifier of pipeline to show

Example:
    %[1]s pipeline show --id "6BD40BD6-7AF6-FB4E-C1C1-23700A0E68DE"
`, os.Args[0])
}

func pipelineProcessingUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] pipeline processing -id STRING

List all processing configurations of a pipeline given its ID
    -id STRING: Identifier of pipeline

Example:
    %[1]s pipeline processing --id "799C598D-49D1-F53A-0ED6-818482028C45"
`, os.Args[0])
}

// batchUsage displays the usage of the batch command and its subcommands.
func batchUsage() {
	fmt.Fprintf(os.Stderr, `The batch service manages batches of collections.
Usage:
    %[1]s [globalflags] batch COMMAND [flags]

COMMAND:
    submit: Submit a new batch
    status: Retrieve status of current batch operation.
    hints: Retrieve form hints

Additional help:
    %[1]s batch COMMAND --help
`, os.Args[0])
}
func batchSubmitUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] batch submit -body JSON

Submit a new batch
    -body JSON: 

Example:
    %[1]s batch submit --body '{
      "completed_dir": "Laboriosam odit.",
      "path": "Sit nihil.",
      "pipeline": "Necessitatibus vel aut deleniti quia qui.",
      "processing_config": "Vel voluptatem.",
      "retention_period": "Sed perferendis illum illum omnis et officiis."
   }'
`, os.Args[0])
}

func batchStatusUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] batch status

Retrieve status of current batch operation.

Example:
    %[1]s batch status
`, os.Args[0])
}

func batchHintsUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] batch hints

Retrieve form hints

Example:
    %[1]s batch hints
`, os.Args[0])
}

// collectionUsage displays the usage of the collection command and its
// subcommands.
func collectionUsage() {
	fmt.Fprintf(os.Stderr, `The collection service manages packages being transferred to Archivematica.
Usage:
    %[1]s [globalflags] collection COMMAND [flags]

COMMAND:
    monitor: Monitor implements monitor.
    list: List all stored collections
    show: Show collection by ID
    delete: Delete collection by ID
    cancel: Cancel collection processing by ID
    retry: Retry collection processing by ID
    workflow: Retrieve workflow status by ID
    download: Download collection by ID
    decide: Make decision for a pending collection by ID
    bulk: Bulk operations (retry, cancel...).
    bulk-status: Retrieve status of current bulk operation.

Additional help:
    %[1]s collection COMMAND --help
`, os.Args[0])
}
func collectionMonitorUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] collection monitor

Monitor implements monitor.

Example:
    %[1]s collection monitor
`, os.Args[0])
}

func collectionListUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] collection list -name STRING -original-id STRING -transfer-id STRING -aip-id STRING -pipeline-id STRING -earliest-created-time STRING -latest-created-time STRING -status STRING -cursor STRING

List all stored collections
    -name STRING: 
    -original-id STRING: 
    -transfer-id STRING: 
    -aip-id STRING: 
    -pipeline-id STRING: 
    -earliest-created-time STRING: 
    -latest-created-time STRING: 
    -status STRING: 
    -cursor STRING: 

Example:
    %[1]s collection list --name "Laudantium eos fugiat iure sit ea." --original-id "Et dolor ullam consequatur dignissimos." --transfer-id "1576BB1F-D21C-05AD-6677-73725A387FA6" --aip-id "7F641448-35F7-7B05-AFB4-8DFCC48CED66" --pipeline-id "D358BA94-28F1-C0F1-EA4E-E3BFBD8A6AE2" --earliest-created-time "1996-05-01T14:23:24Z" --latest-created-time "1987-02-15T23:56:43Z" --status "error" --cursor "Sit et inventore et."
`, os.Args[0])
}

func collectionShowUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] collection show -id UINT

Show collection by ID
    -id UINT: Identifier of collection to show

Example:
    %[1]s collection show --id 6649367811978689086
`, os.Args[0])
}

func collectionDeleteUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] collection delete -id UINT

Delete collection by ID
    -id UINT: Identifier of collection to delete

Example:
    %[1]s collection delete --id 14411764229641892412
`, os.Args[0])
}

func collectionCancelUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] collection cancel -id UINT

Cancel collection processing by ID
    -id UINT: Identifier of collection to remove

Example:
    %[1]s collection cancel --id 12015603943555843617
`, os.Args[0])
}

func collectionRetryUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] collection retry -id UINT

Retry collection processing by ID
    -id UINT: Identifier of collection to retry

Example:
    %[1]s collection retry --id 3222974273559150037
`, os.Args[0])
}

func collectionWorkflowUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] collection workflow -id UINT

Retrieve workflow status by ID
    -id UINT: Identifier of collection to look up

Example:
    %[1]s collection workflow --id 8684027374223699867
`, os.Args[0])
}

func collectionDownloadUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] collection download -id UINT

Download collection by ID
    -id UINT: Identifier of collection to look up

Example:
    %[1]s collection download --id 9570527133616476063
`, os.Args[0])
}

func collectionDecideUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] collection decide -body JSON -id UINT

Make decision for a pending collection by ID
    -body JSON: 
    -id UINT: Identifier of collection to look up

Example:
    %[1]s collection decide --body '{
      "option": "Nihil officiis et enim."
   }' --id 3923434489665488889
`, os.Args[0])
}

func collectionBulkUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] collection bulk -body JSON

Bulk operations (retry, cancel...).
    -body JSON: 

Example:
    %[1]s collection bulk --body '{
      "operation": "cancel",
      "size": 15492129042979125718,
      "status": "in progress"
   }'
`, os.Args[0])
}

func collectionBulkStatusUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] collection bulk-status

Retrieve status of current bulk operation.

Example:
    %[1]s collection bulk-status
`, os.Args[0])
}
