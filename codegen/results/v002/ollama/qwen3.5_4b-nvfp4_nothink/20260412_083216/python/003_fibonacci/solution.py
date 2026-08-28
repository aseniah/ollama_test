import sys

def main():
    # Accept a single integer argument N from the command line
    if len(sys.argv) != 2:
        print("Usage: python script.py <N>", file=sys.stderr)
        sys.exit(1)

    try:
        n = int(sys.argv[1])
    except ValueError:
        print("Invalid input", file=sys.stderr)
        sys.exit(1)

    # If N < 1, print nothing
    if n < 1:
        return

    fibonacci_sequence = [1]  # Start with the first '1'
    
    # Generate Fibonacci numbers until the next one would exceed N
    while True:
        next_fib = fibonacci_sequence[-1] + fibonacci_sequence[-2]
        
        if next_fib > n:
            break
        
        fibonacci_sequence.append(next_fib)

    # Print all Fibonacci numbers up to and including the largest that does not exceed N
    for num in fibonacci_sequence:
        print(num)

if __name__ == "__main__":
    main()