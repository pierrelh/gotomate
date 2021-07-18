// Setting the class of an Input
class Input {

    constructor(data){
        this.ElementType = data.ElementType;
        this.Value       = data.Value;
        this.Options     = data.Options;
        this.Name        = data.Name;
        this.InnerHTML   = data.InnerHTML;
        this.Type        = data.Type;
        this.ID          = data.ID;
        this.Disabled    = data.Disabled;
        this.IsVar       = data.IsVar;
    }

    // Create the input
    Create() {
        var input = document.createElement(this.ElementType);

        if (this.Type != null) {
            input.type = this.Type;
        }

        if (this.Value != null) {
            input.value = this.Value;
        } else if(this.Type == "number") {
            input.value = "0";
        }

        if (this.ID != null) {
            input.id = this.ID;
        }

        if (this.Disabled != null) {
            input.disabled = this.Disabled;
        }
        
        if (this.IsVar) {
            input.classList.add("variable-field");
        }

        if (this.Options != null) {
            var selected = -1;
            for (let i = 0; i < this.Options.length; i++) { // Add all the select's options
                var opt = document.createElement('option');
                opt.value = this.Options[i].Value;
                opt.innerHTML = this.Options[i].Name;
                if (this.Name == this.Options[i].Name && this.Value == this.Options[i].Value) {
                    selected = i;
                }
                input.appendChild(opt);
            }
            input.selectedIndex = selected;
        }

        if (this.InnerHTML != null) {
            input.InnerHTML = this.InnerHTML;
        }

        return input;
    }
};

// Setting the class of a Label
class Label {

    constructor(datas) {
        this.InnerHTML = datas.InnerHTML;
        this.HtmlFor   = datas.HtmlFor;
    }

    // Create the label
    Create() {
        var label = document.createElement("label")
        label.innerHTML = this.InnerHTML != null ? this.InnerHTML : "";
        label.htmlFor = this.HtmlFor;
        label.classList = "labels"
        return label
    }
};

// Setting the class of a VarToggler
class VarToggler {

    constructor(datas) {
        this.ID                = datas.ID;
        this.Checked           = datas.Checked;
        this.Disabled          = datas.Disabled;
        this.AssignTo          = datas.AssignTo;
        this.OriginalFieldType = datas.Type;
    }

    // Creating the label of the VarToggler
    CreateLabel() {
        var label = document.createElement("label");
        label.innerHTML = "Is a var ?";
        label.classList = "check-labels";
        label.htmlFor = this.ID;
        return label;
    }

    // Creating the checkbox of the VarToggler
    CreateCheckbox() {
        var checkbox = document.createElement("input");
        checkbox.type = "checkbox";
        checkbox.checked = this.Checked != null ? this.Checked : false;
        checkbox.id = this.ID;
        checkbox.disabled = this.Disabled != null ? this.Disabled : false;
        var assignTo = this.AssignTo;
        var originalType = this.OriginalFieldType;

        // Adding the event when the checkbox status change
        checkbox.addEventListener("change", function() {
            document.getElementById(assignTo).classList.toggle("variable-field");
            if (this.checked) {
                document.getElementById(assignTo).type = "text";
            } else {
                document.getElementById(assignTo).type = originalType;
            }
        })
        return checkbox;
    }
};