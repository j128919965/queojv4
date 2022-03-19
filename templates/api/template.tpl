
syntax = "v1"

info(
	title: // TODO: add title
	desc: // TODO: add description
	author: "QueZr"
	email: "853681401@qq.com"
)

type request {
	// TODO: add members here and delete this comment
}

type response {
	// TODO: add members here and delete this comment
}

service {{.serviceName}} {
	@handler GetUser // TODO: set handler name and delete this comment
	get /users/id/:userId(request) returns(response)

	@handler CreateUser // TODO: set handler name and delete this comment
	post /users/create(request)
}
