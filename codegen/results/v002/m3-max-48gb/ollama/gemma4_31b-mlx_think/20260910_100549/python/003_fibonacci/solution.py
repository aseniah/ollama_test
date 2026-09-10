import sys

def main():
    """
    Prints Fibonacci numbers up to and including the largest Fibonacci number 
    that does not exceed N.
    """
    # Check if the command line argument is provided
    if len(sys.argv) < 2:
        return

    try:
        # Parse the integer N from the command line argument
        n = int(sys.argv[1])
    except ValueError:
        # If the argument is not an integer, we simply exit
        return

    # If N < 1, print nothing as per requirements
    if n < 1:
        return

    # Initialize the first two Fibonacci numbers
    a, b = 1, 1
    
    # Loop while the current Fibonacci number is less than or equal to N
    while a <= n:
        print(a)
        # Calculate the next number in the sequence
        # a becomes the current b, b becomes the sum of the previous two
        a, b = b, a + b

if __name__ == "__main__":
    main()