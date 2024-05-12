# Gophercraft Wizard Commands

1. [Connection](#Connection)
    1. [connect](#connect)
    2. [disconnect](#disconnect)
    3. [forget](#forget)
2. [Logging in](#Logging-in)
    1. [login](#login)
    1. [session](#session)
    2. [logout](#logout)
2. [Account management](#account-management)
    1. [account status](#account-status)
    2. [account create](#account-create)
3. [Administration](#Administration)
    1. [admin account](#admin-account)
        1. [admin account lock](#admin-account-lock)
        2. [admin account ban](#admin-account-ban)
        3. [admin account suspend](#admin-account-suspend)
        4. [admin account unlock](#admin-account-unlock)
        5. [admin account unban](#admin-account-unban)
        6. [admin account unsuspend](#admin-account-unsuspend)
    2. [admin gameaccount](#admin-gameaccount)
        1. [admin gameaccount ban](#admin-gameaccount-ban)
        2. [admin gameaccount suspend](#admin-gameaccount-suspend)
        3. [admin gameaccount unban](#admin-gameaccount-unban)
        4. [admin gameaccount unsuspend](#admin-gameaccount-unsuspend)
4. [World management](#world-management)
    1. [world create](#world-create)
# Connection

## connect

Connect to a Home server running at `localhost:32777`

```bash
gophercraft connect
```

Or connect to a known server host:port: 

```bash
gophercraft connect some-gophercraft-instance.your-host.org:32777
```

## disconnect

Disconnect from whatever server you're currently connected to:

```bash
gophercraft disconnect
```

## forget

Forget the fingerprint of a server you've previously connected to:

```
gophercraft forget localhost:32777
```

# Logging in

## login

You must be [connected](#connect) to log in.

```bash
gophercraft login
```

## logout

You must be [logged in](#login) to log out.

If connected, this command will attempt to invalidate your web token server-side.

```bash
gophercraft logout
```

## session

You must be [logged in](#login) to generate a session file.

Gophercraft has a special way of handling VERY OLD clients that use .ses files.

While other servers require you to enter your credentials into a file (Which is uploaded to the world server in PLAINTEXT), this command requests a login ticket from the home server.

The world server only sees your ticket in plaintext, and never gains full access to your account.

```bash
gophercraft session
```

# Account Management

## account create

You must be logged in as a moderator or administrator to create an account like this.

```
gophercraft account create
```

## account status

Get the status of your logged in account:

```bash
gophercraft account status
```

If you're an admin, you can also view the status of another account you know the ID of, in this case account #1:

```bash
gophercraft account status 1
```

# Administration

You must be logged in as a moderator or administrator to use tools such as these.

## admin account

These accounts modify the state of an account. Modifying this has consequences for everything under a specific account username string/ID number.

If you only want to affect a game account belonging to a particular user, refer to [admin gameaccount](#admin-gameaccount).

## admin account lock

To prevent an account's ability to log into the game indefinitely, but not to use the home server, in this example account #1:

```bash
gophercraft admin account lock 1
```

## admin account ban

To remove an account's ability to log into the home server indefinitely, in this example account #1:

```bash
gophercraft admin account ban 1
```

## admin account suspend

To suspend an account's ability to log into the game, but still allowing them to use the home server, and expiring their suspension after a fixed duration.

In this example, account #1 is suspended, with the suspension expiring after a period of 1 month (730 hours):

```bash
gophercraft admin account suspend 1 730h
```

## admin account unlock

To remove the lock state of an account, in this example account #1:

```bash
gophercraft admin account unlock 1
```

## admin account unban

To lift the ban on an account, in this example account #1:

```bash
gophercraft admin account unban 1
```

## admin account unsuspend

To lift the suspension of an account, in this example account #1:

```bash
gophercraft admin account unsuspend 1
```

## admin gameaccount

These accounts modify the state of a game account. Note that accounts and gameaccounts are not the same thing!

## admin gameaccount ban

To remove a game account's ability to log into game indefinitely, in this example game account #2:

```bash
gophercraft admin gameaccount ban 2
```

## admin gameaccount suspend

To suspend a game account's (game account #2) ability to log into the game for a fixed duration

In this example, game account #1 is suspended, with the suspension expiring after a period of 1 month (730 hours):

```bash
gophercraft admin gameaccount suspend 2 730h
```

## admin gameaccount unban

To lift the ban on a game account, in this example game account #2:

```bash
gophercraft admin gameaccount unban 2
```

## admin gameaccount unsuspend

To lift the suspension of a game account, in this example game account #2:

```bash
gophercraft admin gameaccount unsuspend 2
```

# World management

## world create

This command creates a new world and registers it with the home server.

It is highly recommended, though not strictly necessary (if you know what you are doing), to supply a game volume for the wizard to extract from.

This will auto-detect the build ID number of your game, as well as save datapacks necessary to run the game.

Choosing to extract may take a very long time.

```bash
gophercraft world create --realm-name "World One" --game-volume /media/device/game/
```

If you'd rather supply the needed datapacks yourself, you can use a different command:

```bash
gophercraft world create --realm-name "World One" --build $BUILD_ID_NUMBER
```
