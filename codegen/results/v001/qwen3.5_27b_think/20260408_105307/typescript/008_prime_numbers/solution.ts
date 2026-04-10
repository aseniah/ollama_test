#!/usr/bin/env tsx

/**
 * Check if a number is prime
 */
function isPrime(num: number): boolean {
    if (num < 2) return false;
    if (num === 2) return true;
    if (num % 2 === 0) return false;
    
    const sqrt = Math.floor(Math.sqrt(num));
    for (let i = 3; i <= sqrt; i += 2) {
        if (num % i === 0) return false;
    }
    return true;
}

/**
 * Main function
 */
function main(): void {
    const nStr = process.argv[2];
    
    if (nStr === undefined) {
        return;
    }
    
    const n = parseInt(nStr, 10);
    
    if (isNaN(n)) {
        return;
    }
    
    for (let i = 2; i <= n; i++) {
        if (isPrime(i)) {
            console.log(i);
        }
    }
}

main();