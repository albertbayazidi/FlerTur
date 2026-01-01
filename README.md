<p align='center'>

  <img src='public/logo.png'>
  
</p>

## Oversikt
[Fler TUR](https://bayazidi.xyz/) er en nettjeneste som prøver å finne de billigste billettprisene en måned fram i tid.
Verdiene som blir framstilt kan avvike litt. 

## Hvordan funker Fler Tur
Fler TUR henter automatisk inn reiseinformasjon fra EnTUR ved å sjekke ulike kombinasjoner av stasjoner. Dataene samles inn fortløpende og lagres, slik at de er tilgjengelige for brukerne uten at de selv trenger å gjøre noe.

Nettsiden henter automatisk reisedata for de neste to ukene. Siden det å hente informasjon direkte fra EnTUR kan være tidkrevende, sjekkes kun de sju største rutene i begge retninger. Dette gjør prosessen rask og effektiv

Tallene oppdateres én gang i timen. Det betyr at priser og tilgjengelighet i enkelte tilfeller kan avvike noe fra sanntidsinformasjonen hos EnTUR.

## Kjøre lokalt
Om man har forslag til forbedringer, kan man teste dem lokalt ved å gjøre følgende.
``` bash 
git clone https://github.com/albertbayazidi/FlerTur.git
cd FlerTur
npm/bun i
npm/bun run dev
```

<img src="https://cdn.buymeacoffee.com/buttons/default-orange.png" alt="Buy Me A Coffee" height="41" width="174" />

