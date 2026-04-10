import sys

def print_fibonacci_up_to_n(n):
    """
    Prints all Fibonacci numbers up to and including the largest 
    Fibonacci number that does not exceed n.
    The sequence starts: 1, 1, 2, 3, 5, ...
    """
    if n < 1:
        return

    # Initialize the first two numbers of the sequence as requested
    a, b = 1, 1
    
    while a <= n:
        print(a)
        # Update a and b to the next numbers in the sequence
        a, b = b, a + b

def main():
    # Check if an argument was provided
    if len(sys.argv) != 2:
        # If the number of arguments is not exactly 1 (excluding script name), exit.
        return

    try:
        # Parse the command line argument as an integer
        n = int(sys.argv[1])
        print_fibonacci_up_to_n(n)
    except ValueError:
        # If the argument is not a valid integer, exit silently.
        return

if __name__ == "__main__":
    main()