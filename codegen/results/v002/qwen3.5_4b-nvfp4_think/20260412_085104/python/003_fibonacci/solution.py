import sys

def main():
    # Get N from command line argument
    try:
        n = int(sys.argv[1])
    except IndexError:
        print("Usage: python script.py <integer>")
        return
    
    # Handle cases where N is invalid or less than 1
    if n < 1:
        return
    elif not isinstance(n, int) or (not isinstance(n, float) and not isinstance(n, int)):
        try:
            n = int(float(sys.argv[1]))
        except ValueError:
            print("Error: Invalid integer")
            return
    
    # Generate Fibonacci sequence up to N
    a, b = 1, 1
    while True:
        if a > n:
            break
        print(a)
        a, b = b, a + b

if __name__ == "__main__":
    main()