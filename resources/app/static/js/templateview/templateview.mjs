import { Input }		from './components/input.mjs'
import { Label }		from './components/label.mjs'
import { VarToggler }	from './components/vartoggler.mjs'

// Setting the TemplateView's class
export const _TemplateView = new class {
	constructor() {
		this.View	= document.getElementById('Templateview');
		this.Params = document.getElementById('Parameters');
		this.Cancel = document.getElementById('CancelParameters');
		this.Close	= document.getElementById('CloseTemplateview');
		this.Form	= document.getElementById('TemplateForm');
		this.Data 	= null;
		this.DataID = null;

		// Add event listener when cancel button is pressed or close cross clicked
		this.Close.addEventListener('click', evt => this.Hide(evt))
		this.Cancel.addEventListener('click', evt => {
			evt.preventDefault();
			this.Hide(evt);
		});

		// Add event listener when parameters are sended
		this.Form.addEventListener('submit', evt => {
			evt.preventDefault();
			this.SendData(evt);
		});
	}

	// Hydrate the template view with the received callback
	Hydrate(callback, id) {
		if (typeof callback == undefined && callback.Content == null)
			return;

		this.DataID = id; // Set the id of the instruction
		this.Params.innerHTML = ''; // Delete the content of the section

		let fields = callback.Content.Template; // Get the fields of callback
		let databinder = callback.Content.Databinder
		for (let index = 0; index < fields.length; index++) {
			let ul = document.createElement('ul'); // Create a new field list
			let isVar = false;
			let fieldID = 'Field' + index;
			let varToggler = undefined;

			if (fields[index].VariableToggler != null) {
				if (fields[index].VariableToggler.Checked || databinder[fields[index].VariableToggler.Bind])
					isVar = true;

				varToggler = new VarToggler({
					ID			: 'VariableToggler' + index,
					Checked		: isVar,
					Disabled	: fields[index].VariableToggler.Disabled,
					AssignTo	: fieldID,
					Type		: fields[index].Input.Type,
				}).Create();
			}

			// Create a label if needed
			if (fields[index].Label != null) {

				// Set the label parameters
				let label = new Label({
					InnerHTML: fields[index].Label.Text,
					HtmlFor: fieldID
				}).Create();

				ul.appendChild(label);
			}

			if (fields[index].Input != null) { // Create an Input if needed

				let inputValue = fields[index].Input.Value;

				if (isVar && databinder[fields[index].Input.BindVariable] != undefined)
					inputValue = databinder[fields[index].Input.BindVariable];						
				else if (!isVar && databinder[fields[index].Input.Bind] != undefined)
					inputValue = databinder[fields[index].Input.Bind];


				let inputData = {
					ElementType	: fields[index].Input.ElementType,
					ID			: fieldID,
					IsVar		: isVar,
					Disabled	: fields[index].Input.Disabled,
					Value		: inputValue,
				};

				switch (fields[index].Input.ElementType) { // Search wich input to create
					case 'input': // Create an input of type input set the parameters
						inputData.Type = fields[index].Input.Type;
						break;

					case 'select': // Create an input of type select & set the parameters
						inputData.Options = fields[index].Input.Options;
						break;

					default:
						break;
				}
				let input = new Input(inputData).Create();
				ul.appendChild(input);
			}

			if (varToggler != undefined)
				ul.appendChild(varToggler);

			this.Params.appendChild(ul);
		}
		// Setting the datas to send if the form is send
		this.Data = callback.Content;
		this.Show();
	}

	// Show the template view if hidded
	Show() {
		if (!this.View.classList.contains('templateview-show'))
			this.View.classList.add('templateview-show');
	}

	// Hide the template view if showed
	Hide() {
		if (this.View.classList.contains('templateview-show')) {
			this.View.classList.remove('templateview-show');
			this.DataID = null;
		}
	}

	// Clean the form
	Clean() {
		this.Params.innerHTML = '';
	}

	// Send the setted datas in the template view to Go
	SendData() {
		for (let index = 0; index < this.Data.Template.length; index++) {
			let thisInput = this.Data.Template[index];
			let isVar = false;
			if (thisInput.VariableToggler != null) {
				if (this.Data.Databinder[thisInput.VariableToggler.Bind] != undefined)
					this.Data.Databinder[thisInput.VariableToggler.Bind] = this.Params.children[index].querySelector('.check').children[1].checked;
				isVar = this.Params.children[index].querySelector('.check').children[1].checked;
			}

			let input = this.Params.children[index].querySelector('.input').children[0]
			
			if (isVar) {
				if (this.Data.Databinder[thisInput.Input.BindVariable] != undefined)
					this.Data.Databinder[thisInput.Input.BindVariable] = input.value;
			} else {
				
				switch (thisInput.Input.ElementType) {
					// Set the required values of the form data
					case 'input':
						if (this.Data.Databinder[thisInput.Input.Bind] != undefined) {
							if (thisInput.Input.Type == 'number')
								this.Data.Databinder[thisInput.Input.Bind] = parseInt(input.value);
							else
								this.Data.Databinder[thisInput.Input.Bind] = input.value;
						}
						break;
	
					case 'select':
						if (this.Data.Databinder[thisInput.Input.Bind] != undefined)
							this.Data.Databinder[thisInput.Input.Bind] = input.value;
						if (input.options[input.selectedIndex] != undefined && this.Data.Databinder[thisInput.Input.BindName] != undefined)
							this.Data.Databinder[thisInput.Input.BindName] = input.options[input.selectedIndex].text;
						break;
	
					case 'textarea':
						if (this.Data.Databinder[thisInput.Input.Bind] != undefined)
							this.Data.Databinder[thisInput.Input.Bind] = input.innerHTML;
						break;
	
					default:
						break;
				}
			}
		}

		astilectron.sendMessage({
			Identifier	: 'SetDatabinderDatas',
			Content		: {
				ID			: parseInt(this.DataID),
				Databinder	: this.Data.Databinder,
			}
		});

		this.Hide();
	}

	// GetTimeOut Return the time out to show the template view if hided or showed
	GetTimeOut() {
		this.Hide();
		return !this.View.classList.contains('templateview-show') ? 0 : 300;
	}
}