using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Text;

// Read the file content
var text = File.ReadAllText("input/text.txt");

// Convert to lowercase
text = text.ToLower();

// Extract words - keep only letters
var words = new List<string>();
var wordBuilder = new StringBuilder();

foreach (char c in text)
{
    if (char.IsLetter(c))
    {
        wordBuilder.Append(c);
    }
    else
    {
        if (wordBuilder.Length > 0)
        {
            words.Add(wordBuilder.ToString());
            wordBuilder.Clear();
        }
    }
}

// Add the last word if exists
if (wordBuilder.Length > 0)
{
    words.Add(wordBuilder.ToString());
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
var sortedWords = frequency.OrderByDescending(k => k.Value)
                          .ThenBy(k => k.Key);

// Output results
foreach (var kvp in sortedWords)
{
    Console.WriteLine($"{kvp.Key}: {kvp.Value}");
}