// Code generated by goagen v1.3.0, DO NOT EDIT.
//
// API "cluster": Application User Types
//
// Command:
// $ goagen
// --design=github.com/fabric8-services/fabric8-cluster/design
// --out=$(GOPATH)/src/github.com/fabric8-services/fabric8-cluster-client
// --pkg=cluster
// --version=v1.3.0

package cluster

import "github.com/goadesign/goa"

// clusterData user type.
type clusterData struct {
	// API URL
	APIURL *string `form:"api-url,omitempty" json:"api-url,omitempty" xml:"api-url,omitempty"`
	// User application domain name in the cluster
	AppDNS *string `form:"app-dns,omitempty" json:"app-dns,omitempty" xml:"app-dns,omitempty"`
	// Cluster is full if set to 'true'
	CapacityExhausted *bool `form:"capacity-exhausted,omitempty" json:"capacity-exhausted,omitempty" xml:"capacity-exhausted,omitempty"`
	// Web console URL
	ConsoleURL *string `form:"console-url,omitempty" json:"console-url,omitempty" xml:"console-url,omitempty"`
	// Logging URL
	LoggingURL *string `form:"logging-url,omitempty" json:"logging-url,omitempty" xml:"logging-url,omitempty"`
	// Metrics URL
	MetricsURL *string `form:"metrics-url,omitempty" json:"metrics-url,omitempty" xml:"metrics-url,omitempty"`
	// Cluster name
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
}

// Validate validates the clusterData type instance.
func (ut *clusterData) Validate() (err error) {
	if ut.Name == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "name"))
	}
	if ut.ConsoleURL == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "console-url"))
	}
	if ut.MetricsURL == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "metrics-url"))
	}
	if ut.APIURL == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "api-url"))
	}
	if ut.LoggingURL == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "logging-url"))
	}
	if ut.AppDNS == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "app-dns"))
	}
	if ut.CapacityExhausted == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "capacity-exhausted"))
	}
	return
}

// Publicize creates ClusterData from clusterData
func (ut *clusterData) Publicize() *ClusterData {
	var pub ClusterData
	if ut.APIURL != nil {
		pub.APIURL = *ut.APIURL
	}
	if ut.AppDNS != nil {
		pub.AppDNS = *ut.AppDNS
	}
	if ut.CapacityExhausted != nil {
		pub.CapacityExhausted = *ut.CapacityExhausted
	}
	if ut.ConsoleURL != nil {
		pub.ConsoleURL = *ut.ConsoleURL
	}
	if ut.LoggingURL != nil {
		pub.LoggingURL = *ut.LoggingURL
	}
	if ut.MetricsURL != nil {
		pub.MetricsURL = *ut.MetricsURL
	}
	if ut.Name != nil {
		pub.Name = *ut.Name
	}
	return &pub
}

// ClusterData user type.
type ClusterData struct {
	// API URL
	APIURL string `form:"api-url" json:"api-url" xml:"api-url"`
	// User application domain name in the cluster
	AppDNS string `form:"app-dns" json:"app-dns" xml:"app-dns"`
	// Cluster is full if set to 'true'
	CapacityExhausted bool `form:"capacity-exhausted" json:"capacity-exhausted" xml:"capacity-exhausted"`
	// Web console URL
	ConsoleURL string `form:"console-url" json:"console-url" xml:"console-url"`
	// Logging URL
	LoggingURL string `form:"logging-url" json:"logging-url" xml:"logging-url"`
	// Metrics URL
	MetricsURL string `form:"metrics-url" json:"metrics-url" xml:"metrics-url"`
	// Cluster name
	Name string `form:"name" json:"name" xml:"name"`
}

// Validate validates the ClusterData type instance.
func (ut *ClusterData) Validate() (err error) {
	if ut.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "name"))
	}
	if ut.ConsoleURL == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "console-url"))
	}
	if ut.MetricsURL == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "metrics-url"))
	}
	if ut.APIURL == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "api-url"))
	}
	if ut.LoggingURL == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "logging-url"))
	}
	if ut.AppDNS == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "app-dns"))
	}

	return
}

// fullClusterData user type.
type fullClusterData struct {
	// API URL
	APIURL *string `form:"api-url,omitempty" json:"api-url,omitempty" xml:"api-url,omitempty"`
	// User application domain name in the cluster
	AppDNS *string `form:"app-dns,omitempty" json:"app-dns,omitempty" xml:"app-dns,omitempty"`
	// OAuth client default scope
	AuthClientDefaultScope *string `form:"auth-client-default-scope,omitempty" json:"auth-client-default-scope,omitempty" xml:"auth-client-default-scope,omitempty"`
	// OAuth client ID
	AuthClientID *string `form:"auth-client-id,omitempty" json:"auth-client-id,omitempty" xml:"auth-client-id,omitempty"`
	// OAuth client secret
	AuthClientSecret *string `form:"auth-client-secret,omitempty" json:"auth-client-secret,omitempty" xml:"auth-client-secret,omitempty"`
	// Cluster is full if set to 'true'
	CapacityExhausted *bool `form:"capacity-exhausted,omitempty" json:"capacity-exhausted,omitempty" xml:"capacity-exhausted,omitempty"`
	// Web console URL
	ConsoleURL *string `form:"console-url,omitempty" json:"console-url,omitempty" xml:"console-url,omitempty"`
	// Logging URL
	LoggingURL *string `form:"logging-url,omitempty" json:"logging-url,omitempty" xml:"logging-url,omitempty"`
	// Metrics URL
	MetricsURL *string `form:"metrics-url,omitempty" json:"metrics-url,omitempty" xml:"metrics-url,omitempty"`
	// Cluster name
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// Decrypted cluster wide token
	ServiceAccountToken *string `form:"service-account-token,omitempty" json:"service-account-token,omitempty" xml:"service-account-token,omitempty"`
	// Username of the cluster wide user
	ServiceAccountUsername *string `form:"service-account-username,omitempty" json:"service-account-username,omitempty" xml:"service-account-username,omitempty"`
	// Token provider ID
	TokenProviderID *string `form:"token-provider-id,omitempty" json:"token-provider-id,omitempty" xml:"token-provider-id,omitempty"`
}

// Validate validates the fullClusterData type instance.
func (ut *fullClusterData) Validate() (err error) {
	if ut.Name == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "name"))
	}
	if ut.ConsoleURL == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "console-url"))
	}
	if ut.MetricsURL == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "metrics-url"))
	}
	if ut.APIURL == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "api-url"))
	}
	if ut.LoggingURL == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "logging-url"))
	}
	if ut.AppDNS == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "app-dns"))
	}
	if ut.CapacityExhausted == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "capacity-exhausted"))
	}
	if ut.ServiceAccountToken == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "service-account-token"))
	}
	if ut.ServiceAccountUsername == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "service-account-username"))
	}
	if ut.TokenProviderID == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "token-provider-id"))
	}
	if ut.AuthClientID == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "auth-client-id"))
	}
	if ut.AuthClientSecret == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "auth-client-secret"))
	}
	if ut.AuthClientDefaultScope == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "auth-client-default-scope"))
	}
	return
}

// Publicize creates FullClusterData from fullClusterData
func (ut *fullClusterData) Publicize() *FullClusterData {
	var pub FullClusterData
	if ut.APIURL != nil {
		pub.APIURL = *ut.APIURL
	}
	if ut.AppDNS != nil {
		pub.AppDNS = *ut.AppDNS
	}
	if ut.AuthClientDefaultScope != nil {
		pub.AuthClientDefaultScope = *ut.AuthClientDefaultScope
	}
	if ut.AuthClientID != nil {
		pub.AuthClientID = *ut.AuthClientID
	}
	if ut.AuthClientSecret != nil {
		pub.AuthClientSecret = *ut.AuthClientSecret
	}
	if ut.CapacityExhausted != nil {
		pub.CapacityExhausted = *ut.CapacityExhausted
	}
	if ut.ConsoleURL != nil {
		pub.ConsoleURL = *ut.ConsoleURL
	}
	if ut.LoggingURL != nil {
		pub.LoggingURL = *ut.LoggingURL
	}
	if ut.MetricsURL != nil {
		pub.MetricsURL = *ut.MetricsURL
	}
	if ut.Name != nil {
		pub.Name = *ut.Name
	}
	if ut.ServiceAccountToken != nil {
		pub.ServiceAccountToken = *ut.ServiceAccountToken
	}
	if ut.ServiceAccountUsername != nil {
		pub.ServiceAccountUsername = *ut.ServiceAccountUsername
	}
	if ut.TokenProviderID != nil {
		pub.TokenProviderID = *ut.TokenProviderID
	}
	return &pub
}

// FullClusterData user type.
type FullClusterData struct {
	// API URL
	APIURL string `form:"api-url" json:"api-url" xml:"api-url"`
	// User application domain name in the cluster
	AppDNS string `form:"app-dns" json:"app-dns" xml:"app-dns"`
	// OAuth client default scope
	AuthClientDefaultScope string `form:"auth-client-default-scope" json:"auth-client-default-scope" xml:"auth-client-default-scope"`
	// OAuth client ID
	AuthClientID string `form:"auth-client-id" json:"auth-client-id" xml:"auth-client-id"`
	// OAuth client secret
	AuthClientSecret string `form:"auth-client-secret" json:"auth-client-secret" xml:"auth-client-secret"`
	// Cluster is full if set to 'true'
	CapacityExhausted bool `form:"capacity-exhausted" json:"capacity-exhausted" xml:"capacity-exhausted"`
	// Web console URL
	ConsoleURL string `form:"console-url" json:"console-url" xml:"console-url"`
	// Logging URL
	LoggingURL string `form:"logging-url" json:"logging-url" xml:"logging-url"`
	// Metrics URL
	MetricsURL string `form:"metrics-url" json:"metrics-url" xml:"metrics-url"`
	// Cluster name
	Name string `form:"name" json:"name" xml:"name"`
	// Decrypted cluster wide token
	ServiceAccountToken string `form:"service-account-token" json:"service-account-token" xml:"service-account-token"`
	// Username of the cluster wide user
	ServiceAccountUsername string `form:"service-account-username" json:"service-account-username" xml:"service-account-username"`
	// Token provider ID
	TokenProviderID string `form:"token-provider-id" json:"token-provider-id" xml:"token-provider-id"`
}

// Validate validates the FullClusterData type instance.
func (ut *FullClusterData) Validate() (err error) {
	if ut.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "name"))
	}
	if ut.ConsoleURL == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "console-url"))
	}
	if ut.MetricsURL == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "metrics-url"))
	}
	if ut.APIURL == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "api-url"))
	}
	if ut.LoggingURL == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "logging-url"))
	}
	if ut.AppDNS == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "app-dns"))
	}

	if ut.ServiceAccountToken == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "service-account-token"))
	}
	if ut.ServiceAccountUsername == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "service-account-username"))
	}
	if ut.TokenProviderID == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "token-provider-id"))
	}
	if ut.AuthClientID == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "auth-client-id"))
	}
	if ut.AuthClientSecret == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "auth-client-secret"))
	}
	if ut.AuthClientDefaultScope == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "auth-client-default-scope"))
	}
	return
}

// genericData user type.
type genericData struct {
	// UUID of the object
	ID    *string       `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	Links *genericLinks `form:"links,omitempty" json:"links,omitempty" xml:"links,omitempty"`
	Type  *string       `form:"type,omitempty" json:"type,omitempty" xml:"type,omitempty"`
}

// Publicize creates GenericData from genericData
func (ut *genericData) Publicize() *GenericData {
	var pub GenericData
	if ut.ID != nil {
		pub.ID = ut.ID
	}
	if ut.Links != nil {
		pub.Links = ut.Links.Publicize()
	}
	if ut.Type != nil {
		pub.Type = ut.Type
	}
	return &pub
}

// GenericData user type.
type GenericData struct {
	// UUID of the object
	ID    *string       `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	Links *GenericLinks `form:"links,omitempty" json:"links,omitempty" xml:"links,omitempty"`
	Type  *string       `form:"type,omitempty" json:"type,omitempty" xml:"type,omitempty"`
}

// genericLinks user type.
type genericLinks struct {
	Meta    map[string]interface{} `form:"meta,omitempty" json:"meta,omitempty" xml:"meta,omitempty"`
	Related *string                `form:"related,omitempty" json:"related,omitempty" xml:"related,omitempty"`
	Self    *string                `form:"self,omitempty" json:"self,omitempty" xml:"self,omitempty"`
}

// Publicize creates GenericLinks from genericLinks
func (ut *genericLinks) Publicize() *GenericLinks {
	var pub GenericLinks
	if ut.Meta != nil {
		pub.Meta = ut.Meta
	}
	if ut.Related != nil {
		pub.Related = ut.Related
	}
	if ut.Self != nil {
		pub.Self = ut.Self
	}
	return &pub
}

// GenericLinks user type.
type GenericLinks struct {
	Meta    map[string]interface{} `form:"meta,omitempty" json:"meta,omitempty" xml:"meta,omitempty"`
	Related *string                `form:"related,omitempty" json:"related,omitempty" xml:"related,omitempty"`
	Self    *string                `form:"self,omitempty" json:"self,omitempty" xml:"self,omitempty"`
}

// Error objects provide additional information about problems encountered while
// performing an operation. Error objects MUST be returned as an array keyed by errors in the
// top level of a JSON API document.
//
// See. also http://jsonapi.org/format/#error-objects.
type jSONAPIError struct {
	// an application-specific error code, expressed as a string value.
	Code *string `form:"code,omitempty" json:"code,omitempty" xml:"code,omitempty"`
	// a human-readable explanation specific to this occurrence of the problem.
	// Like title, this field’s value can be localized.
	Detail *string `form:"detail,omitempty" json:"detail,omitempty" xml:"detail,omitempty"`
	// a unique identifier for this particular occurrence of the problem.
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// a links object containing the following members:
	// * about: a link that leads to further details about this particular occurrence of the problem.
	Links map[string]*jSONAPILink `form:"links,omitempty" json:"links,omitempty" xml:"links,omitempty"`
	// a meta object containing non-standard meta-information about the error
	Meta map[string]interface{} `form:"meta,omitempty" json:"meta,omitempty" xml:"meta,omitempty"`
	// an object containing references to the source of the error,
	// optionally including any of the following members
	//
	// * pointer: a JSON Pointer [RFC6901] to the associated entity in the request document [e.g. "/data" for a primary data object,
	//            or "/data/attributes/title" for a specific attribute].
	// * parameter: a string indicating which URI query parameter caused the error.
	Source map[string]interface{} `form:"source,omitempty" json:"source,omitempty" xml:"source,omitempty"`
	// the HTTP status code applicable to this problem, expressed as a string value.
	Status *string `form:"status,omitempty" json:"status,omitempty" xml:"status,omitempty"`
	// a short, human-readable summary of the problem that SHOULD NOT
	// change from occurrence to occurrence of the problem, except for purposes of localization.
	Title *string `form:"title,omitempty" json:"title,omitempty" xml:"title,omitempty"`
}

// Validate validates the jSONAPIError type instance.
func (ut *jSONAPIError) Validate() (err error) {
	if ut.Detail == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "detail"))
	}
	return
}

// Publicize creates JSONAPIError from jSONAPIError
func (ut *jSONAPIError) Publicize() *JSONAPIError {
	var pub JSONAPIError
	if ut.Code != nil {
		pub.Code = ut.Code
	}
	if ut.Detail != nil {
		pub.Detail = *ut.Detail
	}
	if ut.ID != nil {
		pub.ID = ut.ID
	}
	if ut.Links != nil {
		pub.Links = make(map[string]*JSONAPILink, len(ut.Links))
		for k2, v2 := range ut.Links {
			pubk2 := k2
			var pubv2 *JSONAPILink
			if v2 != nil {
				pubv2 = v2.Publicize()
			}
			pub.Links[pubk2] = pubv2
		}
	}
	if ut.Meta != nil {
		pub.Meta = ut.Meta
	}
	if ut.Source != nil {
		pub.Source = ut.Source
	}
	if ut.Status != nil {
		pub.Status = ut.Status
	}
	if ut.Title != nil {
		pub.Title = ut.Title
	}
	return &pub
}

// Error objects provide additional information about problems encountered while
// performing an operation. Error objects MUST be returned as an array keyed by errors in the
// top level of a JSON API document.
//
// See. also http://jsonapi.org/format/#error-objects.
type JSONAPIError struct {
	// an application-specific error code, expressed as a string value.
	Code *string `form:"code,omitempty" json:"code,omitempty" xml:"code,omitempty"`
	// a human-readable explanation specific to this occurrence of the problem.
	// Like title, this field’s value can be localized.
	Detail string `form:"detail" json:"detail" xml:"detail"`
	// a unique identifier for this particular occurrence of the problem.
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// a links object containing the following members:
	// * about: a link that leads to further details about this particular occurrence of the problem.
	Links map[string]*JSONAPILink `form:"links,omitempty" json:"links,omitempty" xml:"links,omitempty"`
	// a meta object containing non-standard meta-information about the error
	Meta map[string]interface{} `form:"meta,omitempty" json:"meta,omitempty" xml:"meta,omitempty"`
	// an object containing references to the source of the error,
	// optionally including any of the following members
	//
	// * pointer: a JSON Pointer [RFC6901] to the associated entity in the request document [e.g. "/data" for a primary data object,
	//            or "/data/attributes/title" for a specific attribute].
	// * parameter: a string indicating which URI query parameter caused the error.
	Source map[string]interface{} `form:"source,omitempty" json:"source,omitempty" xml:"source,omitempty"`
	// the HTTP status code applicable to this problem, expressed as a string value.
	Status *string `form:"status,omitempty" json:"status,omitempty" xml:"status,omitempty"`
	// a short, human-readable summary of the problem that SHOULD NOT
	// change from occurrence to occurrence of the problem, except for purposes of localization.
	Title *string `form:"title,omitempty" json:"title,omitempty" xml:"title,omitempty"`
}

// Validate validates the JSONAPIError type instance.
func (ut *JSONAPIError) Validate() (err error) {
	if ut.Detail == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "detail"))
	}
	return
}

// See also http://jsonapi.org/format/#document-links.
type jSONAPILink struct {
	// a string containing the link's URL.
	Href *string `form:"href,omitempty" json:"href,omitempty" xml:"href,omitempty"`
	// a meta object containing non-standard meta-information about the link.
	Meta map[string]interface{} `form:"meta,omitempty" json:"meta,omitempty" xml:"meta,omitempty"`
}

// Publicize creates JSONAPILink from jSONAPILink
func (ut *jSONAPILink) Publicize() *JSONAPILink {
	var pub JSONAPILink
	if ut.Href != nil {
		pub.Href = ut.Href
	}
	if ut.Meta != nil {
		pub.Meta = ut.Meta
	}
	return &pub
}

// See also http://jsonapi.org/format/#document-links.
type JSONAPILink struct {
	// a string containing the link's URL.
	Href *string `form:"href,omitempty" json:"href,omitempty" xml:"href,omitempty"`
	// a meta object containing non-standard meta-information about the link.
	Meta map[string]interface{} `form:"meta,omitempty" json:"meta,omitempty" xml:"meta,omitempty"`
}

// relationGeneric user type.
type relationGeneric struct {
	Data  *genericData           `form:"data,omitempty" json:"data,omitempty" xml:"data,omitempty"`
	Links *genericLinks          `form:"links,omitempty" json:"links,omitempty" xml:"links,omitempty"`
	Meta  map[string]interface{} `form:"meta,omitempty" json:"meta,omitempty" xml:"meta,omitempty"`
}

// Publicize creates RelationGeneric from relationGeneric
func (ut *relationGeneric) Publicize() *RelationGeneric {
	var pub RelationGeneric
	if ut.Data != nil {
		pub.Data = ut.Data.Publicize()
	}
	if ut.Links != nil {
		pub.Links = ut.Links.Publicize()
	}
	if ut.Meta != nil {
		pub.Meta = ut.Meta
	}
	return &pub
}

// RelationGeneric user type.
type RelationGeneric struct {
	Data  *GenericData           `form:"data,omitempty" json:"data,omitempty" xml:"data,omitempty"`
	Links *GenericLinks          `form:"links,omitempty" json:"links,omitempty" xml:"links,omitempty"`
	Meta  map[string]interface{} `form:"meta,omitempty" json:"meta,omitempty" xml:"meta,omitempty"`
}

// relationGenericList user type.
type relationGenericList struct {
	Data  []*genericData         `form:"data,omitempty" json:"data,omitempty" xml:"data,omitempty"`
	Links *genericLinks          `form:"links,omitempty" json:"links,omitempty" xml:"links,omitempty"`
	Meta  map[string]interface{} `form:"meta,omitempty" json:"meta,omitempty" xml:"meta,omitempty"`
}

// Publicize creates RelationGenericList from relationGenericList
func (ut *relationGenericList) Publicize() *RelationGenericList {
	var pub RelationGenericList
	if ut.Data != nil {
		pub.Data = make([]*GenericData, len(ut.Data))
		for i2, elem2 := range ut.Data {
			pub.Data[i2] = elem2.Publicize()
		}
	}
	if ut.Links != nil {
		pub.Links = ut.Links.Publicize()
	}
	if ut.Meta != nil {
		pub.Meta = ut.Meta
	}
	return &pub
}

// RelationGenericList user type.
type RelationGenericList struct {
	Data  []*GenericData         `form:"data,omitempty" json:"data,omitempty" xml:"data,omitempty"`
	Links *GenericLinks          `form:"links,omitempty" json:"links,omitempty" xml:"links,omitempty"`
	Meta  map[string]interface{} `form:"meta,omitempty" json:"meta,omitempty" xml:"meta,omitempty"`
}
