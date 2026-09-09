package main

import (
	"fmt"
	//"database/sql"
	//"fmt"
	//"log"
	//"net/http"
	//"sync"

	//"github.com/gorilla/mux"
	//_ "://github.com" // Нужен драйвер для sqlite3
)

// 1. Структуры и интерфейсы объявляются НА УРОВНЕ ПАКЕТА
type Person struct {
	Name string // В Go принято использовать английские имена полей. 
	Age  int    // Если первая буква заглавная — поле публичное (экспортируемое).
}

type Greeter interface {
	Greet(name string)
}

// 2. Функции тоже объявляются вне main
func sayHello(name string) {
	fmt.Println("Привет, " + name)
}

func test(){
	// Инициализация структуры
	p := Person{Name: "Иван", Age: 30}
	print(p)
	// Восстановление после паники (defer должен стоять в начале функции)
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("Восстановлено от паники: %v\n", err)
		}
	}()

	// Циклы
	for i := 0; i < 3; i++ { // Уменьшил до 3 для краткости вывода
		fmt.Println(i)
	}

	// Вызов внешней функции
	sayHello("Иван")

	// Срезы и карты (добавил нижнее подчеркивание _, чтобы скрыть ошибку "not used")
	_ = []int{1, 2, 3}
	_ = map[string]int{"один": 1, "два": 2}

	// Горутины и Мьютексы
	//mutex := &sync.Mutex{}
	go func() {
		//mutex.Lock()
		fmt.Println("Привет из горутины!")
		//mutex.Unlock()
	}()

	// Переменные
	var name string = "Иван"
	var age int = 30

	print(name)
	if age > 18 {
		fmt.Println("Вы взрослый")
	} else {
		fmt.Println("Вы несовершеннолетний")
	}
	/*
	// Настройка роутера Gorilla Mux (Аналог Express)
	r := mux.NewRouter()
	r.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("API Gateway:", r.URL.Path)
		w.Write([]byte("Привет от API Gateway!"))
	})*/

	// Запуск HTTP-сервера (Закомментировано, чтобы код шел дальше, а не зависал на этой строчке)
	// http.ListenAndServe(":8080", r)
}
