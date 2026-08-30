import sys

def main():
    """
    Solves the task of printing Fibonacci numbers up to and including 
    the largest Fibonacci number that does not exceed N.
    """
    # Check if an argument was provided via the command line
    if len(sys.argv) < 2:
        return

    # Try to parse the first command line argument as an integer
    try:
        n = int(sys.argv[1])
    except ValueError:
        # If the argument is not an integer, we exit silently
        return

    # According to the task, if N < 1, print nothing
    if n < 1:
        return

    # The Fibonacci sequence starts with 1, 1, 2, 3, 5, ...
    # a represents the current Fibonacci number, b represents the next one
    a = 1
    b = 1

    # The first number in the sequence is always 1 (if N >= 1)
    print(a)

    # Check if the second '1' in the sequence is within the limit N
    if b <= n:
        print(b)
        
        # Continue calculating the sequence as long as the next number <= N
        while True:
            next_fib = a + b
            if next_fib <= n:
                print(next_fib)
                # Update a and b to progress the sequence
                a = b
                b = next_fib
            else:
                # Stop the loop when the next Fibonacci number exceeds N
                break

if __name__ == "__main__":
    main()