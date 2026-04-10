import sys

def print_fibonacci_up_to_n(n):
    """
    Prints the Fibonacci sequence starting from 1, 1, 2...
    up to the largest Fibonacci number that does not exceed n.
    """
    # If N < 1, the requirement specifies to print nothing.
    if n < 1:
        return

    # Initialize the first two numbers of the sequence
    a, b = 1, 1
    
    # We handle the first '1' manually or inside the loop.
    # To ensure the sequence starts 1, 1, 2... we print the first 1
    # and then iterate.
    
    current = a
    next_val = b
    
    while current <= n:
        print(current)
        # Calculate the next number in the sequence
        # temp = a + b; a = b; b = temp
        new_val = current + next_val
        current = next_val
        next_val = new_val

def main():
    # Check if an argument was provided
    if len(sys.argv) != 2:
        # If no argument or too many arguments, we exit.
        # Based on instructions, we expect a single integer.
        return

    try:
        # Parse the command line argument to an integer
        n = int(sys.argv[1])
        print_fibonacci_up_to_n(n)
    except ValueError:
        # If the argument is not a valid integer, print nothing or handle error
        pass

if __name__ == "__main__":
    main()