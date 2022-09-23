import { _Dialog } from '../dialog.mjs';

// Element class create an element in the dialog (file or directory) & create it & handle all his actions
class Element {
	constructor(id, data) {
		this.ID			= id;
		this.Element	= document.createElement('li');
		this.Button		= document.createElement('button');
		this.Name		= data.Name;
		this.Selected	= false;
		this.ClassName	= '';
		this.Extension	= this.IsDir ? 'none' : data.Extension;

		this.Button.addEventListener('click', evt => this.Select(evt));
		return this;
	}

	// Create the element
	Create() {
		this.Button.classList.add(this.ClassName);
		this.Button.innerHTML = this.Name;
		this.Element.appendChild(this.Button);
		return this.Element;
	}

	// SetSelected the element
	SetSelected() {
		this.Button.classList.add('selected-content');
		this.Selected = true;
		_Dialog.SelectedElements[this.ID] = this;
	}

	// SetUnselected the element
	SetUnselected() {
		this.Button.classList.remove('selected-content');
		this.Selected = false;
		_Dialog.SelectedElements[this.ID] = null;
	}

	// Select the file
	Select() {
		let isSelected = this.Selected;

		// Unselect the selected elements
		for (const [_, element] of _Dialog.SelectedElements.entries())
			if (element != null)
				element.SetUnselected();

		// Choose either to select or unselect the element
		if (!isSelected)
			this.SetSelected();
	}
}

export class File extends Element {
	constructor(id, data) {
		super(id, data);
		this.ClassName = 'file';
		return this;
	}
}

export class Directory extends Element {
	constructor(id, data) {
		super(id, data);
		this.ClassName = 'directory';
		this.Button.addEventListener('dblclick', evt => this.GoTrough(evt));
	}

	// Action when the directory is dblclicked
	GoTrough() {
		// Send message to Go with the path to follow
		astilectron.sendMessage({
			Identifier	: 'GoTroughFolder',
			Content		: {
				'Path'		: _Dialog.Path + '\\' + this.Name,
				'Extension'	: _Dialog.Extension,
				'Action'	: _Dialog.Action
			}
		});
	}
}