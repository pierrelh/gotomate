package scraping

import (
	"fmt"
	"reflect"
)

// Processing process the functions from screen's package
func Processing(funcName string, instructionData reflect.Value, finished chan bool) int {
	nextID := -1
	switch funcName {
	case "AddAllowedDomain":
		go func() {
			nextID = AddAllowedDomain(instructionData, finished)
		}()
		<-finished
	case "AddDisallowedDomain":
		go func() {
			nextID = AddDisallowedDomain(instructionData, finished)
		}()
		<-finished
	case "IgnoreRobotsTxt":
		go func() {
			nextID = IgnoreRobotsTxt(instructionData, finished)
		}()
		<-finished
	case "NewScraper":
		go func() {
			nextID = NewScraper(instructionData, finished)
		}()
		<-finished
	case "OnFindScrapAttribute":
		go func() {
			nextID = OnFindScrapAttribute(instructionData, finished)
		}()
		<-finished
	case "OnFindScrapChildAttribute":
		go func() {
			nextID = OnFindScrapChildAttribute(instructionData, finished)
		}()
		<-finished
	case "OnFindScrapText":
		go func() {
			nextID = OnFindScrapText(instructionData, finished)
		}()
		<-finished
	case "OnFindScrapChildText":
		go func() {
			nextID = OnFindScrapChildText(instructionData, finished)
		}()
		<-finished
	case "OnFindVisit":
		go func() {
			nextID = OnFindVisit(instructionData, finished)
		}()
		<-finished
	case "OnFindChildVisit":
		go func() {
			nextID = OnFindChildVisit(instructionData, finished)
		}()
		<-finished
	case "ScraperEndCondition":
		go func() {
			nextID = ScraperEndCondition(instructionData, finished)
		}()
		<-finished
	case "ScrapStart":
		go func() {
			nextID = ScrapStart(instructionData, finished)
		}()
		<-finished
	default:
		fmt.Println("FIBER ERROR: This function is not integrated yet: " + funcName)
	}
	return nextID
}
