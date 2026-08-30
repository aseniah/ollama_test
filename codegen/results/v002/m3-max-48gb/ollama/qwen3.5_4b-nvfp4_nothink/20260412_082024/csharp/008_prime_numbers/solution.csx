int n = int.Parse(Args[0]);

// Sieve of Eratosthenes to find primes up to N
var isPrime = new bool[n + 1];
Array.Fill(isPrime, true);
isPrime[0] = false;
if (n >= 1) {
    isPrime[1] = false;
}

// Start sieving from the first prime number
int limit = Math.Min(n, 256); // Optimization: we only need to sieve up to sqrt(N), 
                             // but for very large N we just iterate. 
                             // Actually, the standard sieve goes up to sqrt(n).
                             // However, since we need all primes up to n, and n can be large,
                             // a simple O(n log log n) or even just O(n) trial division might be needed
                             // if N is huge (though typical for simple challenges). 
                             // Let's do the true sieve: we only need to mark multiples of p where p <= sqrt(n).

int sqrtN = Math.Sqrt(n);

for (int i = 2; i <= sqrtN; i++) {
    if (isPrime[i]) {
        // Mark multiples of i
        for (int j = i * i; j <= n; j += i) {
            isPrime[j] = false;
        }
    }
}

// Collect and print primes
System.IO.WriteLine("Found prime numbers:");
for (int i = 2; i <= n; i++) {
    if (isPrime[i]) {
        System.IO.WriteLine(i);
    }
}