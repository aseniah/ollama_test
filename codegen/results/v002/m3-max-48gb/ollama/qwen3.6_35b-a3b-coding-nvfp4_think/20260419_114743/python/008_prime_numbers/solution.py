import sys

def main():
    # Parse the integer N from command line arguments
    n = int(sys.argv[1])
    
    def is_prime(num):
        """Check if a number is prime."""
        if num < 2:
            return False
        # Check divisibility up to the square root of num
        for i in range(2, int(num**0.5) + 1):
            if num % i == 0:
                return False
        return True

    # Print all prime numbers up to and including N
    for num in range(2, n + 1):
        if is_prime(num):
            print(num)

if __name__ == "__main__":
    main()