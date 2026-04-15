Read `input/input.csv` and produce a JSON array to stdout.
The file `input/expected_format.json` shows the expected output —
use it to infer the transformation rules and output format.
Calculate ages as of July 1, 2025.
Do not output anything other than the JSON array.

--- input/expected_format.json ---
[
  {
    "FirstName": "John",
    "LastName": "Lennon",
    "Birthday": "1940-10-09",
    "Age": 40,
    "Relatives": [
      {
        "FirstName": "Alfred",
        "LastName": "Lennon",
        "Relationship": "Father"
      },
      {
        "FirstName": "Julia",
        "LastName": "Stanley",
        "Relationship": "Mother"
      }
    ]
  },
  {
    "FirstName": "James",
    "LastName": "McCartney",
    "Birthday": "1942-06-18",
    "Age": 83,
    "Relatives": [
      {
        "FirstName": "Jim",
        "LastName": "McCartney",
        "Relationship": "Father"
      },
      {
        "FirstName": "Mary",
        "LastName": "McCartney",
        "Relationship": "Mother"
      },
      {
        "FirstName": "Mike",
        "LastName": "McGear",
        "Relationship": "Brother"
      }
    ]
  },
  {
    "FirstName": "Ringo",
    "LastName": "Starr",
    "Birthday": "1940-07-07",
    "Age": 84,
    "Relatives": [
      {
        "FirstName": "Richard",
        "LastName": "Starkey",
        "Relationship": "Father"
      },
      {
        "FirstName": "Elsie",
        "LastName": "Gleave",
        "Relationship": "Mother"
      },
      {
        "FirstName": "Marie",
        "LastName": "Maguire",
        "Relationship": "Sister"
      }
    ]
  },
  {
    "FirstName": "George",
    "LastName": "Harrison",
    "Birthday": "1943-02-25",
    "Age": 58,
    "Relatives": [
      {
        "FirstName": "Harold",
        "LastName": "Harrison",
        "Relationship": "Father"
      },
      {
        "FirstName": "Louise",
        "LastName": "French",
        "Relationship": "Mother"
      },
      {
        "FirstName": "Peter",
        "LastName": "Harrison",
        "Relationship": "Brother"
      },
      {
        "FirstName": "Louise",
        "LastName": "Harrison",
        "Relationship": "Sister"
      }
    ]
  }
]

--- input/input.csv ---
Name,Birthday,Died,Father,Mother,Brother,Sister
John Winston Lennon,10/9/1940,12/8/1980,Alfred Lennon,Julia Stanley,null,null
James Paul McCartney,6/18/1942,null,Jim McCartney,Mary McCartney,Mike McGear,null
Ringo Starr,7/7/1940,null,Richard Starkey,Elsie Gleave,null,Marie Maguire
George Harrison,2/25/1943,11/29/2001,Harold Harrison,Louise French,Peter Harrison,Louise Harrison
