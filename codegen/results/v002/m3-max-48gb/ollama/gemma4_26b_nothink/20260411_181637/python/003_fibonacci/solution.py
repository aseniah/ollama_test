import sys

def print_fibonacci_up_to_n(n):
    """
    Prints the Fibonacci sequence starting from 1, 1, 2, 3, 5...
    up to the largest Fibonacci number that does not exceed n.
    """
    if n < 1:
        return

    a, b = 1, 1
    
    # The sequence starts with 1. We print the first 1 manually 
    # to handle the logic of the loop smoothly.
    print(a)
    
    # We use a loop to calculate the next numbers in the sequence.
    # 'a' represents the current number being printed.
    # 'b' represents the next number in the sequence.
    while b <= n:
        print(b)
        # Update a and b: a becomes the old b, b becomes the sum of old a and b
        a, b = b, a + b

if __name__ == "__main__":
    # Check if an argument was provided via command line
    if len(sys.argv) != 2:
        # If no argument or too many arguments, we can't proceed.
        # The prompt implies a single integer will be provided.
        sys.exit(1)

    try:
        # Parse the command line argument as an integer
        n_value = int(sys.argv[1])
        print_fibonacci_up_to_n(n_value)
    except ValueError:
        # If the argument is not a valid integer, exit silently or handle error
        sys.exit(1)