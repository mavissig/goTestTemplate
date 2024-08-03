package template_tests_and_benchmarks

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type TestAndBench struct {
	name          string
	performAction func(*ExampleType, string)
	verifyResult  func(testing.TB, bool, *ExampleType, string)
}

/*
Враппер для запуска общей тестовой функции если запускаются тесты
*/
func Test_MainTemplate(t *testing.T) {
	TBMainTemplate(t, nil)
}

/*
Враппер для запуска общей тестовой функции если запускаются тесты
*/
func Benchmark_MainTemplate(b *testing.B) {
	TBMainTemplate(nil, b)
}

/*
Шаблон общей тестовой функции, которая позволяет запускать как тесты так и бенчмарки
*/
func TBMainTemplate(t *testing.T, b *testing.B) {
	testData := map[string]any{
		// Тестовые значения
	}

	expectedData := map[string]any{
		// Ожидаемые значения
	}

	tests := []TestAndBench{
		{
			name: "test name",
			verifyResult: func(t testing.TB, bench bool, p *ExampleType, testName string) {
				// вызов тестируемого кода
				if !bench {
					// Проверка результатов теста если запущен НЕ БЕНЧМАРК
				}
			},
		},
	}

	for _, test := range tests {
		p := New()
		if t != nil {
			t.Run(test.name, func(t *testing.T) {
				test.verifyResult(t, false, p, test.name)
			})
		} else if b != nil {
			b.ResetTimer()
			b.Run(test.name, func(b *testing.B) {
				b.StopTimer()
				b.StartTimer()
				for i := 0; i < b.N; i++ {
					test.verifyResult(b, true, p, test.name)
				}
			})
		}
	}
}

// Пример использования шаблона для тестирования сеттера строки

func Test_SetFieldString(t *testing.T) {
	TBSetFieldString(t, nil)
}

func Benchmark_SetFieldString(b *testing.B) {
	TBSetFieldString(nil, b)
}

func TBSetFieldString(t *testing.T, b *testing.B) {
	testData := map[string][]string{
		"valid": {
			"valid string",
			"1234567890",
		},
		"invalid": {
			"",
			"gg",
			"invalid",
		},
	}

	expectedData := map[string][]string{
		"valid": {
			"valid string",
			"1234567890",
		},
		"invalid": {},
	}

	tests := []TestAndBench{
		{
			name: "valid",
			verifyResult: func(t testing.TB, bench bool, p *ExampleType, testName string) {
				for i := range testData[testName] {
					err := p.SetFieldString(testData[testName][i])
					if err != nil {
						t.Errorf("err: %v\n", err)
					}

					if !bench {
						assert.Equal(t, expectedData[testName][i], p.FieldString)
					}
				}
			},
		},
		{
			name: "invalid",
			verifyResult: func(t testing.TB, bench bool, p *ExampleType, testName string) {
				for i := range testData[testName] {
					err := p.SetFieldString(testData[testName][i])
					if !bench {
						assert.Error(t, err)
						assert.Equal(t, "", p.FieldString)
					}
				}
			},
		},
	}

	for _, test := range tests {
		p := New()
		if t != nil {
			t.Run(test.name, func(t *testing.T) {
				test.verifyResult(t, false, p, test.name)
			})
		} else if b != nil {
			b.ResetTimer()
			b.Run(test.name, func(b *testing.B) {
				b.StopTimer()
				b.StartTimer()
				for i := 0; i < b.N; i++ {
					test.verifyResult(b, true, p, test.name)
				}
			})
		}
	}
}
