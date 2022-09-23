// Setting the class of a VarToggler
export class VarToggler {
	constructor(datas) {
		this.ID					= datas.ID;
		this.Checked			= datas.Checked;
		this.Disabled			= datas.Disabled;
		this.AssignTo			= datas.AssignTo;
		this.OriginalFieldType	= datas.Type;
		this.Label				= document.createElement('label');
		this.Checkbox			= document.createElement('input');
		this.Element			= document.createElement('li');
	}

	// Creating the VariableToggler element
	Create() {
		this.Element.classList = 'check';
		this.Element.appendChild(this.CreateLabel());
		this.Element.appendChild(this.CreateCheckbox());
		return this.Element;
	}

	// Creating the label of the VarToggler
	CreateLabel() {
		this.Label.innerHTML = 'Is a var ?';
		this.Label.classList = 'check-labels';
		this.Label.htmlFor = this.ID;
		return this.Label;
	}

	// Creating the checkbox of the VarToggler
	CreateCheckbox() {
		this.Checkbox.type = 'checkbox';
		this.Checkbox.checked = this.Checked != null ? this.Checked : false;
		this.Checkbox.id = this.ID;
		this.Checkbox.disabled = this.Disabled != null ? this.Disabled : false;
		const self = this;

		// Adding the event when the checkbox status change
		this.Checkbox.addEventListener('change', function() {
			document.getElementById(self.AssignTo).classList.toggle('variable-field');
			document.getElementById(self.AssignTo).type = this.checked ? 'text' : self.OriginalFieldType;
			if (this.checked)
				document.getElementById(self.AssignTo).value = '';
		})
		return this.Checkbox;
	}
};