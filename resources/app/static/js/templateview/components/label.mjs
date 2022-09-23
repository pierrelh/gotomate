// Setting the class of a Label
export class Label {
	constructor(datas) {
		this.InnerHTML	= datas.InnerHTML;
		this.HtmlFor	= datas.HtmlFor;
		this.Label		= document.createElement('label')
		this.Element	= document.createElement('li');
	}

	// Create the label
	Create() {
		this.Element.classList = 'label';
		this.Label.innerHTML = this.InnerHTML != null ? this.InnerHTML : '';
		this.Label.htmlFor = this.HtmlFor;
		this.Label.classList = 'labels';
		this.Element.appendChild(this.Label)
		return this.Element;
	}
};