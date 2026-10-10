package main

import (
	"fmt"
	"time"

	"3-struct/api"
	"3-struct/bins"
	"3-struct/config"
	"3-struct/file"
	"3-struct/storage"
)

func main() {

	cfg := config.NewEncrypter() // reads KEY from the environment, panics if it's missing
	apiClient := api.NewApi(cfg) // hands it to api
	if apiClient.GetKey() != "" {
		fmt.Println("Ключ загружен")
	}

	// file.JsonDb -> storage.Storage -> bins.Service
	repo, err := file.NewJsonDb(storage.DefaultFileName)
	if err != nil {
		fmt.Println("Ошибка создания файлового хранилища:", err)
		return
	}

	store, err := storage.New(repo)
	if err != nil {
		fmt.Println("Ошибка создания storage:", err)
		return
	}

	service, err := bins.NewService(store)
	if err != nil {
		fmt.Println("Ошибка создания сервиса корзин:", err)
		return
	}

	// 1. Создаём несколько корзин
	if _, err := service.Create("bin-001", true, time.Now(), "Private Bin"); err != nil {
		fmt.Println("Ошибка создания bin1:", err)
		return
	}
	if _, err := service.Create("bin-002", false, time.Now(), "Public Documents"); err != nil {
		fmt.Println("Ошибка создания bin2:", err)
		return
	}

	// 2. Сохраняем список в файл через хранилище
	fmt.Println("--- Сохранение данных ---")
	if err := service.Save(); err != nil {
		fmt.Println("Ошибка сохранения:", err)
		return
	}
	fmt.Println("Данные успешно сохранены.")

	// 3. Читаем данные обратно
	fmt.Println("\n--- Чтение данных ---")
	loadedList, err := service.Load()
	if err != nil {
		fmt.Println("Ошибка чтения:", err)
		return
	}

	// 4. Выводим результат, проверяя, что всё восстановилось
	fmt.Printf("Успешно прочитано корзин: %d\n", len(loadedList.Bins))
	for _, b := range loadedList.Bins {
		fmt.Printf("- ID: %s, Название: %s, Приватный: %t, Создан: %s\n",
			b.Id, b.Name, b.Private, b.CreatedAt.Format(time.RFC3339))
	}
}
