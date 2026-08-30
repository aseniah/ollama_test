var N = int.Parse(Args[0]);

if (N < 2) {
    return; // No primes <= N
}

// Sieve of Eratosthenes
int[] sieve = new int[N + 1];

for (int i = 2; i <= N; i++) {
    sieve[i] = 1; // Assume prime initially
}

int sqrtN = (int)Math.Sqrt(N);

for (int p = 2; p <= sqrtN; p++) {
    if (sieve[p]) {
        int start = p * p;
        if (start > N) continue;
        
        // Start marking multiples starting from p*p to avoid redundant work
        for (int i = start; i <= N; i += p) {
            sieve[i] = 0;
        }
    }
}

// Print primes
for (int i = 2; i <= N; i++) {
    if (sieve[i]) {
        System.Console.WriteLine(i);
    }
}