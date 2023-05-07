# PokerCalc

Calculates the probability of winning a poker hand directly from the command line blazingly fast.

## What is PokerCalc?

PokerCalc is a command-line tool designed to help poker players quickly calculate their chances of winning. Since a full-blown computer is not allowed at the table, PokerCalc can be run on a smartphone using an app like Termux to execute simple commands for calculating the probability of having a winning hand. PokerCalc is textable when using Telegram. This means that you can text your hand to a Telegram bot and get the results back in seconds. For your poker table this just looks like normal texting with a relative or friend or whatever. PokerCalc is fast. It can calculate the probability of winning a hand with 5 players and 5 cards on the table in less than 0.1 seconds on a modern smartphone.

## How to use PokerCalc?

To use PokerCalc, provide your hand using the following notation for suits: black spades (BS), red hearts (RH), black diamonds (BD), and red clubs (RC). For example, a hand with a king of black spades and a 10 of black diamonds would be entered as "BSK BD10". The first argument should always be the number of players (between 2 and 8), followed by your hand. After that, you can optionally enter 0 to 5 community cards that are already face-up on the table. Based on this information, PokerCalc will calculate your chances of winning.



## Examples
1. A hand with a king of black spades and a 10 of black diamonds, with 3 players and 3 cards on the table.
```
pokerc 3 "BSK BD10" "RC9 RH3 BDA"
```
2. A hand with a jack of red hearts and a 10 of black diamonds, with 5 players and 0 cards on the table.
```
pokerc 5 "RHJ BD10"
```
3. A hand with an Ace of red hearts and an Ace of black spades, with 2 players and 5 cards on the table.
```
pokerc 2 "RHA BSA" "RH3 RC3 BD3 BDA BSK"
```
3. A hand with an queen of red hearts and an Ace of black spades, with 3 players and 5 cards on the table.
```
pokerc 3 "RHQ BSA" "RH3 RC3 BD3 BDA BSK"
```

## Gaps

PokerCalc currently does not check for double card entries in the players own hand. This means that if you enter something like "RH3 RH3" as your hand, PokerCalc will not complain. 

## Download and Run
You can download the appropriate binary for your operating system from the dist folder in this repository. The available binaries are:

- `pokerc_macos` for x86-64 macOS
- `pokerc_linux` for x86-64 Linux
- `pokerc_windows.exe` for x86-64 Windows
- `pokerc_macos_arm` for ARM macOS (Apple M1)
- `pokerc_linux_arm` for ARM Linux
- `pokerc_windows_arm.exe` for ARM Windows (Surface Pro X or similar)


## Running the Binary
### macOS and Linux

Download the binary file for your operating system from the dist folder.
Open a terminal window and navigate to the directory where you downloaded the binary.
Make the binary executable by running chmod +x <binary_name>, replacing <binary_name> with the appropriate file name (e.g., pokerc_macos or pokerc_linux).
Run the binary by executing ./<binary_name> followed by the command line arguments, as described in the "How to use PokerCalc?" section.

### Windows

Download the .exe file for your operating system from the dist folder.
Open a Command Prompt or PowerShell window and navigate to the directory where you downloaded the .exe file.
Run the binary by executing .\<binary_name>.exe followed by the command line arguments, as described in the "How to use PokerCalc?" section.