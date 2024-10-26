package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	fmt.Println("Jogo da adivinhação!")
	fmt.Println("Advinhe o número de 1 a 100!")

	rand.Seed(time.Now().UnixNano())
	x := rand.Intn(100) + 1

	scanner := bufio.NewScanner(os.Stdin)
	chutes := [10]int64{}

	for i := range chutes {
		fmt.Println("Qual seu chute?")
		scanner.Scan()

		chute := strings.TrimSpace(scanner.Text())
		chuteInt, err := strconv.ParseInt(chute, 10, 64)
		if err != nil {
			fmt.Println("Por favor, insira um número válido!")
			i--
			continue
		}

		// Comparação e resposta ao jogador
		switch {
		case chuteInt < int64(x):
			fmt.Println("Você errou! Seu chute foi menor do que o número.")
		case chuteInt > int64(x):
			fmt.Println("Você errou! Seu chute foi maior do que o número.")
		case chuteInt == int64(x):
			fmt.Printf("Parabéns! Você acertou! O número era %d.\n", x)
			fmt.Printf("Você acertou em %d tentativa(s).\n", i+1)
			fmt.Printf("Sua história de chutes: %v\n", chutes[:i+1])
			return
		}

		chutes[i] = chuteInt
	}

	// Se todas as tentativas forem usadas sem sucesso
	fmt.Printf("Infelizmente, você não acertou! O número era %d.\n", x)
	fmt.Printf("História de chutes: %v\n", chutes)
}
