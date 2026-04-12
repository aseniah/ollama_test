import sys

def print_fibonacci_up_to_n(n):
    """
    Prints the Fibonacci sequence starting from 1, 1, 2, 3, 5...
    up to the largest Fibonacci number that does not exceed n.
    If n < 1, nothing is printed.
    """
    if n < 1:
        return

    # Initialize the first two numbers of the sequence
    a, b = 1, 1
    
    # We use a loop to generate numbers. 
    # The logic starts with a=1, b=1. 
    # On the first iteration, 1 is printed.
    # Then a becomes 1, b becomes 2.
    # On the second iteration, 1 is printed.
    # This ensures the sequence starts 1, 1, ... as requested.
    while a <= n:
        print(a)
        # Calculate the next Fibonacci number
        a, b = b, a + b

def main():
    # Ensure there is exactly one argument provided in the command line
    if len(sys.argv) != 2:
        return

    try:
        # Convert the command line argument to an integer
        n = int(sys.argv[1])
        print_fibonacci_up_to_n(n)
    except ValueError:
        # If the argument is not a valid integer, do nothing
        pass

if __name__ == "__main__":
    main()