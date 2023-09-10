package telegram

const msgHelp = `I can save and keep you pages. Also I can offer you them to read.
In order to save the page, just send me all link to it.

In order to get a random page from your list, send me command /random
Caution! After that, this page will be removed from your list!`

const msgHello = "Hi there! 🤟🏻 \n\n" + msgHelp

const (
	msgUnknownCommand = "Unknown command 🤷🏻‍"
	msgNoSavedPages   = "You have no saved pages 😢"
	msgSaved          = "Saved! ✅"
	msgAlreadyExists  = "You have already have this page in your list 😐"
	msgListLinks      = "📄Your saved links:\n"
)
