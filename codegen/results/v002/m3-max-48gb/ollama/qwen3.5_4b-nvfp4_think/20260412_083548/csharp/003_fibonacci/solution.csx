int n = int.Parse(Args[0]);

if (n < 1) {
    return;
}

long a = 1;
long b = 1;

Console.WriteLine(a);
Console.WriteLine(b);

while (b <= n && a + b > n) {
    long next = a + b;
    if (next <= n) {
        Console.WriteLine(next);
    }
    a = b;
    b = next;
}