package resource

import (
	"fmt"
	"os/exec"
	"strings"
)

// GetResources returns a list of terraform resources from state
func GetResources() ([]string, error) {
	cmd := exec.Command("terraform", "state", "list")
	output, err := cmd.CombinedOutput()

	// stateファイルがない場合はエラー
	if err != nil {
		return nil, fmt.Errorf("could not get resources from state")
	}

	var resources []string
	for _, line := range strings.Split(string(output), "\n") {
		if line = strings.TrimSpace(line); line != "" {
			resources = append(resources, line)
		}
	}

	return resources, nil
}

// GroupResourcesByType groups resources by their resource type
func GroupResourcesByType(resources []string) map[string][]string {
	groups := make(map[string][]string)

	for _, resource := range resources {
		parts := strings.SplitN(resource, ".", 2)
		if len(parts) < 2 {
			continue
		}

		resourceType := parts[0]
		if _, exists := groups[resourceType]; !exists {
			groups[resourceType] = []string{}
		}

		groups[resourceType] = append(groups[resourceType], resource)
	}

	return groups
}
