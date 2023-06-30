package main

import (
	"errors"
	"fmt"
	"math"
	"math/rand"
	"sort"
	"strings"
	"time"
	"unicode"
)

// q.1

func calcularMedia(slice []int) float64 {
	soma := 0
	for _, valor := range slice {
		soma += valor
	}

	media := float64(soma) / float64(len(slice))
	return media
}

func main() {
	var tamanho int
	fmt.Print("Digite o tamanho do slice: ")
	fmt.Scanln(&tamanho)

	slice := make([]int, tamanho)
	for i := 0; i < tamanho; i++ {
		fmt.Printf("Digite o valor %d: ", i+1)
		fmt.Scanln(&slice[i])
	}

	media := calcularMedia(slice)
	fmt.Println("A média é:", media)
}

// q.2

func contarVogais(str string) int {
	vogais := []rune{'a', 'e', 'i', 'o', 'u'}
	contador := 0

	for _, char := range str {
		for _, vogal := range vogais {
			if char == vogal || char == vogal-32 {
				contador++
				break
			}
		}
	}

	return contador
}

func main() {
	var texto string
	fmt.Print("Digite uma frase: ")
	fmt.Scanln(&texto)

	quantidade := contarVogais(texto)
	fmt.Println("A quantidade de vogais é:", quantidade)
}

// q.3

func concatenarStrings(slice []string) string {
	resultado := strings.Join(slice, "")
	return resultado
}

func main() {
	var tamanho int
	fmt.Print("Digite o tamanho do slice: ")
	fmt.Scanln(&tamanho)

	slice := make([]string, tamanho)
	for i := 0; i < tamanho; i++ {
		fmt.Printf("Digite a string %d: ", i+1)
		fmt.Scanln(&slice[i])
	}

	concatenacao := concatenarStrings(slice)
	fmt.Println("A concatenação é:", concatenacao)
}

// q.4

func jogarJogo() int {
	resposta := rand.Intn(100) + 1
	tentativas := 0

	for {
		var valor int
		fmt.Print("Digite um número entre 1 e 100: ")
		fmt.Scanf("%d", &valor)
		tentativas++

		if valor == resposta {
			fmt.Println("Parabéns, você acertou!")
			fmt.Println("Você utilizou", tentativas, "tentativas.")
			break
		} else if valor < resposta {
			fmt.Println("O número é maior que", valor, ".")
		} else {
			fmt.Println("O número é menor que", valor, ".")
		}
	}

	return tentativas
}

func jogarNovamente() bool {
	var resposta string
	fmt.Print("Você deseja jogar novamente? (s/n): ")
	fmt.Scanf("%s", &resposta)
	return resposta == "s" || resposta == "S"
}

func exibirTodasAsJogadas(jogadas []int) {
	for i, tentativas := range jogadas {
		fmt.Printf("Você utilizou %d tentativas na %dª jogada.\n", tentativas, i+1)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	fmt.Println("Bem-vindo ao jogo da adivinhação!")

	jogadas := []int{}

	for {
		tentativas := jogarJogo()
		jogadas = append(jogadas, tentativas)

		if !jogarNovamente() {
			break
		}
	}

	fmt.Println()
	exibirTodasAsJogadas(jogadas)
}

// q.5

func EncontrarPosicao(slice []int, valor int) int {
	for i, num := range slice {
		if num == valor {
			return i
		}
	}
	return -1
}

func main() {
	slice := []int{10, 20, 30, 40, 50}
	valor := 30
	posicao := EncontrarPosicao(slice, valor)
	fmt.Printf("Valor %d encontrado na posição %d\n", valor, posicao)

	valor = 60
	posicao = EncontrarPosicao(slice, valor)
	fmt.Printf("Valor %d encontrado na posição %d\n", valor, posicao)
}

// q.6

func ConcatenarStrings(slice []string) (string, error) {
	if len(slice) == 0 {
		return "", errors.New("Slice vazio")
	}

	resultado := strings.Join(slice, ",")
	return resultado, nil
}

func main() {
	slice := []string{"Maçã", "Banana", "Laranja"}
	resultado, err := ConcatenarStrings(slice)
	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Println("Resultado:", resultado)
	}

	sliceVazio := []string{}
	resultadoVazio, errVazio := ConcatenarStrings(sliceVazio)
	if errVazio != nil {
		fmt.Println("Erro:", errVazio)
	} else {
		fmt.Println("Resultado:", resultadoVazio)
	}
}

// q.7

type FuncaoInt func(int) int

func AplicarFuncao(slice []int, funcao FuncaoInt) ([]int, error) {
	if len(slice) == 0 {
		return nil, errors.New("Slice vazio")
	}

	resultado := make([]int, len(slice))
	for i, num := range slice {
		resultado[i] = funcao(num)
	}

	return resultado, nil
}

func main() {
	slice := []int{1, 2, 3, 4, 5}
	dobroSlice, err := AplicarFuncao(slice, func(x int) int {
		return x * 2
	})
	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Println("Dobro do slice:", dobroSlice)
	}

	sliceVazio := []int{}
	resultadoVazio, errVazio := AplicarFuncao(sliceVazio, func(x int) int {
		return x * 2
	})
	if errVazio != nil {
		fmt.Println("Erro:", errVazio)
	} else {
		fmt.Println("Resultado:", resultadoVazio)
	}
}

// q.8

func FiltrarNumerosPares(slice []int) ([]int, error) {
	if len(slice) == 0 {
		return nil, errors.New("Slice vazio")
	}

	numerosPares := make([]int, 0)
	for _, num := range slice {
		if num%2 == 0 {
			numerosPares = append(numerosPares, num)
		}
	}

	return numerosPares, nil
}

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	paresSlice, err := FiltrarNumerosPares(slice)
	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Println("Números pares:", paresSlice)
	}

	sliceVazio := []int{}
	resultadoVazio, errVazio := FiltrarNumerosPares(sliceVazio)
	if errVazio != nil {
		fmt.Println("Erro:", errVazio)
	} else {
		fmt.Println("Resultado:", resultadoVazio)
	}
}

// q.9

func ExtrairPalavras(str string) ([]string, error) {
	if str == "" {
		return nil, errors.New("String vazia")
	}

	palavras := strings.Fields(str)
	return palavras, nil
}

func main() {
	str := "Olá, mundo! Como vai você?"
	palavras, err := ExtrairPalavras(str)
	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Println("Palavras:", palavras)
	}

	strVazia := ""
	resultadoVazio, errVazio := ExtrairPalavras(strVazia)
	if errVazio != nil {
		fmt.Println("Erro:", errVazio)
	} else {
		fmt.Println("Resultado:", resultadoVazio)
	}
}

// q.10

func OrdenarSlice(slice []int) ([]int, error) {
	if len(slice) == 0 {
		return nil, errors.New("Slice vazio")
	}

	sliceOrdenado := make([]int, len(slice))
	copy(sliceOrdenado, slice)
	sort.Ints(sliceOrdenado)

	return sliceOrdenado, nil
}

func main() {
	slice := []int{5, 2, 8, 1, 9, 3}
	sliceOrdenado, err := OrdenarSlice(slice)
	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Println("Slice ordenado:", sliceOrdenado)
	}

	sliceVazio := []int{}
	resultadoVazio, errVazio := OrdenarSlice(sliceVazio)
	if errVazio != nil {
		fmt.Println("Erro:", errVazio)
	} else {
		fmt.Println("Resultado:", resultadoVazio)
	}
}

// q.11

func VerificarPrimo(numero int) (bool, error) {
	if numero < 0 {
		return false, errors.New("Número negativo")
	}

	if numero < 2 {
		return false, nil
	}

	limite := int(math.Sqrt(float64(numero)))
	for i := 2; i <= limite; i++ {
		if numero%i == 0 {
			return false, nil
		}
	}

	return true, nil
}

func main() {
	numero := 13
	primo, err := VerificarPrimo(numero)
	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Printf("%d é primo: %t\n", numero, primo)
	}

	numeroNegativo := -7
	primoNegativo, errNegativo := VerificarPrimo(numeroNegativo)
	if errNegativo != nil {
		fmt.Println("Erro:", errNegativo)
	} else {
		fmt.Printf("%d é primo: %t\n", numeroNegativo, primoNegativo)
	}
}

// q.12

func FiltrarStringsMaiusculas(slice []string) (string, error) {
	if len(slice) == 0 {
		return "", errors.New("Slice vazio")
	}

	var result strings.Builder
	for _, str := range slice {
		if len(str) > 0 && unicode.IsUpper(rune(str[0])) {
			result.WriteString(str)
			result.WriteString(" ")
		}
	}

	return strings.TrimSpace(result.String()), nil
}

func main() {
	slice := []string{"Hello", "world", "OpenAI", "GPT-3"}
	stringsMaiusculas, err := FiltrarStringsMaiusculas(slice)
	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Println("Strings maiúsculas:", stringsMaiusculas)
	}

	sliceVazio := []string{}
	resultadoVazio, errVazio := FiltrarStringsMaiusculas(sliceVazio)
	if errVazio != nil {
		fmt.Println("Erro:", errVazio)
	} else {
		fmt.Println("Resultado:", resultadoVazio)
	}
}

// q.13

func SomarDigitos(numero int) (int, error) {
	if numero < 0 {
		return 0, errors.New("Número negativo")
	}

	soma := 0
	for numero != 0 {
		digito := numero % 10
		soma += digito
		numero /= 10
	}

	return soma, nil
}

func main() {
	numero := 12345
	soma, err := SomarDigitos(numero)
	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Printf("A soma dos dígitos de %d é: %d\n", numero, soma)
	}

	numeroNegativo := -9876
	somaNegativa, errNegativa := SomarDigitos(numeroNegativo)
	if errNegativa != nil {
		fmt.Println("Erro:", errNegativa)
	} else {
		fmt.Printf("A soma dos dígitos de %d é: %d\n", numeroNegativo, somaNegativa)
	}
}

// q.14

func IntersecaoSlices(slice1, slice2 []int) ([]int, error) {
	if len(slice1) == 0 || len(slice2) == 0 {
		return nil, errors.New("Um dos slices está vazio")
	}

	intersecao := make([]int, 0)
	ocorrencias := make(map[int]bool)

	for _, num := range slice1 {
		ocorrencias[num] = true
	}

	for _, num := range slice2 {
		if ocorrencias[num] {
			intersecao = append(intersecao, num)
		}
	}

	return intersecao, nil
}

func main() {
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{4, 5, 6, 7, 8}
	intersecao, err := IntersecaoSlices(slice1, slice2)
	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Println("Interseção dos slices:", intersecao)
	}

	sliceVazio := []int{}
	resultadoVazio, errVazio := IntersecaoSlices(slice1, sliceVazio)
	if errVazio != nil {
		fmt.Println("Erro:", errVazio)
	} else {
		fmt.Println("Resultado:", resultadoVazio)
	}
}

// q.15

func VerificarPresenca(numero int, slice []int) (bool, error) {
	if len(slice) == 0 {
		return false, errors.New("Slice vazio")
	}

	for _, num := range slice {
		if num == numero {
			return true, nil
		}
	}

	return false, nil
}

func main() {
	numero := 5
	slice := []int{1, 2, 3, 4, 5}
	presente, err := VerificarPresenca(numero, slice)
	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Printf("O número %d está presente no slice: %t\n", numero, presente)
	}

	sliceVazio := []int{}
	resultadoVazio, errVazio := VerificarPresenca(numero, sliceVazio)
	if errVazio != nil {
		fmt.Println("Erro:", errVazio)
	} else {
		fmt.Println("Resultado:", resultadoVazio)
	}
}

// q.16

func SubstituirVogaisPorAsterisco(str string) (string, error) {
	if len(str) == 0 {
		return "", errors.New("String vazia")
	}

	vogais := []string{"a", "e", "i", "o", "u", "A", "E", "I", "O", "U"}

	for _, vogal := range vogais {
		str = strings.ReplaceAll(str, vogal, "*")
	}

	return str, nil
}

func main() {
	str := "Hello World"
	novaStr, err := SubstituirVogaisPorAsterisco(str)
	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Println("Nova string:", novaStr)
	}

	strVazia := ""
	resultadoVazio, errVazio := SubstituirVogaisPorAsterisco(strVazia)
	if errVazio != nil {
		fmt.Println("Erro:", errVazio)
	} else {
		fmt.Println("Resultado:", resultadoVazio)
	}
}

// q.17

func OrdenarStrings(slice []string) (string, error) {
	if len(slice) == 0 {
		return "", errors.New("Slice vazio")
	}

	sort.Strings(slice)
	novaStr := strings.Join(slice, " ")

	return novaStr, nil
}

func main() {
	slice := []string{"banana", "abacaxi", "laranja", "maçã"}
	novaStr, err := OrdenarStrings(slice)
	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Println("Nova string:", novaStr)
	}

	sliceVazio := []string{}
	resultadoVazio, errVazio := OrdenarStrings(sliceVazio)
	if errVazio != nil {
		fmt.Println("Erro:", errVazio)
	} else {
		fmt.Println("Resultado:", resultadoVazio)
	}
}

// q.18

func OrdenarStrings(slice []string) (string, error) {
	if len(slice) == 0 {
		return "", errors.New("Slice vazio")
	}

	sort.Strings(slice)
	novaStr := strings.Join(slice, " ")

	return novaStr, nil
}

func main() {
	slice := []string{"banana", "abacaxi", "laranja", "maçã"}
	novaStr, err := OrdenarStrings(slice)
	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Println("Nova string:", novaStr)
	}

	sliceVazio := []string{}
	resultadoVazio, errVazio := OrdenarStrings(sliceVazio)
	if errVazio != nil {
		fmt.Println("Erro:", errVazio)
	} else {
		fmt.Println("Resultado:", resultadoVazio)
	}
}

// q.19

func isPrimo(numero int) bool {
	if numero < 2 {
		return false
	}

	for i := 2; i*i <= numero; i++ {
		if numero%i == 0 {
			return false
		}
	}

	return true
}

func NumerosPrimos(numero int) ([]int, error) {
	if numero < 0 {
		return nil, errors.New("Número negativo")
	}

	primos := []int{}
	for i := 2; i <= numero; i++ {
		if isPrimo(i) {
			primos = append(primos, i)
		}
	}

	return primos, nil
}

func main() {
	numero := 20
	primos, err := NumerosPrimos(numero)
	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Println("Números primos até", numero, ":", primos)
	}

	numeroNegativo := -5
	primosNegativos, errNegativos := NumerosPrimos(numeroNegativo)
	if errNegativos != nil {
		fmt.Println("Erro:", errNegativos)
	} else {
		fmt.Println("Números primos até", numeroNegativo, ":", primosNegativos)
	}
}

// q.20

func FiltrarStrings(slice []string) ([]string, error) {
	if len(slice) == 0 {
		return nil, errors.New("Slice vazio")
	}

	filtradas := []string{}
	for _, str := range slice {
		if len(str) > 5 {
			filtradas = append(filtradas, str)
		}
	}

	return filtradas, nil
}

func main() {
	slice := []string{"apple", "banana", "orange", "kiwi", "mango"}
	filtradas, err := FiltrarStrings(slice)
	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Println("Strings filtradas:", filtradas)
	}

	sliceVazio := []string{}
	filtradasVazias, errVazio := FiltrarStrings(sliceVazio)
	if errVazio != nil {
		fmt.Println("Erro:", errVazio)
	} else {
		fmt.Println("Strings filtradas:", filtradasVazias)
	}
}
