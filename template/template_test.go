package template

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Test struct {
	name         string
	verifyResult func(*testing.T, *ExampleType, string)
}

type Bench struct {
	name         string
	verifyResult func(*testing.B, *ExampleType, string)
}

/*
Основной шаблон для написания тестов
*/
func Test_MainTemplate(t *testing.T) {
	testData := map[string][]string{
		// тестовые значения
	}

	expectedData := map[string][]string{
		// ожидаемые значения
	}

	tests := []Test{
		{
			name: "test name",
			verifyResult: func(t *testing.T, p *ExampleType, testName string) {
				// сценарий теста и сравнение результатов с данными из expectedData
			},
		},
	}

	for _, test := range tests {
		p := New()
		t.Run(test.name, func(t *testing.T) {
			test.verifyResult(t, p, test.name)
		})
	}
}

/*
Основной шаблон для написания бенчмарков
*/
func Benchmark_MainTemplate(b *testing.B) {
	testData := map[string][]string{
		// тестовые значения
	}

	tests := []Bench{
		{
			name: "bench name",
			verifyResult: func(b *testing.B, p *ExampleType, testName string) {
				// вызов тестируемой функции для замера быстродействия
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
				test.verifyResult(b, p, test.name)
			}
		})
	}
}

/*
Пример использования шаблона для написания теста
*/
func Test_SetFieldString(t *testing.T) {
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

	tests := []Test{
		{
			name: "valid",
			verifyResult: func(t *testing.T, p *ExampleType, testName string) {
				for i := range testData[testName] {
					err := p.SetFieldString(testData[testName][i])
					if err != nil {
						t.Error(err)
					}
					assert.Equal(t, expectedData[testName][i], p.FieldString)
				}
			},
		},
		{
			name: "invalid",
			verifyResult: func(t *testing.T, p *ExampleType, testName string) {
				for i := range testData[testName] {
					err := p.SetFieldString(testData[testName][i])
					assert.Error(t, err)
					assert.Equal(t, "", p.FieldString)
				}
			},
		},
	}

	for _, test := range tests {
		p := New()
		t.Run(test.name, func(t *testing.T) {
			test.verifyResult(t, p, test.name)
		})
	}
}

/*
Пример использования шаблона для написания бенчмарков
*/
func Benchmark_SetFieldString(b *testing.B) {
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

	tests := []Bench{
		{
			name: "valid",
			verifyResult: func(b *testing.B, p *ExampleType, testName string) {
				for i := range testData[testName] {
					p.SetFieldString(testData[testName][i])
				}
			},
		},
		{
			name: "invalid",
			verifyResult: func(b *testing.B, p *ExampleType, testName string) {
				for i := range testData[testName] {
					p.SetFieldString(testData[testName][i])
				}
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
				test.verifyResult(b, p, test.name)
			}
		})
	}
}
