var packages

class SideBar {
    constructor() {
        this.View      = document.getElementById("PackagesList");
        this.FiberName = document.getElementById("FiberName")
        this.Packages  = []

        this.FiberName.addEventListener("keyup", evt => this.SaveFiberName())
    }
    Init(content) {
        packages = content
        for (let x = 0; x < packages.length; x++) {
            var newPackage = new Packages(x, packages[x].Name).Create()
            this.Packages.push(newPackage)
            PackagesList.appendChild(newPackage)
        }
    }
    SaveFiberName() {
        data.Identifier = "UpdateFiberName";
        data.Content = {"name": this.FiberName.value};
        astilectron.sendMessage(data);
    }
}

class Packages {
    constructor(id, name) {
        this.ID        = id;
        this.Name      = name;
        this.Element   = document.createElement("li");
        this.Span      = document.createElement("span")
        this.FuncList  = document.createElement("ul");
        this.Collapsed = true;
    }
    Create() {
        this.Span.classList = "branch caret"
        this.Span.innerHTML = this.Name
        this.Span.addEventListener("click", evt => this.OpenList(evt));
        this.Element.appendChild(this.Span)

        this.FuncList.classList = "nested"
        for (var x = 0; x < packages[this.ID].Functions.length; x++) {
            var func = new Function(x, this.Name, packages[this.ID].Functions[x]).Create(this.ID)
            this.FuncList.appendChild(func)
        }
        this.Element.appendChild(this.FuncList)
        return this.Element
    }
    OpenList() {
        this.Span.classList.toggle("caret-down");
          
        if(this.Collapsed) {
            var sectionHeight = this.FuncList.scrollHeight;
            this.FuncList.style.height = sectionHeight + 'px';
            this.Collapsed = false
        } else {
            var sectionHeight = this.FuncList.scrollHeight;    
            var elementTransition = this.FuncList.style.transition;
            this.FuncList.style.transition = '';
        
            requestAnimationFrame(function() {
                this.FuncList.style.height = sectionHeight + 'px';
                this.FuncList.style.transition = elementTransition;
                requestAnimationFrame(function() {
                    this.FuncList.style.height = 0 + 'px';
                }.bind(this));
            }.bind(this));
            this.Collapsed = true
        }
    }
}

class Function {
    constructor(id, packageName, funcName) {
        this.ID          = id
        this.PackageName = packageName;
        this.FuncName    = funcName;
        this.Element     = document.createElement("li");
    }
    Create(parentID) {
        var functionSpan = document.createElement("span")
        functionSpan.innerHTML = packages[parentID].Functions[this.ID]
        this.Element.appendChild(functionSpan)
        this.Element.addEventListener("click", evt => this.Open(evt));
        return this.Element
    }
    Open() {
        data.Identifier = "CreateInstruction";
        data.Content = {
            "PackageName": this.PackageName,
            "FuncName": this.FuncName
        };
        astilectron.sendMessage(data, function(callback) {
            new Instruction(callback)
        });
    }
}

sideBar = new SideBar