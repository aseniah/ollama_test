import sys

def main():
    # Read N from command line
    args = sys.argv[1:]
    if not args:
        return
    
    try:
        N = int(args[0])
    except ValueError:
        return
    
    if N < 1:
        return
    
    # Generate Fibonacci numbers up to N
    a, b = 1, 1
    while a <= N:
        print(a)
        a, b = b, a + b

if __name__ == "__main__":
    main()