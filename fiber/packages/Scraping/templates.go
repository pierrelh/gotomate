package scraping

import "github.com/lxn/walk/declarative"

// AddAllowedDomainTemplate Dialog's AddAllowedDomain Template
var AddAllowedDomainTemplate = []declarative.Widget{
	declarative.GroupBox{
		Title:  "Scraper",
		Layout: declarative.HBox{},
		Children: []declarative.Widget{
			declarative.TextEdit{
				Text:          declarative.Bind("ScraperVarName"),
				CompactHeight: true,
			},
			declarative.CheckBox{
				Text:    "Is a Var",
				Checked: true,
				Enabled: false,
			},
		},
	},
	declarative.GroupBox{
		Title:  "Path to allow",
		Layout: declarative.HBox{},
		Children: []declarative.Widget{
			declarative.TextEdit{
				Text:          declarative.Bind("PathVarName"),
				Visible:       declarative.Bind("PathIsAVar.Checked"),
				CompactHeight: true,
			},
			declarative.TextEdit{
				Text:    declarative.Bind("Path"),
				Visible: declarative.Bind("!PathIsAVar.Checked"),
			},
			declarative.CheckBox{
				Name:    "PathIsAVar",
				Text:    "Is a Var",
				Checked: declarative.Bind("PathIsVar"),
			},
		},
	},
}

// AddDisallowedDomainTemplate Dialog's AddDisallowedDomain Template
var AddDisallowedDomainTemplate = []declarative.Widget{
	declarative.GroupBox{
		Title:  "Scraper",
		Layout: declarative.HBox{},
		Children: []declarative.Widget{
			declarative.TextEdit{
				Text:          declarative.Bind("ScraperVarName"),
				CompactHeight: true,
			},
			declarative.CheckBox{
				Text:    "Is a Var",
				Checked: true,
				Enabled: false,
			},
		},
	},
	declarative.GroupBox{
		Title:  "Path to disallow",
		Layout: declarative.HBox{},
		Children: []declarative.Widget{
			declarative.TextEdit{
				Text:          declarative.Bind("PathVarName"),
				Visible:       declarative.Bind("PathIsAVar.Checked"),
				CompactHeight: true,
			},
			declarative.TextEdit{
				Text:    declarative.Bind("Path"),
				Visible: declarative.Bind("!PathIsAVar.Checked"),
			},
			declarative.CheckBox{
				Name:    "PathIsAVar",
				Text:    "Is a Var",
				Checked: declarative.Bind("PathIsVar"),
			},
		},
	},
}

// IgnoreRobotsTxtTemplate Dialog's IgnoreRobotsTxt Template
var IgnoreRobotsTxtTemplate = []declarative.Widget{
	declarative.GroupBox{
		Title:  "Scraper",
		Layout: declarative.HBox{},
		Children: []declarative.Widget{
			declarative.TextEdit{
				Text:          declarative.Bind("ScraperVarName"),
				CompactHeight: true,
			},
			declarative.CheckBox{
				Text:    "Is a Var",
				Checked: true,
				Enabled: false,
			},
		},
	},
	declarative.GroupBox{
		Title:  "Ignore robots.txt ?",
		Layout: declarative.HBox{},
		Children: []declarative.Widget{
			declarative.TextEdit{
				Text:               declarative.Bind("IgnoreVarName"),
				Visible:            declarative.Bind("IgnoreIsAVar.Checked"),
				AlwaysConsumeSpace: true,
				CompactHeight:      true,
			},
			declarative.CheckBox{
				Text:    "True ?",
				Visible: declarative.Bind("!IgnoreIsAVar.Checked"),
				Checked: declarative.Bind("Ignore"),
			},
			declarative.CheckBox{
				Name:    "IgnoreIsAVar",
				Text:    "Is a Var",
				Checked: declarative.Bind("IgnoreIsVar"),
			},
		},
	},
}

// NewScraperTemplate Dialog's NewScraper Template
var NewScraperTemplate = []declarative.Widget{
	declarative.Label{
		Text: "Output var:",
	},
	declarative.TextEdit{
		Text:          declarative.Bind("Output"),
		CompactHeight: true,
	},
}

// OnFindScrapAttributeTemplate Dialog's OnFindScrapAttribute Template
var OnFindScrapAttributeTemplate = []declarative.Widget{
	declarative.GroupBox{
		Title:  "Scraper",
		Layout: declarative.HBox{},
		Children: []declarative.Widget{
			declarative.TextEdit{
				Text:          declarative.Bind("ScraperVarName"),
				CompactHeight: true,
			},
			declarative.CheckBox{
				Text:    "Is a Var",
				Checked: true,
				Enabled: false,
			},
		},
	},
	declarative.GroupBox{
		Title:  "Element to scrap",
		Layout: declarative.HBox{},
		Children: []declarative.Widget{
			declarative.TextEdit{
				Text:          declarative.Bind("ElementVarName"),
				Visible:       declarative.Bind("ElementIsAVar.Checked"),
				CompactHeight: true,
			},
			declarative.TextEdit{
				Text:    declarative.Bind("Element"),
				Visible: declarative.Bind("!ElementIsAVar.Checked"),
			},
			declarative.CheckBox{
				Name:    "ElementIsAVar",
				Text:    "Is a Var",
				Checked: declarative.Bind("ElementIsVar"),
			},
		},
	},
	declarative.GroupBox{
		Title:  "Attribute to scrap",
		Layout: declarative.HBox{},
		Children: []declarative.Widget{
			declarative.TextEdit{
				Text:          declarative.Bind("AttributeVarName"),
				Visible:       declarative.Bind("AttributeIsAVar.Checked"),
				CompactHeight: true,
			},
			declarative.TextEdit{
				Text:    declarative.Bind("Attribute"),
				Visible: declarative.Bind("!AttributeIsAVar.Checked"),
			},
			declarative.CheckBox{
				Name:    "AttributeIsAVar",
				Text:    "Is a Var",
				Checked: declarative.Bind("AttributeIsVar"),
			},
		},
	},
	declarative.GroupBox{
		Title:  "Array to store the data",
		Layout: declarative.HBox{},
		Children: []declarative.Widget{
			declarative.TextEdit{
				Text:          declarative.Bind("TabVarName"),
				CompactHeight: true,
			},
			declarative.CheckBox{
				Text:    "Is a Var",
				Checked: true,
				Enabled: false,
			},
		},
	},
}

// OnFindScrapChildAttributeTemplate Dialog's OnFindScrapChildAttribute Template
var OnFindScrapChildAttributeTemplate = []declarative.Widget{
	declarative.GroupBox{
		Title:  "Scraper",
		Layout: declarative.HBox{},
		Children: []declarative.Widget{
			declarative.TextEdit{
				Text:          declarative.Bind("ScraperVarName"),
				CompactHeight: true,
			},
			declarative.CheckBox{
				Text:    "Is a Var",
				Checked: true,
				Enabled: false,
			},
		},
	},
	declarative.GroupBox{
		Title:  "Parent's child to scrap",
		Layout: declarative.HBox{},
		Children: []declarative.Widget{
			declarative.TextEdit{
				Text:          declarative.Bind("ElementVarName"),
				Visible:       declarative.Bind("ElementIsAVar.Checked"),
				CompactHeight: true,
			},
			declarative.TextEdit{
				Text:    declarative.Bind("Element"),
				Visible: declarative.Bind("!ElementIsAVar.Checked"),
			},
			declarative.CheckBox{
				Name:    "ElementIsAVar",
				Text:    "Is a Var",
				Checked: declarative.Bind("ElementIsVar"),
			},
		},
	},
	declarative.GroupBox{
		Title:  "Child name to scrap",
		Layout: declarative.HBox{},
		Children: []declarative.Widget{
			declarative.TextEdit{
				Text:          declarative.Bind("ChildAttributeVarName"),
				Visible:       declarative.Bind("ChildAttributeIsAVar.Checked"),
				CompactHeight: true,
			},
			declarative.TextEdit{
				Text:    declarative.Bind("ChildAttribute"),
				Visible: declarative.Bind("!ChildAttributeIsAVar.Checked"),
			},
			declarative.CheckBox{
				Name:    "ChildAttributeIsAVar",
				Text:    "Is a Var",
				Checked: declarative.Bind("ChildAttributeIsVar"),
			},
		},
	},
	declarative.GroupBox{
		Title:  "Attribute to scrap",
		Layout: declarative.HBox{},
		Children: []declarative.Widget{
			declarative.TextEdit{
				Text:          declarative.Bind("AttributeVarName"),
				Visible:       declarative.Bind("AttributeIsAVar.Checked"),
				CompactHeight: true,
			},
			declarative.TextEdit{
				Text:    declarative.Bind("Attribute"),
				Visible: declarative.Bind("!AttributeIsAVar.Checked"),
			},
			declarative.CheckBox{
				Name:    "AttributeIsAVar",
				Text:    "Is a Var",
				Checked: declarative.Bind("AttributeIsVar"),
			},
		},
	},
	declarative.GroupBox{
		Title:  "Array to store the data",
		Layout: declarative.HBox{},
		Children: []declarative.Widget{
			declarative.TextEdit{
				Text:          declarative.Bind("TabVarName"),
				CompactHeight: true,
			},
			declarative.CheckBox{
				Text:    "Is a Var",
				Checked: true,
				Enabled: false,
			},
		},
	},
}

// OnFindScrapTextTemplate Dialog's OnFindScrapText Template
var OnFindScrapTextTemplate = []declarative.Widget{
	declarative.GroupBox{
		Title:  "Scraper",
		Layout: declarative.HBox{},
		Children: []declarative.Widget{
			declarative.TextEdit{
				Text:          declarative.Bind("ScraperVarName"),
				CompactHeight: true,
			},
			declarative.CheckBox{
				Text:    "Is a Var",
				Checked: true,
				Enabled: false,
			},
		},
	},
	declarative.GroupBox{
		Title:  "Element to scrap",
		Layout: declarative.HBox{},
		Children: []declarative.Widget{
			declarative.TextEdit{
				Text:          declarative.Bind("ElementVarName"),
				Visible:       declarative.Bind("ElementIsAVar.Checked"),
				CompactHeight: true,
			},
			declarative.TextEdit{
				Text:    declarative.Bind("Element"),
				Visible: declarative.Bind("!ElementIsAVar.Checked"),
			},
			declarative.CheckBox{
				Name:    "ElementIsAVar",
				Text:    "Is a Var",
				Checked: declarative.Bind("ElementIsVar"),
			},
		},
	},
	declarative.GroupBox{
		Title:  "Array to store the data",
		Layout: declarative.HBox{},
		Children: []declarative.Widget{
			declarative.TextEdit{
				Text:          declarative.Bind("TabVarName"),
				CompactHeight: true,
			},
			declarative.CheckBox{
				Text:    "Is a Var",
				Checked: true,
				Enabled: false,
			},
		},
	},
}

// OnFindScrapChildTextTemplate Dialog's OnFindScrapChildText Template
var OnFindScrapChildTextTemplate = []declarative.Widget{
	declarative.GroupBox{
		Title:  "Scraper",
		Layout: declarative.HBox{},
		Children: []declarative.Widget{
			declarative.TextEdit{
				Text:          declarative.Bind("ScraperVarName"),
				CompactHeight: true,
			},
			declarative.CheckBox{
				Text:    "Is a Var",
				Checked: true,
				Enabled: false,
			},
		},
	},
	declarative.GroupBox{
		Title:  "Parent's child to scrap",
		Layout: declarative.HBox{},
		Children: []declarative.Widget{
			declarative.TextEdit{
				Text:          declarative.Bind("ElementVarName"),
				Visible:       declarative.Bind("ElementIsAVar.Checked"),
				CompactHeight: true,
			},
			declarative.TextEdit{
				Text:    declarative.Bind("Element"),
				Visible: declarative.Bind("!ElementIsAVar.Checked"),
			},
			declarative.CheckBox{
				Name:    "ElementIsAVar",
				Text:    "Is a Var",
				Checked: declarative.Bind("ElementIsVar"),
			},
		},
	},
	declarative.GroupBox{
		Title:  "Child name to scrap",
		Layout: declarative.HBox{},
		Children: []declarative.Widget{
			declarative.TextEdit{
				Text:          declarative.Bind("ChildAttributeVarName"),
				Visible:       declarative.Bind("ChildAttributeIsAVar.Checked"),
				CompactHeight: true,
			},
			declarative.TextEdit{
				Text:    declarative.Bind("ChildAttribute"),
				Visible: declarative.Bind("!ChildAttributeIsAVar.Checked"),
			},
			declarative.CheckBox{
				Name:    "ChildAttributeIsAVar",
				Text:    "Is a Var",
				Checked: declarative.Bind("ChildAttributeIsVar"),
			},
		},
	},
	declarative.GroupBox{
		Title:  "Array to store the data",
		Layout: declarative.HBox{},
		Children: []declarative.Widget{
			declarative.TextEdit{
				Text:          declarative.Bind("TabVarName"),
				CompactHeight: true,
			},
			declarative.CheckBox{
				Text:    "Is a Var",
				Checked: true,
				Enabled: false,
			},
		},
	},
}

// OnFindVisitTemplate Dialog's OnFindVisit Template
var OnFindVisitTemplate = []declarative.Widget{
	declarative.GroupBox{
		Title:  "Scraper",
		Layout: declarative.HBox{},
		Children: []declarative.Widget{
			declarative.TextEdit{
				Text:          declarative.Bind("ScraperVarName"),
				CompactHeight: true,
			},
			declarative.CheckBox{
				Text:    "Is a Var",
				Checked: true,
				Enabled: false,
			},
		},
	},
	declarative.GroupBox{
		Title:  "Element to visit",
		Layout: declarative.HBox{},
		Children: []declarative.Widget{
			declarative.TextEdit{
				Text:          declarative.Bind("ElementVarName"),
				Visible:       declarative.Bind("ElementIsAVar.Checked"),
				CompactHeight: true,
			},
			declarative.TextEdit{
				Text:    declarative.Bind("Element"),
				Visible: declarative.Bind("!ElementIsAVar.Checked"),
			},
			declarative.CheckBox{
				Name:    "ElementIsAVar",
				Text:    "Is a Var",
				Checked: declarative.Bind("ElementIsVar"),
			},
		},
	},
}

// OnFindChildVisitTemplate Dialog's OnFindChildVisit Template
var OnFindChildVisitTemplate = []declarative.Widget{
	declarative.GroupBox{
		Title:  "Scraper",
		Layout: declarative.HBox{},
		Children: []declarative.Widget{
			declarative.TextEdit{
				Text:          declarative.Bind("ScraperVarName"),
				CompactHeight: true,
			},
			declarative.CheckBox{
				Text:    "Is a Var",
				Checked: true,
				Enabled: false,
			},
		},
	},
	declarative.GroupBox{
		Title:  "Parent's child to visit",
		Layout: declarative.HBox{},
		Children: []declarative.Widget{
			declarative.TextEdit{
				Text:          declarative.Bind("ElementVarName"),
				Visible:       declarative.Bind("ElementIsAVar.Checked"),
				CompactHeight: true,
			},
			declarative.TextEdit{
				Text:    declarative.Bind("Element"),
				Visible: declarative.Bind("!ElementIsAVar.Checked"),
			},
			declarative.CheckBox{
				Name:    "ElementIsAVar",
				Text:    "Is a Var",
				Checked: declarative.Bind("ElementIsVar"),
			},
		},
	},
	declarative.GroupBox{
		Title:  "Child name to visit",
		Layout: declarative.HBox{},
		Children: []declarative.Widget{
			declarative.TextEdit{
				Text:          declarative.Bind("ChildAttributeVarName"),
				Visible:       declarative.Bind("ChildAttributeIsAVar.Checked"),
				CompactHeight: true,
			},
			declarative.TextEdit{
				Text:    declarative.Bind("ChildAttribute"),
				Visible: declarative.Bind("!ChildAttributeIsAVar.Checked"),
			},
			declarative.CheckBox{
				Name:    "ChildAttributeIsAVar",
				Text:    "Is a Var",
				Checked: declarative.Bind("ChildAttributeIsVar"),
			},
		},
	},
}

// ScraperEndConditionTemplate Dialog's ScraperEndCondition Template
var ScraperEndConditionTemplate = []declarative.Widget{
	declarative.GroupBox{
		Title:  "Scraper",
		Layout: declarative.HBox{},
		Children: []declarative.Widget{
			declarative.TextEdit{
				Text:          declarative.Bind("ScraperVarName"),
				CompactHeight: true,
			},
			declarative.CheckBox{
				Text:    "Is a Var",
				Checked: true,
				Enabled: false,
			},
		},
	},
	declarative.GroupBox{
		Title:  "Iterations numbers",
		Layout: declarative.HBox{},
		Children: []declarative.Widget{
			declarative.TextEdit{
				Text:          declarative.Bind("EndVarName"),
				Visible:       declarative.Bind("EndIsAVar.Checked"),
				CompactHeight: true,
			},
			declarative.NumberEdit{
				Value:    declarative.Bind("End"),
				Visible:  declarative.Bind("!EndIsAVar.Checked"),
				Decimals: 0,
			},
			declarative.CheckBox{
				Name:    "EndIsAVar",
				Text:    "Is a Var",
				Checked: declarative.Bind("EndIsVar"),
			},
		},
	},
}

// ScrapStartTemplate Dialog's ScrapStart Template
var ScrapStartTemplate = []declarative.Widget{
	declarative.GroupBox{
		Title:  "Scraper",
		Layout: declarative.HBox{},
		Children: []declarative.Widget{
			declarative.TextEdit{
				Text:          declarative.Bind("ScraperVarName"),
				CompactHeight: true,
			},
			declarative.CheckBox{
				Text:    "Is a Var",
				Checked: true,
				Enabled: false,
			},
		},
	},
	declarative.GroupBox{
		Title:  "Url to scrap",
		Layout: declarative.HBox{},
		Children: []declarative.Widget{
			declarative.TextEdit{
				Text:          declarative.Bind("UrlVarName"),
				Visible:       declarative.Bind("UrlIsAVar.Checked"),
				CompactHeight: true,
			},
			declarative.TextEdit{
				Text:    declarative.Bind("Url"),
				Visible: declarative.Bind("!UrlIsAVar.Checked"),
			},
			declarative.CheckBox{
				Name:    "UrlIsAVar",
				Text:    "Is a Var",
				Checked: declarative.Bind("UrlIsVar"),
			},
		},
	},
}
