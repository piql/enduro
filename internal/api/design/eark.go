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
		Description("Submit a new E-Ark Workflow")
		Result(EarkResult)
		Error("not_available")
		Error("not_valid")
		HTTP(func() {
			POST("/")
			Response(StatusAccepted)
			Response("not_available", StatusConflict)
			Response("not_valid", StatusBadRequest)
		})
	})
	Method("aip_gen_status", func() {
		Description("Retrieve status of current E-Ark Workflow operation.")
		Result(EarkStatusResult)
		HTTP(func() {
			GET("/")
			Response(StatusOK)
		})
	})
	Method("create_dips", func() {
		Description("Create a DIP from E-Ark AIP")
		Result(EarkDIPResult)
		Error("not_available")
		Error("not_valid")
		HTTP(func() {
			POST("/")
			Response(StatusAccepted)
			Response("not_available", StatusConflict)
			Response("not_valid", StatusBadRequest)
		})
	})
})

var EarkResult = Type("EarkResult", func() {
	Attribute("workflow_id", String)
	Attribute("run_id", String)
	Required("workflow_id", "run_id")
})

var EarkStatusResult = Type("EarkStatusResult", func() {
	Attribute("running", Boolean)
	Attribute("status", String)
	Attribute("workflow_id", String)
	Attribute("run_id", String)
	Required("running")
})

var EarkDIPResult = Type("EarkDIPResult", func() {
	Attribute("success", Boolean)
	Attribute("aip_name", String)
	Attribute("dip_name", String)
	Required("success")
})
