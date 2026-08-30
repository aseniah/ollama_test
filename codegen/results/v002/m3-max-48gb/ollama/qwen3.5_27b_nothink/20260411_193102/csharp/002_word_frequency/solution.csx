using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Text;

// Read the file content
var content = File.ReadAllText("input/text.txt");

// Dictionary to store word frequencies
var wordCounts = new Dictionary<string, int>();

// Normalize the text: convert to lowercase
var normalizedText = content.ToLowerInvariant();

// Split the text into potential tokens using non-letter characters as delimiters
// We replace any character that is not a letter or whitespace with a space
var cleanedText = "";
foreach (char c in normalizedText)
{
    if (char.IsLetter(c))
    {
        cleanedText += c;
    }
    else
    {
        cleanedText += " ";
    }
}

// Split by whitespace and filter empty strings
var words = cleanedText.Split(new[] { ' ', '\r', '\n', '\t' }, StringSplitOptions.RemoveEmptyEntries);

// Count frequencies
foreach (var word in words)
{
    if (wordCounts.ContainsKey(word))
    {
        wordCounts[word]++;
    }
    else
    {
        wordCounts.Add(word, 1);
    }
}

// Sort the results:
// 1. By count descending
// 2. By word alphabetically ascending for ties
var sortedWords = wordCounts
    .OrderByDescending(kv => kv.Value)
    .ThenBy(kv => kv.Key)
    .ToList();

// Output the results
foreach (var pair in sortedWords)
{
    Console.WriteLine($"{pair.Key}: {pair.Value}");
}