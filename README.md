# lesparse - парсер глав и уроков из MarkDown

## Использование

	p := lesparse.NewParser()

	f, err := os.Open("example.md")
	if err != nil {
		panic(err)
	}

	content, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}

	chapter := p.Chapter(string(content))

	b, _ := json.MarshalIndent(chapter, "  ", "  ")
	fmt.Println(string(b))

`p.Chapter(string(content))` возвращает объект главы. 
В дальнейших инструкциях под именем 'chapter' будет определяться именно он.

## Формат MarkDown

Каждый файл '*.md' - это глава.

### Заголовок

Глава должна содержать заголовок, обрамленный знаками '---' с новой строки.
Сам заголовок имеет следующую структуру:

* title - название главы
* title_meta - мета тег
* description - описание главы
* attachments - список прикрепленных файлов
* slides_link - список слайдов, показываемых после прохождения главы
* free_preview - определяет доступность главы для гостя

Пример:

    ---
    title_meta: 'Глава 1'
    title: 'Основы Python'
    description: 'В этой главе мы расскажем вам о Python. Вы узнаете насколько он клевый.'
    attachments:
    slides_link: 'https://s3.amazonaws.com/assets.datacamp.com/production/course_735/slides/chapter1.pdf'
    free_preview: true
    ---

После осуществления парсинга все поля заголовка доступны в:

    chapter.Header 

### Урок

Глава состоит из набора уроков. Урок имеет человекочитаемый заголовок, например:

    ## Интерфейс Python

После вызова парсинга можно получить доступ к оригинальному телу Markdown:

    chapter.Lessions[0].Markdown

#### Урок: заголовок

Каждый урок имеет свой заголовок, который обрамляется "```". 
Сразу после символов этих символов, без пробелов, встречается ключевое слово "lesson".
Структура заголовка урока должна содержать следующие данные:

* type - тип урока (ID для среды исполнения)
* lang - язык программирования (платформа)
* xp - количество получаемого опыта за прохождение урока
* skills - уровень сложности урока

Пример:

```lesson
type: NormalExercise
lang: python
xp: 100
skills: 2
```

Заголовок урока доступен в объекте "Lessons":

    chapter.Lessons[0].Lang
    chapter.Lessons[0].XP
    ...

#### Урок: основной текст

После заголовка можно использовать любой текст в формате MarkDown.


#### Урок: инструкции

Блок инструкций определяется ключевым словом:

    `@instructions`

Окончание блока определяется по наличию пустой строки после.

Например:

    `@instructions`
    - Experiment in the IPython Shell; type `5 / 8`, for example.
    - Add another line of code to the Python script on the top-right (not in the Shell): `print(7 + 10)`.
    - Hit _Submit Answer_ to execute the Python script and receive feedback.

Инструкции доступны в объекте "Lesson" уже отрендеренные в HTML:

    chapter.Lessons[0].HTML.Instructions


#### Урок: подсказки

Блок подсказок определяется ключевым словом:

    `@hint`

Окончание блока определяется по наличию пустой строки после.
Каждая новая строка считается новой подсказкой!

Например:

    `@hint`
    Simply add `print(7 + 10)` in the script on the top-right (not in the Shell) and hit 'Submit Answer'.

Подсказки доступны в объекте "Lesson" собранные в слайсы и отрендеренные в HTML:

    chapter.Lessons[0].HTML.Hints[0]


#### Урок: предопределенный код

Предопределенный код выполняется всегда до кода пользователя и до кода тестов.

Каждый блок кода предопределенного кода определяется по ключевому:

    `@pre_exercise_code`


Сразу после идентификатора блока должны находиться границы блока кода, обрамленные "```".

Конец блока определяются по повторному использованию "```".
Также, после этих символов позволяется использовать идентификатор языка в фигурных скобках "{python}", но он никак не используется объектами.

Пример:

    `@pre_exercise_code`
    ```{python}
    a = 1
    ```

Структура кода находится в объекте Code:

    chapter.Lessons[0].Code.PEC


#### Урок: пример кода

Каждый блок кода примера кода определяется по ключевому:

    `@sample_code`


Сразу после идентификатора блока должны находиться границы блока кода, обрамленные "```".

Конец блока определяются по повторному использованию "```".
Также, после этих символов позволяется использовать идентификатор языка в фигурных скобках "{python}", но он никак не используется объектами.

Пример:

    `@sample_code`
    ```{python}
    # Example, do not modify!
    print(5 / 8)
    
    # Print the sum of 7 and 10
    
    ```

Структура кода находится в объекте Code:

    chapter.Lessons[0].Code.Sample


#### Урок: решение теста

Каждый блок кода решения теста определяется по ключевому:

    `@solution`


Сразу после идентификатора блока должны находиться границы блока кода, обрамленные "```".

Конец блока определяются по повторному использованию "```".
Также, после этих символов позволяется использовать идентификатор языка в фигурных скобках "{python}", но он никак не используется объектами.

Пример:

    `@solution`
    ```{python}
    # Example, do not modify!
    print(5 / 8)
    
    # Put code below here
    print(7 + 10)
    ```

Структура кода находится в объекте Code:

    chapter.Lessons[0].Code.Solution


#### Урок: тест


Каждый блок теста определяется по ключевому:

    `@test`


Сразу после идентификатора блока должны находиться границы блока кода, обрамленные "```".

Конец блока определяются по повторному использованию "```".
Также, после этих символов позволяется использовать идентификатор языка в фигурных скобках "{python}", но он никак не используется объектами.

Пример:
    
    `@test`
    ```{python}
    Ex().has_printout(1, not_printed_msg = "__JINJA__:Have you used `{{sol_call}}` to print out the sum of 7 and 10?")
    success_msg("Great! On to the next one!")
    ```

Структура кода находится в объекте Code:

    chapter.Lessons[0].Code.Test


#### Урок: вопросы и ответ

Каждый блок теста определяется по ключевому:

    `@possible_answers`

Блок индекса ответа (индекс от 1):

    `@answer`

Конец блока определяются по пустой строке после.

Пример:

    `@possible_answers`
    - You want to do some quick calculations.
    - For your new business, you want to develop a database-driven website.
    - Your boss asks you to clean and analyze the results of the latest satisfaction survey.
    - All of the above.
    
    `@answer`
    3

Структура возможных ответов отрендерена в:

    chapter.Lessons[0].HTML.PossibleAnswers

Индекс ответа в:

    chapter.Lessons[0].AnswerIndex
