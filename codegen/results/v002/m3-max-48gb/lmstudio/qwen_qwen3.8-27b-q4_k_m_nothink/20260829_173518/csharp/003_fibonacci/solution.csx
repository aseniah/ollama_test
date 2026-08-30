using System;

// Read the integer N from command line
if (Args.Count == 0)
{
    return;
}

if (!int.TryParse(Args[0], out int n))
{
    return;
}

if (n < 1)
{
    return;
}

long a = 1;
long b = 1;

if (a <= n)
    Console.WriteLine(a);
if (b <= n && b > a)
    Console.WriteLine(b);

while (true)
{
    long c = a + b;
    if (c > n)
        break;
    Console.WriteLine(c);
    a = b;
    b = c;
}