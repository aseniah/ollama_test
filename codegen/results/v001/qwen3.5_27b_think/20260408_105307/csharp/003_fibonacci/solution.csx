using System;

if (Args.Count != 1 || !int.TryParse(Args[0], out int n))
{
    return;
}

if (n < 1)
{
    return;
}

long a = 1;
long b = 1;

Console.WriteLine(a);
Console.WriteLine(b);

while (true)
{
    long next = a + b;
    if (next > n)
    {
        break;
    }
    Console.WriteLine(next);
    a = b;
    b = next;
}