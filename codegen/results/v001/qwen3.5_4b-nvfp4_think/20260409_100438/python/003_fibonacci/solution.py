import sys

if len(sys.argv) != 2 or not sys.argv[1].isdigit():
    print("Usage: python fibonacci.py <number>")
    sys.exit(1)

n = int(sys.argv[1])

if n < 1:
    pass  # No output for negative or zero N
else:
    a, b = 1, 1
    fib_numbers = [a] if n >= 0 else []
    while b <= n:
        fib_numbers.append(b)
        a, b = b, a + b
        if b > n:
            break
    
    for num in fib_numbers:
        print(num)