package edge

import "encoding/xml"

// Edge struct
type Edge struct {
	XMLName  xml.Name `xml:"edge"`
	ObjectID string   `xml:"objectId,omitempty"`
	Version  string   `xml:"version,omitempty"`
	Type     string   `xml:"type,omitempty"`
	Name     string   `xml:"name"`
	FQDN     string   `xml:"fqdn"`
	Status   string   `xml:"status"`
	Tenant   string   `xml:"tenant"`
}

// PagedEdgeList struct
type PagedEdgeList struct {
	EdgePage Page `xml:"edgePage,omitempty"`
}

// Page struct
type Page struct {
	PagingInfo  PagingInfo `xml:"pagingInfo,omitempty"`
	EdgeSummary []Summary  `xml:"edgeSummary,omitempty"`
}

// PagingInfo struct
type PagingInfo struct {
	XMLName            xml.Name `xml:"pagingInfo"`
	PageSize           int      `xml:"pageSize,omitempty"`
	StartIndex         int      `xml:"startIndex"`
	TotalCount         int      `xml:"totalCount"`
	SortOrderAscending bool     `xml:"sortOrderAscending,omitempty"`
	SortBy             string   `xml:"sortBy,omitempty"`
}

// Summary struct
type Summary struct {
	XMLName        xml.Name `xml:"edgeSummary"`
	ObjectID       string   `xml:"objectId,omitempty"`
	ObjectTypeName string   `xml:"objectTypeName,omitempty"`
	Revision       string   `xml:"revision,omitempty"`
	Type           string   `xml:"type,omitempty>typeName,omitempty"`
	Name           string   `xml:"name"`
}
