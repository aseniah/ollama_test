import sys

def main():
    # Check if an argument was provided
    if len(sys.argv) < 2:
        return

    try:
        # Parse the input argument N as an integer
        n = int(sys.argv[1])
    except ValueError:
        # If the argument is not a valid integer, do nothing
        return

    # If N < 1, the program should print nothing as per requirements
    if n < 1:
        return

    # Start the sequence: 1, 1, 2, 3, 5...
    a, b = 1, 1
    
    # We print the first number manually to handle the sequence logic
    # while ensuring we don't print if n < 1 (already checked)
    print(a)
    
    # Calculate subsequent Fibonacci numbers
    while b <= n:
        print(b)
        # Update values to the next pair in the sequence
        a, b = b, a + b

if __name__ == "__main__":
    main()