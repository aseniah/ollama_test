import sys

def fibonacci(n):
    a, b = 1, 1
    while a <= n:
        yield a
        a, b = b, a + b

if len(sys.argv) != 2:
    print("Usage: python script.py <integer>")
else:
    try:
        N = int(sys.argv[1])
        if N < 1:
            pass
        else:
            for num in fibonacci(N):
                print(num)
    except ValueError:
        print("Please provide a valid integer.")