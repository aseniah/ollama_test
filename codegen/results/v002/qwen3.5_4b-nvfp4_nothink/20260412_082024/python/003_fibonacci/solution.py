import sys

def main():
    # Accept a single integer argument N from the command line
    if len(sys.argv) != 2:
        print("Usage: python script.py <integer>")
        return

    try:
        n = int(sys.argv[1])
    except ValueError:
        print("Error: Argument must be an integer.")
        return

    # If N < 1, print nothing
    if n < 1:
        return

    fib_sequence = []
    a, b = 1, 1
    
    while True:
        # Append the current Fibonacci number if it's within bounds
        if a <= n:
            fib_sequence.append(a)
        else:
            break
        
        # Move to the next Fibonacci number
        a, b = b, a + b
    
    # Print all numbers up to and including N, one per line
    for num in fib_sequence:
        print(num)

if __name__ == "__main__":
    main()