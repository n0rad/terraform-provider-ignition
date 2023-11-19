package ignition

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceLuks() *schema.Resource {
	return &schema.Resource{
		Exists: resourceLuksExists,
		Read:   resourceLuksRead,
		Schema: map[string]*schema.Schema{
			"clevis": {
				Type:     schema.TypeList,
				Optional: true,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"tang": {
							Type:     schema.TypeList,
							Optional: true,
							ForceNew: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"url": {
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
									},
									"thumbprint": {
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
									},
									"advertisement": {
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
									},
								},
							},
						},
						"tpm2": {
							Type:     schema.TypeBool,
							ForceNew: true,
							Optional: true,
						},
						"threshold": {
							Type:     schema.TypeInt,
							ForceNew: true,
							Optional: true,
						},
						"custom": {
							Type:     schema.TypeList,
							ForceNew: true,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"pin": {
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
									},
									"config": {
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
									},
									"needs_network": {
										Type:     schema.TypeBool,
										Optional: true,
										ForceNew: true,
									},
								},
							},
						},
					},
				},
			},
			"device": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"key_file": {
				Type:     schema.TypeList,
				Optional: true,
				ForceNew: true,
				Elem:     configReferenceResource,
			},
			"label": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"open_options": {
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				ForceNew: true,
			},
			"options": {
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				ForceNew: true,
			},
			"uuid": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"wipe_volume": {
				Type:     schema.TypeBool,
				Optional: true,
				ForceNew: true,
			},
			"rendered": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceLuksRead(d *schema.ResourceData, meta interface{}) error {
	id, err := buildLuks(d)
	if err != nil {
		return err
	}

	d.SetId(id)
	return nil
}

func resourceLuksExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	id, err := buildLuks(d)
	if err != nil {
		return false, err
	}

	return id == d.Id(), nil
}

func buildLuks(d *schema.ResourceData) (string, error) {
	//TODO
	//return hash(string(b)), nil
	return "", nil
}
