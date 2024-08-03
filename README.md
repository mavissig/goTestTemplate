# Go Test Template

## Content

- [Preamble](#preamble)
- [Introduction](#introduction)
- [Template](#template)
- [Template for scripts](#template-for-scripts)
- [Template for test and benchmark](#template-for-test-and-benchmark)

## Preamble

![gtt_intro_img](misc/images/intro.PNG)

В этом репозитории я хотел бы показать шаблон для тестов, которым я
пользуюсь в своей повседневной жизни.

Подробная статья о шаблоне присутствует на Habr, её вы можете [прочитать здесь](https://habr.com/ru/articles/833448/).
## Introduction

### Проблема тестирования

Основные проблемы большего количества тестов, которые я наблюдаю, 
заключаются в том, что тесты не имеют четкой структуры, которую можно было бы легко понять и выделить 
каждую часть. Присутствует сложность чтения, особенно страшно выглядят тесты когда перед тестом исполняется 
какой-либо массивный сценарий, а так же тесты не являются модульными и как следствие их тяжело поддерживать и расширять.

### Как я к этому пришел

Сразу скажу, я очень люблю когда код можно легко масштабировать. В тестах я решил не отступать от своих взглядов и 
принял решение, что паттерн должен легко использоваться не только для тестирования простых функций и методов, но и 
для сложных сценариев, а так же хотел сделать вариант с использованием одного блока кода для тестов и бенчмарков.

Требования, которые выдвигались мной при выборе архитектуры тестов:
- максимально возможная читаемость и простота кода
- возможность переиспользования кода
- модульность и гибкость
- удобство при отладке
- хорошая шаблонная структура

Нужно понимать, что общий вид тестов может меняться в зависимости от потребностей, но в целом он сохраняет 
свою архитектуру.

### Какие плюсы от использования этого паттерна
- **Повторное использование кода.**
Возможность повторного использования общих действий и проверок между разными тестами. Это уменьшает дублирование 
кода и упрощает внесение изменений.
- **Модульность и гибкость.**
Модульный подход позволяет легко добавлять новые тесты с минимальными изменениями в существующем коде. Например, 
добавление новых действий или проверок без затрагивания основной логики тестирования.
- **Улучшение читаемости.**
Тесты становятся более читаемыми и понятными, так как каждая часть теста отвечает за конкретную задачу. Это 
облегчает обзор и анализ тестов, особенно при большом количестве различных сценариев.
- **Структурированное тестирование.**
Этот паттерн помогает структурировать тесты, что особенно важно в больших проектах с множеством сценариев. 
Хорошо организованные тесты облегчают их поддержку и расширение.
- **Повышение надежности тестов.**
Четкое разделение действий и проверок снижает риск ошибок в тестах. Это помогает избежать ситуаций, когда одно 
действие случайно влияет на проверку другого сценария.
- **Хорошая организация кода.**
Такой подход способствует лучшей организации кода, что облегчает его сопровождение и улучшение. Структурированные 
тесты легче модифицировать и расширять.

### Структура документации

В README показаны шаблоны в зависимости от требуемого функционала. Примеры использования каждого шаблона 
присутствуют в `./<template-name>/pattern_test.go`

## Template

Директория с базовым шаблоном [ЗДЕСЬ](./template/template_test.go).

Основной шаблон реализует тесты согласно структуре 

```go
type Test struct {
	name         string
	verifyResult func(*testing.T, *ExampleType, string)
}
```

Бенчмарки реализуются согласно структуре:

```go
type Bench struct {
	name          string
	verifyResult  func(*testing.B, *ExampleType, string)
}
```

Базовый шаблон на котором можно реализовать большинство тестовых кейсов выглядит так:
```go
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
```

Бенчмарки на этом шаблоне пишутся следующим образом:
```go
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
```

## Template for scripts

Директория с шаблоном для тестов со сценарием [ЗДЕСЬ](./template-for-scripts/template_test.go).

Шаблон для тестов с предварительным сценарием использует структуру:

```go
type Test struct {
	name          string
	performAction func(*ExampleType, string)
	verifyResult  func(*testing.T, *ExampleType, string)
}
```

И структура для бенчмарков:

```go
type Bench struct {
	name          string
	performAction func(*ExampleType, string)
	verifyResult  func(*testing.B, *ExampleType, string)
}
```

Сам шаблон для тестов со сценарием:

```go
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
```

Для бенчмарков он тоже подходит идеально:

```go
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
```

## Template for test and benchmark

Директория с универсальным шаблоном [ЗДЕСЬ](./template-tests-and-benchmarks/template_test.go).

Шаблон для запуска тестов и бенчмарков на базе одного кода использует структуру:

```go
type TestAndBench struct {
	name          string
	performAction func(*ExampleType, string)
	verifyResult  func(testing.TB, bool, *ExampleType, string)
}
```

Сам же шаблон выглядит так:

```go
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
```