package models

type SearchLineDataRequest struct {
	LanguageName string `json:"languageName"`
	SearchedText string `json:"searchedText"`
}

type Stops struct {
	Id       int64       `json:"Id"`
	Code     string      `json:"Code"`
	Name     string      `json:"Name"`
	Type     int         `json:"Type"`
	TypeName string      `json:"TypeName"`
	Routes   []Routes    `json:"Routes"`
	Distance interface{} `json:"Distance"`
}

type Routes struct {
	StopStatus               int            `json:"StopStatus"`
	StopStatusBaseName       string         `json:"StopStatusBaseName"`
	StopStatusName           string         `json:"StopStatusName"`
	StopStatusDefinition     interface{}    `json:"StopStatusDefinition"`
	StopStatusDefinitionName interface{}    `json:"StopStatusDefinitionName"`
	StopCausal               interface{}    `json:"StopCausal"`
	StopCausalName           interface{}    `json:"StopCausalName"`
	StopServices             []StopServices `json:"StopServices"`
	Accessibility            bool           `json:"Accessibility"`
	ResolutionTime           interface{}    `json:"ResolutionTime"`
	ResolutionTimeName       interface{}    `json:"ResolutionTimeName"`
	LastUpdate               string         `json:"LastUpdate"`
	UpdatedBy                interface{}    `json:"UpdatedBy"`
	Languages                []Languages    `json:"Languages"`
	Note                     interface{}    `json:"Note"`
	StatusText               interface{}    `json:"StatusText"`
	StatusDefinitionText     interface{}    `json:"StatusDefinitionText"`
	CausalText               interface{}    `json:"CausalText"`
	ResolutionTimeText       interface{}    `json:"ResolutionTimeText"`
	RouteStopId              int64          `json:"RouteStopId"`
	RouteStopOrder           int64          `json:"RouteStopOrder"`
	RouteId                  int64          `json:"RouteId"`
	RouteCode                string         `json:"RouteCode"`
	RouteName                string         `json:"RouteName"`
	RouteDescription         string         `json:"RouteDescription"`
	RouteStatus              int            `json:"RouteStatus"`
	RouteStatusName          string         `json:"RouteStatusName"`
	RouteType                interface{}    `json:"RouteType"`
	RouteTypeName            interface{}    `json:"RouteTypeName"`
	RouteSection             interface{}    `json:"RouteSection"`
	RouteSectionName         interface{}    `json:"RouteSectionName"`
	ViewRouteUrl             interface{}    `json:"ViewRouteUrl"`
	Schedule                 Schedule       `json:"Schedule"`
	Primary                  bool           `json:"Primary"`
	LineId                   int64          `json:"LineId"`
	LineName                 string         `json:"LineName"`
	LineCode                 string         `json:"LineCode"`
	LineType                 int            `json:"LineType"`
	LineTypeName             string         `json:"LineTypeName"`
}

type StopServices struct {
	Key     int    `json:"Key"`
	Name    string `json:"Name"`
	Enabled bool   `json:"Enabled"`
}

type Languages struct {
	Id                   int64       `json:"Id"`
	LanguageName         string      `json:"LanguageName"`
	Note                 interface{} `json:"Note"`
	StatusText           interface{} `json:"StatusText"`
	StatusDefinitionText interface{} `json:"StatusDefinitionText"`
	CausalText           interface{} `json:"CausalText"`
	ResolutionTimeText   interface{} `json:"ResolutionTimeText"`
	Current              bool        `json:"Current"`
}

type Schedule struct {
	ActiveDays []int `json:"ActiveDays"`
	IsOn       bool  `json:"IsOn"`
}
