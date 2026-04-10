using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Text;

string content = File.ReadAllText("input/text.txt");

// Convert to lowercase first
content = content.ToLower();

// Extract words: keep only letters, ignore punctuation and whitespace
var sb = new StringBuilder();
foreach (char c in content)
{
    if (char.IsLetter(c))
    {
        sb.Append(c);
    }
    else
    {
        if (sb.Length > 0)
        {
            yieldWords(sb.ToString());
            sb.Clear();
        }
    }
}

if (sb.Length > 0)
{
    yieldWords(sb.ToString());
}

void yieldWords(string word)
{
    // This helper just pushes into the dictionary for counting
    // We need a way to collect these, so we'll use a dictionary directly instead
}

// Re-implement logic without closure issues in script context
var counts = new Dictionary<string, int>();

sb.Clear();
foreach (char c in content)
{
    if (char.IsLetter(c))
    {
        sb.Append(c);
    }
    else
    {
        if (sb.Length > 0)
        {
            string w = sb.ToString();
            if (counts.ContainsKey(w))
                counts[w]++;
            else
                counts[w] = 1;
            sb.Clear();
        }
    }
}

if (sb.Length > 0)
{
    string w = sb.ToString();
    if (counts.ContainsKey(w))
        counts[w]++;
    else
        counts[w] = 1;
}

var sorted = counts
    .OrderByDescending(kvp => kvp.Value)
    .ThenBy(kvp => kvp.Key)
    .ToList();

foreach (var kvp in sorted)
{
    Console.WriteLine($"{kvp.Key}: {kvp.Value}");
}