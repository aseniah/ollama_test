function main(): void {
    const n = parseInt(process.argv[2], 10);
    
    if (isNaN(n) || n < 2) {
        return;
    }
    
    // Sieve of Eratosthenes
    const isPrime = new Array(n + 1).fill(true);
    isPrime[0] = false;
    isPrime[1] = false;
    
    for (let i = 2; i * i <= n; i++) {
        if (isPrime[i]) {
            for (let j = i * i; j <= n; j += i) {
                isPrime[j] = false;
            }
        }
    }
    
    for (let i = 2; i <= n; i++) {
        if (isPrime[i]) {
            console.log(i);
        }
    }
}

main();