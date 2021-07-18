// Setting the TemplateView's class 
class TemplateView {
	constructor() {
		this.View = document.getElementById("Templateview");
		this.Params = document.getElementById("Parameters");
		this.Cancel = document.getElementById("CancelParameters");
		this.Close = document.getElementById("CloseTemplateview");
		this.Form = document.getElementById("TemplateForm");
		this.Data = null;
		this.DataID = null;

		// Add event listener when cancel button is pressed or close cross clicked
		this.Close.addEventListener("click", evt => this.Hide(evt))
		this.Cancel.addEventListener("click", evt => {
			evt.preventDefault();
			this.Hide(evt);
		});

		// Add event listener when parameters are sended
		this.Form.addEventListener("submit", evt => {
			evt.preventDefault();
			this.SendData(evt);
		});
	}

	// Hydrate the template view with the received callback
	Hydrate(callback, id) {
		if (callback != undefined && callback.Content != null) {
			console.log(callback);
			this.DataID = id; // Set the id of the instruction
			this.Params.innerHTML = ""; // Delete the content of the section

			var fields = callback.Content; // Get the fields of callback
			for (let index = 0; index < fields.length; index++) {
				var ul = document.createElement("ul"); // Create a new field list
				var isVar = false;
				var fieldID = "Field" + index;
				if (fields[index].VariableToggler != null && fields[index].VariableToggler.Checked) {
					isVar = true;
				}

				// Create a label if needed
				if (fields[index].Label != null && fields[index].Label.ElementType != "") {
					var labelLi = document.createElement("li");
					labelLi.classList = "label";

					// Set the label parameters
					var label = new Label({
						InnerHTML: fields[index].Label.Text,
						HtmlFor: fieldID
					});
					labelLi.appendChild(label.Create());

					ul.appendChild(labelLi);
				}

				if (fields[index].Input != null && fields[index].Input.ElementType != "") { // Create an Input if needed
					var li = document.createElement("li");
					li.classList = "input";

					var inputData = {
						ElementType: fields[index].Input.ElementType,
						ID: fieldID,
						IsVar: isVar
					};

					switch (fields[index].Input.ElementType) { // Search wich input to create
						case "input": // Create an input of type input set the parameters
							inputData.Value = fields[index].Input.Value;
							inputData.Type = fields[index].Input.Type;
							inputData.Disabled = fields[index].Input.Disabled;
							break;

						case "select": // Create an input of type select & set the parameters
							inputData.Options = fields[index].Input.Options;
							inputData.Value = fields[index].Input.Value;
							inputData.Name = fields[index].Input.Name;
							break;

						case "textArea": // Create an input of type textarea & set the parameters
							inputData.InnerHTML = fields[index].Input.Value;
							break;

						default:
							break;
					}
					var input = new Input(inputData);
					li.appendChild(input.Create());
					ul.appendChild(li);
				}

				// Create a variable toggler if needed
				if (fields[index].VariableToggler != null && fields[index].VariableToggler.ElementType != "") {
					var checkboxLi = document.createElement("li");
					checkboxLi.classList = "check";

					var varToggler = new VarToggler({
						ID: "VariableToggler" + index,
						Checked: fields[index].VariableToggler.Checked,
						Disabled: fields[index].VariableToggler.Disabled,
						AssignTo: fieldID,
						Type: fields[index].Input.Type,
					});

					// Create & set the variable toggler label parameters
					checkboxLi.appendChild(varToggler.CreateLabel());

					// Create & set the variable toggler checkbox parameters
					checkboxLi.appendChild(varToggler.CreateCheckbox());

					ul.appendChild(checkboxLi);
				}

				this.Params.appendChild(ul);
			}
			// Setting the datas to send if the form is send
			this.Data = callback.Content;
			this.Show();
		}
	}

	// Show the template view if hidded
	Show() {
		if (!this.View.classList.contains("templateview-show")) {
			this.View.classList.add("templateview-show");
		}
	}

	// Hide the template view if showed
	Hide() {
		if (this.View.classList.contains("templateview-show")) {
			this.View.classList.remove("templateview-show");
			this.DataID = null;
		}
	}

	// Clean the form
	Clean() {
		this.Params.innerHTML = "";
	}

	// Send the setted datas in the template view to Go
	SendData() {
		for (let index = 0; index < this.Data.length; index++) {
			if (this.Data[index].Input.ElementType != "") {
				var input = this.Params.children[index].querySelector('.input').children[0]
				switch (this.Data[index].Input.ElementType) {
					// Set the required values of the form data
					case "input":
						if (this.Data[index].Input.Type == "number" && input.value == "") {
							this.Data[index].Input.Value = "0";
						}else {
							this.Data[index].Input.Value = input.value;
						}
						break;

					case "select":
						this.Data[index].Input.Value = input.value;
						if (input.options[input.selectedIndex] != undefined) {
							this.Data[index].Input.Name = input.options[input.selectedIndex].text;
						}
						break;

					case "textarea":
						this.Data[index].Input.Value = input.innerHTML;
						break;

					default:
						break;
				}
			}

			if (this.Data[index].VariableToggler != null) {
				this.Data[index].VariableToggler.Checked = this.Params.children[index].querySelector('.check').children[1].checked;
			}
		}
		data.Identifier = "SetDatabinderDatas";
		data.Content = {
			ID: parseInt(this.DataID),
			Template: this.Data,
		};
		console.log(data);
		// Sending the data to Go
		astilectron.sendMessage(data);
		this.Hide();
	}

	// GetTimeOut Return the time out to show the template view if hided or showed
	GetTimeOut() {
		var time;
		if (!this.View.classList.contains("templateview-show")) {
			time = 0;
		} else {
			time = 300;
		}
		this.Hide();
		return time;
	}
}

templateView = new TemplateView;