import sys

def main():
    # Read N from command line
    if len(sys.argv) != 2:
        print(f"Usage: {sys.argv[0]} <integer>")
        return
    
    try:
        n = int(sys.argv[1])
    except ValueError:
        print("Error: Argument must be an integer")
        return

    # If N < 1, print nothing as per requirements
    if n < 1:
        return

    # Generate Fibonacci numbers up to N
    fib_list = []
    a, b = 1, 1  # First two Fibonacci numbers
    
    # Add the first Fibonacci number (1)
    if a <= n:
        fib_list.append(a)
    
    # Continue generating subsequent Fibonacci numbers
    while True:
        next_fib = a + b
        if next_fib > n:
            break
        # Only add if it's <= n
        if next_fib <= n:
            fib_list.append(next_fib)
        a, b = b, next_fib
    
    # Print each Fibonacci number on a new line
    for num in fib_list:
        print(num)

if __name__ == "__main__":
    main()