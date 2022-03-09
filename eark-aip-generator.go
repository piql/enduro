package main

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"go.uber.org/cadence/activity"
	"go.uber.org/cadence/worker"
	"go.uber.org/cadence/workflow"
)

const (
	EarkAipGeneratorWorkflowName        = "eark-aip-generator"
	ListSipPackagesActivityName         = "sip:list-packages"
	ValidateSipPackagesActivityName     = "sip:validate-packages"
	SipValidationReportActivityName     = "sip:validation-report"
	PrepareAMTransferActivityName       = "am-aip:prepare-transfer"
	ExecuteAMTransferActivityName       = "am-aip:execute-transfer"
	WaitForBatchActivityName            = "am-aip:wait-for-batch"
	CollectProcessingDataActivityName   = "am-aip:collect-processing-data"
	GenerateEarkAipActivityName         = "eark-aip:gernerate-eark-aip"
	WaitForAMProcessActivityName        = "am-aip:wait-for-am-process"
	DownloadAndPlaceAMAIPActivityName   = "am-aip:download-and-place-am-aip"
	UpdatePreservationMetsActivityName  = "eark-aip:update-preservation-mets"
	ValidateEarkAipPackagesActivityName = "eark-aip:validare-eark-aip"
	EarkAipValidationReportActivityName = "eark-aip:validation-report"
	ValidateTaskListName                = "global"
)

// Logger
var (
	WarningLogger *log.Logger
	InfoLogger    *log.Logger
	ErrorLogger   *log.Logger
)

func registerEarkAipGeneratorWorkflowActivities(w worker.Worker) {
	//Register workflow
	w.RegisterWorkflowWithOptions(
		EarkAipGeneratorWorkflow,
		workflow.RegisterOptions{Name: EarkAipGeneratorWorkflowName},
	)
	w.RegisterActivityWithOptions(
		ListSipPackagesActivity,
		activity.RegisterOptions{Name: ListSipPackagesActivityName},
	)
	w.RegisterActivityWithOptions(
		ValidateSipPackagesActivity,
		activity.RegisterOptions{Name: ValidateSipPackagesActivityName},
	)
	w.RegisterActivityWithOptions(
		SipValidationReportActivity,
		activity.RegisterOptions{Name: SipValidationReportActivityName},
	)
	w.RegisterActivityWithOptions(
		PrepareAMTransferActivity,
		activity.RegisterOptions{Name: PrepareAMTransferActivityName},
	)
	w.RegisterActivityWithOptions(
		ExecuteAMTransferActivity,
		activity.RegisterOptions{Name: ExecuteAMTransferActivityName},
	)
	w.RegisterActivityWithOptions(
		WaitForBatchActivity,
		activity.RegisterOptions{Name: WaitForBatchActivityName},
	)
	w.RegisterActivityWithOptions(
		CollectProcessingDataActivity,
		activity.RegisterOptions{Name: CollectProcessingDataActivityName},
	)
	w.RegisterActivityWithOptions(
		GenerateEarkAipActivity,
		activity.RegisterOptions{Name: GenerateEarkAipActivityName},
	)
	w.RegisterActivityWithOptions(
		WaitForAMProcessActivity,
		activity.RegisterOptions{Name: WaitForAMProcessActivityName},
	)
	w.RegisterActivityWithOptions(
		DownloadAndPlaceAMAIPActivity,
		activity.RegisterOptions{Name: DownloadAndPlaceAMAIPActivityName},
	)
	w.RegisterActivityWithOptions(
		UpdatePreservationMetsActivity,
		activity.RegisterOptions{Name: UpdatePreservationMetsActivityName},
	)
	w.RegisterActivityWithOptions(
		ValidateEarkAipPackagesActivity,
		activity.RegisterOptions{Name: ValidateEarkAipPackagesActivityName},
	)
	w.RegisterActivityWithOptions(
		EarkAipValidationReportActivity,
		activity.RegisterOptions{Name: EarkAipValidationReportActivityName},
	)
}

func EarkAipGeneratorWorkflow(ctx workflow.Context) error {

	var package_details []PackageDetails
	var batch_submission_data BatchData

	// Logger
	// If the file doesn't exist, create it or append to the file
	if _, err := os.Stat("logs/logs.txt"); err == nil {
		e := os.Remove("logs/logs.txt")
		if e != nil {
			ErrorLogger.Println(err)
			return err
		}
	} else {
		err := os.MkdirAll("logs", os.ModePerm)
		if err != nil {
			return err
		}
	}
	file, err := os.OpenFile("logs/logs.txt", os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	InfoLogger = log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	WarningLogger = log.New(file, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)

	// List SIP Packages
	{
		activityOptions := workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
			TaskList:               ValidateTaskListName,
			ScheduleToStartTimeout: time.Second * 10,
			StartToCloseTimeout:    time.Minute,
		})

		future := workflow.ExecuteActivity(activityOptions, ListSipPackagesActivityName)

		err := future.Get(ctx, &package_details)
		if err != nil {
			ErrorLogger.Println(err)
			return err
		}
	}

	//Validate SIP Packages
	{
		activityOptions := workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
			TaskList:               ValidateTaskListName,
			ScheduleToStartTimeout: time.Second * 10,
			StartToCloseTimeout:    time.Minute,
		})

		future := workflow.ExecuteActivity(activityOptions, ValidateSipPackagesActivityName, package_details)

		err := future.Get(ctx, &package_details)
		if err != nil {
			ErrorLogger.Println(err)
			return err
		}
	}

	// SIP Validation Report
	{
		activityOptions := workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
			TaskList:               ValidateTaskListName,
			ScheduleToStartTimeout: time.Second * 10,
			StartToCloseTimeout:    time.Minute,
		})

		future := workflow.ExecuteActivity(activityOptions, SipValidationReportActivityName, package_details)

		err := future.Get(ctx, &package_details)
		if err != nil {
			ErrorLogger.Println(err)
			return err
		}
	}

	// Prepare AM Transfer
	{
		activityOptions := workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
			TaskList:               ValidateTaskListName,
			ScheduleToStartTimeout: time.Second * 10,
			StartToCloseTimeout:    time.Minute,
		})

		future := workflow.ExecuteActivity(activityOptions, PrepareAMTransferActivityName, package_details)

		err := future.Get(ctx, &package_details)
		if err != nil {
			ErrorLogger.Println(err)
			return err
		}
	}

	// Execute AM Transfer

	{
		activityOptions := workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
			TaskList:               ValidateTaskListName,
			ScheduleToStartTimeout: time.Second * 10,
			StartToCloseTimeout:    time.Minute,
		})

		future := workflow.ExecuteActivity(activityOptions, ExecuteAMTransferActivityName)

		err := future.Get(ctx, &batch_submission_data)
		if err != nil {
			ErrorLogger.Println(err)
			return err
		}
	}

	// Wait for batch process to complete
	{
		activityOptions := workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
			TaskList:               ValidateTaskListName,
			ScheduleToStartTimeout: time.Second * 10,
			StartToCloseTimeout:    time.Minute,
		})

		future := workflow.ExecuteActivity(activityOptions, WaitForBatchActivityName, batch_submission_data)

		err := future.Get(ctx, nil)
		if err != nil {
			ErrorLogger.Println(err)
			return err
		}
	}

	// Collect processing data
	{
		activityOptions := workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
			TaskList:               ValidateTaskListName,
			ScheduleToStartTimeout: time.Second * 10,
			StartToCloseTimeout:    time.Minute,
		})

		future := workflow.ExecuteActivity(activityOptions, CollectProcessingDataActivityName, batch_submission_data, package_details)

		err := future.Get(ctx, &package_details)
		if err != nil {
			ErrorLogger.Println(err)
			return err
		}
	}

	// Generate EARK AIP
	{
		activityOptions := workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
			TaskList:               ValidateTaskListName,
			ScheduleToStartTimeout: time.Second * 10,
			StartToCloseTimeout:    time.Minute,
		})

		future := workflow.ExecuteActivity(activityOptions, GenerateEarkAipActivityName, package_details)

		err := future.Get(ctx, &package_details)
		if err != nil {
			ErrorLogger.Println(err)
			return err
		}
	}

	// Wait for AM process
	{
		activityOptions := workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
			TaskList:               ValidateTaskListName,
			ScheduleToStartTimeout: time.Second * 10,
			StartToCloseTimeout:    time.Minute * 15 * time.Duration(len(package_details)),
		})

		future := workflow.ExecuteActivity(activityOptions, WaitForAMProcessActivityName, package_details)

		err := future.Get(ctx, &package_details)
		if err != nil {
			ErrorLogger.Println(err)
			return err
		}
	}

	// Download and place am aip
	{
		activityOptions := workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
			TaskList:               ValidateTaskListName,
			ScheduleToStartTimeout: time.Second * 10,
			StartToCloseTimeout:    time.Minute,
		})

		future := workflow.ExecuteActivity(activityOptions, DownloadAndPlaceAMAIPActivityName, package_details)

		err := future.Get(ctx, &package_details)
		if err != nil {
			ErrorLogger.Println(err)
			return err
		}
	}

	// Update preservation mets
	{
		activityOptions := workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
			TaskList:               ValidateTaskListName,
			ScheduleToStartTimeout: time.Second * 10,
			StartToCloseTimeout:    time.Minute,
		})

		future := workflow.ExecuteActivity(activityOptions, UpdatePreservationMetsActivityName, package_details)

		err := future.Get(ctx, nil)
		if err != nil {
			ErrorLogger.Println(err)
			return err
		}
	}

	// Validate eark aip
	{
		activityOptions := workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
			TaskList:               ValidateTaskListName,
			ScheduleToStartTimeout: time.Second * 10,
			StartToCloseTimeout:    time.Minute,
		})

		future := workflow.ExecuteActivity(activityOptions, ValidateEarkAipPackagesActivityName, package_details)

		err := future.Get(ctx, &package_details)
		if err != nil {
			ErrorLogger.Println(err)
			return err
		}
	}

	// EARK AIP Validation Report
	{
		activityOptions := workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
			TaskList:               ValidateTaskListName,
			ScheduleToStartTimeout: time.Second * 10,
			StartToCloseTimeout:    time.Minute,
		})

		future := workflow.ExecuteActivity(activityOptions, EarkAipValidationReportActivityName, package_details)

		err := future.Get(ctx, nil)
		if err != nil {
			ErrorLogger.Println(err)
			return err
		}
	}

	return nil
}

func ListSipPackagesActivity(ctx context.Context) ([]PackageDetails, error) {

	InfoLogger.Println("Starting: List Packages")

	var package_details []PackageDetails

	files, err := ioutil.ReadDir("./sips")
	if err != nil {
		ErrorLogger.Println(err)
		return nil, err
	}
	for _, file := range files {
		if file.IsDir() {
			package_details = append(package_details, PackageDetails{Sip_name: file.Name()})
		}
	}
	return package_details, nil
}

func ValidateSipPackagesActivity(ctx context.Context, package_details []PackageDetails) ([]PackageDetails, error) {

	InfoLogger.Println("Starting: List Packages")

	// Run the validator on each package to ensure each SIP is valid
	for i, pkg := range package_details {
		cmd := exec.Command("java", "-jar", "scripts/commons-ip2-cli-2.0.1.jar", "validate", "-i", "sips/"+pkg.Sip_name)
		stdout, err := cmd.Output()
		if err != nil {
			ErrorLogger.Println(err)
			return nil, err
		}

		// Get the report location from the validation output string
		path := strings.Replace(string(stdout), "\n", "", -1)

		jsonFile, err := os.Open(path)
		if err != nil {
			ErrorLogger.Println(err)
			return nil, err
		}
		defer jsonFile.Close()

		byteValue, _ := ioutil.ReadAll(jsonFile)
		var data CommonsValidatorData
		json.Unmarshal([]byte(byteValue), &data)
		// pkg is a copy of package details so must set variables using index
		package_details[i].Sip_valid = data.Summary.Result == "VALID"
	}
	return package_details, nil
}

func SipValidationReportActivity(ctx context.Context, package_details []PackageDetails) ([]PackageDetails, error) {

	InfoLogger.Println("Starting: SIP Validation")

	// Remove invalid sip packages from the list
	for i, pkg := range package_details {
		if !pkg.Sip_valid {
			package_details = append(package_details[:i], package_details[i+1:]...)
		}
	}

	// If no packages are valid we terminate the process
	if len(package_details) == 0 {
		err := errors.New("No valid sip packages")
		ErrorLogger.Println(err)
		return nil, err
	}
	return package_details, nil
}

func PrepareAMTransferActivity(ctx context.Context, package_details []PackageDetails) ([]PackageDetails, error) {

	InfoLogger.Println("Starting: Prepare AM Transfers")

	// Delete existing (previous) content from 'am_transfers' directory
	// os.IsExist is blind to empty files
	if _, err := os.Stat("am_transfers"); !os.IsNotExist(err) {
		err := RemoveContents("am_transfers")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}

	// Create Archivematica transfers for each sip
	// Creates a transfer directory for each representation in the SIP
	for _, pkg := range package_details {
		InfoLogger.Println("Package:", pkg.Sip_name)
		cmd := exec.Command("python3.9", "scripts/sip_to_am_transfer/sip_to_am_transfer.py", "-i", "sips/"+pkg.Sip_name, "-o", "am_transfers")
		_, err := cmd.Output()
		//InfoLogger.Println("SipToAmTransfer Output:\n", string(op))
		if err != nil {
			ErrorLogger.Println(err)
			return nil, err
		}
	}
	// Append transfer file names to package details am_transfers[]
	files, err := ioutil.ReadDir("am_transfers")
	if err != nil {
		ErrorLogger.Println((err))
		return nil, err
	}
	// Update package details with newly generated am transfer files useing efficient search function
	// package_details, err = transfer_search(package_details, 0, files, 0)
	package_details, err = tsearch(package_details, files)
	if err != nil {
		ErrorLogger.Println((err))
		return nil, err
	}
	return package_details, nil
}

// This function efficiently searches through the files[] appending each am_transfer file to its respective SIP (package_details.Am_transfers[])
// It does this recursively, using indexes for the sips[] and files[]. Only increasing files_index when a match has been found
// (i.e. file name starts with sip name) and appended. And only increasing the sips_index when a match isn't found and therefore must be an
// am_transfer for the next SIP. This rides on the fect that ioutil.ReadDir returns and array in alphabetical order and package_package details
// is also alphabetical on Sip_name.
func transfer_search(sips []PackageDetails, sip_index int, files []fs.FileInfo, file_index int) ([]PackageDetails, error) {
	if strings.HasPrefix(files[file_index].Name(), sips[sip_index].Sip_name) {
		sips[sip_index].Am_transfers = append(sips[sip_index].Am_transfers, AmTransferDetails{Name: files[file_index].Name()})
		file_index += 1
	} else {
		sip_index += 1
	}
	if sip_index == len(sips) || file_index == len(files) {
		WarningLogger.Println(len(sips), sip_index, len(files), file_index)
		return sips, nil
	} else {
		WarningLogger.Println("Looping.")
		transfer_search(sips, sip_index, files, file_index)
	}
	WarningLogger.Println("Made it to end.")
	return sips, nil
}

func tsearch(package_details []PackageDetails, files []fs.FileInfo) ([]PackageDetails, error) {
	var file_i = 0
	for i, pkg := range package_details {
		if strings.HasPrefix(files[file_i].Name(), pkg.Sip_name) {
			package_details[i].Am_transfers = append(package_details[i].Am_transfers, AmTransferDetails{Name: files[file_i].Name()})
			file_i += 1
		} else {
			err := errors.New("Mismatch in transfer search")
			return nil, err
		}
	}
	return package_details, nil
}

func ExecuteAMTransferActivity(ctx context.Context) (BatchData, error) {

	InfoLogger.Println("Starting: Execute AM Batch Transfer")

	var batch_submission_output map[string]interface{}
	var batch_submission_data BatchData

	// Prepare and execute API call
	working_dir, err := os.Getwd()
	if err != nil {
		return batch_submission_data, err
	}
	postBody, _ := json.Marshal(map[string]string{
		"path":     working_dir + "/am_transfers/",
		"pipeline": "am",
	})
	responseBody := bytes.NewBuffer(postBody)
	submission_time := time.Now().Format("2006-01-02T15:04:05Z")
	WarningLogger.Println("API call made at:", submission_time)
	resp, err := http.Post("http://localhost:9000/batch", "application/json", responseBody)
	if err != nil {
		return batch_submission_data, err
	}
	defer resp.Body.Close()

	// Read the repsonse to our http post request
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return batch_submission_data, err
	}

	err = json.Unmarshal(body, &batch_submission_output)
	if err != nil {
		return batch_submission_data, err
	}

	// If it fails, determine why
	// If it succeeds, store the run ID and the submission time to BatchData and return it
	if fault, ok := batch_submission_output["fault"]; ok {
		var fault_str string
		if fault == true {
			fault_str = "Bad Request Response"
		} else {
			fault_str = "Conflict Response"
		}
		err = errors.New("Error: batch submission failed: " + fault_str)
		ErrorLogger.Println(err)
		return batch_submission_data, err
	} else {
		if workflow_id, ok := batch_submission_output["workflow_id"]; ok {
			if workflow_id == "batch-workflow" {
				run_id := batch_submission_output["run_id"].(string)
				batch_submission_data = BatchData{run_id, submission_time}
				return batch_submission_data, nil
			} else {
				err = errors.New("Error: batch submission failed")
				ErrorLogger.Println(err)
				return batch_submission_data, err
			}
		}
	}
	err = errors.New("Error: batch submission failed: Unknown")
	ErrorLogger.Println(err)
	return batch_submission_data, err
}

func WaitForBatchActivity(ctx context.Context, batch_data BatchData) error {

	InfoLogger.Println("Starting: Wait For Batch Completion")

	var batch_status map[string]interface{}

	// Loops until the batch process is complete or fails
	for ok := true; ok; ok = true {
		resp, err := http.Get("http://localhost:9000/batch")
		if err != nil {
			ErrorLogger.Println(err)
			return err
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			ErrorLogger.Println(err)
			return err
		}
		err = json.Unmarshal(body, &batch_status)
		if err != nil {
			ErrorLogger.Println(err)
			return err
		}

		// If the batch process is still running we wait 5 seconds before checking again
		// If it not running it is completed or has failed
		if running, ok := batch_status["running"]; ok {
			if running == true {
				time.Sleep(time.Second * 5)
			} else if running == false {
				if status, ok := batch_status["status"]; ok {
					if status == "completed" {
						InfoLogger.Println("Batch workflow complete.")
						return nil
					} else {
						err = errors.New("Error: batch workflow not completed")
						ErrorLogger.Println(err)
						return err
					}
				}
			} else {
				err = errors.New("Error reading batch status")
				ErrorLogger.Println(err)
				return err
			}
		} else {
			err = errors.New("Error reading batch status")
			ErrorLogger.Println(err)
			return err
		}
	}
	return nil
}

func CollectProcessingDataActivity(ctx context.Context, batch_data BatchData, package_details []PackageDetails) ([]PackageDetails, error) {

	InfoLogger.Println("Starting: Collect Processing Data")

	var collection_output CollectionData
	submission_time := batch_data.Time

	resp, err := http.Get("http://localhost:9000/collection?earliest_created_time=" + submission_time)
	if err != nil {
		ErrorLogger.Println(err)
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		ErrorLogger.Println(err)
		return nil, err
	}
	err = json.Unmarshal(body, &collection_output)
	if err != nil {
		ErrorLogger.Println(err)
		return nil, err
	}

	// To efficiently assign process IDs to the am_transfers we must first sort the returned list of collection items alphabetically on the item name.
	sort.Slice(collection_output.Items[:], func(i, j int) bool {
		return collection_output.Items[i].Name < collection_output.Items[j].Name
	})

	// Here we loop through the packages_details and their am-transfers to assign the ID of each.
	// This process rides on collection_output.Items being in alphabetical order on name and package details and their am_transfer are also sorted alphabetically.
	// With every loop the am_transfer[am_t_i].Name should match collection_items[col_i].Name
	var col_i = 0
	collection_items := collection_output.Items
	for pkg_i, pkg := range package_details {
		if strings.HasPrefix(collection_items[col_i].Name, pkg.Sip_name) {
			for am_t_i, am_trans := range pkg.Am_transfers {
				if am_trans.Name == collection_items[col_i].Name {
					package_details[pkg_i].Am_transfers[am_t_i].Id = collection_items[col_i].Id
					col_i = col_i + 1
				} else {
					err = errors.New("ERROR: Failure in CollectProcessingDataActivity. " + am_trans.Name + ": " + collection_items[col_i].Name)
					ErrorLogger.Println(package_details, collection_output)
					ErrorLogger.Println(err)
					return nil, err
				}
			}
		} else {
			err = errors.New("ERROR: Failure in CollectProcessingDataActivity. " + pkg.Sip_name + ": " + collection_items[col_i].Name)
			ErrorLogger.Println(err)
			return nil, err
		}
	}
	return package_details, nil
}

func GenerateEarkAipActivity(ctx context.Context, package_details []PackageDetails) ([]PackageDetails, error) {

	InfoLogger.Println("Starting: Generate EARK AIP")

	// For each SIP package generate a new EARK AIP from it's contents.
	// The script should output the name of the resulting AIP which we set in package_details
	for i, pkg := range package_details {
		cmd := exec.Command("python3.9", "scripts/sip_to_eark_aip/sip_to_eark_aip.py", "sips/"+pkg.Sip_name, "eark_aips")
		op, err := cmd.Output()
		split_output := strings.Split(string(op), "\n")
		aip_name := split_output[len(split_output)-2]
		InfoLogger.Println("New AIP Name:", aip_name)
		package_details[i].Aip_name = aip_name
		if err != nil {
			ErrorLogger.Println(err.Error())
			return nil, err
		}
	}
	return package_details, nil
}

func WaitForAMProcessActivity(ctx context.Context, package_details []PackageDetails) error {

	InfoLogger.Println("Starting: Wait For Am Process")

	// For each am_transfer we check the status of its process through Archivematica.
	// If the process is still running we wait for 30 seconds before checking again.
	// We wait until all processes are complete. If one fails, we abort.
	for _, pkg := range package_details {
		for _, am_trans := range pkg.Am_transfers {
			completed_process := false
			for ok := true; ok; ok = !completed_process {
				var collection_item EnduroItem
				resp, err := http.Get(fmt.Sprint("http://localhost:9000/collection/", am_trans.Id))

				if err != nil {
					ErrorLogger.Println(err)
					return err
				}
				defer resp.Body.Close()
				body, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					ErrorLogger.Println(err)
					return err
				}
				err = json.Unmarshal(body, &collection_item)
				if err != nil {
					ErrorLogger.Println(err)
					return err
				}
				// InfoLogger.Printf("Collection Item: %+v", collection_item)
				if collection_item.Status == "error" {
					err = errors.New("Error: AM Process Failed: " + collection_item.Workflow_id)
					ErrorLogger.Println(err)
					return err
				}
				completed_process = collection_item.Status == "done"
				if !completed_process {
					time.Sleep(time.Second * 30)
				} else {
					InfoLogger.Println(am_trans.Name + " AM process completed.")
				}
			}
		}
	}
	return nil
}

func DownloadAndPlaceAMAIPActivity(ctx context.Context, package_details []PackageDetails) error {

	InfoLogger.Println("Starting: Download and Place AM AIP")

	// Iterate through all packages and their am_transfers to download the completed AIPs from Archivematica.
	// Place the downloaded AIPs in their respective directory within the EARK AIP representations folder
	for _, pkg := range package_details {
		for j, am_trans := range pkg.Am_transfers {
			resp, err := http.Get(fmt.Sprint("http://localhost:9000/collection/", am_trans.Id, "/download"))
			if err != nil {
				ErrorLogger.Println(err)
				return err
			}
			defer resp.Body.Close()

			if resp.StatusCode != 200 {
				return errors.New("Error: Unsuccesful download request")
			}

			var rep_num = fmt.Sprintf("%02d", j+1)
			preservation_file := "eark_aips/" + pkg.Aip_name + "/representations/rep" + rep_num + ".1/data/" + am_trans.Name + ".zip"

			// Create the file
			out, err := os.Create(preservation_file)
			if err != nil {
				ErrorLogger.Println(err)
				return err
			}
			defer out.Close()

			// Write the body to file
			_, err = io.Copy(out, resp.Body)
			if err != nil {
				ErrorLogger.Println(err)
				return err
			}
		}
	}
	return nil
}

func UpdatePreservationMetsActivity(ctx context.Context, package_details []PackageDetails) error {

	InfoLogger.Println("Starting: Update Preservation METS")

	// For each preservation rep we need to update the preservation METS as well as the root mets.
	// This script does both.
	for _, pkg := range package_details {
		for j := range pkg.Am_transfers {
			var rep_num = fmt.Sprintf("%02d", j+1)
			location := "eark_aips/" + pkg.Aip_name + "/representations/rep" + rep_num + ".1"
			cmd := exec.Command("python3.9", "scripts/sip_to_eark_aip/update_rep_mets.py", location)
			_, err := cmd.Output()
			// InfoLogger.Println(string(op))
			if err != nil {
				ErrorLogger.Println(err.Error())
				return err
			}
		}
	}
	return nil
}

func ValidateEarkAipPackagesActivity(ctx context.Context, package_details []PackageDetails) ([]PackageDetails, error) {

	InfoLogger.Println("Starting: Validate EARK AIP")

	// Iterate over every package running the resulting EARK AIP through the commomns IP validator
	for i, pkg := range package_details {
		cmd := exec.Command("java", "-jar", "scripts/commons-ip2-cli-2.0.1.jar", "validate", "-i", "eark_aips/"+pkg.Aip_name)
		stdout, err := cmd.Output()

		if err != nil {
			ErrorLogger.Println(err)
			return nil, err
		}

		path := strings.Replace(string(stdout), "\n", "", -1)

		jsonFile, err := os.Open(path)
		if err != nil {
			ErrorLogger.Println(err)
			return nil, err
		}
		defer jsonFile.Close()

		byteValue, _ := ioutil.ReadAll(jsonFile)
		var data CommonsValidatorData
		json.Unmarshal([]byte(byteValue), &data)
		package_details[i].Aip_valid = data.Summary.Result == "VALID"
	}
	return package_details, nil
}

func EarkAipValidationReportActivity(ctx context.Context, package_details []PackageDetails) error {

	InfoLogger.Println("Starting: EARK AIP Validation Report")

	// Display in the logs every EARK AIP that fails verification
	for _, pkg := range package_details {
		if !pkg.Aip_valid {
			WarningLogger.Println("AIP Package:", pkg.Sip_name, pkg.Aip_name, "failed validation")
		}
	}
	return nil
}

// Removes the contents of a directory.
func RemoveContents(dir string) error {
	d, err := os.Open(dir)
	if err != nil {
		return err
	}
	defer d.Close()
	names, err := d.Readdirnames(-1)
	if err != nil {
		return err
	}
	for _, name := range names {
		err = os.RemoveAll(filepath.Join(dir, name))
		if err != nil {
			return err
		}
	}
	return nil
}

// Used throughout the process to track details pertaining to each package
type PackageDetails struct {

	// Name of the original SIP
	Sip_name string

	// Validation status of the SIP
	Sip_valid bool

	// List of Archivematica transfers generated. One for each representation in the SIP
	Am_transfers []AmTransferDetails

	// Name of the EARK AIP generated and placed in 'eark-aips/' from running sip_to_eark_aip.py
	Aip_name string

	// Validation status of the EARK AIP
	Aip_valid bool
}

// Stores an Archivematica transfer name and process ID
type AmTransferDetails struct {

	// The name of the Archivematica transfer folder
	Name string

	// The Enduro item ID
	// Alllows monitoring of the enduro process and download of the final Archivematica AIP
	Id int
}

// Stores the data needed to monitor the batch API call and gether its resulting processes
type BatchData struct {

	// The batch run_id - used to monitor the progress of the batch process
	Run_id string `json:"run_id"`

	// The submission time of the API call - used to gather the collection of processes generated after this time, thus generated by the batch API call. (99.9% Effective - needs improving)
	Time string `json:"time"`
}

// Used to read the response from an enduro API collections http request
type CollectionData struct {

	// List of stored items in collection
	Items []EnduroItem `json:"items"`

	// Not used
	Next_cursor string `json:"next_cursor"`
}

// Stores the enduro processing workflow details
// Allows for monitoring of the process
type EnduroItem struct {
	Aip_id       string `json:"aip_id"`
	Completed_at string `json:"completed_at"` // time
	Created_at   string `json:"created_at"`   // time
	Id           int    `json:"id"`
	Name         string `json:"name"`
	Original_id  string `json:"original_id"`
	Pipeline_id  string `json:"pipeline_id"`
	Run_id       string `json:"run_id"`
	Started_at   string `json:"started_at"` // time

	// Status is used to check if the process is still running and whether it is complete/failed
	Status string `json:"status"`

	Transfer_id string `json:"transfer_id"`
	Workflow_id string `json:"workflow_id"`
}

// Used to read data returned from the Commons IP validator
type CommonsValidatorData struct {
	Header     map[string]interface{}  `json:"header"`
	Validation map[string]interface{}  `json:"validation"`
	Summary    CommonsValidatorSummary `json:"summary"`
}

// The summary is the most important bit as we only need the final validation result.
type CommonsValidatorSummary struct {
	Success  int `json:"success"`
	Warnings int `json:"warnings"`
	Errors   int `json:"errors"`
	Skipped  int `json:"skipped"`
	Notes    int `json:"notes"`

	// Result should read "VALID" if sucessful, otherwise invalid
	Result string `json:"result"`
}
