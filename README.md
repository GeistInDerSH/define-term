# Def

Application written in Golang to allow for the user to define a given term.

## Usage

```
def [-l language_code] [-s] word
```
### Flags

#### -l
The `-l` flag allows the user to specify the language of the word that they would like defined. If this flag is not provided, it defaults to 'en'. All language codes and their language are specified below.

##### Language Codes
| Code | Language |
| ---- | -------- |
| ar   | Arabic |
| de   | German |
| en   | English |
| es   | Spanish |
| fr   | French |
| hi   | Hindi |
| it   | Italian |
| ja   | Japanese |
| ko   | Korean |
| pt-BR| Brazilian Portuguese |
| ru   | Russian |
| tr   | Turkish
| zh-CN| Chinese (Simplified) |

#### -s
If the `-s` flag is enabled, then synonyms will be shown for each definition. By default, this is set to false.
