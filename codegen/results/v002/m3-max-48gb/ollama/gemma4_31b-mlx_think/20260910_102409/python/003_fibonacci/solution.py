import sys

def main():
    """
    Accepts an integer N from the command line and prints Fibonacci numbers
    up to and including the largest Fibonacci number that does not exceed N.
    """
    # Ensure at least one argument is provided
    if len(sys.argv) < 2:
        return

    try:
        # Parse the command line argument N
        n = int(sys.argv[1])
    except ValueError:
        # If the argument is not a valid integer, simply exit
        return

    # If N < 1, print nothing as per requirements
    if n < 1:
        return

    # Initialize the first two Fibonacci numbers
    a, b = 1, 1
    
    # Print the first number in the sequence
    print(a)
    
    # Continue generating and printing numbers as long as they do not exceed N
    while b <= n:
        print(b)
        # Calculate the next Fibonacci number
        a, b = b, a + b

if __name__ == "__main__":
    main()