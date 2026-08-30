using System;
using System.Text;

if (Args.Count != 1 || !int.TryParse(Args[0], out int n))
{
    Console.Error.WriteLine("Usage: <script> <N>");
    Environment.Exit(1);
}

if (n < 2)
{
    return;
}

var sb = new StringBuilder();
for (int i = 2; i <= n; i++)
{
    bool isPrime = true;
    for (int d = 2; d * d <= i; d++)
    {
        if (i % d == 0)
        {
            isPrime = false;
            break;
        }
    }
    if (isPrime)
    {
        sb.Append(i).Append('\n');
    }
}

Console.Write(sb.ToString());