using System;
using System.IO;
using System.Text.Json;
using System.Text.Json.Linq;

class Program
{
    static void Main(string[] args)
    {
        string filePath = "input/text.txt";
        string[] lines = File.ReadAllLines(filePath);

        var wordCount = new Dictionary<string, int>();

        foreach (var line in lines)
        {
            string[] words = line.ToLower().Split(new[] { ' ', '.', ',', ':', ';', '?', '!', '(', ')', '[', ']', '{', '}', '-', '_' }, StringSplitOptions.RemoveEmptyEntries);

            foreach (var word in words)
            {
                var charPattern = new Regex(@"[a-zA-Z]");
                string cleanedWord = charPattern.Replace(word, string.Empty);
                wordCount[cleanedWord] = wordCount.ContainsKey(cleanedWord) ? wordCount[cleanedWord] + 1 : 1;
            }
        }

        string[] sortedWords = wordCount.OrderByDescending(w => w.Value).ThenBy(w => w.Key).ToArray();

        foreach (var entry in sortedWords)
        {
            Console.WriteLine($"{entry.Key}: {entry.Value}");
        }
    }
}