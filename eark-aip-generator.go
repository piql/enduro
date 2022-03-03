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
	// var sips_list []string
	// var sip_validation_results PackageValidationResults
	var package_details []PackageDetails

	// var valid_packages []string
	var batch_submission_data BatchData
	// var collection_data CollectionData
	var eark_validation_results PackageValidationResults

	// Logger
	// If the file doesn't exist, create it or append to the file

	err := os.MkdirAll("logs", os.ModePerm)
	if err != nil {
		return err
	}
	if _, err := os.Stat("logs/logs.txt"); err == nil {
		e := os.Remove("logs/logs.txt")
		if e != nil {
			ErrorLogger.Println(err)
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

		err := future.Get(ctx, nil)
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

		err := future.Get(ctx, &eark_validation_results)
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

		future := workflow.ExecuteActivity(activityOptions, EarkAipValidationReportActivityName, eark_validation_results)

		err := future.Get(ctx, &package_details)
		if err != nil {
			ErrorLogger.Println(err)
			return err
		}
	}

	return nil
}

func ListSipPackagesActivity(ctx context.Context) ([]PackageDetails, error) {
	// var sip_packages []string
	var package_details []PackageDetails

	files, err := ioutil.ReadDir("./sips")
	if err != nil {
		ErrorLogger.Println(err)
		return nil, err
	}
	for _, file := range files {
		if file.IsDir() {
			package_details = append(package_details, PackageDetails{Sip_name: file.Name()})
			// sip_packages = append(sip_packages, file.Name())
		}
	}
	return package_details, nil
}

func ValidateSipPackagesActivity(ctx context.Context, package_details []PackageDetails) ([]PackageDetails, error) {
	// validation_results := PackageValidationResults{}

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

	// Remove invalid sip packages
	for i, pkg := range package_details {
		if !pkg.Sip_valid {
			package_details = append(package_details[:i], package_details[i+1:]...)
		}
	}

	if len(package_details) == 0 {
		err := errors.New("No valid sip packages")
		ErrorLogger.Println(err)
		return nil, err
	}
	return package_details, nil
}

func PrepareAMTransferActivity(ctx context.Context, package_details []PackageDetails) ([]PackageDetails, error) {

	// Delete previous content from 'am_transfers' directory
	// os.IsExist is blind to empty files
	if _, err := os.Stat("am_transfers"); !os.IsNotExist(err) {
		err := RemoveContents("am_transfers")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}

	for _, pkg := range package_details {
		InfoLogger.Println("Package:", pkg.Sip_name)
		cmd := exec.Command("python3.9", "scripts/sip_to_am_transfer/sip_to_am_transfer.py", "-i", "sips/"+pkg.Sip_name, "-o", "am_transfers")
		op, err := cmd.Output()
		InfoLogger.Println("SipToAmTransfer Output:\n", string(op))
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
	//Update package details with newly generated am transfer files usding efficient search function
	package_details, err = transfer_search(package_details, 0, files, 0)
	if err != nil {
		ErrorLogger.Println((err))
		return nil, err
	}
	return package_details, nil
}

func transfer_search(sips []PackageDetails, sip_index int, files []fs.FileInfo, file_index int) ([]PackageDetails, error) {
	if strings.HasPrefix(files[file_index].Name(), sips[sip_index].Sip_name) {
		sips[sip_index].Am_transfers = append(sips[sip_index].Am_transfers, AmTransferDetails{Name: files[file_index].Name()})
		// sips[sip_index].Am_transfers = append(sips[sip_index].Am_transfers, files[file_index].Name())
		file_index += 1
	} else {
		sip_index += 1
	}
	if sip_index == len(sips) || file_index == len(files) {
		return sips, nil
	} else {
		transfer_search(sips, sip_index, files, file_index)
	}
	//err := errors.New("Error in transfer search. Didn't end properly.")
	return sips, nil
}

func ExecuteAMTransferActivity(ctx context.Context) (BatchData, error) {

	var batch_submission_output map[string]interface{}
	//var batch_submission_data []string
	var batch_submission_data BatchData

	working_dir, err := os.Getwd()
	if err != nil {
		return batch_submission_data, err
	}
	InfoLogger.Println("Beginning API Call")
	postBody, _ := json.Marshal(map[string]string{
		"path":     working_dir + "/am_transfers/",
		"pipeline": "am",
	})
	responseBody := bytes.NewBuffer(postBody)
	resp, err := http.Post("http://localhost:9000/batch", "application/json", responseBody)
	submission_time := time.Now().Format("2006-01-02T15:04:05Z")
	if err != nil {
		return batch_submission_data, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return batch_submission_data, err
	}

	err = json.Unmarshal(body, &batch_submission_output)
	if err != nil {
		return batch_submission_data, err
	}

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
				// InfoLogger.Printf("%+v", batch_submission_data)
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

	var batch_status map[string]interface{}

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

	// The output will be in reverse order - with the last process first
	// So we will iterate collection output in reverse.
	// This is to attempt to efficitently assign am trasnfer ids without exponentially timed nested loops
	var col_i = len(collection_output.Items) - 1
	for pkg_i, pkg := range package_details {
		if strings.HasPrefix(collection_output.Items[col_i].Name, pkg.Sip_name) {
			for am_t_i, am_trans := range pkg.Am_transfers {
				if am_trans.Name == collection_output.Items[col_i].Name {
					package_details[pkg_i].Am_transfers[am_t_i].Id = collection_output.Items[col_i].Id
					col_i = col_i - 1
				} else {
					err = errors.New("ERROR: Failure in CollectProcessingDataActivity. " + am_trans.Name + ": " + collection_output.Items[col_i].Name)
					ErrorLogger.Println(package_details, collection_output)
					ErrorLogger.Println(err)
					return nil, err
				}
			}
		} else {
			err = errors.New("ERROR: Failure in CollectProcessingDataActivity. " + pkg.Sip_name + ": " + collection_output.Items[col_i].Name)
			ErrorLogger.Println(err)
			return nil, err
		}
	}
	return package_details, nil
}

func GenerateEarkAipActivity(ctx context.Context, package_details []PackageDetails) ([]PackageDetails, error) {
	for i, pkg := range package_details {
		cmd := exec.Command("python3.9", "scripts/sip_to_eark_aip/sip_to_eark_aip.py", "sips/"+pkg.Sip_name, "eark_aips")
		op, err := cmd.Output()
		split_output := strings.Split(string(op), "\n")
		aip_name := split_output[len(split_output)-2]
		InfoLogger.Println(split_output) 
		InfoLogger.Println("SipToEarkAip Output:\n", aip_name)
		package_details[i].Aip_name = aip_name
		if err != nil {
			ErrorLogger.Println(err.Error())
			return nil, err
		}
	}
	return package_details, nil
}

func WaitForAMProcessActivity(ctx context.Context, package_details []PackageDetails) error {
	for i, pkg := range package_details {
		for j, am_trans := range pkg.Am_transfers {
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
					time.Sleep(time.Minute)
				} else {
					package_details[i].Am_transfers[j].Am_aip_name = collection_item.Aip_id
					InfoLogger.Println(am_trans.Name + "is completed.")
					InfoLogger.Println(am_trans)
				}
			}
		}
	}
	return nil
}

func DownloadAndPlaceAMAIPActivity(ctx context.Context, package_details []PackageDetails) error {

	for _, pkg := range package_details {
		for _, am_trans := range pkg.Am_transfers {
			resp, err := http.Get(fmt.Sprint("http://localhost:9000/collection/", am_trans.Id, "/download"))
			if err != nil {
				ErrorLogger.Println(err)
				return err
			}
			defer resp.Body.Close()

			if resp.StatusCode != 200 {
				return errors.New("Error: Unsuccesful download request")
			}

			preservation_file := "eark_aips/" + pkg.Sip_name + "/representations/rep01.1/data/" + am_trans.Name + ".zip"

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

			InfoLogger.Println(preservation_file + "downloaded and placed.")
		}
	}
	return nil
}

func UpdatePreservationMetsActivity(ctx context.Context, valid_packages []string) error {
	for _, pkg := range valid_packages {
		location := "eark_aips/" + pkg + "/representations/rep01.1"
		cmd := exec.Command("python3.9", "scripts/sip_to_eark_aip/update_rep_mets.py", location)
		op, err := cmd.Output()
		InfoLogger.Println(string(op))
		if err != nil {
			ErrorLogger.Println(err.Error())
			return err
		}
	}
	return nil
}

func ValidateEarkAipPackagesActivity(ctx context.Context, eark_packages []string) (PackageValidationResults, error) {
	validation_results := PackageValidationResults{}

	for pkg := range eark_packages {
		cmd := exec.Command("java", "-jar", "scripts/java/commons-ip2-cli-2.0.1.jar", "validate", "-i", "eark_aips/"+eark_packages[pkg])
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
		validation_results[eark_packages[pkg]] = data.Summary.Result == "VALID"
	}
	return validation_results, nil
}

func EarkAipValidationReportActivity(ctx context.Context, validation_results PackageValidationResults) ([]string, error) {
	var valid_packages []string

	for identifier, passed := range validation_results {
		if passed {
			valid_packages = append(valid_packages, identifier)
		}
	}
	if len(valid_packages) == 0 {
		err := errors.New("No valid eark packages")
		ErrorLogger.Println(err)
		return nil, err
	}
	return valid_packages, nil
}

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

type PackageDetails struct {
	Sip_name     string
	Sip_valid    bool
	Am_transfers []AmTransferDetails
	Aip_name     string
}

type AmTransferDetails struct {
	Name        string
	Id          int
	Am_aip_name string
}

type BatchData struct {
	Run_id string `json:"run_id"`
	Time   string `json:"time"`
}

type CollectionData struct {
	Items       []EnduroItem `json:"items"`
	Next_cursor string       `json:"next_cursor"`
}

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
	Status       string `json:"status"`
	Transfer_id  string `json:"transfer_id"`
	Workflow_id  string `json:"workflow_id"`
}

type PackageValidationResults map[string]bool

type CommonsValidatorData struct {
	Header     map[string]interface{}  `json:"header"`
	Validation map[string]interface{}  `json:"validation"`
	Summary    CommonsValidatorSummary `json:"summary"`
}

type CommonsValidatorSummary struct {
	Success  int    `json:"success"`
	Warnings int    `json:"warnings"`
	Errors   int    `json:"errors"`
	Skipped  int    `json:"skipped"`
	Notes    int    `json:"notes"`
	Result   string `json:"result"`
}
