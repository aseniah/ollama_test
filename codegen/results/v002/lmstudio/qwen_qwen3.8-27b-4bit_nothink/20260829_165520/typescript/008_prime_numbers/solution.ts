import * as path from 'path';

function main() {
    const n = parseInt(process.argv[2]);
    
    if (isNaN(n)) {
        return;
    }
    
    if (n < 2) {
        return;
    }
    
    // Sieve of Eratosthenes to find all primes up to N
    const sieve: boolean[] = new Array(n + 1).fill(true);
    sieve[0] = false;
    if (n >= 1) sieve[1] = false;
    
    for (let i = 2; i * i <= n; i++) {
        if (sieve[i]) {
            for (let j = i * i; j <= n; j += i) {
                sieve[j] = false;
            }
        }
    }
    
    for (let i = 2; i <= n; i++) {
        if (sieve[i]) {
            console.log(i);
        }
    }
}

main();