# Requirements

- A game directory (5875-12340)
- 1GB free RAM (recommended)

# Setting up a home server

### 1. Create a configuration directory and go into it

By default, Gophercraft Home will read a configuration from whichever directory it starts in.

Make a directory like this:

```bash
mkdir ~/gophercraft_home/
cd ~/gophercraft_home
```

### 2. Create a home server configuration

Use the `gophercraft` command to generate a default configuration.

```bash
gophercraft home create
```

Take a look in your current directory. If you see `home.txt, key.pem, cert.pem`, you successfully created a home server.

For the first run, it's highly recommended to look through `home.txt` and comment out unneeded services to minimize attack surface.

### 3. Register an offline admin account

A newly created home will have no admin accounts. You will need an admin account to administrate your home server.

*Note that If the home server is running and using the database, this command **WILL NOT WORK**. This is because you're directly modifying the database, which may already be opened by the home server process.*

To register an admin account while the home server is running, you must `gophercraft login` as admin and then run `gophercraft account create`.


```bash
gophercraft_home register
```
You will be prompted to enter a username and password through standard input.

You can use this command whenever you need to reset an account password.

### 4. Launch the home server

To run the home server, you just need to run the "gophercraft_home" executable while inside the configuration directory.

```bash
gophercraft_home
```

### 5. Connect to the home server

In this step, you create a connection to an external home server.

```bash
gophercraft connect localhost:32777
```

If you're just connecting to localhost on the default port, you can omit the address:

```bash
gophercraft connect
```

The first time you connect to a host you will be asked whether you trust the public key fingerprint.

### 6. Log in to your home server

Now that you have a trusted home server, you can attempt to log in using the account details you registered with `gophercraft_home register`

```
gophercraft login
```

# Setting up a world server

A game directory is required to set up a world server.

### 1. Log in (optional)

If you are setting up a world server on a different device/user profile than your home server, you will need to connect and log in to your admin account again.

```
gophercraft connect <external IP address of home server>:32777
gophercraft login
```

### 1. Create a directory and go into it

Like the home server, a world server reads configuration from its working directory.

```bash
mkdir ~/gophercraft_world_1/
cd ~/gophercraft_world_1
```

### 2. Create a world server configuration

It is highly recommended, though not strictly necessary (if you know what you are doing), to supply a game volume for the wizard to extract from. This will auto-detect the build ID number of your game, as well as save datapacks necessary to run the game. This may take a very long time. (Highly recommended)

```bash
gophercraft world create --realm-name "World One" --game-volume /media/device/game/
```

If you'd rather supply the needed datapacks yourself, you can use a different command:

```bash
gophercraft world create --realm-name "World One" --build $BUILD_ID_NUMBER
```

- `--realm-name`
    - The name of the world server as appears in the realmlist
- `--build`
    - The version number of the client you want the server to emulate
- `--game-volume`
    - the folder containing the game executable and data files

As Gophercraft receives updates, the format of the datapacks will evolve.

This will inevitably break compatibility with previously generated datapacks, so it's sometimes useful to refresh only the "base" datapacks (datapacks such as !db.zip, !maps.zip)

```bash
gophercraft world refresh --game-volume /media/device/game
```

### 3. Install additional required datapacks

The world server also needs what are called "content" datapacks. These contain information about:

- NPCs/NPC dialogue
- Quests
- Items
- Spawn items
- Teleport locations

A few are hosted here on the Gophercraft GitHub for different versions of the game:

- [Datapack for 0.5.3.3368 (Alpha)](https://github.com/Gophercraft/datapack-alpha)
- [Datapack for 1.12.1.5875 (Vanilla)](https://github.com/Gophercraft/datapack-vanilla)

To install these, you can `git clone` them directly into the `datapacks/` subdirectory inside your world server configuration directory.

You can also download the ZIP-archived version of them into the same directory.

### 4. Launch the world server

A correctly configured world server will continuously attempt to communicate back and forth with the home server, updating the realm server as necessary.

```
gophercraft_world
```



