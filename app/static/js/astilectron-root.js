var data = {
    Identifier: null,
    Content: null,
};

document.addEventListener('astilectron-ready', function() {
    astilectron.onMessage(function(message) {
        switch (message.Identifier) {
            case "InitPackages":
                sideBar.Init(message.Content.Packages)                
                break;
            case "InitFiber":
                scrollView.Clean()
                new Instruction(message.Content)
                break;
            case "CreateInstruction":
                new Instruction(message.Content)
                break;
            case "InitNewFiber":
                dialog.Set(0)
                break;
            case "FiberUnamed":
                dialog.Set(1)
                break;
            default:
                console.log("Unknown Identifier received")
                break;
        }
    });
})