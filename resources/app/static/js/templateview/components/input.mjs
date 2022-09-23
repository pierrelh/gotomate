// Setting the class of an Input
export class Input {
	constructor(data) {
		this.ElementType	= data.ElementType;
		this.Value			= data.Value;
		this.Options		= data.Options;
		this.Type			= data.Type;
		this.ID				= data.ID;
		this.Disabled		= data.Disabled;
		this.IsVar			= data.IsVar;
		this.Input			= document.createElement(this.ElementType);
		this.Element		= document.createElement('li');
	}

	// Create the input
	Create() {
		this.Element.classList = 'input';

		this.Input.id = this.ID;

		if (this.ElementType == 'textarea')
			this.Input.InnerHTML = this.Value;
		else
			this.Input.value = this.Value;

		if (this.Disabled != null)
			this.Input.disabled = this.Disabled;

		if (this.Type != null)
			this.Input.type = this.IsVar ? 'text' : this.Type;
		
		if (this.IsVar)
			this.Input.classList.add('variable-field');

		if (this.Options != null) {
			let selected = -1;
			for (let i = 0; i < this.Options.length; i++) { // Add all the select's options
				let opt = document.createElement('option');
				opt.value = this.Options[i].Value;
				opt.innerHTML = this.Options[i].Name;
				if (this.Value == this.Options[i].Value)
					selected = i;
				this.Input.appendChild(opt);
			}
			this.Input.selectedIndex = selected;
		}

		this.Element.appendChild(this.Input)
		return this.Element;
	}
};