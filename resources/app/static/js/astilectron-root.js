import { _SideBar }		from './sidebar/sidebar.mjs';
import { _Dialog }		from './dialog/dialog.mjs';
import { _InfoWindow }	from './info-window/info-window.mjs';
import { _View }		from './view/view.mjs';
import { Instruction }	from './view/instruction/instruction.mjs'

document.addEventListener('astilectron-ready', function() {
	astilectron.onMessage(function(message) {
		switch (message.Identifier) {
			case 'InitPackages':
				_SideBar.Init(message.Content.Packages);
				break;

			case 'NewFiber':
				_View.Clean();
				_SideBar.SetFiberName(message.Content.Name);
				for (let i = 0; i < message.Content.Instructions.length; i++)
					new Instruction(message.Content.Instructions[i]);
				
				// scrollView.InitLines();
				break;

			case 'CreateInstruction':
				new Instruction(message.Content);
				break;

			case 'InitNewFiber':
				_InfoWindow.Set(0);
				break;

			case 'FiberUnamed':
				_InfoWindow.Set(1);
				break;

			case 'EraseUnamedFiber':
				_InfoWindow.Set(2, message.Content.Path);
				break;

			case 'ImportFiber':
				_Dialog.SetAction('ImportFiber');
				_Dialog.Hydrate(message.Content);
				break;

			case 'ExportFiber':
				_Dialog.SetAction('ExportFiber');
				_Dialog.Hydrate(message.Content);
				break;

			case 'ImportPackage':
				_Dialog.SetAction('ImportPackage');
				_Dialog.Hydrate(message.Content);
				break;

			case 'InitImportFiber':
				_InfoWindow.Set(3);
				break;

			default:
				break;
		}
	});
});