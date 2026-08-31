const N: number = parseInt(process.argv[2]);

if (isNaN(N) || N < 2) {
    // No primes up to N
    process.exit(0);
}

const isPrime = (num: number): boolean => {
    if (num < 2) return false;
    if (num === 2) return true;
    if (num % 2 === 0) return false;
    for (let i = 3; i * i <= num; i += 2) {
        if (num % i === 0) return false;
    }
    return true;
};

for (let i = 2; i <= N; i++) {
    if (isPrime(i)) {
        console.log(i);
    }
}