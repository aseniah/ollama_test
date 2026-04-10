int n = int.Parse(Args[0]);
var primes = new List<int>();
for (int i = 2; i <= n; i++) {
    bool isPrime = true;
    for (int j = 2; j * j <= i; j++) {
        if (i % j == 0) {
            isPrime = false;
            break;
        }
    }
    if (isPrime) {
        primes.Add(i);
    }
}
using (var ms = new Microsoft.CodeAnalysis.Text.TextSpanWriter(System.Console.Out)) {
    foreach (var p in primes) {
        ms.WriteLine(p.ToString());
    }
}