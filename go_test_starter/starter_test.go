package starter_test

import (
	"io"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	starter "github.com/williaminfante/go_test_starter"
)

// Os testes desse arquivo representam a fase "Red", onde asseguramos que a função
// ainda não implementada ou não testada se comporte conforme o esperado.

// TestSayHello testa a função SayHello, garantindo que ela retorne a saudação correta.
// Aqui, verificamos se o retorno da função SayHello corresponde ao esperado para entradas diferentes.
func TestSayHello(t *testing.T) {
	greeting := starter.SayHello("William")
	assert.Equal(t, "Hello William. Welcome!", greeting)
// O teste acima falharia se a implementação de SayHello não retornasse a saudação correta.
	another_greeting := starter.SayHello("asdf ghjkl")
	assert.Equal(t, "Hello asdf ghjkl. Welcome!", another_greeting)
}

// TestOddOrEven testa a função OddOrEven para verificar se ela classifica corretamente números como pares ou ímpares. 
func TestOddOrEven(t *testing.T) {
// O método t.Run permite a execução de subtestes, cada um testando diferentes casos de uso. Isso é bem útil no TDD para cobrir múltiplos cenários com um único teste principal.
	t.Run("Check Non Negative Numbers", func(t *testing.T) {
		assert.Equal(t, "45 is an odd number", starter.OddOrEven(45))
		assert.Equal(t, "42 is an even number", starter.OddOrEven(42))
		assert.Equal(t, "0 is an even number", starter.OddOrEven(0))
	})
// Cada asserção verifica se a função OddOrEven retorna a string correta para números não negativos.
// O teste acima falharia se a implementação de OddOrEven não classificasse corretamente os números como pares ou ímpares.
	t.Run("Check Negative Numbers", func(t *testing.T) {
		assert.Equal(t, "-45 is an odd number", starter.OddOrEven(-45))
		assert.Equal(t, "-42 is an even number", starter.OddOrEven(-42))
	})
// O teste acima verifica se a função OddOrEven lida corretamente com números negativos.
}

// TestCheckhealth testa a função Checkhealth, que simula uma verificação de saúde de um serviço.
// Esse teste usa um servidor HTTP simulado para testar o endpoint, verificando tanto o conteúdo da resposta quanto o status HTTP retornado.
func TestCheckhealth(t *testing.T) {
	t.Run("Check health status", func(t *testing.T) {
		req := httptest.NewRequest("GET", "http://mysite.com/example", nil)
// O método httptest.NewRequest cria uma nova solicitação HTTP com o método GET e o URL fornecido.
		writer := httptest.NewRecorder()
// O método httptest.NewRecorder cria um gravador de resposta HTTP simulado.
		starter.Checkhealth(writer, req)
// A função Checkhealth é chamada com o gravador e a solicitação simulada.
		response := writer.Result()
		body, err := io.ReadAll(response.Body)

		assert.Equal(t, "health check passed", string(body))
// A resposta da função Checkhealth é verificada para garantir que o corpo da resposta seja "health check passed".
		assert.Equal(t, 200, response.StatusCode)
// verifica se o status para garantir que seja 200.
		assert.Equal(t,
			"text/plain; charset=utf-8",
			response.Header.Get("Content-Type"))
// verifica se não tem erro no corpo da respota.

		assert.Equal(t, nil, err)
	})
}