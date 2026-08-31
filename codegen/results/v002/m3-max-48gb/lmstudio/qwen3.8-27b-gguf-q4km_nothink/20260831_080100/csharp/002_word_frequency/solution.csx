using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Text;

var content = File.ReadAllText("input/text.txt");

// Convert to lowercase
content = content.ToLowerInvariant();

// Extract words (only letters)
var words = new List<string>();
var currentWord = new StringBuilder();

foreach (char c in content)
{
    if (char.IsLetter(c))
    {
        currentWord.Append(c);
    }
    else
    {
        if (currentWord.Length > 0)
        {
            words.Add(currentWord.ToString());
            currentWord.Clear();
        }
    }
}

// Don't forget the last word
if (currentWord.Length > 0)
{
    words.Add(currentWord.ToString());
}

// Count frequencies
var frequency = new Dictionary<string, int>();
foreach (var word in words)
{
    if (frequency.ContainsKey(word))
    {
        frequency[word]++;
    }
    else
    {
        frequency[word] = 1;
    }
}

// Sort by count descending, then alphabetically ascending for ties
var sorted = frequency.OrderByDescending(kvp => kvp.Value).ThenBy(kvp => kvp.Key).ToList();

// Output
var output = new StringBuilder();
foreach (var kvp in sorted)
{
    output.AppendLine($"{kvp.Key}: {kvp.Value}");
}

Console.Write(output.ToString());