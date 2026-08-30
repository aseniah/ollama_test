import sys

def print_fibonacci_up_to_n(n):
    """
    Prints the Fibonacci sequence up to the largest number that does not exceed N.
    Sequence starts: 1, 1, 2, 3, 5, ...
    """
    if n < 1:
        return

    a, b = 1, 1
    
    # The problem specifies the sequence starts 1, 1, 2...
    # We handle the first 1 manually to ensure logic works for N=1
    print(a)
    
    # If N is at least 1 and we have our first 1, check if we can print the second 1
    if n >= 1:
        # Since the sequence is 1, 1, 2... and N >= 1, the second 1 is always valid.
        # However, we must check if N allows for the second 1 (which it does if N >= 1).
        # We print it and move to calculating the next terms.
        print(b)
        
        # Calculate subsequent numbers
        while True:
            next_fib = a + b
            if next_fib > n:
                break
            print(next_fib)
            a, b = b, next_fib

def main():
    # Check if an argument was provided
    if len(sys.argv) != 2:
        return

    try:
        # Parse the input argument as an integer
        n = int(sys.argv[1])
        print_fibonacci_up_to_n(n)
    except ValueError:
        # If the argument is not a valid integer, exit silently or handle error
        pass

if __name__ == "__main__":
    main()