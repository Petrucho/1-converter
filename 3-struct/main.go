package main

import (
	"fmt"
	"time"

	"3-struct/bins"
	"3-struct/storage"
)

func main() {
	// 1. Создаем несколько корзин
	bin1, err := bins.CreateBin("bin-001", true, time.Now(), "Private Bin")
	if err != nil {
		fmt.Println("Ошибка создания bin1:", err)
		return
	}

	bin2, err := bins.CreateBin("bin-002", false, time.Now(), "Public Documents")
	if err != nil {
		fmt.Println("Ошибка создания bin2:", err)
		return
	}

	// 2. Инициализируем список и добавляем их туда
	list := bins.NewBinList()
	list.Add(bin1)
	list.Add(bin2)

	// 3. Сохраняем список в файл через пакет storage
	fmt.Println("--- Сохранение данных ---")
	err = storage.SaveBins(list)
	if err != nil {
		fmt.Println("Ошибка сохранения:", err)
		return
	}
	fmt.Println("Данные успешно сохранены.")

	// 4. Читаем данные обратно из файла
	fmt.Println("\n--- Чтение данных ---")
	loadedList, err := storage.ReadBins()
	if err != nil || loadedList == nil {
		fmt.Println("Ошибка чтения:", err)
		return
	}

	// 5. Выводим результат, проверяя, что всё восстановилось
	fmt.Printf("Успешно прочитано корзин: %d\n", len(loadedList.Bins))
	for _, b := range loadedList.Bins {
		fmt.Printf("- ID: %s, Название: %s, Приватный: %t, Создан: %s\n",
			b.Id, b.Name, b.Private, b.CreatedAt.Format(time.RFC3339))
	}
}
