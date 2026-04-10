import sys

def print_fibonacci_up_to_n(n):
    """
    Prints Fibonacci numbers starting from 1, 1, 2... 
    up to the largest Fibonacci number <= n.
    """
    if n < 1:
        return

    a, b = 1, 1
    
    # The sequence starts with 1, 1. We handle the first two 1s explicitly
    # or via a loop logic that checks the current number against n.
    while a <= n:
        print(a)
        # Calculate next Fibonacci number
        a, b = b, a + b

def main():
    # Check if exactly one argument is provided (sys.argv[0] is the script name)
    if len(sys.argv) != 2:
        return

    try:
        # Parse the command line argument as an integer
        n = int(sys.argv[1])
        print_fibonacci_up_to_n(n)
    except ValueError:
        # If the argument is not a valid integer, we exit silently 
        # or handle error as per requirements.
        pass

if __name__ == "__main__":
    main()