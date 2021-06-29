package scraping

import (
	"fmt"
	"gotomate/fiber/variable"
	"reflect"

	"github.com/gocolly/colly"
)

// AddAllowedDomain Setting allowed Domains for scraper
func AddAllowedDomain(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: Starting new process ...")

	scraper, err := variable.GetValue(instructionData, "ScraperVarName")
	if err != nil {
		finished <- true
		return -1
	}

	path, err := variable.GetValue(instructionData, "PathVarName", "PathIsVar", "Path")
	if err != nil {
		finished <- true
		return -1
	}

	scraper.(*colly.Collector).AllowedDomains = append(scraper.(*colly.Collector).AllowedDomains, path.(string))
	variable.SetVariable(instructionData.FieldByName("ScraperVarName").Interface().(string), scraper.(*colly.Collector))
	finished <- true
	return -1
}

// AddDisallowedDomain Setting disallowed Domains for scraper
func AddDisallowedDomain(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: Starting new process ...")

	scraper, err := variable.GetValue(instructionData, "ScraperVarName")
	if err != nil {
		finished <- true
		return -1
	}

	path, err := variable.GetValue(instructionData, "PathVarName", "PathIsVar", "Path")
	if err != nil {
		finished <- true
		return -1
	}

	scraper.(*colly.Collector).DisallowedDomains = append(scraper.(*colly.Collector).DisallowedDomains, path.(string))
	variable.SetVariable(instructionData.FieldByName("ScraperVarName").Interface().(string), scraper.(*colly.Collector))
	finished <- true
	return -1
}

// IgnoreRobotsTxt Setting ignoreRobotsTxt Domains for scraper
func IgnoreRobotsTxt(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: Starting new process ...")

	scraper, err := variable.GetValue(instructionData, "ScraperVarName")
	if err != nil {
		finished <- true
		return -1
	}

	ignore, err := variable.GetValue(instructionData, "IgnoreVarName", "IgnoreIsVar", "Ignore")
	if err != nil {
		finished <- true
		return -1
	}

	scraper.(*colly.Collector).IgnoreRobotsTxt = ignore.(bool)
	variable.SetVariable(instructionData.FieldByName("ScraperVarName").Interface().(string), scraper.(*colly.Collector))
	finished <- true
	return -1
}

// NewScraper Create a new scraper
func NewScraper(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: Starting new process ...")

	variable.SetVariable(instructionData.FieldByName("Output").Interface().(string), colly.NewCollector())
	finished <- true
	return -1
}

// OnFindScrapAttribute Scrap a target's attribute when the wanted element if found
func OnFindScrapAttribute(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: Starting new process ...")

	scraper, err := variable.GetValue(instructionData, "ScraperVarName")
	if err != nil {
		finished <- true
		return -1
	}

	element, err := variable.GetValue(instructionData, "ElementVarName", "ElementIsVar", "Element")
	if err != nil {
		finished <- true
		return -1
	}

	attribute, err := variable.GetValue(instructionData, "AttributeVarName", "AttributeIsVar", "Attribute")
	if err != nil {
		finished <- true
		return -1
	}

	tab, err := variable.GetValue(instructionData, "TabVarName")
	if err != nil {
		finished <- true
		return -1
	}

	scraper.(*colly.Collector).OnHTML(element.(string), func(e *colly.HTMLElement) {
		value := e.Attr(attribute.(string))
		tab = append(tab.([]string), value)
		variable.SetVariable(instructionData.FieldByName("TabVarName").Interface().(string), tab.([]string))
	})

	variable.SetVariable(instructionData.FieldByName("ScraperVarName").Interface().(string), scraper.(*colly.Collector))
	finished <- true
	return -1
}

// OnFindScrapChildAttribute Scrap a target's children attribute when the wanted element if found
func OnFindScrapChildAttribute(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: Starting new process ...")

	scraper, err := variable.GetValue(instructionData, "ScraperVarName")
	if err != nil {
		finished <- true
		return -1
	}

	element, err := variable.GetValue(instructionData, "ElementVarName", "ElementIsVar", "Element")
	if err != nil {
		finished <- true
		return -1
	}

	attribute, err := variable.GetValue(instructionData, "AttributeVarName", "AttributeIsVar", "Attribute")
	if err != nil {
		finished <- true
		return -1
	}

	childAttribute, err := variable.GetValue(instructionData, "ChildAttributeVarName", "ChildAttributeIsVar", "ChildAttribute")
	if err != nil {
		finished <- true
		return -1
	}

	tab, err := variable.GetValue(instructionData, "TabVarName")
	if err != nil {
		finished <- true
		return -1
	}

	scraper.(*colly.Collector).OnHTML(element.(string), func(e *colly.HTMLElement) {
		value := e.ChildAttr(childAttribute.(string), attribute.(string))
		tab = append(tab.([]string), value)
		variable.SetVariable(instructionData.FieldByName("TabVarName").Interface().(string), tab.([]string))
	})

	variable.SetVariable(instructionData.FieldByName("ScraperVarName").Interface().(string), scraper.(*colly.Collector))
	finished <- true
	return -1
}

// OnFindScrapText Scrap a target's text when the wanted element if found
func OnFindScrapText(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: Starting new process ...")

	scraper, err := variable.GetValue(instructionData, "ScraperVarName")
	if err != nil {
		finished <- true
		return -1
	}

	element, err := variable.GetValue(instructionData, "ElementVarName", "ElementIsVar", "Element")
	if err != nil {
		finished <- true
		return -1
	}

	tab, err := variable.GetValue(instructionData, "TabVarName")
	if err != nil {
		finished <- true
		return -1
	}

	scraper.(*colly.Collector).OnHTML(element.(string), func(e *colly.HTMLElement) {
		value := e.Text
		tab = append(tab.([]string), value)
		variable.SetVariable(instructionData.FieldByName("TabVarName").Interface().(string), tab.([]string))
	})
	variable.SetVariable(instructionData.FieldByName("ScraperVarName").Interface().(string), scraper.(*colly.Collector))
	finished <- true
	return -1
}

// OnFindScrapChildText Scrap a target's children text when the wanted element if found
func OnFindScrapChildText(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: Starting new process ...")

	scraper, err := variable.GetValue(instructionData, "ScraperVarName")
	if err != nil {
		finished <- true
		return -1
	}

	element, err := variable.GetValue(instructionData, "ElementVarName", "ElementIsVar", "Element")
	if err != nil {
		finished <- true
		return -1
	}

	childAttribute, err := variable.GetValue(instructionData, "ChildAttributeVarName", "ChildAttributeIsVar", "ChildAttribute")
	if err != nil {
		finished <- true
		return -1
	}

	tab, err := variable.GetValue(instructionData, "TabVarName")
	if err != nil {
		finished <- true
		return -1
	}

	scraper.(*colly.Collector).OnHTML(element.(string), func(e *colly.HTMLElement) {
		value := e.ChildText(childAttribute.(string))
		tab = append(tab.([]string), value)
		variable.SetVariable(instructionData.FieldByName("TabVarName").Interface().(string), tab.([]string))
	})
	variable.SetVariable(instructionData.FieldByName("ScraperVarName").Interface().(string), scraper.(*colly.Collector))
	finished <- true
	return -1
}

// OnFindVisit Visit a target when the wanted element if found
func OnFindVisit(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: Starting new process ...")

	scraper, err := variable.GetValue(instructionData, "ScraperVarName")
	if err != nil {
		finished <- true
		return -1
	}

	element, err := variable.GetValue(instructionData, "ElementVarName", "ElementIsVar", "Element")
	if err != nil {
		finished <- true
		return -1
	}

	scraper.(*colly.Collector).OnHTML(element.(string), func(e *colly.HTMLElement) {
		t := e.Attr("href")
		scraper.(*colly.Collector).Visit(t)
	})

	variable.SetVariable(instructionData.FieldByName("ScraperVarName").Interface().(string), scraper.(*colly.Collector))
	finished <- true
	return -1
}

// OnFindChildVisit Visit a target's children when the wanted element if found
func OnFindChildVisit(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: Starting new process ...")

	scraper, err := variable.GetValue(instructionData, "ScraperVarName")
	if err != nil {
		finished <- true
		return -1
	}

	element, err := variable.GetValue(instructionData, "ElementVarName", "ElementIsVar", "Element")
	if err != nil {
		finished <- true
		return -1
	}

	childAttribute, err := variable.GetValue(instructionData, "ChildAttributeVarName", "ChildAttributeIsVar", "ChildAttribute")
	if err != nil {
		finished <- true
		return -1
	}

	scraper.(*colly.Collector).OnHTML(element.(string), func(e *colly.HTMLElement) {
		t := e.ChildAttr(childAttribute.(string), "href")
		scraper.(*colly.Collector).Visit(t)
	})

	variable.SetVariable(instructionData.FieldByName("ScraperVarName").Interface().(string), scraper.(*colly.Collector))
	finished <- true
	return -1
}

// ScraperEndCondition Create a new scraper
func ScraperEndCondition(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: Starting new process ...")

	scraper, err := variable.GetValue(instructionData, "ScraperVarName")
	if err != nil {
		finished <- true
		return -1
	}

	scraper.(*colly.Collector).OnResponse(func(r *colly.Response) {

		end, err := variable.GetValue(instructionData, "EndVarName", "EndIsVar", "End")
		if err != nil {
			finished <- true
			return
		}

		end = end.(int) - 1

		if end.(int) <= 0 {
			panic("Scrap finished")
		}
		if endIsVar := instructionData.FieldByName("EndIsVar").Interface().(bool); endIsVar {
			variable.SetVariable(instructionData.FieldByName("EndVarName").Interface().(string), end.(int))
		} else {
			instructionData.FieldByName("End").SetInt(end.(int64))
		}

	})

	variable.SetVariable(instructionData.FieldByName("ScraperVarName").Interface().(string), scraper.(*colly.Collector))
	finished <- true
	return -1
}

// ScrapStart Visit a url to start a scraping
func ScrapStart(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: Starting new process ...")

	scraper, err := variable.GetValue(instructionData, "ScraperVarName")
	if err != nil {
		finished <- true
		return -1
	}

	url, err := variable.GetValue(instructionData, "UrlVarName", "UrlIsVar", "Url")
	if err != nil {
		finished <- true
		return -1
	}

	scraper.(*colly.Collector).Visit(url.(string))
	scraper.(*colly.Collector).Wait()
	finished <- true
	return -1
}
