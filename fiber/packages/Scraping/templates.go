package scraping

import (
	"gotomate-astilectron/fiber/template"
)

// AddAllowedDomainTemplate Dialog's AddAllowedDomain Template
var AddAllowedDomainTemplate = &template.InstructionTemplate{
	template.Field{
		Label: template.Label{
			Text: "Scraper:",
		},
		Input: template.TextInput{
			BindVariable: "ScraperVarName",
		},
		VariableToggler: template.VariableToggler{
			Checked:  true,
			Disabled: true,
		},
	}.Build(),
	template.Field{
		Label: template.Label{
			Text: "Path to allow:",
		},
		Input: template.UrlInput{
			Bind:         "Path",
			BindVariable: "PathVarName",
		},
		VariableToggler: template.VariableToggler{
			Bind: "PathIsVar",
		},
	}.Build(),
}

// AddDisallowedDomainTemplate Dialog's AddDisallowedDomain Template
var AddDisallowedDomainTemplate = &template.InstructionTemplate{
	template.Field{
		Label: template.Label{
			Text: "Scraper:",
		},
		Input: template.TextInput{
			BindVariable: "ScraperVarName",
		},
		VariableToggler: template.VariableToggler{
			Checked:  true,
			Disabled: true,
		},
	}.Build(),
	template.Field{
		Label: template.Label{
			Text: "Path to disallow:",
		},
		Input: template.UrlInput{
			Bind:         "Path",
			BindVariable: "PathVarName",
		},
		VariableToggler: template.VariableToggler{
			Bind: "PathIsVar",
		},
	}.Build(),
}

// IgnoreRobotsTxtTemplate Dialog's IgnoreRobotsTxt Template
var IgnoreRobotsTxtTemplate = &template.InstructionTemplate{
	template.Field{
		Label: template.Label{
			Text: "Scraper:",
		},
		Input: template.TextInput{
			BindVariable: "ScraperVarName",
		},
		VariableToggler: template.VariableToggler{
			Checked:  true,
			Disabled: true,
		},
	}.Build(),
	template.Field{
		Label: template.Label{
			Text: "Ignore robots.txt ?:",
		},
		Input: template.CheckboxInput{
			Bind:         "Ignore",
			BindVariable: "IgnoreVarName",
		},
		VariableToggler: template.VariableToggler{
			Bind: "IgnoreIsVar",
		},
	}.Build(),
}

// NewScraperTemplate Dialog's NewScraper Template
var NewScraperTemplate = &template.InstructionTemplate{
	template.Field{
		Label: template.Label{
			Text: "Output:",
		},
		Input: template.TextInput{
			Bind: "Output",
		},
	}.Build(),
}

// OnFindScrapAttributeTemplate Dialog's OnFindScrapAttribute Template
var OnFindScrapAttributeTemplate = &template.InstructionTemplate{
	template.Field{
		Label: template.Label{
			Text: "Scraper:",
		},
		Input: template.TextInput{
			BindVariable: "ScraperVarName",
		},
		VariableToggler: template.VariableToggler{
			Checked:  true,
			Disabled: true,
		},
	}.Build(),
	template.Field{
		Label: template.Label{
			Text: "Element to scrap:",
		},
		Input: template.TextInput{
			Bind:         "Element",
			BindVariable: "ElementVarName",
		},
		VariableToggler: template.VariableToggler{
			Bind: "ElementIsVar",
		},
	}.Build(),
	template.Field{
		Label: template.Label{
			Text: "Attribute to scrap:",
		},
		Input: template.TextInput{
			Bind:         "Attribute",
			BindVariable: "AttributeVarName",
		},
		VariableToggler: template.VariableToggler{
			Bind: "AttributeIsVar",
		},
	}.Build(),
	template.Field{
		Label: template.Label{
			Text: "Array to store the data:",
		},
		Input: template.TextInput{
			BindVariable: "TabVarName",
		},
		VariableToggler: template.VariableToggler{
			Checked:  true,
			Disabled: true,
		},
	}.Build(),
}

// OnFindScrapChildAttributeTemplate Dialog's OnFindScrapChildAttribute Template
var OnFindScrapChildAttributeTemplate = &template.InstructionTemplate{
	template.Field{
		Label: template.Label{
			Text: "Scraper:",
		},
		Input: template.TextInput{
			BindVariable: "ScraperVarName",
		},
		VariableToggler: template.VariableToggler{
			Checked:  true,
			Disabled: true,
		},
	}.Build(),
	template.Field{
		Label: template.Label{
			Text: "Parent's child to scrap:",
		},
		Input: template.TextInput{
			Bind:         "Element",
			BindVariable: "ElementVarName",
		},
		VariableToggler: template.VariableToggler{
			Bind: "ElementIsVar",
		},
	}.Build(),
	template.Field{
		Label: template.Label{
			Text: "Child name to scrap:",
		},
		Input: template.TextInput{
			Bind:         "ChildAttribute",
			BindVariable: "ChildAttributeVarName",
		},
		VariableToggler: template.VariableToggler{
			Bind: "ChildAttributeIsVar",
		},
	}.Build(),
	template.Field{
		Label: template.Label{
			Text: "Attribute to scrap:",
		},
		Input: template.TextInput{
			Bind:         "Attribute",
			BindVariable: "AttributeVarName",
		},
		VariableToggler: template.VariableToggler{
			Bind: "AttributeIsVar",
		},
	}.Build(),
	template.Field{
		Label: template.Label{
			Text: "Array to store the data:",
		},
		Input: template.TextInput{
			BindVariable: "TabVarName",
		},
		VariableToggler: template.VariableToggler{
			Checked:  true,
			Disabled: true,
		},
	}.Build(),
}

// OnFindScrapTextTemplate Dialog's OnFindScrapText Template
var OnFindScrapTextTemplate = &template.InstructionTemplate{
	template.Field{
		Label: template.Label{
			Text: "Scraper:",
		},
		Input: template.TextInput{
			BindVariable: "ScraperVarName",
		},
		VariableToggler: template.VariableToggler{
			Checked:  true,
			Disabled: true,
		},
	}.Build(),
	template.Field{
		Label: template.Label{
			Text: "Element to scrap:",
		},
		Input: template.TextInput{
			Bind:         "Element",
			BindVariable: "ElementVarName",
		},
		VariableToggler: template.VariableToggler{
			Bind: "ElementIsVar",
		},
	}.Build(),
	template.Field{
		Label: template.Label{
			Text: "Array to store the data:",
		},
		Input: template.TextInput{
			BindVariable: "TabVarName",
		},
		VariableToggler: template.VariableToggler{
			Checked:  true,
			Disabled: true,
		},
	}.Build(),
}

// OnFindScrapChildTextTemplate Dialog's OnFindScrapChildText Template
var OnFindScrapChildTextTemplate = &template.InstructionTemplate{
	template.Field{
		Label: template.Label{
			Text: "Scraper:",
		},
		Input: template.TextInput{
			BindVariable: "ScraperVarName",
		},
		VariableToggler: template.VariableToggler{
			Checked:  true,
			Disabled: true,
		},
	}.Build(),
	template.Field{
		Label: template.Label{
			Text: "Parent's child to scrap:",
		},
		Input: template.TextInput{
			Bind:         "Element",
			BindVariable: "ElementVarName",
		},
		VariableToggler: template.VariableToggler{
			Bind: "ElementIsVar",
		},
	}.Build(),
	template.Field{
		Label: template.Label{
			Text: "Child name to scrap:",
		},
		Input: template.TextInput{
			Bind:         "ChildAttribute",
			BindVariable: "ChildAttributeVarName",
		},
		VariableToggler: template.VariableToggler{
			Bind: "ChildAttributeIsVar",
		},
	}.Build(),
	template.Field{
		Label: template.Label{
			Text: "Array to store the data:",
		},
		Input: template.TextInput{
			BindVariable: "TabVarName",
		},
		VariableToggler: template.VariableToggler{
			Checked:  true,
			Disabled: true,
		},
	}.Build(),
}

// OnFindVisitTemplate Dialog's OnFindVisit Template
var OnFindVisitTemplate = &template.InstructionTemplate{
	template.Field{
		Label: template.Label{
			Text: "Scraper:",
		},
		Input: template.TextInput{
			BindVariable: "ScraperVarName",
		},
		VariableToggler: template.VariableToggler{
			Checked:  true,
			Disabled: true,
		},
	}.Build(),
	template.Field{
		Label: template.Label{
			Text: "Element to visit:",
		},
		Input: template.TextInput{
			Bind:         "Element",
			BindVariable: "ElementVarName",
		},
		VariableToggler: template.VariableToggler{
			Bind: "ElementIsVar",
		},
	}.Build(),
}

// OnFindChildVisitTemplate Dialog's OnFindChildVisit Template
var OnFindChildVisitTemplate = &template.InstructionTemplate{
	template.Field{
		Label: template.Label{
			Text: "Scraper:",
		},
		Input: template.TextInput{
			BindVariable: "ScraperVarName",
		},
		VariableToggler: template.VariableToggler{
			Checked:  true,
			Disabled: true,
		},
	}.Build(),
	template.Field{
		Label: template.Label{
			Text: "Parent's child to scrap:",
		},
		Input: template.TextInput{
			Bind:         "Element",
			BindVariable: "ElementVarName",
		},
		VariableToggler: template.VariableToggler{
			Bind: "ElementIsVar",
		},
	}.Build(),
	template.Field{
		Label: template.Label{
			Text: "Child name to scrap:",
		},
		Input: template.TextInput{
			Bind:         "ChildAttribute",
			BindVariable: "ChildAttributeVarName",
		},
		VariableToggler: template.VariableToggler{
			Bind: "ChildAttributeIsVar",
		},
	}.Build(),
}

// ScraperEndConditionTemplate Dialog's ScraperEndCondition Template
var ScraperEndConditionTemplate = &template.InstructionTemplate{
	template.Field{
		Label: template.Label{
			Text: "Scraper:",
		},
		Input: template.TextInput{
			BindVariable: "ScraperVarName",
		},
		VariableToggler: template.VariableToggler{
			Checked:  true,
			Disabled: true,
		},
	}.Build(),
	template.Field{
		Label: template.Label{
			Text: "Iterations numbers:",
		},
		Input: template.NumberInput{
			Bind:         "End",
			BindVariable: "EndVarName",
			Decimals:     0,
		},
		VariableToggler: template.VariableToggler{
			Bind: "EndIsVar",
		},
	}.Build(),
}

// ScrapStartTemplate Dialog's ScrapStart Template
var ScrapStartTemplate = &template.InstructionTemplate{
	template.Field{
		Label: template.Label{
			Text: "Scraper:",
		},
		Input: template.TextInput{
			BindVariable: "ScraperVarName",
		},
		VariableToggler: template.VariableToggler{
			Checked:  true,
			Disabled: true,
		},
	}.Build(),
	template.Field{
		Label: template.Label{
			Text: "Url to scrap:",
		},
		Input: template.UrlInput{
			Bind:         "Url",
			BindVariable: "UrlVarName",
		},
		VariableToggler: template.VariableToggler{
			Bind: "UrlIsVar",
		},
	}.Build(),
}
