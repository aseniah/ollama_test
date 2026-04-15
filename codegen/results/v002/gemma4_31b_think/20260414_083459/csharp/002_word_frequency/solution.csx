using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Text;

// Read the input file
string filePath = "input/text.txt";
if (!File.Exists(filePath))
{
    return;
}

string content = File.ReadAllText(filePath);

// Split by whitespace to get potential word tokens
string[] tokens = content.Split(new char[] { ' ', '\t', '\n', '\r' }, StringSplitOptions.RemoveEmptyEntries);

Dictionary<string, int> frequencies = new Dictionary<string, int>();

foreach (string token in tokens)
{
    // Convert to lowercase and strip all characters that are not letters
    StringBuilder sb = new StringBuilder();
    foreach (char c in token.ToLower())
    {
        if (char.IsLetter(c))
        {
            sb.Append(c);
        }
    }

    string word = sb.ToString();

    // Only count if the resulting word is not empty
    if (!string.IsNullOrEmpty(word))
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
}

// Sort by count descending, then by word alphabetically ascending
var sortedList = frequencies
    .OrderByDescending(pair => pair.Value)
    .ThenBy(pair => pair.Key);

// Output the results in the format 'word: count'
foreach (var pair in sortedList)
{
    Console.WriteLine($"{pair.Key}: {pair.Value}");
}