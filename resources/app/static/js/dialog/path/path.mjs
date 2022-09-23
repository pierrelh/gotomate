import { _Dialog } from '../dialog.mjs';

// Path class create the dialog path Ariane's string
export class Path {
	constructor(fullpath, name) {
		this.Element	= document.createElement('li');
		this.Button		= document.createElement('button');
		this.Name		= name;
		this.Path		= fullpath.split(name)[0] + name + '\\';

		this.Button.addEventListener('click', evt => this.Select(evt));
		return this;
	}

	// Create the path element & return it
	Create() {
		this.Button.innerHTML = this.Name;
		this.Element.appendChild(this.Button);
		return this.Element;
	}

	// Select when the user click on this path
	Select() {
		// Send message to Go with the path to follow
		astilectron.sendMessage({
			Identifier	: 'GoTroughFolder',
			Content		: {
				'Path'		: this.Path,
				'Extension'	: _Dialog.Extension,
				'Action'	: _Dialog.Action
			}
		});
	}
}