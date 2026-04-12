#r "System.Linq"

using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Text;

// Read the file
string content = File.ReadAllText("input/text.txt");

// Extract words (only letters), convert to lowercase
var wordList = new List<string>();
StringBuilder sb = new StringBuilder();

foreach (char c in content.ToLowerInvariant())
{
    if (char.IsLetter(c))
    {
        sb.Append(c);
    }
    else if (sb.Length > 0)
    {
        wordList.Add(sb.ToString());
        sb.Clear();
    }
}
if (sb.Length > 0)
{
    wordList.Add(sb.ToString());
}

// Count frequencies
var frequencyDict = new Dictionary<string, int>();
foreach (string word in wordList)
{
    if (!frequencyDict.ContainsKey(word))
    {
        frequencyDict[word] = 0;
    }
    frequencyDict[word]++;
}

// Sort by count descending, then alphabetically ascending
var sorted = frequencyDict.OrderByDescending(kv => kv.Value)
    .ThenBy(kv => kv.Key)
    .ToList();

// Output
foreach (var kv in sorted)
{
    Console.WriteLine($"{kv.Key}: {kv.Value}");
}