package scraping

import (
	"gotomate/fiber/variable"
	"gotomate/log"

	"github.com/gocolly/colly"
)

// AddAllowedDomain Setting allowed Domains for scraper
func AddAllowedDomain(instructionData interface{}, finished chan bool) int {
	log.FiberInfo("Adding allowed domain to the scraper")

	scraper, err1 := variable.Keys{VarName: "ScraperVarName"}.GetValue(instructionData)
	path, err2 := variable.Keys{VarName: "PathVarName", IsVarName: "PathIsVar", Name: "Path"}.GetValue(instructionData)

	if err1 != nil || err2 != nil {
		finished <- true
		return -1
	}

	scraper.(*colly.Collector).AllowedDomains = append(scraper.(*colly.Collector).AllowedDomains, path.(string))
	variable.SetVariable(instructionData, "ScraperVarName", scraper.(*colly.Collector))
	finished <- true
	return -1
}

// AddDisallowedDomain Setting disallowed Domains for scraper
func AddDisallowedDomain(instructionData interface{}, finished chan bool) int {
	log.FiberInfo("Adding disallowed domain to the scraper")

	scraper, err1 := variable.Keys{VarName: "ScraperVarName"}.GetValue(instructionData)
	path, err2 := variable.Keys{VarName: "PathVarName", IsVarName: "PathIsVar", Name: "Path"}.GetValue(instructionData)

	if err1 != nil || err2 != nil {
		finished <- true
		return -1
	}

	scraper.(*colly.Collector).DisallowedDomains = append(scraper.(*colly.Collector).DisallowedDomains, path.(string))
	variable.SetVariable(instructionData, "ScraperVarName", scraper.(*colly.Collector))
	finished <- true
	return -1
}

// IgnoreRobotsTxt Setting ignoreRobotsTxt Domains for scraper
func IgnoreRobotsTxt(instructionData interface{}, finished chan bool) int {
	log.FiberInfo("Setting the ignore robots.txt")

	scraper, err1 := variable.Keys{VarName: "ScraperVarName"}.GetValue(instructionData)
	ignore, err2 := variable.Keys{VarName: "IgnoreVarName", IsVarName: "IgnoreIsVar", Name: "Ignore"}.GetValue(instructionData)
	if err1 != nil || err2 != nil {
		finished <- true
		return -1
	}

	scraper.(*colly.Collector).IgnoreRobotsTxt = ignore.(bool)
	variable.SetVariable(instructionData, "ScraperVarName", scraper.(*colly.Collector))
	finished <- true
	return -1
}

// NewScraper Create a new scraper
func NewScraper(instructionData interface{}, finished chan bool) int {
	log.FiberInfo("Creating a new scraper")

	variable.SetVariable(instructionData, "Output", colly.NewCollector())
	finished <- true
	return -1
}

// OnFindScrapAttribute Scrap a target's attribute when the wanted element is found
func OnFindScrapAttribute(instructionData interface{}, finished chan bool) int {
	log.FiberInfo("Setting the scraper's on find scraping attribute")

	scraper, err1 := variable.Keys{VarName: "ScraperVarName"}.GetValue(instructionData)
	element, err2 := variable.Keys{VarName: "ElementVarName", IsVarName: "ElementIsVar", Name: "Element"}.GetValue(instructionData)
	attribute, err3 := variable.Keys{VarName: "AttributeVarName", IsVarName: "AttributeIsVar", Name: "Attribute"}.GetValue(instructionData)
	tab, err4 := variable.Keys{VarName: "TabVarName"}.GetValue(instructionData)

	if err1 != nil || err2 != nil || err3 != nil || err4 != nil {
		finished <- true
		return -1
	}

	scraper.(*colly.Collector).OnHTML(element.(string), func(e *colly.HTMLElement) {
		value := e.Attr(attribute.(string))
		tab = append(tab.([]string), value)
		variable.SetVariable(instructionData, "TabVarName", tab.([]string))
	})

	variable.SetVariable(instructionData, "ScraperVarName", scraper.(*colly.Collector))
	finished <- true
	return -1
}

// OnFindScrapChildAttribute Scrap a target's children attribute when the wanted element is found
func OnFindScrapChildAttribute(instructionData interface{}, finished chan bool) int {
	log.FiberInfo("Setting the scraper's on find scraping child's attribute")

	scraper, err1 := variable.Keys{VarName: "ScraperVarName"}.GetValue(instructionData)
	element, err2 := variable.Keys{VarName: "ElementVarName", IsVarName: "ElementIsVar", Name: "Element"}.GetValue(instructionData)
	attribute, err3 := variable.Keys{VarName: "AttributeVarName", IsVarName: "AttributeIsVar", Name: "Attribute"}.GetValue(instructionData)
	childAttribute, err4 := variable.Keys{VarName: "ChildAttributeVarName", IsVarName: "ChildAttributeIsVar", Name: "ChildAttribute"}.GetValue(instructionData)
	tab, err5 := variable.Keys{VarName: "TabVarName"}.GetValue(instructionData)

	if err1 != nil || err2 != nil || err3 != nil || err4 != nil || err5 != nil {
		finished <- true
		return -1
	}

	scraper.(*colly.Collector).OnHTML(element.(string), func(e *colly.HTMLElement) {
		value := e.ChildAttr(childAttribute.(string), attribute.(string))
		tab = append(tab.([]string), value)
		variable.SetVariable(instructionData, "TabVarName", tab.([]string))
	})

	variable.SetVariable(instructionData, "ScraperVarName", scraper.(*colly.Collector))
	finished <- true
	return -1
}

// OnFindScrapText Scrap a target's text when the wanted element is found
func OnFindScrapText(instructionData interface{}, finished chan bool) int {
	log.FiberInfo("Setting the scraper's on find scraping text")

	scraper, err1 := variable.Keys{VarName: "ScraperVarName"}.GetValue(instructionData)
	element, err2 := variable.Keys{VarName: "ElementVarName", IsVarName: "ElementIsVar", Name: "Element"}.GetValue(instructionData)
	tab, err3 := variable.Keys{VarName: "TabVarName"}.GetValue(instructionData)

	if err1 != nil || err2 != nil || err3 != nil {
		finished <- true
		return -1
	}

	scraper.(*colly.Collector).OnHTML(element.(string), func(e *colly.HTMLElement) {
		value := e.Text
		tab = append(tab.([]string), value)
		variable.SetVariable(instructionData, "TabVarName", tab.([]string))
	})
	variable.SetVariable(instructionData, "ScraperVarName", scraper.(*colly.Collector))
	finished <- true
	return -1
}

// OnFindScrapChildText Scrap a target's children text when the wanted element is found
func OnFindScrapChildText(instructionData interface{}, finished chan bool) int {
	log.FiberInfo("Setting the scraper's on find scraping child's text")

	scraper, err1 := variable.Keys{VarName: "ScraperVarName"}.GetValue(instructionData)
	element, err2 := variable.Keys{VarName: "ElementVarName", IsVarName: "ElementIsVar", Name: "Element"}.GetValue(instructionData)
	childAttribute, err3 := variable.Keys{VarName: "ChildAttributeVarName", IsVarName: "ChildAttributeIsVar", Name: "ChildAttribute"}.GetValue(instructionData)
	tab, err4 := variable.Keys{VarName: "TabVarName"}.GetValue(instructionData)

	if err1 != nil || err2 != nil || err3 != nil || err4 != nil {
		finished <- true
		return -1
	}

	scraper.(*colly.Collector).OnHTML(element.(string), func(e *colly.HTMLElement) {
		value := e.ChildText(childAttribute.(string))
		tab = append(tab.([]string), value)
		variable.SetVariable(instructionData, "TabVarName", tab.([]string))
	})
	variable.SetVariable(instructionData, "ScraperVarName", scraper.(*colly.Collector))
	finished <- true
	return -1
}

// OnFindVisit Visit a target when the wanted element is found
func OnFindVisit(instructionData interface{}, finished chan bool) int {
	log.FiberInfo("Setting the scraper's on find visit")

	scraper, err1 := variable.Keys{VarName: "ScraperVarName"}.GetValue(instructionData)
	element, err2 := variable.Keys{VarName: "ElementVarName", IsVarName: "ElementIsVar", Name: "Element"}.GetValue(instructionData)

	if err1 != nil || err2 != nil {
		finished <- true
		return -1
	}

	scraper.(*colly.Collector).OnHTML(element.(string), func(e *colly.HTMLElement) {
		t := e.Attr("href")
		scraper.(*colly.Collector).Visit(t)
	})

	variable.SetVariable(instructionData, "ScraperVarName", scraper.(*colly.Collector))
	finished <- true
	return -1
}

// OnFindChildVisit Visit a target's children when the wanted element is found
func OnFindChildVisit(instructionData interface{}, finished chan bool) int {
	log.FiberInfo("Setting the scraper's on find child's visit")

	scraper, err1 := variable.Keys{VarName: "ScraperVarName"}.GetValue(instructionData)
	element, err2 := variable.Keys{VarName: "ElementVarName", IsVarName: "ElementIsVar", Name: "Element"}.GetValue(instructionData)
	childAttribute, err3 := variable.Keys{VarName: "ChildAttributeVarName", IsVarName: "ChildAttributeIsVar", Name: "ChildAttribute"}.GetValue(instructionData)

	if err1 != nil || err2 != nil || err3 != nil {
		finished <- true
		return -1
	}

	scraper.(*colly.Collector).OnHTML(element.(string), func(e *colly.HTMLElement) {
		t := e.ChildAttr(childAttribute.(string), "href")
		scraper.(*colly.Collector).Visit(t)
	})

	variable.SetVariable(instructionData, "ScraperVarName", scraper.(*colly.Collector))
	finished <- true
	return -1
}

// ScraperEndCondition Setting the scraper's end condition
func ScraperEndCondition(instructionData interface{}, finished chan bool) int {
	log.FiberInfo("Setting the scraper's end condition")

	scraper, err1 := variable.Keys{VarName: "ScraperVarName"}.GetValue(instructionData)
	end, err2 := variable.Keys{VarName: "EndVarName", IsVarName: "EndIsVar", Name: "End"}.GetValue(instructionData)
	endIsVar, err3 := variable.Keys{Name: "EndIsVar"}.GetValue(instructionData)
	endVarName, err4 := variable.Keys{Name: "EndVarName"}.GetValue(instructionData)

	if err1 != nil || err2 != nil || err3 != nil || err4 != nil {
		finished <- true
		return -1
	}

	scraper.(*colly.Collector).OnResponse(func(r *colly.Response) {

		end = end.(int) - 1

		if end.(int) <= 0 {
			panic("Scrap finished")
		}

		if endIsVar.(bool) {
			variable.SetVariable(instructionData, endVarName.(string), end.(int))
		} else {
			variable.SetVariable(instructionData, "End", end.(int64))
		}

	})

	variable.SetVariable(instructionData, "ScraperVarName", scraper.(*colly.Collector))
	finished <- true
	return -1
}

// ScrapStart Visit a url to start a scraping
func ScrapStart(instructionData interface{}, finished chan bool) int {
	log.FiberInfo("Starting the Scraper")

	scraper, err1 := variable.Keys{VarName: "ScraperVarName"}.GetValue(instructionData)
	url, err2 := variable.Keys{VarName: "UrlVarName", IsVarName: "UrlIsVar", Name: "Url"}.GetValue(instructionData)

	if err1 != nil || err2 != nil {
		finished <- true
		return -1
	}

	scraper.(*colly.Collector).Visit(url.(string))
	scraper.(*colly.Collector).Wait()
	finished <- true
	return -1
}
