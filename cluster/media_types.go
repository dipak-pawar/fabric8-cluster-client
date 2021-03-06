// Code generated by goagen v1.3.0, DO NOT EDIT.
//
// API "cluster": Application Media Types
//
// Command:
// $ goagen
// --design=github.com/fabric8-services/fabric8-cluster/design
// --out=$(GOPATH)/src/github.com/fabric8-services/fabric8-cluster-client
// --pkg=cluster
// --version=v1.3.0

package cluster

import (
	"github.com/goadesign/goa"
	"net/http"
)

// Holds the response to a cluster list request (default view)
//
// Identifier: application/vnd.clusterlist+json; view=default
type ClusterList struct {
	Data []*ClusterData `form:"data" json:"data" xml:"data"`
	// An array of mixed types
	Included []interface{} `form:"included,omitempty" json:"included,omitempty" xml:"included,omitempty"`
}

// Validate validates the ClusterList media type instance.
func (mt *ClusterList) Validate() (err error) {
	if mt.Data == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "data"))
	}
	for _, e := range mt.Data {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// DecodeClusterList decodes the ClusterList instance encoded in resp body.
func (c *Client) DecodeClusterList(resp *http.Response) (*ClusterList, error) {
	var decoded ClusterList
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// Holds the response to a full cluster list request (default view)
//
// Identifier: application/vnd.fullclusterlist+json; view=default
type FullClusterList struct {
	Data []*FullClusterData `form:"data" json:"data" xml:"data"`
	// An array of mixed types
	Included []interface{} `form:"included,omitempty" json:"included,omitempty" xml:"included,omitempty"`
}

// Validate validates the FullClusterList media type instance.
func (mt *FullClusterList) Validate() (err error) {
	if mt.Data == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "data"))
	}
	for _, e := range mt.Data {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// DecodeFullClusterList decodes the FullClusterList instance encoded in resp body.
func (c *Client) DecodeFullClusterList(resp *http.Response) (*FullClusterList, error) {
	var decoded FullClusterList
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// JSONAPIErrors media type (default view)
//
// Identifier: application/vnd.jsonapierrors+json; view=default
type JSONAPIErrors struct {
	Errors []*JSONAPIError `form:"errors" json:"errors" xml:"errors"`
}

// Validate validates the JSONAPIErrors media type instance.
func (mt *JSONAPIErrors) Validate() (err error) {
	if mt.Errors == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "errors"))
	}
	for _, e := range mt.Errors {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// DecodeJSONAPIErrors decodes the JSONAPIErrors instance encoded in resp body.
func (c *Client) DecodeJSONAPIErrors(resp *http.Response) (*JSONAPIErrors, error) {
	var decoded JSONAPIErrors
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// The status of the current running instance (default view)
//
// Identifier: application/vnd.status+json; view=default
type Status struct {
	// The time when built
	BuildTime string `form:"buildTime" json:"buildTime" xml:"buildTime"`
	// Commit SHA this build is based on
	Commit string `form:"commit" json:"commit" xml:"commit"`
	// The status of the used configuration. 'OK' or an error message if there is something wrong with the configuration used by service.
	ConfigurationStatus string `form:"configurationStatus" json:"configurationStatus" xml:"configurationStatus"`
	// The status of Database connection. 'OK' or an error message is displayed.
	DatabaseStatus string `form:"databaseStatus" json:"databaseStatus" xml:"databaseStatus"`
	// 'True' if the Developer Mode is enabled
	DevMode *bool `form:"devMode,omitempty" json:"devMode,omitempty" xml:"devMode,omitempty"`
	// The time when started
	StartTime string `form:"startTime" json:"startTime" xml:"startTime"`
}

// Validate validates the Status media type instance.
func (mt *Status) Validate() (err error) {
	if mt.Commit == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "commit"))
	}
	if mt.BuildTime == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "buildTime"))
	}
	if mt.StartTime == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "startTime"))
	}
	if mt.DatabaseStatus == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "databaseStatus"))
	}
	if mt.ConfigurationStatus == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "configurationStatus"))
	}
	return
}

// DecodeStatus decodes the Status instance encoded in resp body.
func (c *Client) DecodeStatus(resp *http.Response) (*Status, error) {
	var decoded Status
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}
