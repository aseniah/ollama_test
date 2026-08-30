import sys

def main():
    if len(sys.argv) != 2:
        return
    
    try:
        n = int(sys.argv[1])
    except ValueError:
        return
    
    if n < 1:
        return
    
    # Generate Fibonacci numbers up to and including the largest one <= n
    # Start with 1, 1, 2, 3, 5, ...
    fib1 = 1
    fib2 = 1
    
    print(fib1)
    
    if n >= 1:
        print(fib2)
    
    a, b = 1, 1
    while True:
        a, b = b, a + b
        if b > n:
            break
        print(b)

if __name__ == "__main__":
    main()