import sys

def main():
    if len(sys.argv) != 2:
        return
        
    n = int(sys.argv[1])
    if n < 2:
        return

    sieve = [True] * (n + 1)
    sieve[0] = sieve[1] = False
    
    for i in range(2, int(n**0.5) + 1):
        if sieve[i]:
            # Calculate how many multiples to mark to avoid creating a temporary list
            count = (n - i*i) // i + 1
            sieve[i*i:n+1:i] = [False] * count
            
    for num in range(2, n + 1):
        if sieve[num]:
            print(num)

if __name__ == "__main__":
    main()