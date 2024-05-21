import { ScrapeData } from './Scraper.js';


//ScrapeData('https://entur.no/reiseresultater?transportModes=rail%2Ctram%2Cbus%2Ccoach%2Cwater%2Ccar_ferry%2Cmetro%2Cflytog%2Cflybuss&date=1717725600000&tripMode=oneway&walkSpeed=1.3&minimumTransferTime=120&timepickerMode=departAfter&allowFlexible=false&startId=NSR%3AStopPlace%3A420&startLabel=Lillehammer%20stasjon&startLat=61.114912&startLon=10.461479&stopId=NSR%3AStopPlace%3A59872&stopLabel=Oslo%20S&stopLat=59.910357&stopLon=10.753051')
//.then(data => console.log(data), err => console.error(err));

const stations = new Map([
    ["Oslo%20S", ["3A59872", "59.910357", "10.753051"]],
    ["Trondheim%20S", ["3A59977", "63.436279", "10.399123"]],
    ["Lillehammer%20stasjon", ["3A420", "61.114912", "10.461479"]],
    ["Stavanger%20stasjon", ["3A61291", "58.966568", "5.732616"]],
    ["Bergen%20stasjon", ["3A59983", "60.390434", "5.333511"]],
    ["Lillestrøm%20stasjon", ["3A62339", "59.952915", "11.045364&"]],
    ["Fredrikstad%20stasjon", ["3A58794", "59.208805", "10.950282"]],
    ["Kristiansand%20stasjon", ["3A61608", "58.14559", "7.988067"]]
]);


function process(start_loc,end_loc){
    const days = 30;

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

    var today = new Date();
    var core_date = '2024-06-01'
    var test_date = '2024-06-15'    //TEST VARIABLE

    var new_date = addDaysToDate(test_date); // CHANGE TO TODAY

    var diff_in_day = getDifferenceInDays(core_date,new_date)

    const times = [];

    console.log('updated day', new_date);
    console.log('days between',diff_in_day);

    for (let i = days-1; i > 0; i--){
        times.push(convertDateToUrl(diff_in_day - i));
    }

    let base_url = 'https://entur.no/reiseresultater?transportModes=rail%2Ctram%2Cbus%2Ccoach%2Cwater%2Ccar_ferry%2Cmetro%2Cflytog%2Cflybuss&date=';

    let trip_mode = 'tripMode=oneway&walkSpeed=1.3&minimumTransferTime=120&timepickerMode=departAfter&allowFlexible=false&';

    let start_id = 'startId=NSR%3AStopPlace%';

    let stop_id = 'stopId=NSR%3AStopPlace%';

    for (let i = 0; i < times.length; i++){
        let url_kopi = base_url;
        url_kopi += times[i] + '&' + trip_mode;

        url_kopi += start_id + stations.get(start_loc)[0] + '&startLabel=' + start_loc + '&startLat=' + stations.get(start_loc)[1] + '&startLon=' + stations.get(start_loc)[2] + '&';

        url_kopi += stop_id + stations.get(end_loc)[0] + '&stopLabel=' + end_loc + '&stopLat=' + stations.get(end_loc)[1] + '&stopLon=' + stations.get(end_loc)[2];
        var new_url = url_kopi;
    }
    console.log(new_url)


}
process('Lillehammer%20stasjon','Fredrikstad%20stasjon');