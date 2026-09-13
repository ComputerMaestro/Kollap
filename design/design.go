package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = Service("kollap", func() {
	Description("A backend service powering a Collaborative Knowledge Platform")

	Method("createWorkspace", func() {
		Payload(String, "test1")
		Result(String, "Created")

		HTTP(func() {
			GET("/workspace/{name}")
		})
	})
})
