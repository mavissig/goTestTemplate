package template_for_scripts

import (
	"testing"
)

type Test struct {
	name          string
	performAction func(*ExampleType, string)
	verifyResult  func(*testing.T, *ExampleType, string)
}

type Bench struct {
	name          string
	performAction func(*ExampleType, string)
	verifyResult  func(*testing.B, *ExampleType, string)
}

/*
Основной шаблон для написания тестов с предварительным сценарием
*/
func Test_MainTemplate(t *testing.T) {
	testData := map[string]any{
		// Тестовые значения
	}

	expectedData := map[string]any{
		// Ожидаемые значения
	}

	tests := []Test{
		{
			name: "test name",
			performAction: func(*ExampleType, string) {
				/*
					Выполнение предварительных действий
				*/
			},
			verifyResult: func(t *testing.T, p *ExampleType, testName string) {
				/*
					Проверка результатов теста
				*/
			},
		},
	}

	for _, test := range tests {
		p := New()
		t.Run(test.name, func(t *testing.T) {
			test.performAction(p, test.name)
			test.verifyResult(t, p, test.name)
		})
	}
}

/*
Основной шаблон для написания бенчмарков с предварительным сценарием
*/
func Benchmark_MainTemplate(b *testing.B) {
	testData := map[string]any{
		// Тестовые значения
	}

	expectedData := map[string]any{
		// Ожидаемые значения
	}

	tests := []Bench{
		{
			name: "bench name",
			performAction: func(*ExampleType, string) {
				/*
					Выполнение предварительных действий
				*/
			},
			verifyResult: func(b *testing.B, p *ExampleType, testName string) {
				/*
					Вызов тестируемой функции для замера быстродействия
				*/
			},
		},
	}

	for _, test := range tests {
		p := New()
		b.ResetTimer()
		b.Run(test.name, func(b *testing.B) {
			b.StopTimer()
			b.StartTimer()
			for i := 0; i < b.N; i++ {
				test.performAction(p, test.name)
				test.verifyResult(b, p, test.name)
			}
		})
	}
}
