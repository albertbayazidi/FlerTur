import { Result } from 'postcss';
import { ScrapeData } from './Scraper.js';


const stations = new Map([
    ["Oslo%20S", ["3A59872", "59.910357", "10.753051"]],
    ["Trondheim%20S", ["3A59977", "63.436279", "10.399123"]],
    ["Lillehammer%20stasjon", ["3A420", "61.114912", "10.461479"]],
    ["Stavanger%20stasjon", ["3A61291", "58.966568", "5.732616"]],
    ["Bergen%20stasjon", ["3A59983", "60.390434", "5.333511"]],
    ["Lillestrøm%20stasjon", ["3A62339", "59.952915", "11.045364&"]],
    ["Fredrikstad%20stasjon", ["3A58794", "59.208805", "10.950282"]],
    ["Kristiansand%20stasjon", ["3A61608", "58.14559", "7.988067"]],
    ["Kongsvinger%20stasjon", ["3A58808", "60.187497", "12.004063"]],
]);

const base_url = 'https://entur.no/reiseresultater?transportModes=rail%2Ctram%2Cbus%2Ccoach%2Cwater%2Ccar_ferry%2Cmetro%2Cflytog%2Cflybuss&date=';

const trip_mode = 'tripMode=oneway&walkSpeed=1.3&minimumTransferTime=120&timepickerMode=departAfter&allowFlexible=false&';

const start_id = 'startId=NSR%3AStopPlace%';

const stop_id = 'stopId=NSR%3AStopPlace%';

const days = 30;

const times = [];

const store_promis = [];

const store_urls = [];

async function process(start_loc,end_loc){
    
    function addDaysToDate(date) {
        const resultDate = new Date(date);
        resultDate.setDate(resultDate.getDate() + days);

        return resultDate;
    }

    function getDifferenceInDays(date1, date2){
        const parsedDate1 = new Date(date1);
        const parsedDate2 = new Date(date2);
    
        const differenceInMillis = Math.abs(parsedDate2 - parsedDate1);
        const differenceInDays = Math.ceil(differenceInMillis / (1000 * 60 * 60 * 24));
    
        return differenceInDays;
    }

    function convertDateToUrl(date){
        let data_intersept = 1717207199999
        let data_tangent = 86400000

        return data_intersept + date * data_tangent;
    }

    const core_date = '2024-06-01'
    const test_date = '2024-06-15'    //this variable should be gotten from the user

    var new_date = addDaysToDate(test_date); 

    var diff_in_day = getDifferenceInDays(core_date,new_date)

    for (let i = days; i > 0; i--){
        times.push(convertDateToUrl(diff_in_day - i));
    }

    for (let i = 0; i < days; i++){
        let url_kopi = base_url;
        url_kopi += times[i] + '&' + trip_mode;

        url_kopi += start_id + stations.get(start_loc)[0] + '&startLabel=' + start_loc + '&startLat=' + stations.get(start_loc)[1] + '&startLon=' + stations.get(start_loc)[2] + '&';

        url_kopi += stop_id + stations.get(end_loc)[0] + '&stopLabel=' + end_loc + '&stopLat=' + stations.get(end_loc)[1] + '&stopLon=' + stations.get(end_loc)[2];
        let new_url = url_kopi;
        store_urls.push(new_url);
        store_promis.push(ScrapeData(new_url));
    }

    const storage = await Promise.all(store_promis);

    return [storage, store_urls];
}
var start = performance.now()
await process('Oslo%20S','Trondheim%20S')
.then((result) => {console.log(result)})
.catch((err) => {console.log(err)});
var en = performance.now()
console.log("prosses tok ", en-start,"sek")