class Dialog {
    constructor() {
        this.View    = document.getElementById("Dialog");
        this.Title   = document.getElementById("DLGTitle");
        this.Confirm = document.getElementById("DLGConfirm");
        this.Cancel  = document.getElementById("DLGCancel");
        this.OK      = document.getElementById("DLGOK");
        this.Options = {
            0: {
                identifier: "CreateNewFiber",
                title: "This fiber isn't saved, do you really want to open a new one ?",
            },
            1: {
                identifier: false,
                title: "Give a name to your fiber to save her.",
            }
        }
        
        this.Cancel.addEventListener("click", evt => this.Hide(evt))
        this.OK.addEventListener("click", evt => this.Hide(evt))
    }
    Set(id) {
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
                this.Confirm.addEventListener("click", evt => this.SendDlg(id))
            }
            this.Show()
        }
    }
    SendDlg(id) {
        data.Identifier = this.Options[id].identifier;
        data.Content = null;
        astilectron.sendMessage(data);
        this.Hide();
    }
    CancelDlg() {
        if (this.View.classList.contains("dialog-show")) {
            this.View.classList.remove("dialog-show")
        }
    }
    Show() {
        if (!this.View.classList.contains("dialog-show")) {
            this.View.classList.add("dialog-show")
        }
    }
    Hide() {
        if (this.View.classList.contains("dialog-show")) {
            this.View.classList.remove("dialog-show")
        }
    }
}

var dialog = new Dialog