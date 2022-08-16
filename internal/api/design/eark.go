package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = Service("eark", func() {
	Description("The eark service manages earkes of collections.")
	HTTP(func() {
		Path("/eark")
	})
	Method("submit", func() {
		Description("Submit a new eark")
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
	Method("status", func() {
		Description("Retrieve status of current eark operation.")
		Result(EarkStatusResult)
		HTTP(func() {
			GET("/")
			Response(StatusOK)
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
