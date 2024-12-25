package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func main() {
	superIncrOrder := generateSuperIncreasingSequence(8)
	sum := 0
	for i := range superIncrOrder {
		sum += superIncrOrder[i]
	}
	q := sum
	for {
		if isPrime(q) {
			break
		}
		q++
	}

	r := rand.Intn(sum)

	fmt.Println("Enter the text")

	beta := mod(superIncrOrder, q, r)

	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	text = strings.TrimSpace(text)

	newText := encrypt(text, beta)
	decryption(newText, superIncrOrder, q, r)
}

func generateSuperIncreasingSequence(n int) []int {
	sequence := make([]int, n)
	sum := 0

	for i := 0; i < n; i++ {
		// Генерируем случайное число больше суммы предыдущих элементов + 1
		nextValue := sum + rand.Intn(100) + 1 // +1 чтобы гарантировать, что следующее число больше суммы
		sequence[i] = nextValue
		sum += nextValue
	}

	return sequence
}

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func mod(array []int, q int, r int) []int {
	newarray := make([]int, len(array))
	for i := range array {
		newarray[i] = array[i] * r % q
	}
	return newarray
}

func encrypt(text string, beta []int) []int {
	array := binary(text)
	encryptedText := make([]int, len(text))

	for i, str := range array {
		a := 0
		for j, char := range str {

			b := int(char - '0')
			a = a + int(b)*beta[j]
		}
		encryptedText[i] = a
	}
	fmt.Println("Encrypted text: ", encryptedText)
	return encryptedText
}

func decryption(encryptedText, superIncrOrder []int, q, r int) {
	a, _ := modInverse(r, q)
	arr := mod(encryptedText, q, a)

	arrayOfPlace := make([][]int, len(encryptedText))
	for i := range arr {
		currentIndex := 0
		for arr[i] > 0 {
			smalEl, place := SmallerElement(superIncrOrder, arr[i])
			arrayOfPlace[i] = append(arrayOfPlace[i], place)
			arr[i] -= smalEl
			currentIndex++
		}
	}
	decryptedText := make([]string, 0, 10)

	for i := 0; i < len(encryptedText); i++ {
		decryptedText = append(decryptedText, "00000000")
	}
	for i, row := range decryptedText {
		runes := []rune(row)
		for _, j := range arrayOfPlace[i] {
			runes[j] = '1'
		}
		newnum := string(runes)
		num, _ := strconv.ParseInt(newnum, 2, 64)
		char := string(rune(num))
		decryptedText[i] = char

	}
	result := strings.Join(decryptedText, "")
	fmt.Println("Decrypted text: ", result)
}

// Функция для нахождения НОД и обратного элемента
func extendedGCD(a, b int) (gcd, x, y int) {
	if b == 0 {
		return a, 1, 0
	}
	gcd, x1, y1 := extendedGCD(b, a%b)
	x = y1
	y = x1 - (a/b)*y1
	return
}

// Функция для нахождения мультипликативного обратного
func modInverse(a, m int) (int, error) {
	gcd, x, _ := extendedGCD(a, m)
	if gcd != 1 {
		return 0, fmt.Errorf("обратного элемента не существует")
	}
	return (x%m + m) % m, nil
}

func SmallerElement(elements []int, curEl int) (int, int) {
	for i := range elements {
		if elements[i] > curEl {
			//fmt.Println(elements[i-1], i-1)
			return elements[i-1], i - 1
		}

	}
	return elements[len(elements)-1], len(elements) - 1
}

func binary(s string) []string {
	array := make([]string, len(s))
	for i, c := range s {
		array[i] = fmt.Sprintf("%s%.8b", "", c)
	}
	return array
}
