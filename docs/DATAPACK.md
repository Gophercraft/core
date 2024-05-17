# What's a datapack?

All datapacks are essentially folders containing text files. The folder can be archived as a ZIP, but if so it must use forward-slashes '/'.

The most important datapacks are **Base Datapacks.** These are extracted from a game content archive using the Gophercraft Wizard.

The base datapack `!db.zip` contains fundamental information about a game client. Without this you won't get very far. It is a relatively small zip file.

The base datapack `!maps.zip` contains bulk map geometry, and is of course much larger. You don't need to extract this, but when you do, creature AI gets smarter and your world becomes more robust against cheaters.

You'll also need a **Content Datapack**. These packs contain creature, item and gameobject templates, spawn positions and scripts.

[Content datapack for Alpha](https://github.com/Gophercraft/datapack-alpha)

[Content datapack for Vanilla](https://github.com/Gophercraft/datapack-vanilla)

You can clone or download these repositories as zips and put them in your world's Datapacks folder.

# Text files

Gophercraft makes use of [a custom text format inspired by JSON and CSV](https://github.com/Gophercraft/text).

## Pack.txt

All datapacks contain a Pack.txt, a file which contains basic metadata.

Example:

```c
{
  ID "example_pack"
  Name "Name your pack"
  Version 1
  MinimumCoreVersion 0.7.1 
  Description "Yada yada tell us about what this does"
  Repository "put URL to project Git repository or ZIP file (optional)"
  Authors
  {
    "Credit your authors here"
  }
  Dependencies
  {
    {
      ID "put the ID of the needed datapack"
      MinimumVersion 1
    }
  }
  OverrideTables
  {
    "Name of table that you want to override"
    "Override means to replace every previous record"
  }
}
```

## Text Databases

Datapacks can contain a `DB/` folder, which holds your text databases. Imagine a db record structure like this:

```go
type RecordExample struct {
  ID uint32
  Text string
}
```

This could be represented in one file (DB/RecordExample.txt):

```c
{
  ID 1
  Text One
}
{
  ID 2
  Text Two
}
```

Alternatively you can split this file into two files with arbitrary names, just by placing them in a folder with the appropriate record name:

(DB/RecordExample/One.txt)
```c
{
  ID 1
  Text One
}
```

(DB/RecordExample/Two.txt)
```c
{
  ID 2
  Text Two
}
```



