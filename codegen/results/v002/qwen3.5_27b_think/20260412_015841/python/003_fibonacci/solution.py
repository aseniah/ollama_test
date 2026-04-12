import sys

def main():
    # Get the command line argument
    N = int(sys.argv[1])
    
    # If N < 1, print nothing
    if N < 1:
        return
    
    # Generate and print Fibonacci numbers up to N
    a, b = 1, 1
    while a <= N:
        print(a)
        a, b = b, a + b

if __name__ == "__main__":
    main()