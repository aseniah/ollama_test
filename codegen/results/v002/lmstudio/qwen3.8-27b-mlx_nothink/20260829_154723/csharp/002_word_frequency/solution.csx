using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;

string path = "input/text.txt";
if (!File.Exists(path))
{
    // In case the working directory is different, try current directory
    return;
}

string[] lines = File.ReadAllLines(path);
var wordCounts = new Dictionary<string, int>();

foreach (string line in lines)
{
    // Convert to lowercase and extract words
    string lower = line.ToLowerInvariant();
    int i = 0;
    int len = lower.Length;
    while (i < len)
    {
        // Skip non-letters
        while (i < len && !char.IsLetter(lower[i]))
            i++;
        // Collect letters
        if (i < len)
        {
            int start = i;
            while (i < len && char.IsLetter(lower[i]))
                i++;
            string word = lower.Substring(start, i - start);
            if (word.Length > 0)
            {
                if (wordCounts.ContainsKey(word))
                    wordCounts[word]++;
                else
                    wordCounts[word] = 1;
            }
        }
    }
}

var sorted = wordCounts
    .OrderByDescending(kv => kv.Value)
    .ThenBy(kv => kv.Key, StringComparer.Ordinal)
    .ToList();

foreach (var kv in sorted)
{
    Console.WriteLine($"{kv.Key}: {kv.Value}");
}