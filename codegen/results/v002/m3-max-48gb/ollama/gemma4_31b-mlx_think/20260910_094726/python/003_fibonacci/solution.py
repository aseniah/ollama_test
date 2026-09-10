import sys

def main():
    # Ensure an argument is provided
    if len(sys.argv) < 2:
        return

    try:
        # Accept a single integer argument N from the command line
        n = int(sys.argv[1])
    except ValueError:
        # If the argument is not an integer, we simply exit
        return

    # If N < 1, print nothing
    if n < 1:
        return

    # Initialize the start of the sequence: 1, 1, 2, 3, 5...
    a, b = 1, 1
    
    # Print Fibonacci numbers as long as the current number does not exceed N
    while a <= n:
        print(a)
        # Calculate the next number in the sequence
        a, b = b, a + b

if __name__ == "__main__":
    main()