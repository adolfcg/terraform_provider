// resource_server.go
package main

import (
        "io"
        "log"
        "net/http"
        "github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func newOrderResource() *schema.Resource {
        return &schema.Resource{
                Create: newOrderCreate,
                Read:   newOrderRead,
                Update: newOrderUpdate,
                Delete: newOrderDelete,

                Schema: map[string]*schema.Schema{
                        "style": &schema.Schema{
                                Type:     schema.TypeString,
                                Required: true,
                        },
                        "size": &schema.Schema{
                                Type:     schema.TypeString,
                                Required: true,
                        },
                },
        }
}

func newOrderCreate(d *schema.ResourceData, m interface{}) error {
        resp, err := http.Get("https://www.uuidtools.com/api/generate/v1/count/1")
        if err != nil {
                log.Fatal(err)
        }
        defer resp.Body.Close()
        if resp.StatusCode == http.StatusOK {
                bodyBytes, err := io.ReadAll(resp.Body)
                if err != nil {
                        log.Fatal(err)
                }
                bodyString := string(bodyBytes)
                d.SetId(bodyString)
        }

        return newOrderRead(d, m)
}

func newOrderRead(d *schema.ResourceData, m interface{}) error {
        return nil
}

func newOrderUpdate(d *schema.ResourceData, m interface{}) error {
        return newOrderRead(d, m)
}

func newOrderDelete(d *schema.ResourceData, m interface{}) error {
        d.SetId("")
        return nil
}
