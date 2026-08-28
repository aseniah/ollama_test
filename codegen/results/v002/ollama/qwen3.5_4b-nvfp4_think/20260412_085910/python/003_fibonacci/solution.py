import sys

def main():
    try:
        n = int(sys.argv[1])
    except (IndexError, ValueError):
        return
    
    if n < 1:
        return
    
    # Generate Fibonacci sequence
    fib_a, fib_b = 1, 1
    
    while True:
        print(fib_a)
        next_fib = fib_a + fib_b
        if next_fib > n:
            break
        fib_a, fib_b = fib_b, next_fib

if __name__ == "__main__":
    main()