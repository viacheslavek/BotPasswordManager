package tgEvents

const (
	msgHelp = `
		Мои команды:
		/set <site> <login> <password> добавляет логин и пароль ресурса 
		/get <site> получает все аккаунты для нужного ресурса
		/del <site> удаляет все аккаунты ресурса
		/help справка
		/start приветствие

	`
	msgStart = `
		Приветствие.
		Здесь ты можешь сохранить свои пароли, но они будут в большей опасности,
		чем если ты запишешь их себе в статус в телеграмм. Кто-то вообще читает статусы в телеграмм?
		Когда станет безопасно, я тебе сообщу

	`
	msgAdd = "success save account!\n"
	msgDel = "success delete account!\n"
	msgGet = `
		success get account!
		Your account:

	`
	msgUnknownCommand = "Unknown command.\nSent /help to see opportunities"
)
