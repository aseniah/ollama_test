import sys

def main():
    # Read the command line argument
    if len(sys.argv) != 2:
        print("Usage: python script.py N")
        return
    
    try:
        N = int(sys.argv[1])
    except ValueError:
        print("Please provide a valid integer")
        return
    
    # If N < 1, print nothing
    if N < 1:
        return
    
    # Generate Fibonacci numbers up to N
    a, b = 1, 1
    
    # Print the first Fibonacci number
    if a <= N:
        print(a)
    
    # Continue generating and printing Fibonacci numbers
    while b <= N:
        print(b)
        # Calculate next Fibonacci number
        a, b = b, a + b

if __name__ == "__main__":
    main()