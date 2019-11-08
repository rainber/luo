//Домашняя работа 1
package main

import "fmt"

//1.Напишите функцию multiply с двумя вещественными аргументами, которая возвращает их произведение. 
//  При этом, разрешается использовать только операции сложения и присваивания.
func multiply(x, y float64)float64{
    sum := 0.0
    //j := int(y)
    //for i := 0; i < j; i++{
    for i:= 0; i < int(y); i++{
        sum += x 
    }
    return sum + x * (y - float64(int(y))) 
}

//2.Напишите функцию fibonacci1 с одним целочисленным аргументом n, которая с помощью цикла находит и возвращает n-ый член последовательности Фибоначчи.
func fibonacci1(n int)int{
    a, b := []int{1}, []int{1}
	for i := 0; i < n; i++ {
		a = append(a, b[i])
		b = append(b, a[i]+b[i])
	}
	return a[n]
}

//3.Напишите функцию fibonacci2 с одним целочисленным аргументом n, которая с помощью рекурсии находит возвращает n-ый член последовательности Фибоначчи.
func fibonacci2(n int)int{
    if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else if n > 1 {
		return fibonacci2(n-1) + fibonacci2(n-2)
	} else {
		return -1
	}
}

//4.Напишите фунцию bubble_sort с одним аргументом типа ссылка на массив вещественных чисел. 
//  Функция должна отсортировать массив методом пузырька. 
//  Функция не должна ничего возвращать.
func bubble_sort(arr[]int){
    length := len(arr)
	for i := 0; i < length-1; i++ {
		flag := true 
		for j := 0; j < length-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				flag = false
			}
		}
		if flag {
			break
		}
	}
}

//5.Напишите фунцию unique_count с одним аргументом типа ссылка на массив целых чисел. 
//  Функция должна вернуть целое число — количество различных значений. 
//  Например для массива [1, 2, 3, 4, 1, 2, 2, 3, 2], правильный ответ — 4.
func unique_count(arr[]int)int{
    result := []int{}  
    for i := range slc{
        flag := true
        for j := range result{
            if slc[i] == result[j] {
                flag = false  
                break
            }
        }
        if flag {  
            result = append(result, slc[i])
        }
    }
    return len(result)
}

func main(){
//task1:
    fmt.Println(multiply(5,1.1))
    
//task2:
    fmt.Println(fibonacci1(8))

//task3:
    fmt.Println(fibonacci2(8))

//task4:
    arr := []int{8, 4, 2, 9, 10, -3, 3, 20, 15, -1}
	bubble_sort(arr)
	fmt.Println(arr)

//task5:
    arr := []int{1, 2, 3, 4, 1, 2, 2, 3, 2}
	fmt.Println(unique_count(arr))
}