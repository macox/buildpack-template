package helloworld

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/buildpacks/libcnb"
)

const (
	fileToDetect = "detect-this.yaml"
)

type Detect struct{}

func (Detect) Detect(context libcnb.DetectContext) (libcnb.DetectResult, error) {
	if _, err := os.Stat(filepath.Join(context.Application.Path, fileToDetect)); err != nil && !os.IsNotExist(err) {
		return libcnb.DetectResult{}, fmt.Errorf("unable to stat %s\n%w", fileToDetect, err)
	}

	fmt.Printf("%s found", fileToDetect)
	result := libcnb.DetectResult{
		Pass: true,
		Plans: []libcnb.BuildPlan{
			{
				Provides: []libcnb.BuildPlanProvide{
					{Name: "helloworld"},
				},
				Requires: []libcnb.BuildPlanRequire{
					{Name: "helloworld"},
				},
			},
		},
	}
	return result, nil
}
