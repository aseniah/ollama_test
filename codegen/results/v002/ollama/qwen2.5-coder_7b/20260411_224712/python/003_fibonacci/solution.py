import sys

def fibonacci(n):
    if n <= 0:
        return []
    elif n == 1:
        return [1]
    
    fibs = [1, 1]
    while True:
        next_fib = fibs[-1] + fibs[-2]
        if next_fib > n:
            break
        fibs.append(next_fib)
    
    return fibs

if __name__ == "__main__":
    if len(sys.argv) != 2:
        print("Usage: python script.py <integer>")
        sys.exit(1)
    
    try:
        N = int(sys.argv[1])
    except ValueError:
        print("Please provide a valid integer.")
        sys.exit(1)
    
    fib_numbers = fibonacci(N)
    for num in fib_numbers:
        print(num)