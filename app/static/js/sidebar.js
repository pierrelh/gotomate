// Setting the SideBar class
class SideBar {

	constructor() {
		this.View      = document.getElementById("PackagesList");
		this.FiberName = document.getElementById("FiberName");
		this.Run       = document.getElementById("FiberRunButton");
		this.Stop      = document.getElementById("FiberStopButton");
		this.Packages  = [];
		this.Data      = null;

		this.FiberName.addEventListener("keyup", evt => this.SaveFiberName());
		this.Run.addEventListener("click", evt => this.FiberRun());
		this.Stop.addEventListener("click", evt => this.FiberStop());
	}

	// Init Fill the sidebar with all the fiber's packages
	Init(content) {
		this.Data = content;
		for (let x = 0; x < this.Data.length; x++) {
			var newPackage = new Packages(x, this.Data[x].Name).Create();
			this.Packages.push(newPackage);
			PackagesList.appendChild(newPackage);
		}
	}

	// FiberRun Send message to run the fiber
	FiberRun() {
		data.Identifier = "RunFiber";
		data.Content = {};
		astilectron.sendMessage(data);
	}

	// FiberStop Send message to stop the fiber
	FiberStop() {
		data.Identifier = "StopFiber";
		data.Content = {};
		astilectron.sendMessage(data);
	}

	// SaveFiberName Save the setted fiber name
	SaveFiberName() {
		data.Identifier = "UpdateFiberName";
		data.Content = {"name": this.FiberName.value};
		astilectron.sendMessage(data);
	}

	// SaveFiberName Set the fiber's name when opening a fiber
	SetFiberName(name) {
		this.FiberName.value = name;
	}
}

sideBar = new SideBar;

// Setting the class of a Packages
class Packages {

	constructor(id, name) {
		this.ID        = id;
		this.Name      = name;
		this.Element   = document.createElement("li");
		this.Span      = document.createElement("span")
		this.FuncList  = document.createElement("ul");
		this.Collapsed = true;
	}

	// Create a package in the Sidebar
	Create() {
		this.Span.classList = "branch caret";
		this.Span.innerHTML = this.Name;
		this.Span.addEventListener("click", evt => this.OpenList(evt));
		this.Element.appendChild(this.Span);

		this.FuncList.classList = "nested";
		for (var x = 0; x < sideBar.Data[this.ID].Functions.length; x++) {
			var func = new Function(x, this.Name, sideBar.Data[this.ID].Functions[x]).Create(this.ID);
			this.FuncList.appendChild(func);
		}
		this.Element.appendChild(this.FuncList);
		return this.Element;
	}

	// Toggle the classlist to show the package's functions
	OpenList() {
		this.Span.classList.toggle("caret-down");

		if(this.Collapsed) {
			var sectionHeight = this.FuncList.scrollHeight;
			this.FuncList.style.height = sectionHeight + 'px';
			this.Collapsed = false;
		} else {
			var sectionHeight = this.FuncList.scrollHeight;
			var elementTransition = this.FuncList.style.transition;
			this.FuncList.style.transition = '';

			requestAnimationFrame(function() {
				this.FuncList.style.height = sectionHeight + 'px';
				this.FuncList.style.transition = elementTransition;
				requestAnimationFrame(function() {
					this.FuncList.style.height = 0 + 'px';
				}.bind(this));
			}.bind(this));
			this.Collapsed = true;
		}
	}
}

// Setting the class of a Function
class Function {

	constructor(id, packageName, funcName) {
		this.ID          = id
		this.PackageName = packageName;
		this.FuncName    = funcName;
		this.Element     = document.createElement("li");
	}

	// Create a function in a package's list
	Create(parentID) {
		var functionSpan = document.createElement("span");
		functionSpan.innerHTML = sideBar.Data[parentID].Functions[this.ID];
		this.Element.appendChild(functionSpan);
		this.Element.addEventListener("click", evt => this.Open(evt));
		return this.Element;
	}

	// Open the function's instruction
	Open() {
		data.Identifier = "CreateInstruction";
		data.Content = {
			"PackageName": this.PackageName,
			"FuncName": this.FuncName
		};
		astilectron.sendMessage(data, function(callback) {
			new Instruction(callback)
		});
	}
}