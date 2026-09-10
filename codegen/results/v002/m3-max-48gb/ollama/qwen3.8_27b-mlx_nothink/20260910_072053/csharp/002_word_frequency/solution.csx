using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Text;

string content = File.ReadAllText("input/text.txt");

// Convert to lowercase
content = content.ToLower();

// Extract words: keep only letters, split on non-letter boundaries
var words = new List<string>();
var current = new StringBuilder();
foreach (char c in content)
{
    if (char.IsLetter(c))
    {
        current.Append(c);
    }
    else
    {
        if (current.Length > 0)
        {
            words.Add(current.ToString());
            current.Clear();
        }
    }
}
if (current.Length > 0)
{
    words.Add(current.ToString());
}

// Count frequency
var freq = new Dictionary<string, int>();
foreach (var word in words)
{
    if (freq.ContainsKey(word))
        freq[word]++;
    else
        freq[word] = 1;
}

// Sort by count descending, then alphabetically ascending
var sorted = freq.OrderByDescending(kv => kv.Value)
                 .ThenBy(kv => kv.Key)
                 .ToList();

// Output
var sb = new StringBuilder();
foreach (var kv in sorted)
{
    sb.AppendLine($"{kv.Key}: {kv.Value}");
}

Console.Write(sb.ToString());