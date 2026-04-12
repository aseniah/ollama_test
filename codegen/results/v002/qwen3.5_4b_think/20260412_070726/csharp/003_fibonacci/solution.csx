int n = int.Parse(Args[0]);

if (n < 1) {
    return;
}

int a = 1;
int b = 1;

if (a <= n) Console.WriteLine(a);

while (b <= n) {
    int next = a + b;
    a = b;
    b = next;
    if (b <= n) Console.WriteLine(b);
}