import sys

def print_fibonacci_up_to_n(n):
    """
    Prints the Fibonacci sequence starting from 1, 1, 2...
    up to the largest Fibonacci number that does not exceed n.
    """
    # If N < 1, the problem specifies to print nothing.
    if n < 1:
        return

    a, b = 1, 1
    
    # Handle the first number in the sequence
    print(a)
    
    # If the first number is the only one that fits N
    if n == 1:
        # We check if there's a second 1 to print (the sequence starts 1, 1...)
        # Since the second 1 also does not exceed 1, we print it.
        print(b)
        return

    # Print the second number
    print(b)

    # Calculate and print subsequent numbers
    while True:
        next_fib = a + b
        if next_fib > n:
            break
        print(next_fib)
        a, b = b, next_fib

def main():
    # Check if an argument was provided
    if len(sys.argv) != 2:
        # In a real production environment, you might print a usage error.
        # Here we exit silently or handle the error as per standard CLI behavior.
        return

    try:
        # Convert the command line argument to an integer
        n = int(sys.argv[1])
        print_fibonacci_up_to_n(n)
    except ValueError:
        # If the input is not a valid integer, we do nothing.
        pass

if __name__ == "__main__":
    main()