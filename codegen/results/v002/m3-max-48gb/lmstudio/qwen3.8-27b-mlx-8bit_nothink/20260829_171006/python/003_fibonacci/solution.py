import sys

def main():
    # Read the integer N from command line arguments
    if len(sys.argv) < 2:
        return
    
    try:
        N = int(sys.argv[1])
    except ValueError:
        return
    
    # If N < 1, print nothing
    if N < 1:
        return
    
    # Generate Fibonacci numbers up to and including the largest one <= N
    # Sequence: 1, 1, 2, 3, 5, ...
    a, b = 1, 1
    while a <= N:
        print(a)
        a, b = b, a + b

if __name__ == "__main__":
    main()