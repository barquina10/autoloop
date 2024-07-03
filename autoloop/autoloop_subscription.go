package autoloop

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"os/exec"
)

func resourceAWSCommand() *schema.Resource {
	return &schema.Resource{
		Create: resourceAWSCommandCreate,
		Read:   resourceAWSCommandRead,
		Update: resourceAWSCommandUpdate,
		Delete: resourceAWSCommandDelete,

		Schema: map[string]*schema.Schema{
			"command": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"output": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceAWSCommandCreate(d *schema.ResourceData, m interface{}) error {
	command := d.Get("command").(string)
	output, err := runAWSCommand(command)
	if err != nil {
		return err
	}
	d.SetId(command)
	d.Set("output", output)
	return resourceAWSCommandRead(d, m)
}

func resourceAWSCommandRead(d *schema.ResourceData, m interface{}) error {
	// No read action required for running commands, just return
	return nil
}

func resourceAWSCommandUpdate(d *schema.ResourceData, m interface{}) error {
	// No update action required for running commands, just return
	return resourceAWSCommandRead(d, m)
}

func resourceAWSCommandDelete(d *schema.ResourceData, m interface{}) error {
	// No delete action required for running commands, just return
	d.SetId("")
	return nil
}

func runAWSCommand(command string) (string, error) {
	out, err := exec.Command("sh", "-c", command).Output()
	if err != nil {
		return "", err
	}
	return string(out), nil
}
