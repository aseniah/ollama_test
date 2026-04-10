import sys

if __name__ == "__main__":
    if len(sys.argv) != 2:
        print("Usage: python script.py N")
        sys.exit(1)
    
    try:
        n = int(sys.argv[1])
    except ValueError:
        print("Error: N must be an integer")
        sys.exit(1)
    
    if n < 1:
        sys.exit(0)
    
    fib1, fib2 = 1, 1
    while fib1 <= n:
        print(fib1)
        fib1, fib2 = fib2, fib1 + fib2