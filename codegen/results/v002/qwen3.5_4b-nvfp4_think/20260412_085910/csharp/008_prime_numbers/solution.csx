var N = int.Parse(Args[0]);

// Check if there are any primes (primes start at 2)
if (N < 2) {
    return;
}

// Find and print prime numbers up to N
for (int i = 2; i <= N; i++) {
    if (IsPrime(i)) {
        Console.WriteLine(i);
    }
}

bool IsPrime(int num) {
    if (num < 2) return false;
    if (num == 2 || num == 3) return true;
    
    // Remove multiples of 2 and 3 early
    if (num % 2 == 0 || num % 3 == 0) return false;
    
    int limit = (int)Math.Sqrt(num);
    for (int i = 5; i <= limit; i += 6) {
        if (num % i == 0 || num % (i + 2) == 0) {
            return false;
        }
    }
    
    return true;
}