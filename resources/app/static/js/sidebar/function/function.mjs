import { Instruction } from '../../view/instruction/instruction.mjs'

// Setting the class of a Function
export class Function {
	constructor(id, packageName, funcName) {
		this.ID				= id
		this.PackageName	= packageName;
		this.FuncName		= funcName;
		this.Element		= document.createElement('li');
	}

	// Create a function in a package's list
	Create(parentID) {
		let functionSpan = document.createElement('span');
		functionSpan.innerHTML = this.FuncName;
		this.Element.appendChild(functionSpan);
		this.Element.addEventListener('click', evt => this.Open(evt));
		return this.Element;
	}

	// Open the function's instruction
	Open() {
		astilectron.sendMessage({
			Identifier	: 'CreateInstruction',
			Content		: {
				'PackageName': this.PackageName,
				'FuncName': this.FuncName
			}
		}, function(callback) {
			new Instruction(callback)
		});
	}
}