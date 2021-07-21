// Setting the class of an Input
class Input {

	constructor(data){
		this.ElementType	= data.ElementType;
		this.Value			= data.Value;
		this.Options		= data.Options;
		this.Type			= data.Type;
		this.ID				= data.ID;
		this.Disabled		= data.Disabled;
		this.IsVar			= data.IsVar;
		this.Input			= document.createElement(this.ElementType);
		this.Element		= document.createElement("li");
	}

	// Create the input
	Create() {
		this.Element.classList = "input";

		this.Input.id = this.ID;

		if (this.ElementType == "textarea") {
			this.Input.InnerHTML = this.Value;
		} else {
			this.Input.value = this.Value;
		}

		if (this.Disabled != null) {
			this.Input.disabled = this.Disabled;
		}

		if (this.Type != null) {
			if (this.IsVar) {
				this.Input.type = "text";				
			} else {
				this.Input.type = this.Type;
			}
		}
		
		if (this.IsVar) {
			this.Input.classList.add("variable-field");
		}

		if (this.Options != null) {
			var selected = -1;
			for (let i = 0; i < this.Options.length; i++) { // Add all the select's options
				var opt = document.createElement('option');
				opt.value = this.Options[i].Value;
				opt.innerHTML = this.Options[i].Name;
				if (this.Value == this.Options[i].Value) {
					selected = i;
				}
				this.Input.appendChild(opt);
			}
			this.Input.selectedIndex = selected;
		}

		this.Element.appendChild(this.Input)
		return this.Element;
	}
};

// Setting the class of a Label
class Label {

	constructor(datas) {
		this.InnerHTML	= datas.InnerHTML;
		this.HtmlFor	= datas.HtmlFor;
		this.Label		= document.createElement("label")
		this.Element	= document.createElement("li");
	}

	// Create the label
	Create() {
		this.Element.classList = "label";

		this.Label.innerHTML = this.InnerHTML != null ? this.InnerHTML : "";
		this.Label.htmlFor = this.HtmlFor;
		this.Label.classList = "labels";

		this.Element.appendChild(this.Label)
		return this.Element;
	}
};

// Setting the class of a VarToggler
class VarToggler {

	constructor(datas) {
		this.ID					= datas.ID;
		this.Checked			= datas.Checked;
		this.Disabled			= datas.Disabled;
		this.AssignTo			= datas.AssignTo;
		this.OriginalFieldType	= datas.Type;
		this.Label				= document.createElement("label");
		this.Checkbox			= document.createElement("input");
		this.Element			= document.createElement("li");
	}

	// Creating the VariableToggler element
	Create() {
		this.Element.classList = "check";
		this.Element.appendChild(this.CreateLabel());
		this.Element.appendChild(this.CreateCheckbox());
		return this.Element;
	}

	// Creating the label of the VarToggler
	CreateLabel() {
		this.Label.innerHTML = "Is a var ?";
		this.Label.classList = "check-labels";
		this.Label.htmlFor = this.ID;
		return this.Label;
	}

	// Creating the checkbox of the VarToggler
	CreateCheckbox() {
		this.Checkbox.type = "checkbox";
		this.Checkbox.checked = this.Checked != null ? this.Checked : false;
		this.Checkbox.id = this.ID;
		this.Checkbox.disabled = this.Disabled != null ? this.Disabled : false;
		var assignTo = this.AssignTo;
		var originalType = this.OriginalFieldType;

		// Adding the event when the checkbox status change
		this.Checkbox.addEventListener("change", function() {
			document.getElementById(assignTo).classList.toggle("variable-field");
			if (this.checked) {
				document.getElementById(assignTo).type = "text";
				document.getElementById(assignTo).value = "";
			} else {
				document.getElementById(assignTo).type = originalType;
			}
		})
		return this.Checkbox;
	}
};