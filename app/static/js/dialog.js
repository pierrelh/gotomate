// Setting the Dialog class
class Dialog {

	constructor() {
		this.View			= document.getElementById("Dialog");
		this.PathSection	= document.getElementById("DirPath");
		this.PathContent	= [];
		this.Path			= undefined;
		this.ContentList	= document.getElementById("ContentList");
		this.Content		= [];
		this.Select			= document.getElementById("DialogSelect");
		this.Cancel			= document.getElementById("DialogCancel");
		this.Extension		= undefined;
		this.Hint			= document.getElementById("FileTypeHint");
		this.Action			= undefined;

		this.Cancel.addEventListener("click", evt => this.Hide(evt));
		this.Select.addEventListener("click", evt => this.Selected(evt));
	}

	// Hiding the Dialog
	Hide() {
		this.Reset();
		if (this.View.classList.contains("dialog-show")) {
			this.View.classList.remove("dialog-show");
		}
	}

	// Setting the datas in the Dialog
	Hydrate(data) {
		this.Reset();

		this.SetExtension(data.Extension);
		this.SetPath(data.Path);
		this.SetContent(data.Files);
		dialog.Show();
	}

	// Trigger when the dialog's Select button is clicked
	Selected() {
		// Checking for each element wich one is selected
		for (let index = 0; index < this.Content.length; index++) {

			// Checking if the element is checked
			if (this.Content[index].Selected) {

				// Checking if the selected element as the right extension
				if (this.Content[index].Extension == this.Extension) {
					console.log("User select file: " + this.Content[index].Name);

					// Send message to Go with selected file
					data.Identifier = this.Action;
					data.Content = {
						"File": this.Path + "\\" + this.Content[index].Name,
					};
					astilectron.sendMessage(data);
					this.Hide();
				}
				return;
			}
		}

		// If the user didn't select anything but the target is a directory, than the file is the current directory
		if (this.Extension == "none") {

			// Send message to Go with selected file
			data.Identifier = this.Action;
			data.Content = {
				"File": this.Path + "\\",
			};
			astilectron.sendMessage(data);
			this.Hide();
		}
	}

	// Setting the content of the files scrollview
	SetContent(files) {

		// Checking if there is files to show
		if (files != null) {
			for (let index = 0; index < files.length; index++) {

				// Creating the new file class
				var file = new File(files[index]);
				this.Content.push(file);

				// Creating the new file elements
				var fileElement = file.Create();
				this.ContentList.appendChild(fileElement);
			}
		}
	}

	// Setting the dialog needed extension & setting the text of the dialog hint
	SetExtension(extension) {
		this.Extension = extension;

		// Checking if the required type is a directory or a file
		if (this.Extension == "none") {
			this.Hint.innerHTML = "Select a directory";
		} else {
			this.Hint.innerHTML = "Select a " + this.Extension + " file";
		}

	}

	// Filling the Ariane's string with the required path
	SetPath(fullPath) {
		this.Path = fullPath;
		var paths = this.Path.split("\\");
		for (let index = 0; index < paths.length; index++) {
			if (paths[index] != "") {

				// Creating the Path element & getting his class & his elements
				var path = new Path(fullPath, paths[index]);
				this.PathContent.push(path);
				var pathElement = path.Create();
				this.PathSection.appendChild(pathElement);

				// Creating the "/" after each path directory
				var slash = document.createElement("li");
				slash.innerHTML = "/";
				this.PathSection.appendChild(slash);
			}
		}
	}

	// Showing the Dialog
	Show() {
		if (!this.View.classList.contains("dialog-show")) {
			this.View.classList.add("dialog-show");
		}
	}

	// Reset the dialog's vars & contents
	Reset() {
		this.ContentList.innerHTML = "";
		this.PathSection.innerHTML = "";

		// Reseting the Files Content of the Dialog
		for (let index = 0; index < this.Content.length; index++) {
			delete this.Content[index];
		}
		this.Content = [];

		// Reseting the Path Content of the Dialog
		for (let index = 0; index < this.PathContent.length; index++) {
			delete this.PathContent[index];
		}
		this.PathContent = [];
	}
}

dialog = new Dialog;

// File class create a file in the dialog (file or directory) & create & handle all his actions
class File {

	constructor(data) {
		this.Element	= document.createElement("li");
		this.Button		= document.createElement("button");
		this.IsDir		= data.IsDir;
		this.Name		= data.Name;
		this.Selected	= false;
		if (this.IsDir) {
			this.Extension = "none";
		} else {
			this.Extension = data.Extension;
		}

		this.Button.addEventListener("click", evt => this.Select(evt));
		if (this.IsDir) {
			this.Button.addEventListener("dblclick", evt => this.GoTrough(evt));
		}
		return this;
	}

	// Create the new file
	Create() {
		if (this.IsDir) {
			this.Button.classList.add("directory");
		}else {
			this.Button.classList.add("file");
		}
		this.Button.innerHTML = this.Name;
		this.Element.appendChild(this.Button);
		return this.Element;
	}

	// Unselect the file
	Unselect() {
		this.Button.classList.remove("selected-content");
		this.Selected = false;
	}

	// Select the file
	Select() {

		// Checking if another file is selected
		for (let index = 0; index < dialog.Content.length; index++) {
			if (dialog.Content[index].Selected) {

				// Unselect the file if it is not this
				if (dialog.Content[index].Name != this.Name) {
					dialog.Content[index].Unselect();
				}
				break;
			}
		}

		// Set selected if not already selectd
		if (!this.Selected) {
			this.Button.classList.add("selected-content");
			this.Selected = true;
		} else {

			// Unselect if already selected
			this.Unselect();
		}
	}

	// Action when the directory is dblclicked
	GoTrough() {
		console.log("Going trough " + dialog.Path + "\\" + this.Name);

		// Send message to Go with the path to follow
		data.Identifier = "GoTroughFolder";
		data.Content = {
			"Path": dialog.Path + "\\" + this.Name,
			"Extension": dialog.Extension,
			"Action": dialog.Action
		};
		astilectron.sendMessage(data);
	}
}

// Path class create the dialog path Ariane's string
class Path {

	constructor(fullpath, name) {
		this.Element	= document.createElement("li");
		this.Button		= document.createElement("button");
		this.Name		= name;
		var path		= fullpath.split(name);
		this.Path		= path[0] + name + "\\";

		this.Button.addEventListener("click", evt => this.Select(evt));
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
		data.Identifier = "GoTroughFolder";
		data.Content = {
			"Path": this.Path,
			"Extension": dialog.Extension,
			"Action": dialog.Action
		};
		astilectron.sendMessage(data);
	}
}