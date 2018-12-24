package checkers

import (
	uuid "github.com/nu7hatch/gouuid"
)

func mustCreatedHandle() (handle string) {
	u4, err := uuid.NewV4()
	if err != nil {
		panic(err)
	}

	handle = u4.String()
	return
}
