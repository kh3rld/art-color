# ASCII Art Color

## Description
This is a simple command-line tool written in Go. It takes a string input entered by the user then returns its graphical representation using the ASCII art characters from the specified banner file.

## Usage
1. **Clone Repository**: Clone this repository to your local machine.
2. **Install go**: Ensure you have Go installed on your machine.
3. **Run Program**: Execute the program by providing the flag, input string, and banner file as command-line arguments. 
    ```bash
    go run . --color=<color> <substring to be colored> "something"
    ```

    Example :

    ```bash
    go run . --color=red kit "a king kitten have kit"
    ```
    This will return the colored graphical representation of the string "kit".

    ```ruby
     _              _   _          
    | |            | | | |         
    | |__     ___  | | | |   ___   
    |  _ \   / _ \ | | | |  / _ \  
    | | | | |  __/ | | | | | (_) | 
    |_| |_|  \___| |_| |_|  \___/  
                                
                                                          
    ```   
Additionally, the program can run with a single string argument even when the user doesn't specify the flag.
```bash
    go run . hello 
```
```bash
     _              _   _          
    | |            | | | |         
    | |__     ___  | | | |   ___   
    |  _ \   / _ \ | | | |  / _ \  
    | | | | |  __/ | | | | | (_) | 
    |_| |_|  \___| |_| |_|  \___/  
                                
                                                          
```   
In this case, the result will be written to the default output file then automatically displayed on the terminal.

### Note
The main_test.go file tests the whole program while ascii_art_test.go tests the functions. To be able to run the main_test.go file successfully you have to first run the main.go file, with "hello" as your input string and "output.txt" as the value for the flag. This will help generate the output file which will be checked by the main_test.go file.

## Collaborators

* **[Tomlee Abila](https://github.com/Tomlee-abila/)**
* **[Kherld Hussein](https://github.com/kherldhussein/)**