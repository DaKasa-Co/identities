package model

import (
	"os"
)

var (
	Version = os.Getenv("APIS_VERSION")
)
