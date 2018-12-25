package main

import (
	"log"

	"github.com/hashicorp/terraform/helper/schema"
)

func resourceServer() *schema.Resource {
	return &schema.Resource{
		Create: resourceServerCreate,
		Read:   resourceServerRead,
		Update: resourceServerUpdate,
		Delete: resourceServerDelete,

		Schema: map[string]*schema.Schema{
			"address": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceServerCreate(d *schema.ResourceData, m interface{}) error {
	log.Printf("Dumm Provider: create")
	return resourceServerRead(d, m)
}

func resourceServerRead(d *schema.ResourceData, m interface{}) error {
	log.Printf("Dumm Provider: read")
	address := d.Get("address").(string)
	log.Println(address)
	d.SetId(address)
	return nil
}

func resourceServerUpdate(d *schema.ResourceData, m interface{}) error {
	log.Printf("Dumm Provider: update")

	return resourceServerRead(d, m)
}

func resourceServerDelete(d *schema.ResourceData, m interface{}) error {
	log.Printf("Dumm Provider: update")
	return nil
}
