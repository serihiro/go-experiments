package errors

import goerrors "errors"

func Is(err, target error) bool {
	return goerrors.Is(err, target)
}
