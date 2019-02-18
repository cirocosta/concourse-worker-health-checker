package healthcheck

import (
	uuid "github.com/nu7hatch/gouuid"
)

func createHandle() (string, error) {
	u4, err := uuid.NewV4()
	if err != nil {
		return "", err
	}

	return "health-check-" + u4.String(), nil
}
