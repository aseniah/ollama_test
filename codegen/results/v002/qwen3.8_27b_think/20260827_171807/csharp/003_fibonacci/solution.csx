if (Args.Count < 1 || !long.TryParse(Args[0], out long N))
    return;

if (N < 1)
    return;

long a = 1, b = 1;
while (a <= N)
{
    Console.WriteLine(a);
    long next = a + b;
    a = b;
    b = next;
}