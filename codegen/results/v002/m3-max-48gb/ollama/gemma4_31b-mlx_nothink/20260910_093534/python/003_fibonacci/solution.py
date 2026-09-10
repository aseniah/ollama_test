import sys

def print_fibonacci_up_to_n():
    # Ensure an argument was provided
    if len(sys.argv) < 2:
        return

    try:
        # Parse the command line argument as an integer
        n = int(sys.argv[1])
    except ValueError:
        # If the input is not an integer, we stop
        return

    # If N < 1, print nothing as per requirements
    if n < 1:
        return

    # Initialize the first two numbers of the Fibonacci sequence
    a, b = 1, 1
    
    # The loop continues as long as the current Fibonacci number is <= N
    while a <= n:
        print(a)
        # Update sequence: next number is the sum of the previous two
        a, b = b, a + b

if __name__ == "__main__":
    print_fibonacci_up_to_n()