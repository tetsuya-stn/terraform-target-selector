package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/tetsuya-stn/terraform-target-selector/internal/resource"
)

func main() {
	// コマンドラインフラグの設定
	var (
		tfCmd   string
		varFile string
	)

	flag.StringVar(&tfCmd, "cmd", "plan", "Terraform command to run (plan, apply, etc.)")
	flag.StringVar(&varFile, "var-file", "", "Path to terraform.tfvars file")
	flag.Parse()

	os.Exit(Run(os.Args))
}

func Run(args []string) int {
	resources, err := resource.GetResources()
	if err != nil {
		fmt.Printf("Error fetching resources: %v\n", err)
		return 1
	}

	if len(resources) == 0 {
		fmt.Println("No resources found. Make sure you are in a Terraform project directory.")
		return 1
	}

	return 0
}
