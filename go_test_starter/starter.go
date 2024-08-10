package starter

import (
	"fmt"
	"math"
	"net/http"
)

// As funções desse arquivo são inicialmente criadas de forma mínima, conforme o ciclo "Green" do TDD, que visa apenas satisfazer as asserções do teste, sem se preocupar com otimizações.

// SayHello é uma função simples que retorna uma saudação personalizada.
func SayHello(name string) string {
	return fmt.Sprintf("Hello %v. Welcome!", name)
// A implementação dessa função foi guiada pelos testes que exigiam uma saudação específica.
}

// OddOrEven verifica se um número é par ou ímpar e retorna uma string correspondente.
// Critérios adicionais ou otimizações podem ser considerados na fase de refatoração (Refactor).
func OddOrEven(num int) string {
	criteria := math.Mod(float64(num), 2)
// Usamos math.Mod para calcular o resto da divisão do número por 2.
// Se o resto for 1 ou -1, o número é ímpar.
	if criteria == 1 || criteria == -1 {
		return fmt.Sprintf("%v is an odd number", num)
	}
// Caso contrário, o número é par.
	return fmt.Sprintf("%v is an even number", num)
}

// Checkhealth é uma função que simula um endpoint de verificação de saúde.
func Checkhealth(writer http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(writer, "health check passed")
// A função responde com a string "health check passed" para indicar que o serviço está funcionando.

}