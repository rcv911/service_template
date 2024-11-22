package pyroscope

import (
	"fmt"

	"github.com/grafana/pyroscope-go"
)

func New(appName, addr string) error {
	_, err := pyroscope.Start(pyroscope.Config{
		ApplicationName: appName,
		ServerAddress:   addr,
		ProfileTypes: []pyroscope.ProfileType{
			pyroscope.ProfileCPU,
			pyroscope.ProfileAllocObjects,
			pyroscope.ProfileAllocSpace,
			pyroscope.ProfileInuseObjects,
			pyroscope.ProfileInuseSpace,
			pyroscope.ProfileGoroutines,
			pyroscope.ProfileMutexCount,
			pyroscope.ProfileMutexDuration,
		},
	})
	if err != nil {
		return fmt.Errorf("failed start pyroscope for addr=%s: %w", addr, err)
	}

	return nil
}
