using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Text;

string inputPath = "input/text.txt";

// Read the file content
string content = File.ReadAllText(inputPath);

// Normalize to lowercase
content = content.ToLowerInvariant();

// Use a dictionary to count frequencies
var frequencyMap = new Dictionary<string, int>();

// Iterate through each character to build words
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
            if (frequencyMap.ContainsKey(word))
            {
                frequencyMap[word]++;
            }
            else
            {
                frequencyMap.Add(word, 1);
            }
            currentWord.Clear();
        }
    }
}

// Handle the last word if the file doesn't end with a non-letter character
if (currentWord.Length > 0)
{
    string word = currentWord.ToString();
    if (frequencyMap.ContainsKey(word))
    {
        frequencyMap[word]++;
    }
    else
    {
        frequencyMap.Add(word, 1);
    }
}

// Sort the words: by count descending, then alphabetically ascending for ties
var sortedWords = frequencyMap
    .OrderByDescending kvp => kvp.Value)
    .ThenBy kvp => kvp.Key)
    .ToList();

// Output results
foreach (var kvp in sortedWords)
{
    Console.WriteLine($"{kvp.Key}: {kvp.Value}");
}