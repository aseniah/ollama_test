using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Text;

// Read the content of the file
string filePath = "input/text.txt";
if (!File.Exists(filePath))
{
    return;
}

string text = File.ReadAllText(filePath).ToLower();

// Extract words consisting only of letters
List<string> words = new List<string>();
StringBuilder currentWord = new StringBuilder();

foreach (char c in text)
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
// Catch the last word if the file doesn't end with a non-letter character
if (currentWord.Length > 0)
{
    words.Add(currentWord.ToString());
}

// Count the frequency of each word
Dictionary<string, int> frequencies = new Dictionary<string, int>();
foreach (string word in words)
{
    if (frequencies.ContainsKey(word))
    {
        frequencies[word]++;
    }
    else
    {
        frequencies[word] = 1;
    }
}

// Sort by count descending, then by word ascending
var sortedWords = frequencies
    .OrderByDescending(kvp => kvp.Value)
    .ThenBy(kvp => kvp.Key);

// Output the results
foreach (var kvp in sortedWords)
{
    Console.WriteLine($"{kvp.Key}: {kvp.Value}");
}