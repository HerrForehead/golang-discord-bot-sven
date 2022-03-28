# golang-discord-bot-sven

## Beschrijving
Een Discord bot gemaakt in Go. Deze Discord bot neemt verschillende commandos en roept hierna data op van een API die deze dan ook teruggeeft. De bot kan worden geconfigureerd doormiddel van de config.json.

## Voorbeeld
https://i.imgur.com/fpMEpJ5.gif

## Hoe te gebruiken

Ten eerste heb je een Discord bot api key nodig. Deze kan je krijgen door op https://discord.com/developers in te loggen met je account. Hier maak je dan een nieuwe
applicatie aan. Als je deze hebt aangemaakt, klik je op de app en ga je naar het tabje "bot". Hier maak je een bot aan. Als deze is aangemaakt, kan je hier de API key
voor je bot krijgen. Zet deze key onder het kopje "Token" in je config.json, tussen de haakjes.
Daarna nodig je deze bot uit naar je Discord server.

Als je de text2img command wil gebruiken, moet je op https://deepai.org/ een account maken en op je dashboard de API key kopiëren naar je config.json, op dezelfde manier
als bij je Discord token. Let wel op dat je, indien nodig, zelf credits moet kopen en toevoegen om de command te gebruiken.

Als de waarden zijn ingevoerd, moet je de bot builden. Dit doe je door in je command line/terminal het volgende commando te gebruiken:
```
go build -o <de naam die jij wil geven>.exe
```
Als je
```
-o <naam>.exe
```
weg laat, default de bot naar de standaard naam.

Hierna kan je de bot als een gewone .exe openen in je terminal en zal de bot activeren met de text "Bot is now running.  Press CTRL-C to exit."
