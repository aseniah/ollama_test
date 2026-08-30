const N = parseInt(process.argv[2], 10);

if (isNaN(N) || N < 2) {
    // No primes up to N
    process.exit(0);
}

// Sieve of Eratosthenes
const isPrime = new Array<boolean>(N + 1).fill(true);
isPrime[0] = false;
if (N >= 1) isPrime[1] = false;

for (let i = 2; i * i <= N; i++) {
    if (isPrime[i]) {
        for (let j = i * i; j <= N; j += i) {
            isPrime[j] = false;
        }
    }
}

for (let i = 2; i <= N; i++) {
    if (isPrime[i]) {
        console.log(i);
    }
}