# Integrator

### Brief
This program implements an algorithm mentioned in the book "Introduction to the parallel problem solving methods" by M. V. Jacobovsky, pages 269-273 !["Введение в параллельные методы решения задач"](https://znanium.ru/catalog/document?id=339999&ysclid=m81lcohejt357176825).

### Problem

The problem is to evaluate:

$\int_a^b{f(x)dx}$

Only one of the endpoint can be a special one.

### Brief description of the method
This method is called "The global stack method". It provides dynamic task balancing between threads (goroutines).
It allows to evaluate the integral fastly and effectively.

1) Thread checks a global stack of tasks for avaliable tasks
2) It checks the interval. 
    - If the interval is too big to provide declared accuracy, it splits it by two parts. One of them is sent to the global stack, for the second one the second step is repeated.
    - If the interval is not to big, the value of integral is counted by trapezoid method.
3) The counted value is added to the global sum variable
4) After all tasks finished the integral value could be read from global sum variable.

