import { Path }				from './path/path.mjs';
import { File, Directory }	from './element/element.mjs';

// Setting the Dialog class
export const _Dialog = new class {
	constructor() {
		this.View				= document.getElementById('Dialog');
		this.PathSection		= document.getElementById('DirPath');
		this.PathContent		= [];
		this.Path				= undefined;
		this.ContentList		= document.getElementById('ContentList');
		this.Content			= [];
		this.Select				= document.getElementById('DialogSelect');
		this.Cancel				= document.getElementById('DialogCancel');
		this.Extension			= undefined;
		this.Hint				= document.getElementById('FileTypeHint');
		this.Action				= undefined;
		this.SelectedElements	= [];

		this.Cancel.addEventListener('click', evt => this.Hide(evt));
		this.Select.addEventListener('click', evt => this.Selected(evt));
	}

	// Hiding the Dialog
	Hide() {
		if (this.View.classList.contains('dialog-show'))
			this.View.classList.remove('dialog-show');
		this.Reset();
	}

	// Setting the datas in the Dialog
	Hydrate(data) {
		this.Reset();
		this.SetExtension(data.Extension);
		this.SetPath(data.Path);
		this.SetContent(data.Files);
		this.Show();
	}

	// Trigger when the dialog's Select button is clicked
	Selected() {
		let file = this.Path + '\\';
		// If the user didn't select anything but the target is a directory, then the file is the current directory
		// Checking for each element wich one is selected
		if (this.Extension != 'none')
			for (const [_, element] of _Dialog.SelectedElements.entries()) {
				// Checking if the checked element as the right extension
				if (element.Extension != this.Extension)
					continue;
				file = this.Path + '\\' + element.Name
				break;
			}
		// Send message to Go with selected file
		astilectron.sendMessage({
			Identifier	: this.Action,
			Content		: {'File': file}
		});
		this.Hide();
	}

	// Setting the content of the files scrollview
	SetContent(files) {

		// Checking if there is files to show
		if (files == null)
			return;

		for (const [id, file] of files.entries()) {
			// Creating the new file class
			let element = file.IsDir ? new Directory(id, file) : new File(id, file);
			this.Content.push(element);

			// Creating the new element elements
			this.ContentList.appendChild(element.Create());
		}
	}

	// Setting the dialog needed extension & setting the text of the dialog hint
	SetExtension(extension) {
		this.Extension = extension;

		// Checking if the required type is a directory or a file
		this.Hint.innerHTML = this.Extension == 'none' ? 'Select a directory' : 'Select a ' + this.Extension + ' file';
	}

	// Filling the Ariane's string with the required path
	SetPath(fullPath) {
		this.Path = fullPath;
		for (const path of this.Path.split('\\')) {
			if (path == '')
				continue;

			// Creating the Path element & getting his class & his elements
			let newPath = new Path(fullPath, path);
			this.PathContent.push(newPath);
			this.PathSection.appendChild(newPath.Create());

			// Creating the '/' after each path directory
			let slash = document.createElement('li');
			slash.innerHTML = '/';
			this.PathSection.appendChild(slash);
		}
	}

	// Showing the Dialog
	Show() {
		if (!this.View.classList.contains('dialog-show'))
			this.View.classList.add('dialog-show');
	}

	// Reset the dialog's vars & contents
	Reset() {
		this.ContentList.innerHTML = this.PathSection.innerHTML = '';

		// Reseting the Files Content of the Dialog
		for (let index = 0; index < this.Content.length; index++)
			delete this.Content[index];

		this.Content = [];

		// Reseting the Path Content of the Dialog
		for (let index = 0; index < this.PathContent.length; index++)
			delete this.PathContent[index];

		this.PathContent = [];
		this.SelectedElements = [];
	}

	SetAction(action) {
		this.Action = action;
	}
}