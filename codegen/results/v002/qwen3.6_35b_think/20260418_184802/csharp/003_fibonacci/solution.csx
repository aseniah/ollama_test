using System;

if (Args.Count < 1 || !int.TryParse(Args[0], out int n) || n < 1)
{
    return;
}

long a = 1;
long b = 1;
while (a <= n)
{
    Console.WriteLine(a);
    long temp = a + b;
    a = b;
    b = temp;
}