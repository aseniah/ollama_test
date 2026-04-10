using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Text;

// Read the file content
string content = File.ReadAllText("input/text.txt");
content = content.ToLower();

// Count word frequencies
var wordFrequencies = new Dictionary<string, int>();
StringBuilder currentWord = new StringBuilder();

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
            string word = currentWord.ToString();
            if (wordFrequencies.ContainsKey(word))
            {
                wordFrequencies[word]++;
            }
            else
            {
                wordFrequencies[word] = 1;
            }
            currentWord.Clear();
        }
    }
}

// Handle last word
if (currentWord.Length > 0)
{
    string word = currentWord.ToString();
    if (wordFrequencies.ContainsKey(word))
    {
        wordFrequencies[word]++;
    }
    else
    {
        wordFrequencies[word] = 1;
    }
}

// Sort by count descending, then alphabetically ascending
var sortedWords = wordFrequencies
    .OrderByDescending(kvp => kvp.Value)
    .ThenBy(kvp => kvp.Key)
    .ToList();

// Output results
foreach (var kvp in sortedWords)
{
    Console.WriteLine($"{kvp.Key}: {kvp.Value}");
}