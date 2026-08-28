import sys

def print_fibonacci_up_to_n(n):
    """
    Prints the Fibonacci sequence starting from 1, 1, 2, 3... 
    up to the largest Fibonacci number that does not exceed n.
    
    Args:
        n (int): The upper limit for the Fibonacci numbers.
    """
    # If N < 1, the requirement states to print nothing.
    if n < 1:
        return

    # Initialize the first two numbers of the sequence as specified.
    a, b = 1, 1
    
    # Iterate while the current Fibonacci number 'a' does not exceed n.
    while a <= n:
        print(a)
        # Calculate the next number in the sequence.
        # a becomes the previous b, and b becomes the sum of the previous a and b.
        a, b = b, a + b

def main():
    # Check if an argument was provided via the command line.
    if len(sys.argv) < 2:
        return

    try:
        # Parse the command line argument as an integer.
        n = int(sys.argv[1])
        print_fibonacci_up_to_n(n)
    except ValueError:
        # If the argument is not a valid integer, we simply exit.
        pass

if __name__ == "__main__":
    main()