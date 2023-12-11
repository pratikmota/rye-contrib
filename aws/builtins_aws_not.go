//go:build !b_aws
// +build !b_aws

package aws

import (
	"rye/env"
)

var Builtins_aws = map[string]*env.Builtin{}
