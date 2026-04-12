function main() {
    // Get N from command line argument
    const n = parseInt(process.argv[2], 10);

    // Validate input: must be a valid integer >= 2 to have primes
    if (isNaN(n) || n < 2) {
        // No primes to print
        return;
    }

    // Sieve of Eratosthenes to find all primes up to N
    const isPrime = new Array<boolean>(n + 1).fill(true);
    isPrime[0] = false;
    isPrime[1] = false;

    for (let i = 2; i * i <= n; i++) {
        if (isPrime[i]) {
            for (let j = i * i; j <= n; j += i) {
                isPrime[j] = false;
            }
        }
    }

    // Print all prime numbers up to and including N
    for (let i = 2; i <= n; i++) {
        if (isPrime[i]) {
            console.log(i);
        }
    }
}

main();