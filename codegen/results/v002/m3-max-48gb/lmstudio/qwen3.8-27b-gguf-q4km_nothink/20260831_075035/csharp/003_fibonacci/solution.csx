if (Args == null || Args.Count == 0 || !int.TryParse(Args[0], out int N))
{
    return;
}

if (N < 1)
{
    return;
}

long a = 1;
long b = 1;

if (N >= 1)
{
    Console.WriteLine(a);
    if (N >= 1)
    {
        Console.WriteLine(b);
    }
}

long next = a + b;
while (next <= N)
{
    Console.WriteLine(next);
    a = b;
    b = next;
    next = a + b;
}