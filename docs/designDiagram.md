# System design interview

## Oppgave: Flertur
Lag et tjeneste som gir brukere evnen til å finne den billigste tog ruten mellom to byer. Må være letterene enn å bruke Entur

1. define problem space
    - scrapene en nettside for tog turerer og priser
    - ta de 8 størrste byene
    - må være raskt 

2. design at high-level
API 

GET https:flertur/api/v1/rute?start=<startStation>end=<endStation>
[
  {
    "startStation": "Oslo s",
    "endStation": "TrondheimS",
    "retrievalTime": "2026-03-06T22:26:25.000Z",
    "pageDataResults": [
      {
        "duration": "7h 20m",
        "startTime": "2026-03-07T08:00:00.000Z",
        "price": 1000,
        "numberOfTrains": 2,
        "trainIds": ["R14", "R2"],
        "url": "url.com"
      },
      {
        "duration": "6h 55m",
        "startTime": "2026-03-07T09:00:00.000Z",
        "price": 229,
        "numberOfTrains": 1,
        "trainIds": ["R46"],
        "url": "url.com"
      }
    ]
  }
]

GET https:flertur/api/v1/rute?start=<startStation>end=<endStation>

high-level diagram
[image](overview-diagram.png)

