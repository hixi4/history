package main

import (
	"fmt"
	"strings"
)

// Персонаж
type Character struct {
	Name      string
	Inventory []string
	Location  string
}

// Гра
type Game struct {
	Character Character
}

// Введення вибору
func (g *Game) getInput() string {
	var input string
	fmt.Scanln(&input)
	return strings.ToLower(input)
}

// Введення вибору з опціями
func (g *Game) getChoice(options []string) string {
	for {
		choice := g.getInput()
		for _, option := range options {
			if choice == option {
				return choice
			}
		}
		fmt.Println("Неправильний вибір. Спробуйте ще раз.")
	}
}

// Основний сценарій гри
func (g *Game) play() {
	fmt.Println("Ви прокидаєтесь в окопі у лісі. Ви все пам'ятаєте, але не знаєте, як тут опинились. Вас звати Василь.")
	fmt.Println("Поряд з вами рюкзак, в якому ви знаходите цигарки, запальничку, автомат і ніж.")
	fmt.Println("Перед вами є два шляхи: 'землянка' і 'ліс'. Куди ви підете?")

	for {
		fmt.Print("Ваш вибір (землянка/ліс): ")
		choice := g.getChoice([]string{"землянка", "ліс"})

		switch choice {
		case "землянка":
			fmt.Println("У землянці дуже темно. Ви використовуєте запальничку, щоб освітити шлях.")
			// Додатковий сценарій для землянки може бути доданий тут
			fmt.Println("Ви вирішили повернутися назад.")
		case "ліс":
			fmt.Println("Ви йдете по лісу.")
			fmt.Println("У лісі ви натрапляєте на мертве тіло ворога.")
			fmt.Println("Ви обираєте оглянути тіло й ідете далі.")
			fmt.Println("Через деякий час ви приходите до кинутих позицій військових.")
			fmt.Println("Ви втомлені і вирішуєте відпочити, а не йти далі.")
			fmt.Println("У найближчому наметі ви знаходите сейф з кодовим замком з трьох чисел.")

			for {
				fmt.Println("Спробуйте відкрити сейф. Введіть код (три числа):")
				var code1, code2, code3 int
				fmt.Print("Перше число: ")
				fmt.Scan(&code1)
				fmt.Print("Друге число: ")
				fmt.Scan(&code2)
				fmt.Print("Третє число: ")
				fmt.Scan(&code3)

				if code1 == 3 && code2 == 5 && code3 == 1 {
					fmt.Println("Сейф відчиняється. Ви бачите секретні документи, і думаєте, як передати їх військовим.")
					fmt.Println("Ви відправляєтесь шукати військових.")
					return
				} else {
					fmt.Println("Неправильний код. Сейф не відкривається.")
					fmt.Println("Чи хочете спробувати ще раз? (так/ні)")
					retry := g.getChoice([]string{"так", "ні"})
					if retry == "ні" {
						fmt.Println("Ви вирішуєте відпочити.")
						return
					}
				}
			}
		}
	}
}

func main() {
	// Створення персонажа
	character := Character{
		Name:      "Василь",
		Inventory: []string{"цигарки", "запальничка", "автомат", "ніж"},
		Location:  "окоп у лісі",
	}

	// Ініціалізація гри
	game := Game{
		Character: character,
	}

	// Запуск гри
	game.play()
}
