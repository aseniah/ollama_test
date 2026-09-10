using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Text;

string filePath = "input/text.txt";

if (!File.Exists(filePath))
{
    return;
}

string text = File.ReadAllText(filePath);
string[] rawWords = text.Split(new char[] { ' ', '\t', '\n', '\r' }, StringSplitOptions.RemoveEmptyEntries);

Dictionary<string, int> frequencies = new Dictionary<string, int>();

foreach (string rawWord in rawWords)
{
    // Convert to lowercase
    string lowercaseWord = rawWord.ToLower();
    
    // Strip all punctuation (keep only letters)
    StringBuilder sb = new StringBuilder();
    foreach (char c in lowercaseWord)
    {
        if (char.IsLetter(c))
        {
            sb.Append(c);
        }
    }
    
    string cleanedWord = sb.ToString();
    
    if (!string.IsNullOrEmpty(cleanedWord))
    {
        if (frequencies.ContainsKey(cleanedWord))
        {
            frequencies[cleanedWord]++;
        }
        else
        {
            frequencies[cleanedWord] = 1;
        }
    }
}

// Sort by count descending, then alphabetically ascending
var sortedWords = frequencies
    .OrderByDescending(pair => pair.Value)
    .ThenBy(pair => pair.Key);

foreach (var pair in sortedWords)
{
    Console.WriteLine($"{pair.Key}: {pair.Value}");
}