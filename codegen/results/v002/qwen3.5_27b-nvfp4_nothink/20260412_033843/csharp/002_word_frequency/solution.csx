#r "System.Text.Json"

using System;
using System.Collections.Generic;
using System.Globalization;
using System.IO;
using System.Linq;
using System.Text.RegularExpressions;

string inputPath = "input/text.txt";

if (!File.Exists(inputPath))
{
    // Fallback for testing if the file path is incorrect or missing during local execution
    // Without this, the script would crash on empty file. 
    // However, per instructions, we assume the file exists as described.
    Console.WriteLine(); 
    return;
}

string content = File.ReadAllText(inputPath);

// Normalize case first
content = content.ToLower(CultureInfo.InvariantCulture);

// Split by non-letter characters to extract words and clean punctuation simultaneously
// Matches any sequence of letters a-z
var matches = Regex.Matches(content, "[a-z]+");

// Dictionary to store word counts
var frequency = new SortedDictionary<string, int>(StringComparer.OrdinalIgnoreCase);

foreach (Match match in matches)
{
    string word = match.Value;
    
    // Double check to ensure no empty strings passed through (though regex [a-z]+ prevents this)
    if (string.IsNullOrEmpty(word)) continue;

    if (frequency.TryGetValue(word, out int count))
    {
        frequency[word] = count + 1;
    }
    else
    {
        frequency[word] = 1;
    }
}

// Sort logic:
// 1. By count descending
// 2. If counts are equal, by word alphabetically ascending
var sortedWords = frequency
    .OrderByDescending(kv => kv.Value)
    .ThenBy(kv => kv.Key)
    .ToList();

foreach (var pair in sortedWords)
{
    Console.WriteLine($"{pair.Key}: {pair.Value}");
}