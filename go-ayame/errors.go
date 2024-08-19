package ayame

import (
	"errors"
	"fmt"
)

var (
	errorInvalidJSON        = errors.New("InvalidJSON")
	errorInvalidMessageType = errors.New("InvalidMessageType")
	ErrorClientDoesNotExist = fmt.Errorf("client does not exist")
)
