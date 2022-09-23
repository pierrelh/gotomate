import { Function } from '../function/function.mjs';

// Setting the class of a Package
export class Package {
	constructor(id, name) {
		this.ID			= id;
		this.Name		= name;
		this.Element	= document.createElement('li');
		this.Span		= document.createElement('span')
		this.FuncList	= document.createElement('ul');
		this.Collapsed	= true;
	}

	// Create a package in the Sidebar
	Create(functions) {
		this.Span.classList = 'branch caret';
		this.Span.innerHTML = this.Name;
		this.Span.addEventListener('click', evt => this.OpenList(evt));
		this.Element.appendChild(this.Span);

		this.FuncList.classList = 'nested';
		for (let [key, func] of functions.entries())
			this.FuncList.appendChild(new Function(key, this.Name, func).Create(this.ID));
		this.Element.appendChild(this.FuncList);
		return this.Element;
	}

	// Toggle the classlist to show the package's functions
	OpenList() {
		this.Span.classList.toggle('caret-down');

		if (this.Collapsed) {
			this.FuncList.style.height = this.FuncList.scrollHeight + 'px';
			this.Collapsed = false;
		} else {
			let elementTransition = this.FuncList.style.transition;
			this.FuncList.style.transition = '';

			requestAnimationFrame(function() {
				this.FuncList.style.height = this.FuncList.scrollHeight + 'px';
				this.FuncList.style.transition = elementTransition;
				requestAnimationFrame(function() {
					this.FuncList.style.height = 0 + 'px';
				}.bind(this));
			}.bind(this));
			this.Collapsed = true;
		}
	}
}