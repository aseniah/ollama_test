import sys

def print_fibonacci_up_to(n):
    """
    Prints the Fibonacci sequence starting from 1, 1, 2, 3, 5...
    up to the largest Fibonacci number that does not exceed n.
    """
    if n < 1:
        return

    # Initialize the first two numbers of the sequence
    a, b = 1, 1
    
    # We use a while loop to check if the current Fibonacci number 'a' 
    # is within the limit N.
    while a <= n:
        print(a)
        # Update a and b to the next numbers in the sequence
        # a becomes the old b, and b becomes the sum of the old a and b
        a, b = b, a + b

def main():
    # Ensure an argument was provided via the command line
    if len(sys.argv) != 2:
        return

    try:
        # Convert the command line argument to an integer
        n = int(sys.argv[1])
        print_fibonacci_up_to(n)
    except ValueError:
        # If the argument is not a valid integer, exit silently
        return

if __name__ == "__main__":
    main()