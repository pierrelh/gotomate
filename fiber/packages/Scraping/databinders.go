package scraping

// AddAllowedDomainDatabinder Define the AddAllowedDomain databinder
type AddAllowedDomainDatabinder struct {
	ScraperVarName string
	Path           string
	PathVarName    string
	PathIsVar      bool
}

// AddDisallowedDomainDatabinder Define the AddDisallowedDomain databinder
type AddDisallowedDomainDatabinder struct {
	ScraperVarName string
	Path           string
	PathVarName    string
	PathIsVar      bool
}

// IgnoreRobotsTxtDatabinder Define the IgnoreRobotsTxt databinder
type IgnoreRobotsTxtDatabinder struct {
	ScraperVarName string
	Ignore         bool
	IgnoreVarName  string
	IgnoreIsVar    bool
}

// NewScraperDatabinder Define the NewScraper databinder
type NewScraperDatabinder struct {
	Output string
}

// OnFindScrapAttributeDatabinder Define the OnFindScrapAttribute databinder
type OnFindScrapAttributeDatabinder struct {
	ScraperVarName   string
	Element          string
	ElementVarName   string
	ElementIsVar     bool
	Attribute        string
	AttributeVarName string
	AttributeIsVar   bool
	TabVarName       string
}

// OnFindScrapChildAttributeDatabinder Define the OnFindScrapChildAttribute databinder
type OnFindScrapChildAttributeDatabinder struct {
	ScraperVarName        string
	Element               string
	ElementVarName        string
	ElementIsVar          bool
	ChildAttribute        string
	ChildAttributeVarName string
	ChildAttributeIsVar   bool
	Attribute             string
	AttributeVarName      string
	AttributeIsVar        bool
	TabVarName            string
}

// OnFindScrapTextDatabinder Define the OnFindScrapText databinder
type OnFindScrapTextDatabinder struct {
	ScraperVarName string
	Element        string
	ElementVarName string
	ElementIsVar   bool
	TabVarName     string
}

// OnFindScrapChildTextDatabinder Define the OnFindScrapChildText databinder
type OnFindScrapChildTextDatabinder struct {
	ScraperVarName        string
	Element               string
	ElementVarName        string
	ElementIsVar          bool
	ChildAttribute        string
	ChildAttributeVarName string
	ChildAttributeIsVar   bool
	TabVarName            string
}

// OnFindVisitDatabinder Define the OnFindVisit databinder
type OnFindVisitDatabinder struct {
	ScraperVarName string
	Element        string
	ElementVarName string
	ElementIsVar   bool
}

// OnFindChildVisitDatabinder Define the OnFindChildVisit databinder
type OnFindChildVisitDatabinder struct {
	ScraperVarName        string
	Element               string
	ElementVarName        string
	ElementIsVar          bool
	ChildAttribute        string
	ChildAttributeVarName string
	ChildAttributeIsVar   bool
}

// ScraperEndConditionDatabinder Define the ScraperEndCondition databinder
type ScraperEndConditionDatabinder struct {
	ScraperVarName string
	End            int64
	EndVarName     string
	EndIsVar       bool
}

// ScrapStartDatabinder Define the ScrapStart databinder
type ScrapStartDatabinder struct {
	ScraperVarName string
	Url            string
	UrlVarName     string
	UrlIsVar       bool
}
