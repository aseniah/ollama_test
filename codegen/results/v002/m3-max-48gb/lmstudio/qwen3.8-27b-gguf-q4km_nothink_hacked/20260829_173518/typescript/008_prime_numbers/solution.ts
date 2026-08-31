function main() {
    const arg = process.argv[2];
    if (arg === undefined) {
        return;
    }
    const N = parseInt(arg, 10);
    if (isNaN(N) || N < 2) {
        return;
    }

    // Sieve of Eratosthenes
    const isPrime = new Array(N + 1).fill(true);
    isPrime[0] = false;
    isPrime[1] = false;

    for (let i = 2; i * i <= N; i++) {
        if (isPrime[i]) {
            for (let j = i * i; j <= N; j += i) {
                isPrime[j] = false;
            }
        }
    }

    const out: string[] = [];
    for (let i = 2; i <= N; i++) {
        if (isPrime[i]) {
            out.push(String(i));
        }
    }
    if (out.length > 0) {
        console.log(out.join('\n'));
    }
}

main();