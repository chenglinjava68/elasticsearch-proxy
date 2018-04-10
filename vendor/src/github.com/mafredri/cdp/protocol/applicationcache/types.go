// Code generated by cdpgen. DO NOT EDIT.

package applicationcache

// Resource Detailed application cache resource information.
type Resource struct {
	URL  string `json:"url"`  // Resource url.
	Size int    `json:"size"` // Resource size.
	Type string `json:"type"` // Resource type.
}

// ApplicationCache Detailed application cache information.
type ApplicationCache struct {
	ManifestURL  string     `json:"manifestURL"`  // Manifest URL.
	Size         float64    `json:"size"`         // Application cache size.
	CreationTime float64    `json:"creationTime"` // Application cache creation time.
	UpdateTime   float64    `json:"updateTime"`   // Application cache update time.
	Resources    []Resource `json:"resources"`    // Application cache resources.
}
