import { Package } from './package/package.mjs';

// Setting the SideBar class
export const _SideBar = new class {
	constructor() {
		this.View		= document.getElementById('PackagesList');
		this.FiberName	= document.getElementById('FiberName');
		this.Run		= document.getElementById('FiberRunButton');
		this.Stop		= document.getElementById('FiberStopButton');
		this.Packages	= [];
		this.Data		= null;

		this.FiberName.addEventListener('keyup', evt => this.SaveFiberName());
		this.Run.addEventListener('click', evt => this.FiberRun());
		this.Stop.addEventListener('click', evt => this.FiberStop());
	}

	// Init Fill the sidebar with all the fiber's packages
	Init(content) {
		this.Data = content;
		for (let [id, pack] of this.Data.entries()) {
			let newPackage = new Package(id, pack.Name).Create(pack.Functions);
			this.Packages.push(newPackage);
			this.View.appendChild(newPackage);
		}
	}

	// FiberRun Send message to run the fiber
	FiberRun() {
		astilectron.sendMessage({
			Identifier	: 'RunFiber',
			Content		: {}
		});

	}

	// FiberStop Send message to stop the fiber
	FiberStop() {
		astilectron.sendMessage({
			Identifier	: 'StopFiber',
			Content		: {}
		});

	}

	// SaveFiberName Save the setted fiber name
	SaveFiberName() {
		astilectron.sendMessage({
			Identifier	: 'UpdateFiberName',
			Content		: {'name' : this.FiberName.value}
		});
		
	}

	// SaveFiberName Set the fiber's name when opening a fiber
	SetFiberName(name) {
		this.FiberName.value = name;
	}
}