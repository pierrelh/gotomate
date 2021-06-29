package scraping

import (
	"fmt"

	"github.com/lxn/walk/declarative"
)

// Build Return the right databinder & the right template for a screen instruction
func Build(function string) (interface{}, []declarative.Widget) {
	switch function {
	case "AddAllowedDomain":
		return new(AddAllowedDomainDatabinder), AddAllowedDomainTemplate
	case "AddDisallowedDomain":
		return new(AddDisallowedDomainDatabinder), AddDisallowedDomainTemplate
	case "IgnoreRobotsTxt":
		return new(IgnoreRobotsTxtDatabinder), IgnoreRobotsTxtTemplate
	case "NewScraper":
		return new(NewScraperDatabinder), NewScraperTemplate
	case "OnFindScrapAttribute":
		return new(OnFindScrapAttributeDatabinder), OnFindScrapAttributeTemplate
	case "OnFindScrapChildAttribute":
		return new(OnFindScrapChildAttributeDatabinder), OnFindScrapChildAttributeTemplate
	case "OnFindScrapText":
		return new(OnFindScrapTextDatabinder), OnFindScrapTextTemplate
	case "OnFindScrapChildText":
		return new(OnFindScrapChildTextDatabinder), OnFindScrapChildTextTemplate
	case "OnFindVisit":
		return new(OnFindVisitDatabinder), OnFindVisitTemplate
	case "OnFindChildVisit":
		return new(OnFindChildVisitDatabinder), OnFindChildVisitTemplate
	case "ScraperEndCondition":
		return new(ScraperEndConditionDatabinder), ScraperEndConditionTemplate
	case "ScrapStart":
		return new(ScrapStartDatabinder), ScrapStartTemplate
	}
	fmt.Println("GOTOMATE ERROR: Unable to find the function for instruction building")
	return nil, nil
}
