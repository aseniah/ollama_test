#r "System.Linq"

using System;
using System.IO;
using System.Linq;
using System.Text;
using System.Collections.Generic;

string content = File.ReadAllText("input/text.txt").ToLower();

var words = new List<string>();
var currentWord = new StringBuilder();

foreach (char c in content)
{
    if (char.IsLetter(c))
        currentWord.Append(c);
    else if (currentWord.Length > 0)
    {
        words.Add(currentWord.ToString());
        currentWord.Clear();
    }
}

if (currentWord.Length > 0)
    words.Add(currentWord.ToString());

var result = words.GroupBy(w => w)
    .Select(g => new { Word = g.Key, Count = g.Count() })
    .OrderByDescending(x => x.Count)
    .ThenBy(x => x.Word);

foreach (var item in result)
{
    Console.WriteLine($"{item.Word}: {item.Count}");
}