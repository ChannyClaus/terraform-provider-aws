package aws

import (
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/hashicorp/terraform/helper/schema"
)

func dataSourceAwsEbsDefaultKmsKey() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceAwsEbsDefaultKmsKeyRead,

		Schema: map[string]*schema.Schema{
			"key_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}
func dataSourceAwsEbsDefaultKmsKeyRead(d *schema.ResourceData, meta interface{}) error {
	conn := meta.(*AWSClient).ec2conn

	res, err := conn.GetEbsDefaultKmsKeyId(&ec2.GetEbsDefaultKmsKeyIdInput{})
	if err != nil {
		return fmt.Errorf("Error reading EBS default KMS key: %q", err)
	}

	d.SetId(time.Now().UTC().String())
	d.Set("key_id", *res.KmsKeyId)

	return nil
}
