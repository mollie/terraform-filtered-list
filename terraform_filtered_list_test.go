package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

// An example of how to test the simple Terraform module in examples/terraform-basic-example using Terratest.
func TestTerraformFilteredListIntersect(t *testing.T) {

	terraformOptions := &terraform.Options{
		// Variables to pass to our Terraform code using -var options
		Vars: map[string]interface{}{
			"input":     []string{"a", "b", "c", "d", "e"},
			"intersect": []string{"d", "e", "f"},
			"exclude":   []string{},
		},

		// Disable colors in Terraform commands so its easier to parse stdout/stderr
		NoColor: true,
	}

	// At the end of the test, run `terraform destroy` to clean up any resources that were created
	defer terraform.Destroy(t, terraformOptions)

	// This will run `terraform init` and `terraform apply` and fail the test if there are any errors
	terraform.InitAndApply(t, terraformOptions)

	// Run `terraform output` to get the values of output variables
	actualFilteredList := terraform.Output(t, terraformOptions, "filtered_list")

	// Verify we're getting back the outputs we expect
	assert.Equal(t, "d,\ne", actualFilteredList)
}

func TestTerraformFilteredListExclude(t *testing.T) {

	terraformOptions := &terraform.Options{
		// Variables to pass to our Terraform code using -var options
		Vars: map[string]interface{}{
			"input":     []string{"a", "b", "c", "d", "e"},
			"intersect": []string{},
			"exclude":   []string{"a", "b", "e"},
		},

		// Disable colors in Terraform commands so its easier to parse stdout/stderr
		NoColor: true,
	}

	// At the end of the test, run `terraform destroy` to clean up any resources that were created
	defer terraform.Destroy(t, terraformOptions)

	// This will run `terraform init` and `terraform apply` and fail the test if there are any errors
	terraform.InitAndApply(t, terraformOptions)

	// Run `terraform output` to get the values of output variables
	actualFilteredList := terraform.Output(t, terraformOptions, "filtered_list")

	// Verify we're getting back the outputs we expect
	assert.Equal(t, "c,\nd", actualFilteredList)
}

func TestTerraformFilteredListIntersectAndExclude(t *testing.T) {

	terraformOptions := &terraform.Options{
		// Variables to pass to our Terraform code using -var options
		Vars: map[string]interface{}{
			"input":     []string{"a", "b", "c", "d", "e"},
			"intersect": []string{"d", "e", "f"},
			"exclude":   []string{"a", "b", "e"},
		},

		// Disable colors in Terraform commands so its easier to parse stdout/stderr
		NoColor: true,
	}

	// At the end of the test, run `terraform destroy` to clean up any resources that were created
	defer terraform.Destroy(t, terraformOptions)

	// This will run `terraform init` and `terraform apply` and fail the test if there are any errors
	terraform.InitAndApply(t, terraformOptions)

	// Run `terraform output` to get the values of output variables
	actualFilteredList := terraform.Output(t, terraformOptions, "filtered_list")

	// Verify we're getting back the outputs we expect
	assert.Equal(t, "d", actualFilteredList)
}
