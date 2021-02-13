---
title_meta: 'Глава 1'
title: 'Основы Python'
description: 'В этой главе мы расскажем вам о Python. Вы узнаете насколько он клевый.'
attachments:
slides_link: 'https://s3.amazonaws.com/assets.datacamp.com/production/course_735/slides/chapter1.pdf'
free_preview: true
---


## Интерфейс Python

```lesson
type: NormalExercise
key: bdc52f0e19
lang: python
xp: 100
skills: 2
```

In the Python script on the right, you can type Python code to solve the exercises. If you hit _Run Code_ or _Submit Answer_, your python script (`script.py`) is executed and the output is shown in the IPython Shell. _Submit Answer_ checks whether your submission is correct and gives you feedback.

You can hit _Run Code_ and _Submit Answer_ as often as you want. If you're stuck, you can click _Get Hint_, and ultimately _Get Solution_.

You can also use the IPython Shell interactively by simply typing commands and hitting Enter. When you work in the shell directly, your code will not be checked for correctness so it is a great way to experiment.

`@instructions`
- Experiment in the IPython Shell; type `5 / 8`, for example.
- Add another line of code to the Python script on the top-right (not in the Shell): `print(7 + 10)`.
- Hit _Submit Answer_ to execute the Python script and receive feedback.

`@hint`
Подсказка 1: Simply add `print(7 + 10)` in the script on the top-right (not in the Shell) and hit 'Submit Answer'.
Подсказка 2: `print(7 + 10)` in the script on the top-right (not in the Shell) and hit 'Submit Answer'.

`@pre_exercise_code`
```{python}
a = 1
```

`@sample_code`
```{python}
# Example, do not modify!
print(5 / 8)

# Print the sum of 7 and 10

```

`@solution`
```{python}
# Example, do not modify!
print(5 / 8)

# Put code below here
print(7 + 10)
```

`@test`
```{python}
Ex().has_printout(1, not_printed_msg = "__JINJA__:Have you used `{{sol_call}}` to print out the sum of 7 and 10?")
success_msg("Great! On to the next one!")
```

---

## Когда нужно использовать Python?

```lesson
type: MultipleChoiceExercise
key: 9703b117fb
lang: python
xp: 50
skills: 2
```

Python is a pretty versatile language. For which applications can you use Python?

`@possible_answers`
- You want to do some quick calculations.
- For your new business, you want to develop a database-driven website.
- Your boss asks you to clean and analyze the results of the latest satisfaction survey.
- All of the above.

`@answer`
3

`@hint`
Hugo mentioned in the video that Python can be used to build practically any piece of software.

`@pre_exercise_code`
```{python}

```

`@test`
```{python}
msg1 = "Incorrect. Python can do simple and quick calculations, but it is much more than that!"
msg2 = "Incorrect. There is a very popular framework to build database-driven websites (Django), but Python can do much more."
msg3 = "Incorrect. Python is a powerful tool to do data analysis, but you can also use it for other ends."
msg4 = "Correct! Python is an extremely versatile language."
Ex().has_chosen(4, [msg1, msg2, msg3, msg4])
```

