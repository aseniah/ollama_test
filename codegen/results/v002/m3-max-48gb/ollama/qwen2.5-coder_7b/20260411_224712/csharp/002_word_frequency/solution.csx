using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Text.RegularExpressions;

var inputFilePath = "input/text.txt";
if (!File.Exists(inputFilePath))
{
    Console.WriteLine("File not found.");
    return;
}

string[] lines = File.ReadAllLines(inputFilePath);
var words = new Dictionary<string, int>();

foreach (var line in lines)
{
    var regex = new Regex("[^a-zA-Z ]");
    var cleanLine = regex.Replace(line.ToLower(), "");
    var wordArray = cleanLine.Split(' ', StringSplitOptions.RemoveEmptyEntries);

    foreach (var word in wordArray)
    {
        if (!string.IsNullOrWhiteSpace(word))
        {
            if (words.ContainsKey(word))
            {
                words[word]++;
            }
            else
            {
                words[word] = 1;
            }
        }
    }
}

var sortedWords = words.OrderByDescending(w => w.Value).ThenBy(w => w.Key);

foreach (var word in sortedWords)
{
    Console.WriteLine($"{word.Key}: {word.Value}");
}