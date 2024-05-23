import puppeteer from 'puppeteer';

const options = {
    args: [
      '--no-sandbox',
      '--disable-setuid-sandbox',
      '--disable-dev-shm-usage',
      '--disable-accelerated-2d-canvas',
      '--no-first-run',
      '--no-zygote',
      '--single-process',
      '--disable-gpu'
    ],
    headless: true
  }

export async function ScrapeData(url) {

	const browser = await puppeteer.launch(options);
	const page = await browser.newPage();
	await page.setDefaultNavigationTimeout(0);
	await page.goto(url);

	// Wait for the page to load
	await new Promise(resolve => setTimeout(resolve, 10000)); // CHECK HOW THIS DEPENDS ON YOUR INTERNET SPEED

	// Get the element handle
	const elementHandles = await page.$$('.transit-result-item__footer'); 

	// Get the data
	const data = await Promise.all(elementHandles.map(async (element) => {
	return await element.evaluate(el => el.innerText);
	}));

	// Close the browser
	await browser.close();

	// Split the data into an array of arrays
	const split_data = data.map((element) => element.split(' '));

	// get every fourth element in the array
	const prices = split_data.map((element) => element[3]);
	//console.log(prices);

	// keep only the numbers in the array
	const array = prices.map((element) => {
		if (element) {
			return element.replace(/[^0-9]/g, '');
		} else {
			return ''; // or any default value you want to assign when element is undefined
		}
	});

	// remove empty strings from the array
	const numbersArray = array.map(item => item === '' ? '10000' : item);

	return Math.min.apply(null,numbersArray);
}


