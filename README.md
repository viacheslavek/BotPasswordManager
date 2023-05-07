# BotPasswordManager

Тестовое задание на разработчика системы кластеризации Tarantool

Мы каждый день пользуемся большим количеством разных сервисов и зачастую, для каждого из них, требуются логин и пароль.
А наш мозг отказывается запомнить их все... 
Поэтому мы предлагаем вам реализовать Telegram бота, который обладает функционалом персонального хранилища паролей.

## Должны быть поддержаны следующие команды (как именно они будут работать, остается на ваше усмотрение):

1. /set - добавляет логин и пароль к сервису
2. /get - получает логин и пароль по названию сервиса
3. /del - удаляет значения для сервиса

## Требования к реализации:

1. Бот должен быть написан на одном из трех языков на выбор: Golang, Python, Lua
2. Чтобы обеспечить небольшую безопасность, пароли не должны оставаться в чате долго,
бот должен их удалять по истечении некоторого времени.
3. Представьте, что у вашего бота очень много пользователей и каждый из них хочет сохранять свои пароли,
чтобы их не видели другие пользователи. Поэтому для каждого из них должно быть свое пространство.

## Будет плюсом:

1. Развернуть бота в облаке и приложить на него ссылку.
2. Использовать Docker
3. Мы не хотим, чтобы наши пользователи расстроились работой нашего сервиса,
поэтому если уборщица выдернула провод из розетки нашего сервера, мы бы не хотели потерять данные.

