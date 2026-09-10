using System;

if (Args.Count < 1)
    return;

if (!int.TryParse(Args[0], out int n))
    return;

if (n < 1)
    return;

int a = 1, b = 1;
while (a <= n)
{
    Console.WriteLine(a);
    int c = a + b;
    a = b;
    b = c;
}