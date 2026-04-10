if (Args.Count == 0) { return; }
int n = int.Parse(Args[0]);
if (n < 1) { return; }
long a = 1;
long b = 1;
Console.WriteLine(a);
while (b <= n) {
    Console.WriteLine(b);
    long next = a + b;
    a = b;
    b = next;
}