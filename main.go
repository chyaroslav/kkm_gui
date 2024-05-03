package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Кассовое место")
	// Создание строки со статусом
	statusBar := widget.NewLabel("Статус: Готово")

	// Создание табличной формы
	table := widget.NewTable(
		func() (int, int) {
			return 5, 3
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("Содержимое")
		},
		func(i widget.TableCellID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText("Значение")
		})
	// Создание компоновщика для разделения окна на верхнюю и нижнюю части
	split := container.NewVSplit(
		table,
		statusBar,
	)

	// Установка размеров долями (10/90), так что таблица будет в нижней половине окна
	//splitOffset := myWindow.Canvas().Size().Height / 2
	split.Offset = 100

	// Добавление компоновщика к главному окну
	myWindow.SetContent(split)
	// Обработчик событий клавиатуры для таблицы

	// Функция обработки событий нажатия клавиши для редактируемой таблицы

	/* table.TypedKey =  func(keyEvent *fyne.KeyEvent) {
		if keyEvent.Name == fyne.KeyF7 {
			// Получить индекс текущей ячейки для определения столбца
			_, column := table.Selected()

			// Проверка, что пользователь находится в нужном столбце
			if column == 1 { // Предполагается, что нужный столбец имеет индекс 1
				// Реализация вызова диалога поиска
				// searchValue()
			}
		}
	} */

	// Добавление кнопки и обработчика события для сохранения данных в БД
	//saveButton := widget.NewButton("Сохранить", func() {
	// Реализация сохранения данных в БД
	//})

	// Функция для отображения модального диалогового окна авторизации
	showLoginDialog := func() {
		entryName := widget.NewEntry()
		entryPassword := widget.NewPasswordEntry()

		dialog.ShowForm("Авторизация", "Войти", "Отмена", []*widget.FormItem{
			{Text: "Имя:", Widget: entryName},
			{Text: "Пароль:", Widget: entryPassword},
		}, func(response bool) {
			if response {
				name := entryName.Text
				password := entryPassword.Text
				// Здесь можно добавить проверку имени и пароля
				// например, сравнение с хранимыми данными
				dialog.ShowInformation("Информация", "Авторизация успешна: "+name+" "+password, myWindow)
			}
		}, myWindow)
	}

	// Показать диалог авторизации при старте программы
	showLoginDialog()

	//myWindow.SetContent(content)
	myWindow.ShowAndRun()
}
