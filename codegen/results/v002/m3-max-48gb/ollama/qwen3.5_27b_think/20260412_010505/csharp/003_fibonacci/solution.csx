if (Args.Count == 0 || !int.TryParse(Args[0], out int n) || n < 1)
{
    return;
}

long a = 1;
long b = 1;

while (a <= n)
{
    Console.WriteLine(a);
    long next = a + b;
    a = b;
    b = next;
}