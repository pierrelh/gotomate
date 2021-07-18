var data = {
    Identifier: null,
    Content: {},
};

document.addEventListener('astilectron-ready', function() {
    astilectron.onMessage(function(message) {
        switch (message.Identifier) {
            case "InitPackages":
                sideBar.Init(message.Content.Packages);
                break;

            case "NewFiber":
                scrollView.Clean();
                sideBar.SetFiberName(message.Content.Name);
                for (let i = 0; i < message.Content.Instructions.length; i++) {
                    new Instruction(message.Content.Instructions[i]);
                }
                break;

            case "CreateInstruction":
                new Instruction(message.Content);
                break;

            case "InitNewFiber":
                infoWindow.Set(0);
                break;

            case "FiberUnamed":
                infoWindow.Set(1);
                break;

            case "EraseUnamedFiber":
                infoWindow.Set(2, message.Content.Path);
                break;

            case "ImportFiber":
                dialog.Hydrate(message.Content);
                dialog.Show();
                break;

            case "InitImportFiber":
                console.log("To Do")
                dialog.Show();
                break;

            default:
                console.log("Unknown Identifier received");
                break;
        }
    });
});