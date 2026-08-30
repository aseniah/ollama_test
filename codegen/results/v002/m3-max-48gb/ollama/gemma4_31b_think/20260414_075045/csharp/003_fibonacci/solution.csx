if (Args.Length == 0)
{
    return;
}

if (!long.TryParse(Args[0], out long n))
{
    return;
}

if (n < 1)
{
    return;
}

long current = 1;
long next = 1;

while (current <= n)
{
    Console.WriteLine(current);
    long temp = current + next;
    current = next;
    next = temp;
}