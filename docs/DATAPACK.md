# What's a datapack?

Datapacks are Gophercraft's method for loading data. They fulfil

All datapacks are essentially folders containing text files. The folder can be archived as a ZIP, but if so it must use forward-slashes '/'.

The most important datapacks are **Base Datapacks.** These are extracted from a game content archive using the Gophercraft Wizard.

The base datapack !db.zip contains fundamental information about a game client. Without this you won't get very far. It is a relatively small zip file.

The base datapack !maps.zip contains bulk map geometry, and is of course much larger. You don't need to extract this, but when you do, creature AI gets smarter and your world becomes more robust against cheaters.

You'll also need a **Content Datapack**. These packs contain creature, item and gameobject templates, spawn positions and scripts.

[Content datapack for Alpha](https://github.com/Gophercraft/datapack-alpha)

[Content datapack for Vanilla](https://github.com/Gophercraft/datapack-vanilla)

You can clone or download these repositories as zips and put them in your world's Datapacks folder.

# Text files

Gophercraft makes use of [a minimal text format inspired by JSON](https://github.com/Gophercraft/text).

## Pack.txt

All datapacks contain a Pack.txt, a file which contains basic metadata.

Example:

```c
{
  Name "Name your pack"
  Author "Credit your authors here"
  Description "Yada yada tell us about what this does"
  MinCoreVersion 0.7.1 // the current Gophercraft core version this is tested against.
  PackVersion 1
  URL "put URL to project home (optional)"
  Depends
  {
    "Name of a pack this pack depends on" 
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



