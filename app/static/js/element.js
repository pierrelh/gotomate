// Setting the class of an Input
class Input {

	constructor(data){
		this.ElementType = data.ElementType;
		this.Value       = data.Value;
		this.Options     = data.Options;
		this.Name        = data.Name;
		this.InnerHTML   = data.InnerHTML;
		this.Type        = data.Type;
		this.ID          = data.ID;
		this.Disabled    = data.Disabled;
		this.IsVar       = data.IsVar;
		this.Input       = document.createElement(this.ElementType);
	}

	// Create the input
	Create() {
		if (this.Type != null) {
			this.Input.type = this.Type;
		}

		if (this.Value != null) {
			this.Input.value = this.Value;
		} else if(this.Type == "number") {
			this.Input.value = "0";
		}

		if (this.ID != null) {
			this.Input.id = this.ID;
		}

		if (this.Disabled != null) {
			this.Input.disabled = this.Disabled;
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
				if (this.Name == this.Options[i].Name && this.Value == this.Options[i].Value) {
					selected = i;
				}
				this.Input.appendChild(opt);
			}
			this.Input.selectedIndex = selected;
		}

		if (this.InnerHTML != null) {
			this.Input.InnerHTML = this.InnerHTML;
		}

		return this.Input;
	}
};

// Setting the class of a Label
class Label {

	constructor(datas) {
		this.InnerHTML = datas.InnerHTML;
		this.HtmlFor   = datas.HtmlFor;
		this.Label     = document.createElement("label")
	}

	// Create the label
	Create() {
		this.Label.innerHTML = this.InnerHTML != null ? this.InnerHTML : "";
		this.Label.htmlFor = this.HtmlFor;
		this.Label.classList = "labels"
		return this.Label
	}
};

// Setting the class of a VarToggler
class VarToggler {

	constructor(datas) {
		this.ID                = datas.ID;
		this.Checked           = datas.Checked;
		this.Disabled          = datas.Disabled;
		this.AssignTo          = datas.AssignTo;
		this.OriginalFieldType = datas.Type;
		this.Label             = document.createElement("label");
		this.Checkbox          = document.createElement("input");
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
			} else {
				document.getElementById(assignTo).type = originalType;
			}
		})
		return this.Checkbox;
	}
};