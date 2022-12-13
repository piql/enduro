package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = Service("eark", func() {
	Description("The eark service manages eark workflows.")
	HTTP(func() {
		Path("/eark")
	})
	Method("gen_eark_aips", func() {
		Description("Submit a new E-Ark AIP Workflow")
		Result(EarkAIPResult)
		Error("not_available")
		Error("not_valid")
		HTTP(func() {
			POST("/gen-aip")
			Response(StatusAccepted)
			Response("not_available", StatusConflict)
			Response("not_valid", StatusBadRequest)
		})
	})
	Method("aip_gen_status", func() {
		Description("Retrieve status of current E-Ark AIP Workflow operation.")
		Result(EarkAIPStatusResult)
		HTTP(func() {
			GET("/aip-gen-status")
			Response(StatusOK)
		})
	})
	Method("gen_eark_dips", func() {
		Description("Submit a new E-Ark DIP Workflow")
		Result(EarkDIPResult)
		Error("not_available")
		Error("not_valid")
		HTTP(func() {
			POST("/gen-dip")
			Response(StatusAccepted)
			Response("not_available", StatusConflict)
			Response("not_valid", StatusBadRequest)
		})
	})
	Method("dip_gen_status", func() {
		Description("Retrieve status of current E-Ark DIP Workflow operation.")
		Result(EarkDIPStatusResult)
		HTTP(func() {
			GET("/dip-gen-status")
			Response(StatusOK)
		})
	})
})

var EarkAIPResult = Type("EarkAIPResult", func() {
	Attribute("workflow_id", String)
	Attribute("run_id", String)
	Required("workflow_id", "run_id")
})

var EarkAIPStatusResult = Type("EarkAIPStatusResult", func() {
	Attribute("running", Boolean)
	Attribute("status", String)
	Attribute("workflow_id", String)
	Attribute("run_id", String)
	Required("running")
})

var EarkDIPResult = Type("EarkDIPResult", func() {
	Attribute("workflow_id", String)
	Attribute("run_id", String)
	Required("workflow_id", "run_id")
})

var EarkDIPStatusResult = Type("EarkDIPStatusResult", func() {
	Attribute("running", Boolean)
	Attribute("status", String)
	Attribute("workflow_id", String)
	Attribute("run_id", String)
	Required("running")
})