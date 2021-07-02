class InfoWindow {
    constructor() {
        this.View    = document.getElementById("InfoWindow");
        this.Title   = document.getElementById("InfoTitle");
        this.Confirm = document.getElementById("InfoConfirm");
        this.Cancel  = document.getElementById("InfoCancel");
        this.OK      = document.getElementById("InfoOK");
        this.Options = {
            0: {
                identifier: "CreateNewFiber",
                title: "This fiber isn't saved, do you really want to open a new one ?",
            },
            1: {
                identifier: false,
                title: "Give a name to your fiber to save her.",
            },
            2: {
                identifier: "OpenSavedFiber",
                title: "Your current fiber isn't saved do you want to proceed ?",
            }
        }
        
        this.Cancel.addEventListener("click", evt => this.Hide(evt))
        this.OK.addEventListener("click", evt => this.Hide(evt))
    }
    Set(id, content) {
        if (this.Options[id] != undefined) {
            this.Title.innerHTML = this.Options[id].title
            if (!this.Options[id].identifier) {
                if (this.OK.classList.contains("hidden")){this.OK.classList.remove("hidden")}
                if (!this.Confirm.classList.contains("hidden")){this.Confirm.classList.add("hidden")}
                if (!this.Cancel.classList.contains("hidden")){this.Cancel.classList.add("hidden")}
            }else {
                if (!this.OK.classList.contains("hidden")){this.OK.classList.add("hidden")}
                if (this.Confirm.classList.contains("hidden")){this.Confirm.classList.remove("hidden")}
                if (this.Cancel.classList.contains("hidden")){this.Cancel.classList.remove("hidden")}
                this.Confirm.addEventListener("click", evt => this.SendDlg(id, content))
            }
            this.Show()
        }
    }
    SendDlg(id, content) {
        data.Identifier = this.Options[id].identifier;
        data.Content = {data: content};
        astilectron.sendMessage(data);
        this.Hide();
    }
    Show() {
        if (!this.View.classList.contains("info-show")) {
            this.View.classList.add("info-show")
        }
    }
    Hide() {
        if (this.View.classList.contains("info-show")) {
            this.View.classList.remove("info-show")
        }
    }
}

var infoWindow = new InfoWindow