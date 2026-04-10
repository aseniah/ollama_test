import sys

def main():
    # Parse the integer argument from command line
    if len(sys.argv) < 2:
        return
    
    try:
        n = int(sys.argv[1])
    except ValueError:
        return
    
    # If N < 1, print nothing
    if n < 1:
        return
    
    # Generate Fibonacci numbers up to and including N
    a, b = 1, 1
    while a <= n:
        print(a)
        a, b = b, a + b

if __name__ == "__main__":
    main()