using System;
using System.Text;

if (Args.Count < 1 || !int.TryParse(Args[0], out int n))
{
    return;
}

if (n < 2)
{
    return;
}

var isPrime = new bool[n + 1];
Array.Fill(isPrime, true);
isPrime[0] = false;
isPrime[1] = false;

for (int i = 2; i * i <= n; i++)
{
    if (isPrime[i])
    {
        for (int j = i * i; j <= n; j += i)
        {
            isPrime[j] = false;
        }
    }
}

var sb = new StringBuilder();
for (int i = 2; i <= n; i++)
{
    if (isPrime[i])
    {
        sb.AppendLine(i.ToString());
    }
}

if (sb.Length > 0)
{
    Console.Write(sb.ToString().TrimEnd('\r', '\n'));
    Console.WriteLine();
}