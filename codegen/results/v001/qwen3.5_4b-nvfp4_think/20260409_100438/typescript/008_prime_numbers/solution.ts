function main() {
    const arg = process.argv[2];
    if (!arg) return;

    let n: number | undefined;
    try {
        n = parseInt(arg);
    } catch (e) {
        // If argument is not a string or invalid, treat as empty input
        return;
    }

    if (n === undefined || n < 0) return;
    
    // Create Sieve of Eratosthenes array
    // size is n + 1 to accommodate index n
    const sieve: boolean[] = new Array<number>(n + 1).fill(true);
    
    // Base cases: 0 and 1 are not prime
    sieve[0] = false;
    sieve[1] = false;

    // Sieve Algorithm
    for (let p = 2; p * p <= n; p++) {
        if (sieve[p]) {
            for (let multiple = p * p; multiple <= n; multiple += p) {
                sieve[multiple] = false;
            }
        }
    }

    // Print primes up to n
    for (let i = 2; i <= n; i++) {
        if (sieve[i]) {
            console.log(i);
        }
    }
}

main();