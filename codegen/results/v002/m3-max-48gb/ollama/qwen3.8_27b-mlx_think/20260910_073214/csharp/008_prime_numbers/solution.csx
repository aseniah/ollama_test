using System;
using System.Text;

int N = int.Parse(Args[0]);

if (N >= 2)
{
    bool[] isComposite = new bool[N + 1];
    var sb = new StringBuilder();

    for (int i = 2; i <= N; i++)
    {
        if (!isComposite[i])
        {
            sb.AppendLine(i.ToString());
            for (long j = (long)i * i; j <= N; j += i)
            {
                isComposite[j] = true;
            }
        }
    }

    System.Console.Write(sb.ToString());
}