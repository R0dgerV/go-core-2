//Пакет для вычисления числа фибоначи
package fibonache

//Вычисляет число фибоначи
func Number(num int) int {
	switch num {
	case 1, 2:
		return 1
	default:
		return Number(num - 1) + Number(num - 2)
	}
}