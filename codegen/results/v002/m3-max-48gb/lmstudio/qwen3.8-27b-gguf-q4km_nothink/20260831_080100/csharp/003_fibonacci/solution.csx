using System;
using System.Collections.Generic;

if (Args.Count < 1)
    return;

if (!int.TryParse(Args[0], out int N))
    return;

if (N < 1)
    return;

long a = 1, b = 1;
Console.WriteLine(a);
if (b <= N) Console.WriteLine(b);
while (true)
{
    long c = a + b;
    if (c > N) break;
    Console.WriteLine(c);
    a = b;
    b = c;
}