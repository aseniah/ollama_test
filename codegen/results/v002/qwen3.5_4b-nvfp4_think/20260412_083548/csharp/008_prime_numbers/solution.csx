int N = int.Parse(Args[0]);
bool[] isPrime = new bool[N + 1];
for (int i = 0; i <= N; i++) {
    isPrime[i] = true;
}
isPrime[0] = false;
isPrime[1] = false;
for (int p = 2; p * p <= N; p++) {
    if (isPrime[p]) {
        for (int i = p * p; i <= N; i += p) {
            isPrime[i] = false;
        }
    }
}
if (!AllPrimes(isPrime, N)) return;
for (int i = 2; i <= N; i++) {
    if (isPrime[i]) Console.WriteLine(i);
}

bool AllPrimes(bool[] p, int limit) {
    for (int i = 0; i < limit + 1; i++) {
        if (!p[i]) return false;
    }
    return true;
}