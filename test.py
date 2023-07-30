import traceback

def fib(n : int) -> int:
    print(n)
    traceback.print_stack()
    if n <= 0:
        return 0
    elif n == 1:
        return 1
    else:
        return fib(n-1) + fib(n-2)

fib(5)