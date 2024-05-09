# Калькулятор для академии Ката

Этот проект представляет собой простой калькулятор, разработанный в рамках тестового задания для академии Ката. Калькулятор выполняет основные арифметические операции (сложение, вычитание, умножение, деление) над целыми числами в диапазоне от 1 до 10. Он также поддерживает операции как с арабскими, так и с римскими числами.

## Функциональность

1. **Сложение (+):** Складывает два числа и выводит результат.
2. **Вычитание (-):** Вычитает второе число из первого и выводит результат.
3. **Умножение (*):** Умножает два числа и выводит результат.
4. **Деление (/):** Делит первое число на второе и выводит результат.

Проект также обрабатывает ошибки ввода, включая:
- Некорректные операции.
- Ввод чисел вне допустимого диапазона (от 1 до 10).
- Некорректные форматы ввода (например, римские числа должны быть в верхнем регистре и не содержать пробелов между символами).


## Требования 
1. Калькулятор умеет выполнять операции сложения, вычитания, умножения и деления с двумя числами: a + b, a - b, a * b, a / b. Данные передаются в одну строку (смотри пример ниже). Решения, в которых каждое число и арифметическая операция передаются с новой строки, считаются неверными.

2. Калькулятор умеет работать как с арабскими (1, 2, 3, 4, 5…), так и с римскими (I, II, III, IV, V…) числами.

3. Калькулятор должен принимать на вход числа от 1 до 10 включительно, не более. На выходе числа не ограничиваются по величине и могут быть любыми.

4. Калькулятор умеет работать только с целыми числами.

5. Калькулятор умеет работать только с арабскими или римскими цифрами одновременно, при вводе пользователем строки вроде 3 + II калькулятор должен выдать панику и прекратить работу.

6. При вводе римских чисел ответ должен быть выведен римскими цифрами, соответственно, при вводе арабских — ответ ожидается арабскими.

7. При вводе пользователем не подходящих чисел приложение выдаёт панику и завершает работу.

8. При вводе пользователем строки, не соответствующей одной из вышеописанных арифметических операций, приложение выдаёт панику и завершает работу.

9. Результатом операции деления является целое число, остаток отбрасывается.

10. Результатом работы калькулятора с арабскими числами могут быть отрицательные числа и ноль. Результатом работы калькулятора с римскими числами могут быть только положительные числа, если результат работы меньше единицы, программа должна выдать панику.
