import sys

def main():
    try:
        N = int(sys.argv[1])
    except (IndexError, ValueError):
        return
    
    if N < 1:
        return
    
    # Generate Fibonacci numbers up to and including the largest <= N
    fibs = []
    a, b = 1, 1
    # First fib is 1
    if a <= N:
        fibs.append(a)
        # next
        a, b = b, a + b
        while a <= N:
            fibs.append(a)
            a, b = b, a + b
    
    for f in fibs:
        print(f)

if __name__ == "__main__":
    main()