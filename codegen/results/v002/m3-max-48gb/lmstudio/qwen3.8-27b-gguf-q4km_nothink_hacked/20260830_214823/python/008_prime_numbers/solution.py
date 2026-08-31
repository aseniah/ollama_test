import sys

def main():
    # Read the integer N from command line argument
    if len(sys.argv) < 2:
        return
    
    try:
        N = int(sys.argv[1])
    except ValueError:
        return
    
    if N < 2:
        return
    
    # Sieve of Eratosthenes to find all primes up to N
    is_prime = [True] * (N + 1)
    is_prime[0] = False
    if N >= 1:
        is_prime[1] = False
    
    for i in range(2, int(N ** 0.5) + 1):
        if is_prime[i]:
            for j in range(i * i, N + 1, i):
                is_prime[j] = False
    
    # Print all primes up to N, one per line
    for i in range(2, N + 1):
        if is_prime[i]:
            print(i)

if __name__ == "__main__":
    main()