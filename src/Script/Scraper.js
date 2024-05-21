import puppeteer from 'puppeteer';

export async function ScrapeData(url) {

	const browser = await puppeteer.launch();
	const page = await browser.newPage();
	await page.goto(url);

	// Wait for the page to load
	await new Promise(resolve => setTimeout(resolve, 7500)); // CHECK HOW THIS DEPENDS ON YOUR INTERNET SPEED

	// Get the element handle
	const elementHandles = await page.$$('.transit-result-item__footer'); 

	// Get the data
	const data = await Promise.all(elementHandles.map(async (element) => {
	return await element.evaluate(el => el.innerText);
	}));

	// Close the browser
	await browser.close();
	return data;
}


