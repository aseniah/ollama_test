import sys

def main():
    # Read the integer N from command line arguments
    if len(sys.argv) < 2:
        return
    try:
        n = int(sys.argv[1])
    except ValueError:
        return
    
    if n < 1:
        return
    
    # Generate Fibonacci numbers up to and including the largest one <= n
    a, b = 1, 1
    print(a)
    if n >= 1:
        a, b = b, a + b
        while a <= n:
            print(a)
            a, b = b, a + b

if __name__ == "__main__":
    main()