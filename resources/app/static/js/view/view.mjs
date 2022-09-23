// Setting the class of the ScrollView
export const _View = new class {
	constructor() {
		this.View			= document.getElementById('ScrollView');
		this.Fiber			= document.getElementById('Fiber');
		this.Lines			= [];
		this.Instructions	= [];

		this.View.addEventListener('scroll', evt => this.Scrolled());
	}

	// Trigger the scroll on the ScrollView
	Scrolled() {
		// Update all the lines's positions
		for (let i = 0; i < this.Lines.length; i++)
			this.Lines[i].position();
	}

	// Clean the ScrollView by removing all the instructions & lines
	Clean() {
		for (let i = 0; i < this.Lines.length; i++)
			this.Lines[i].remove();
		this.Lines = []
		for (let key in this.Instructions)
			delete this.Instructions[key];
		this.Fiber.innerHTML = '';
	}
};