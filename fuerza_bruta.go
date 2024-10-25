package main

import (
	"fmt"
	"math/rand"
	"os/exec"
	"sync"
	"time"
)

// Función para generar todas las combinaciones posibles de contraseñas de una longitud dada
func generatePasswords(length int) []string {
	var passwords []string
	var generate func(current string)

	generate = func(current string) {
		if len(current) == length {
			passwords = append(passwords, current)
			return
		}
		for i := 'a'; i <= 'z'; i++ {
			generate(current + string(i))
		}
	}

	generate("")
	return passwords
}

// decryptFile descifra el archivo utilizando gpg
func decryptFile(password, inputFile, outputFile string) error {
	cmd := exec.Command("gpg", "--batch", "--yes", "--passphrase", password, "-o", outputFile, "-d", inputFile)
	return cmd.Run()
}

// worker es la función que cada goroutine ejecutará
func worker(passwords <-chan string, inputFile, outputFile string, done chan bool, wg *sync.WaitGroup, threadID int) {
	defer wg.Done()

	for {
		select {
		case <-done:
			fmt.Printf("Hilo %d finalizado por señal.\n", threadID)
			return
		case password, ok := <-passwords:
			if !ok {
				fmt.Printf("Hilo %d no hay más contraseñas que probar. Finalizando.\n", threadID)
				return
			}

			fmt.Printf("Hilo %d - Probando contraseña: '%s'\n", threadID, password)

			err := decryptFile(password, inputFile, outputFile)
			if err == nil {
				fmt.Printf("Hilo %d encontró la contraseña correcta: '%s'\n", threadID, password)
				done <- true
				return
			}
		}
	}
}

func main() {
	inputFile := "archive.pdf.gpg"
	outputFile := "archivo_descifrado.pdf"

	pass_max_length := 6
	totalThreads := 12
	var wg sync.WaitGroup
	done := make(chan bool)
	passwords := make(chan string)

	startTime := time.Now()

	for length := 4; length <= pass_max_length; length++ {
		passwordList := generatePasswords(length)

		rand.Seed(time.Now().UnixNano())
		rand.Shuffle(len(passwordList), func(i, j int) {
			passwordList[i], passwordList[j] = passwordList[j], passwordList[i]
		})

		go func() {
			for _, password := range passwordList {
				passwords <- password
			}
			close(passwords)
		}()

		for i := 0; i < totalThreads; i++ {
			wg.Add(1)
			go worker(passwords, inputFile, outputFile, done, &wg, i+1)
		}

		if <-done {
			break
		}

		passwords = make(chan string)
	}

	close(done)
	wg.Wait()

	totalTime := time.Since(startTime)

	fmt.Printf("\nProceso completado. Tiempo total: %.2f segundos.\n", totalTime.Seconds())
}
