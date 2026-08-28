using System;

int n = int.Parse(Args[0]);

if (n < 1)
    return;

Console.WriteLine(1);
Console.WriteLine(1);

long a = 1, b = 1;
while (true)
{
    long c = a + b;
    if (c > n) break;
    Console.WriteLine(c);
    a = b;
    b = c;
}