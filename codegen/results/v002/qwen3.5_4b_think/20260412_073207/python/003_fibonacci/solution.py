import sys

def main():
    # Validate arguments
    if len(sys.argv) != 2:
        print("Usage: python script.py <integer>")
        sys.exit(1)
    
    try:
        n = int(sys.argv[1])
    except ValueError:
        print("Error: Argument must be an integer")
        sys.exit(1)
    
    # If N < 1, print nothing
    if n < 1:
        return
    
    # Generate and print Fibonacci numbers up to N
    a, b = 1, 1
    while a <= n:
        print(a)
        a, b = b, a + b

if __name__ == "__main__":
    main()