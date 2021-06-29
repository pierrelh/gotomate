var click

class ScrollView {
    constructor() {
        this.View         = document.getElementById("ScrollView")
        this.Fiber        = document.getElementById("Fiber");
        this.Lines        = []
        this.Instructions = []

        this.View.addEventListener("scroll", evt => this.Scrolled())
    }
    Scrolled() {
        for (let i = 0; i < this.Lines.length; i++) {
            this.Lines[i].position()
        }
    }
    Clean() {
        for (let i = 0; i < this.Lines.length; i++) {
            this.Lines[i].remove()
        }
        this.Lines = []
        for (let key in this.Instructions) {
            delete this.Instructions[key]
        }
        this.Fiber.innerHTML = ""
    }
}

var scrollView = new ScrollView

class Instruction {
    constructor(data) {
        if (data != undefined) {
            this.Div         = document.createElement("div")
            this.NextIDInput = document.createElement("input");
            this.ID          = data.ID;
            this.FuncName    = data.FuncName;
            this.IconPath    = data.IconPath;
            this.NextID      = data.NextID;
            this.X           = 0;
            this.Y           = 0;
            this.Moved       = false
            this.InLine      = [];
            this.OutLine     = null;
            this.Create()       
        }
    }
    // Create a new Instruction to the fiber
    Create() {
        this.Div.id = "Instruction" + this.ID
        this.Div.className = "instruction"

        var ul = document.createElement("ul")

        var IDLi = document.createElement("li")
        IDLi.className = "id"
        var IDP = document.createElement("p")
        IDP.innerHTML = this.ID
        IDLi.appendChild(IDP)
        ul.appendChild(IDLi)

        var IMGLi = document.createElement("li")
        IMGLi.className = "func-img"
        var img = document.createElement("img")
        img.src = this.IconPath
        IMGLi.appendChild(img)
        ul.appendChild(IMGLi)

        var NameLi = document.createElement("li")
        NameLi.className = "func-name"
        var NameP = document.createElement("p")
        NameP.innerHTML = this.FuncName
        NameLi.appendChild(NameP)
        ul.appendChild(NameLi)

        var NextLi = document.createElement("li")
        NextLi.className = "next-instruction"
        var NextLabel = document.createElement("label")
        NextLabel.innerHTML = "Next:"
        NextLi.appendChild(NextLabel)

        this.NextIDInput.value = this.NextID
        this.NextIDInput.type = "number"
        this.NextIDInput.min = 0
        this.NextIDInput.addEventListener("change", evt => this.CreateOutLine(evt))
        NextLi.appendChild(this.NextIDInput)
        ul.appendChild(NextLi)

        this.Div.appendChild(ul)
        this.Div.addEventListener("mousedown", evt => this.InstructionInteractions(evt));
        scrollView.Instructions[this.Div.id] = this
        scrollView.Fiber.appendChild(this.Div)

        for (var key in scrollView.Instructions) {
            if (scrollView.Instructions[key].NextIDInput.value == this.ID) {
                this.CreateInLine(scrollView.Instructions[key])
                break;
            }
        }
    }
    // Delete the instruction from the fiber & the class
    Delete(callback) {
        if (callback != undefined) {
            if(Object.keys(this.InLine).length != 0) {
                for (var key in this.InLine) {
                    scrollView.Instructions[key].RemoveOutLine()
                    this.InLine[key].remove()
                }
            }

            if(this.OutLine != null) {
                scrollView.Instructions[this.OutLine.end.id].RemoveInLine(this.Div.id)
                this.OutLine.remove()
            }
            
            document.getElementById("Instruction" + this.ID).remove()
            if (templateView.DataID == this.ID) {
                setTimeout(function(){
                    templateView.Clean();
                }, templateView.GetTimeOut());
            }
            delete scrollView.Instructions[this.Div.id];
            delete this
        }
    }
    // Create all the interaction for an instruction
    InstructionInteractions(e) {
        if (e.target.localName != "input") {
            click = e.which
            
            // Set the end of the drag on mouse up
            document.onmouseup = evt => this.CloseDrag(evt);

            if (click == 1) {
                // Get the mouse cursor position at startup:
                this.X = e.clientX;
                this.Y = e.clientY;
                // Make the instruction dragable
                document.onmousemove = evt => this.Drag(evt);
            }
        }
    }       
    // Drag the instruction to the user cursor
    Drag(e) {
        e = e || window.event;
        e.preventDefault();
        // Calculate the new cursor position:
        var posInstructionX = this.X - e.clientX;
        var posInstructionY = this.Y - e.clientY;

        // Set the instruction's new position:
        // Set the instruction's new Y position
        if (posInstructionY > 25 || posInstructionY < -25) {
            if (this.Div.offsetTop - posInstructionY < 0) {
                this.Div.style.top = "0px";
            } else if (this.Div.offsetTop - posInstructionY + this.Div.offsetHeight > scrollView.Fiber.offsetHeight) {
                this.Div.style.top = (scrollView.Fiber.offsetHeight - this.Div.offsetHeight) + "px";
            } else {
                if (posInstructionY > 25) {
                    this.Div.style.top = (this.Div.offsetTop - 25) + "px";
                } else {
                    this.Div.style.top = (this.Div.offsetTop + 25) + "px";
                }
                this.Y = e.clientY;
            }
        }

        // Set the instruction's new X position
        if (posInstructionX > 25 || posInstructionX < -25) {
            if (this.Div.offsetLeft - posInstructionX < 0) {
                this.Div.style.left = "0px";
            } else if (this.Div.offsetLeft - posInstructionX + this.Div.offsetWidth > scrollView.Fiber.offsetWidth) {
                this.Div.style.left = (scrollView.Fiber.offsetWidth - this.Div.offsetWidth) + "px";
            } else {
                if (posInstructionX > 25) {
                    this.Div.style.left = (this.Div.offsetLeft - 25) + "px";
                } else {
                    this.Div.style.left = (this.Div.offsetLeft + 25) + "px";
                }
                this.X = e.clientX;
            }            
        }
        if (this.OutLine != null) {
            this.OutLine.position()
        }
        
        if (Object.keys(this.InLine).length != 0) {
            for (var key in this.InLine) {
                this.InLine[key].position()
            }
        }
        this.Moved = true;
    }
    // End the drag interaction of the instruction
    CloseDrag() {
        // If the instruction hasn't been moved (just clicked) then show the instruction's menu
        if(!this.Moved) {
            if (click == 1 && this.ID != templateView.DataID) {
                setTimeout(function(){
                    // Send request to get Button template in Go
                    data.Identifier = "GetInstructionTemplate";
                    data.Content = {"ID": parseInt(this.ID)};
                    astilectron.sendMessage(data, callback => templateView.Hydrate(callback, this.ID));
                }.bind(this), templateView.GetTimeOut());
            } else if (click == 3 && this.ID != 0) {
                // Send request to delete Button in Go
                data.Identifier = "DeleteInstruction";
                data.Content = {"ID": parseInt(this.ID)};
                astilectron.sendMessage(data, callback => this.Delete(callback));
            }
        } else {
            // Send request to update Button's position in Go
            data.Identifier = "UpdateInstructionPosition";
            data.Content = {
                "ID": parseInt(this.ID),
                "X": parseInt(this.Div.style.left),
                "Y": parseInt(this.Div.style.top)            
            };
            astilectron.sendMessage(data);
        }
        // Unset the mouse up & mouse move event
        document.onmouseup = null;
        document.onmousemove = null;
        this.Moved = false
    }
    // Create a line that came out of this instruction
    CreateOutLine() {
        var nextID = "Instruction" + this.NextIDInput.value
        if (this.OutLine != null) {
            scrollView.Instructions[this.OutLine.end.id].RemoveInLine(this.Div.id)
            this.OutLine.remove()
            this.OutLine = null
        }
        if (document.getElementById(nextID) != undefined && nextID != this.Div.id) {
            var line = new LeaderLine(
                document.getElementById(this.Div.id),
                document.getElementById(nextID), {
                    startSocket: "bottom",
                    endSocket:   "top",
                    path:        "grid",
                    startPlug:   "disc",
                    endPlug:     "disc",
                    color:       "rgba(255, 255, 255, 0.8)",
                }
            );
            
            this.OutLine = line
            scrollView.Lines.push(line)

            for (var key in scrollView.Instructions) {
                if (scrollView.Instructions[key].ID == this.NextIDInput.value) {
                    scrollView.Instructions[key].AddInLine(this.Div.id, line)
                    break;
                }
            }
        }
    }
    // Create a line that came in of this instruction
    CreateInLine(out) {
        var previousID = "Instruction" + out.ID
        if (document.getElementById(previousID) != undefined && previousID != this.Div.id) {
            var line = new LeaderLine(
                document.getElementById(previousID),
                document.getElementById(this.Div.id), {
                    startSocket: "bottom",
                    endSocket:   "top",
                    path:        "grid",
                    startPlug:   "disc",
                    endPlug:     "disc",
                    color:       "rgba(255, 255, 255, 0.8)",
                }
            );
            scrollView.Lines.push(line)
            this.InLine[previousID] = line
            out.AddOutLine(line)
        }
    }
    // Add a new in line
    AddInLine(id, line) {
        this.InLine[id] = line
    }
    // Add a new out line
    AddOutLine(line) {
        this.OutLine = line
    }
    // Remove an out line from the lines & from this element
    RemoveOutLine() {
        for (let i = 0; i < scrollView.Lines.length; i++) {
            if (this.OutLine == scrollView.Lines[i]) {
                scrollView.Lines.splice(i, 1);
                break;
            }            
        }
        this.OutLine = null
    }
    // Remove an in line from the lines & from this element
    RemoveInLine(id) {
        for (let i = 0; i < scrollView.Lines.length; i++) {
            if (scrollView.Lines[i] == this.InLine[id]) {
                scrollView.Lines.splice(i, 1);
                break;
            }            
        }
        delete this.InLine[id];
    }
}