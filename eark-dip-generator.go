package main

import (
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"

	"go.uber.org/cadence/activity"
	"go.uber.org/cadence/worker"
	"go.uber.org/cadence/workflow"
)

const (
	EarkDipGeneratorWorkflowName        = "eark-dip-generator"
	ListAipPackagesActivityName         = "aip:list-aip-packages"
	ValidateAipPackagesActivityName     = "aip:validate-packages"
	AipValidationReportActivityName     = "aip:validation-report"
	GenerateEarkDipActivityName         = "dip:gernerate-eark-dip"
	ValidateEarkDipPackagesActivityName = "dip:validate-eark-dip"
	EarkDipValidationReportActivityName = "dip:validation-report"
	DipGeneratorTaskListName            = "global"

//	CommonsIPValidatorName = "commons-ip2-cli-2.2.1.jar"
)

// Logger
//var (
//	WarningLogger *log.Logger
//	InfoLogger    *log.Logger
//	ErrorLogger   *log.Logger
//)

func registerEarkDipGeneratorWorkflowActivities(w worker.Worker) {
	//Register workflow
	w.RegisterWorkflowWithOptions(
		EarkDipGeneratorWorkflow,
		workflow.RegisterOptions{Name: EarkDipGeneratorWorkflowName},
	)
	w.RegisterActivityWithOptions(
		ListAipPackagesActivity,
		activity.RegisterOptions{Name: ListAipPackagesActivityName},
	)
	w.RegisterActivityWithOptions(
		ValidateAipPackagesActivity,
		activity.RegisterOptions{Name: ValidateAipPackagesActivityName},
	)
	w.RegisterActivityWithOptions(
		AipValidationReportActivity,
		activity.RegisterOptions{Name: AipValidationReportActivityName},
	)
	w.RegisterActivityWithOptions(
		GenerateEarkDipActivity,
		activity.RegisterOptions{Name: GenerateEarkDipActivityName},
	)
	w.RegisterActivityWithOptions(
		ValidateEarkDipPackagesActivity,
		activity.RegisterOptions{Name: ValidateEarkDipPackagesActivityName},
	)
	w.RegisterActivityWithOptions(
		EarkDipValidationReportActivity,
		activity.RegisterOptions{Name: EarkDipValidationReportActivityName},
	)
}

func EarkDipGeneratorWorkflow(ctx workflow.Context) error {

	var package_details []DIPPackageDetails

	// Logger
	// If the file doesn't exist, create it or append to the file
	if _, err := os.Stat("logs/eark-dip-gen.log"); err == nil {
		e := os.Remove("logs/eark-dip-gen.log")
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
	file, err := os.OpenFile("logs/eark-dip-gen.log", os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	InfoLogger = log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	WarningLogger = log.New(file, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)

	// List AIP Packages
	{
		activityOptions := workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
			TaskList:               DipGeneratorTaskListName,
			ScheduleToStartTimeout: time.Second * 10,
			StartToCloseTimeout:    time.Minute,
		})

		future := workflow.ExecuteActivity(activityOptions, ListAipPackagesActivityName)

		err := future.Get(ctx, &package_details)
		if err != nil {
			ErrorLogger.Println(err)
			return err
		}
	}

	//Validate AIP Packages
	{
		activityOptions := workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
			TaskList:               DipGeneratorTaskListName,
			ScheduleToStartTimeout: time.Second * 10,
			StartToCloseTimeout:    time.Minute,
		})

		future := workflow.ExecuteActivity(activityOptions, ValidateAipPackagesActivityName, package_details)

		err := future.Get(ctx, &package_details)
		if err != nil {
			ErrorLogger.Println(err)
			return err
		}
	}

	// AIP Validation Report
	{
		activityOptions := workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
			TaskList:               DipGeneratorTaskListName,
			ScheduleToStartTimeout: time.Second * 10,
			StartToCloseTimeout:    time.Minute,
		})

		future := workflow.ExecuteActivity(activityOptions, AipValidationReportActivityName, package_details)

		err := future.Get(ctx, &package_details)
		if err != nil {
			ErrorLogger.Println(err)
			return err
		}
	}

	// Generate EARK DIP
	{
		activityOptions := workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
			TaskList:               DipGeneratorTaskListName,
			ScheduleToStartTimeout: time.Second * 10,
			StartToCloseTimeout:    time.Minute,
		})

		future := workflow.ExecuteActivity(activityOptions, GenerateEarkDipActivityName, package_details)

		err := future.Get(ctx, &package_details)
		if err != nil {
			ErrorLogger.Println(err)
			return err
		}
	}

	// Validate EARK DIP
	{
		activityOptions := workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
			TaskList:               DipGeneratorTaskListName,
			ScheduleToStartTimeout: time.Second * 10,
			StartToCloseTimeout:    time.Minute,
		})

		future := workflow.ExecuteActivity(activityOptions, ValidateEarkDipPackagesActivityName, package_details)

		err := future.Get(ctx, &package_details)
		if err != nil {
			ErrorLogger.Println(err)
			return err
		}
	}

	// EARK DIP Validation Report
	{
		activityOptions := workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
			TaskList:               DipGeneratorTaskListName,
			ScheduleToStartTimeout: time.Second * 10,
			StartToCloseTimeout:    time.Minute,
		})

		future := workflow.ExecuteActivity(activityOptions, EarkDipValidationReportActivityName, package_details)

		err := future.Get(ctx, nil)
		if err != nil {
			ErrorLogger.Println(err)
			return err
		}
	}

	return nil
}

func ListAipPackagesActivity(ctx context.Context) ([]DIPPackageDetails, error) {

	InfoLogger.Println("Starting: List AIP Packages")

	var package_details []DIPPackageDetails

	files, err := ioutil.ReadDir("./aips")
	if err != nil {
		ErrorLogger.Println(err)
		return nil, err
	}
	for _, file := range files {
		if file.IsDir() {
			package_details = append(package_details, DIPPackageDetails{Dip_name: file.Name()})
		}
	}
	return package_details, nil
}

func ValidateAipPackagesActivity(ctx context.Context, package_details []DIPPackageDetails) ([]DIPPackageDetails, error) {

	InfoLogger.Println("Starting: Validate DIP Packages")

	// Run the validator on each package to ensure each SIP is valid
	for i, pkg := range package_details {
		cmd := exec.Command("java", "-jar", "scripts/"+CommonsIPValidatorName, "validate", "-i", "aips/"+pkg.Aip_name)
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
		package_details[i].Aip_valid = data.Summary.Result == "VALID"
	}
	return package_details, nil
}

func AipValidationReportActivity(ctx context.Context, package_details []DIPPackageDetails) ([]DIPPackageDetails, error) {

	InfoLogger.Println("Starting: AIP Validation Report")

	// Remove invalid sip packages from the list
	for i, pkg := range package_details {
		if !pkg.Aip_valid {
			package_details = append(package_details[:i], package_details[i+1:]...)
		}
	}

	// If no packages are valid we terminate the process
	if len(package_details) == 0 {
		err := errors.New("No valid AIP packages")
		ErrorLogger.Println(err)
		return nil, err
	}
	return package_details, nil
}

func GenerateEarkDipActivity(ctx context.Context, package_details []DIPPackageDetails) ([]DIPPackageDetails, error) {

	InfoLogger.Println("Starting: Generate EARK DIP")

	// For each SIP package generate a new EARK AIP from it's contents.
	// The script should output the name of the resulting AIP which we set in package_details
	for i, pkg := range package_details {
		cmd := exec.Command("python3.9", "scripts/eark_aip_to_dip.py", "-i", "aips/"+pkg.Aip_name, "-o", "dips")
		op, err := cmd.Output()
		split_output := strings.Split(string(op), "\n")
		Dip_name := split_output[len(split_output)-2]
		InfoLogger.Println("New DIP Name:", Dip_name)
		package_details[i].Dip_name = Dip_name
		if err != nil {
			ErrorLogger.Println(err.Error())
			return nil, err
		}
	}
	return package_details, nil
}

func ValidateEarkDipPackagesActivity(ctx context.Context, package_details []DIPPackageDetails) ([]DIPPackageDetails, error) {

	InfoLogger.Println("Starting: Validate EARK DIP")

	// Iterate over every package running the resulting EARK AIP through the commomns IP validator
	for i, pkg := range package_details {
		cmd := exec.Command("java", "-jar", "scripts/"+CommonsIPValidatorName, "validate", "-i", "dips/"+pkg.Dip_name)
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
		package_details[i].Dip_valid = data.Summary.Result == "VALID"
	}
	return package_details, nil
}

func EarkDipValidationReportActivity(ctx context.Context, package_details []DIPPackageDetails) error {

	InfoLogger.Println("Starting: EARK DIP Validation Report")
	mapping := []AIPDIPMapping{}

	// Display in the logs every EARK AIP that fails verification
	for _, pkg := range package_details {
		if !pkg.Dip_valid {
			WarningLogger.Println("DIP Package:", pkg.Aip_name, pkg.Dip_name, "failed validation")
		}
		mapping = append(mapping, AIPDIPMapping{pkg.Aip_name, pkg.Dip_name})
	}
	buf, err := json.Marshal(mapping)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile("logs/AIPDIPMapping.json", buf, 0644)
	if err != nil {
		panic(err)
	}
	return nil
}

// Used throughout the process to track details pertaining to each package
type DIPPackageDetails struct {

	// Name of the original SIP
	Aip_name string

	// Validation status of the SIP
	Aip_valid bool

	// Name of the EARK AIP generated and placed in 'eark-aips/' from running sip_to_eark_aip.py
	Dip_name string

	// Validation status of the EARK AIP
	Dip_valid bool
}

// Used to read data returned from the Commons IP validator
/**
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
**/

type AIPDIPMapping struct {
	SIPName string
	AIPName string
}
