package api

// A few interesting objects for default  values when returning objects in an API.
type EmptyObject struct{}

var Nothing = EmptyObject{}
var EmptyList = []EmptyObject{}
