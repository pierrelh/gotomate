import { _View }			from '../view.mjs';
import { _TemplateView }	from '../../templateview/templateview.mjs';

let click

// Setting the class of an Instruction
export class Instruction {
	constructor(data) {
		if (data != undefined) {
			this.Div			= document.createElement('div');
			this.NextIDInput	= document.createElement('input');
			this.ID				= data.ID;
			this.FuncName		= data.FuncName;
			this.Package		= data.Package;
			this.IconPath		= data.IconPath;
			this.NextID			= data.NextID;
			this.X				= data.X;
			this.Y				= data.Y;
			this.Moved			= false
			this.InLine			= [];
			this.OutLine		= null;
			this.Create();
		} else
			delete this;
	}

	// Create a new Instruction to the fiber
	Create() {
		this.Div.id = 'Instruction' + this.ID;
		this.Div.className = 'instruction';
		this.Div.style.top = this.Y + 'px';
		this.Div.style.left = this.X + 'px';

		let ul = document.createElement('ul');

		// Creating the instruction ID
		let IDLi = document.createElement('li');
		IDLi.className = 'id';
		let IDP = document.createElement('p');
		IDP.innerHTML = this.ID;
		IDLi.appendChild(IDP);
		ul.appendChild(IDLi);

		// Creating the instruction image
		let IMGLi = document.createElement('li');
		IMGLi.className = 'func-img';
		let img = document.createElement('img');
		img.src = '../' + this.IconPath;
		IMGLi.appendChild(img);
		ul.appendChild(IMGLi);

		// Creating the Instruction name
		let NameLi = document.createElement('li');
		NameLi.className = 'func-name';
		let NameP = document.createElement('p');
		NameP.innerHTML = this.FuncName;
		NameLi.appendChild(NameP);
		ul.appendChild(NameLi);

		// Creating the Next ID input if the instruction is not a flow's end instruction
		let NextLi = document.createElement('li');
		NextLi.className = 'next-instruction';
		if (this.Package != 'Flow' || this.FuncName != 'End') {
			let NextLabel = document.createElement('label')
			NextLabel.innerHTML = 'Next:';
			NextLi.appendChild(NextLabel);

			this.NextIDInput.value = this.NextID;
			this.NextIDInput.type = 'number';
			this.NextIDInput.min = 0;
			this.NextIDInput.addEventListener('change', evt => this.CreateOutLine(evt));
			NextLi.appendChild(this.NextIDInput);
		}
		ul.appendChild(NextLi);

		this.Div.appendChild(ul);
		this.Div.addEventListener('mousedown', evt => this.InstructionInteractions(evt));
		_View.Instructions[this.Div.id] = this;
		_View.Fiber.appendChild(this.Div);

		for (let key in _View.Instructions)
			if (_View.Instructions[key].NextIDInput.value == this.ID) {
				this.CreateInLine(_View.Instructions[key]);
				break;
			}
	}

	// Delete the instruction from the fiber & the class
	Delete(callback) {
		if (callback === undefined)
			return;
		
		if (Object.keys(this.InLine).length != 0)
			for (let key in this.InLine) {
				_View.Instructions[key].RemoveOutLine();
				this.InLine[key].remove();
			}

		if (this.OutLine != null) {
			_View.Instructions[this.OutLine.end.id].RemoveInLine(this.Div.id);
			this.OutLine.remove();
		}

		document.getElementById('Instruction' + this.ID).remove();
		if (_TemplateView.DataID == this.ID) {
			setTimeout(function(){
				_TemplateView.Clean();
			}, _TemplateView.GetTimeOut());
		}
		delete _View.Instructions[this.Div.id];
		delete this;
	}

	// Create all the interaction for an instruction
	InstructionInteractions(e) {
		if (e.target.localName != 'input') {
			click = e.which;

			// Set the end of the drag on mouse up
			document.onmouseup = evt => this.CloseDrag(evt);

			if (click == 1) {
				// Get the mouse cursor position at startup:
				this.X = e.clientX;
				this.Y = e.clientY;
				// Make the instruction dragable
				document.onmousemove = evt => this.Drag(evt);
			}
		}
	}

	// Drag the instruction to the user cursor
	Drag(e) {
		e = e || window.event;
		e.preventDefault();
		// Calculate the new cursor position:
		let posInstructionX = this.X - e.clientX;
		let posInstructionY = this.Y - e.clientY;

		// Set the instruction's new Y position
		if (posInstructionY > 25 || posInstructionY < -25) {
			if (this.Div.offsetTop - posInstructionY < 0)
				this.Div.style.top = '0px';
			else if (this.Div.offsetTop - posInstructionY + this.Div.offsetHeight > _View.Fiber.offsetHeight)
				this.Div.style.top = (_View.Fiber.offsetHeight - this.Div.offsetHeight) + 'px';
			else {
				this.Div.style.top = (posInstructionY > 25 ? this.Div.offsetTop - 25 : this.Div.offsetTop + 25) + 'px';
				this.Y = e.clientY;
			}
		}

		// Set the instruction's new X position
		if (posInstructionX > 25 || posInstructionX < -25) {
			if (this.Div.offsetLeft - posInstructionX < 0)
				this.Div.style.left = '0px';
			else if (this.Div.offsetLeft - posInstructionX + this.Div.offsetWidth > _View.Fiber.offsetWidth)
				this.Div.style.left = (_View.Fiber.offsetWidth - this.Div.offsetWidth) + 'px';
			else {
				this.Div.style.left = (posInstructionX > 25 ? this.Div.offsetLeft - 25 : this.Div.offsetLeft + 25) + 'px';
				this.X = e.clientX;
			}
		}
		if (this.OutLine != null)
			this.OutLine.position();

		if (Object.keys(this.InLine).length != 0)
			for (let key in this.InLine)
				this.InLine[key].position();
		this.Moved = true;
	}

	// End the drag interaction of the instruction
	CloseDrag() {
		// If the instruction hasn't been moved (just clicked) then show the instruction's menu
		if (!this.Moved) {
			if (click == 1 && this.ID != _TemplateView.DataID) {
				setTimeout(function(){
					// Send request to get Button template in Go
					astilectron.sendMessage({
							Identifier	: 'GetInstructionTemplate',
							Content		: {'ID': parseInt(this.ID)}
						},
						callback => _TemplateView.Hydrate(callback, this.ID)
					);
				}.bind(this), _TemplateView.GetTimeOut());
			} else if (click == 3 && this.ID != 0) {
				// Send request to delete Button in Go
				astilectron.sendMessage({
						Identifier	: 'DeleteInstruction',
						Content		: {'ID': parseInt(this.ID)}
					},
					callback => this.Delete(callback)
				);				
			}
		} else {
			// Send request to update Button's position in Go
			astilectron.sendMessage({
				Identifier	: 'UpdateInstructionPosition',
				Content		: {
					'ID'	: parseInt(this.ID),
					'X'		: parseInt(this.Div.style.left),
					'Y'		: parseInt(this.Div.style.top)
				}
			});
		}
		// Unset the mouse up & mouse move event
		document.onmouseup = null;
		document.onmousemove = null;
		this.Moved = false;
	}

	// Create a line that came out of this instruction
	CreateOutLine() {
		let nextID = 'Instruction' + this.NextIDInput.value;
		if (this.OutLine != null) {
			_View.Instructions[this.OutLine.end.id].RemoveInLine(this.Div.id);
			this.OutLine.remove();
			this.OutLine = null;
		}
		// Creating a new line if the instruction with the id exist & if it's not this
		if (document.getElementById(nextID) != undefined && nextID != this.Div.id) {
			let line = new LeaderLine(
				document.getElementById(this.Div.id),
				document.getElementById(nextID), {
					startSocket	: 'bottom',
					endSocket	: 'top',
					path		: 'grid',
					startPlug	: 'disc',
					endPlug		: 'disc',
					color		: 'rgba(255, 255, 255, 0.8)',
				}
			);

			this.OutLine = line;
			_View.Lines.push(line);

			for (let key in _View.Instructions)
				if (_View.Instructions[key].ID == this.NextIDInput.value) {
					_View.Instructions[key].AddInLine(this.Div.id, line);
					break;
				}
		}
		// Sending the data to Go
		astilectron.sendMessage({
			Identifier	: 'UpdateInstructionNextID',
			Content		: {
				ID: parseInt(this.ID),
				NextID: parseInt(this.NextIDInput.value),
			}
		});
	}

	// Create a line that came in of this instruction
	CreateInLine(out) {
		let previousID = 'Instruction' + out.ID;
		if (document.getElementById(previousID) != undefined && previousID != this.Div.id) {
			let line = new LeaderLine(
				document.getElementById(previousID),
				document.getElementById(this.Div.id), {
					startSocket	: 'bottom',
					endSocket	: 'top',
					path		: 'grid',
					startPlug	: 'disc',
					endPlug		: 'disc',
					color		: 'rgba(255, 255, 255, 0.8)',
				}
			);
			_View.Lines.push(line);
			this.InLine[previousID] = line;
			out.AddOutLine(line);
		}
	}

	// Add a new in line
	AddInLine(id, line) {
		this.InLine[id] = line;
	}

	// Add a new out line
	AddOutLine(line) {
		this.OutLine = line;
	}

	// Remove an out line from the lines & from this element
	RemoveOutLine() {
		for (let i = 0; i < _View.Lines.length; i++)
			if (this.OutLine == _View.Lines[i]) {
				_View.Lines.splice(i, 1);
				break;
			}
		this.OutLine = null;
	}

	// Remove an in line from the lines & from this element
	RemoveInLine(id) {
		for (let i = 0; i < _View.Lines.length; i++)
			if (_View.Lines[i] == this.InLine[id]) {
				_View.Lines.splice(i, 1);
				break;
			}
		delete this.InLine[id];
	}
}