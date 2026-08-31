using System;
using System.Text;

if (Args == null || Args.Count == 0)
    return;

if (!int.TryParse(Args[0], out int N))
    return;

if (N < 2)
    return;

var sb = new StringBuilder();

for (int i = 2; i <= N; i++)
{
    bool isPrime = true;
    for (int j = 2; j * j <= i; j++)
    {
        if (i % j == 0)
        {
            isPrime = false;
            break;
        }
    }
    if (isPrime)
    {
        sb.AppendLine(i.ToString());
    }
}

if (sb.Length > 0)
    Console.Write(sb.ToString().TrimEnd('\n'));