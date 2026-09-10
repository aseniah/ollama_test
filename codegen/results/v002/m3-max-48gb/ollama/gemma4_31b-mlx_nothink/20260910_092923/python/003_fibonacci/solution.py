import sys

def print_fibonacci_up_to_n():
    # Check if the argument N is provided
    if len(sys.argv) < 2:
        return

    try:
        # Convert the command line argument to an integer
        n = int(sys.argv[1])
    except ValueError:
        # If the argument is not a valid integer, do nothing/exit
        return

    # If N < 1, print nothing as per requirements
    if n < 1:
        return

    # Initialize the first two numbers of the sequence
    a, b = 1, 1
    
    # The sequence starts: 1, 1, 2, 3, 5...
    # We print 'a' and then calculate the next number
    while a <= n:
        print(a)
        # Update a and b to the next numbers in the sequence
        a, b = b, a + b

if __name__ == "__main__":
    print_fibonacci_up_to_n()