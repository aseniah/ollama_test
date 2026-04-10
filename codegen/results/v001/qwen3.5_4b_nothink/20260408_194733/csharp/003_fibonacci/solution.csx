if (Args.Count < 1) { return; }
int n = int.Parse(Args[0]);
if (n < 1) { return; }
int a = 1;
int b = 1;
int maxFib = a;
Console.WriteLine(maxFib);

while (maxFib + a <= n) {
    int next = b + a;
    a = b;
    b = next;
    if (b <= n) {
        Console.WriteLine(b);
    } else {
        Console.WriteLine(a);
        return;
    }
}

Console.WriteLine(a);