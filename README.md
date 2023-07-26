# Założenia

W języku programowania Golang napisz program, który będzie:
1. Wysyłał żądanie GET na zadany host;
2. Mierzył czas od wysłania żądania do czasu otrzymania odpowiedzi;
3. Sprawdzał kod odpowiedzi HTTP;
4. Sprawdzał czy odpowiedź to JSON (Content-Type);
5. Walidował czy JSON z odpowiedzi ma prawidłową składnię;
6. Wykonywał 10 sprawdzeń, co 5 sekund.
Wynik każdego sprawdzenia powinien być raportowany w osobnym wierszu na konsoli oraz dopisywany do pliku log.txt i zawierać informację o czasie zdarzenia oraz dane z punktów 1-6.
[Zadanie dodatkowe] Program ma sprawdzić czy ceny EUR w 100 ostatnich notowaniach mieściły się w zakresie od 4,5 do 4,7 zł, jeśli się nie mieściły to wykazać w jakich dniach.
Host: http://api.nbp.pl/api/exchangerates/rates/a/eur/last/100/?format=json
X: 10
Y: 5