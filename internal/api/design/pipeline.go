package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = Service("pipeline", func() {
	Description("The pipeline service manages Archivematica pipelines.")
	HTTP(func() {
		Path("/pipeline")
	})
	Method("list", func() {
		Description("List all known pipelines")
		Payload(func() {
			Attribute("name", String)
		})
		Result(ArrayOf(StoredPipeline))
		HTTP(func() {
			GET("/")
			Response(StatusOK)
			Params(func() {
				Param("name")
			})
		})
	})
	Method("show", func() {
		Description("Show pipeline by ID")
		Payload(func() {
			Attribute("id", String, "Identifier of pipeline to show", func() { Format(FormatUUID) })
			Required("id")
		})
		Result(StoredPipeline)
		Error("not_found", PipelineNotFound, "Pipeline not found")
		HTTP(func() {
			GET("/{id}")
			Response(StatusOK)
			Response("not_found", StatusNotFound)
		})
	})
	Method("processing", func() {
		Description("List all processing configurations of a pipeline given its ID")
		Payload(func() {
			Attribute("id", String, "Identifier of pipeline", func() { Format(FormatUUID) })
			Required("id")
		})
		Result(ArrayOf(String))
		Error("not_found", PipelineNotFound, "Pipeline not found")
		HTTP(func() {
			GET("/{id}/processing")
			Response(StatusOK)
			Response("not_found", StatusNotFound)
		})
	})
})

var Pipeline = Type("Pipeline", func() {
	Description("Pipeline describes an Archivematica pipeline.")
	Attribute("id", String, "Identifier of the pipeline", func() { Format(FormatUUID) })
	Attribute("name", String, "Name of the pipeline")
	Attribute("capacity", Int64, "Maximum concurrent transfers")
	Attribute("current", Int64, "Current transfers")
	Required("name")
})

var StoredPipeline = ResultType("application/vnd.enduro.stored-pipeline", func() {
	Description("StoredPipeline describes a pipeline retrieved by this service.")
	Reference(Pipeline)
	Attributes(func() {
		Attribute("id")
		Attribute("name")
		Attribute("capacity")
		Attribute("current")
	})
	View("default", func() {
		Attribute("id")
		Attribute("name")
		Attribute("capacity")
		Attribute("current")
	})
	Required("name")
})

var PipelineNotFound = Type("PipelineNotFound", func() {
	Description("Pipeline not found.")
	Attribute("message", String, "Message of error", func() {
		Meta("struct:error:name")
	})
	Attribute("id", String, "Identifier of missing pipeline")
	Required("message", "id")
})
