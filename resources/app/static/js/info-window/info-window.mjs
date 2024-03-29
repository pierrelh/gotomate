// Setting the InfoWindow class
export const _InfoWindow = new class {
	constructor() {
		this.View		= document.getElementById('InfoWindow');
		this.Title		= document.getElementById('InfoTitle');
		this.Confirm	= document.getElementById('InfoConfirm');
		this.Cancel		= document.getElementById('InfoCancel');
		this.OK			= document.getElementById('InfoOK');
		this.Options 	=	{
								0: {
									identifier	: 'IWCreateNewFiber',
									title		: 'This fiber isn\'t saved, do you really want to open a new one ?',
								},
								1: {
									identifier	: false,
									title		: 'Give a name to your fiber to save her.',
								},
								2: {
									identifier	: 'IWOpenSavedFiber',
									title		: 'Your current fiber isn\'t saved do you want to proceed ?',
								},
								3: {
									identifier	: 'IWImportFiber',
									title		: 'Your current fiber isn\'t saved do you want to proceed ?',
								}
							};

		this.Cancel.addEventListener('click', evt => this.Hide(evt));
		this.OK.addEventListener('click', evt => this.Hide(evt));
	}

	// Setting the datas in the info-window
	Set(id, content) {
		if (this.Options[id] === undefined)
			return;
		this.Title.innerHTML = this.Options[id].title;
		if (!this.Options[id].identifier) {
			if (this.OK.classList.contains('hidden'))
				this.OK.classList.remove('hidden')
			if (!this.Confirm.classList.contains('hidden'))
				this.Confirm.classList.add('hidden')
			if (!this.Cancel.classList.contains('hidden'))
				this.Cancel.classList.add('hidden')
		} else {
			if (!this.OK.classList.contains('hidden'))
				this.OK.classList.add('hidden')
			if (this.Confirm.classList.contains('hidden'))
				this.Confirm.classList.remove('hidden')
			if (this.Cancel.classList.contains('hidden'))
				this.Cancel.classList.remove('hidden')
			this.Confirm.addEventListener('click', evt => this.SendDlg(id, content));
		}
		this.Show();
	}

	// Send the data when the info-window is confirmed
	SendDlg(id, content) {
		astilectron.sendMessage({
			Identifier	: this.Options[id].identifier,
			Content		: {data: content}
		});
		this.Hide();
	}

	// Showing the info-window
	Show() {
		if (!this.View.classList.contains('info-show'))
			this.View.classList.add('info-show');
	}

	// Hiding the info-window
	Hide() {
		if (this.View.classList.contains('info-show'))
			this.View.classList.remove('info-show');
	}
}