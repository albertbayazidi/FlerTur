package rod_utils

import (
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/proto"
)

var containerSelector = ".transit-result__list__container"
var buttonSelector = "//button[contains(., 'Senere avganger') or contains(., 'Flere avganger') or contains(., 'Load more')]"

func captureLiElements(page *rod.Page) rod.Elements {
	liElements, err := page.Elements("li.transit-result-item")
	if err != nil {
		return rod.Elements{}
	}

	return liElements
}

func pressButton(button *rod.Element, page *rod.Page) bool {
	containers, err := page.Elements(containerSelector)
	if err == nil && len(containers) > 1 {
		return true
	}

	liElements := captureLiElements(page)
	initialCount := len(liElements)

	err = button.Click(proto.InputMouseButtonLeft, 1)
	if err != nil {
		return true
	}

	retries := 0
	for len(liElements) <= initialCount && retries < 20 {
		time.Sleep(250 * time.Millisecond)
		liElements = captureLiElements(page)

		containers, err = page.Elements(containerSelector)
		if err == nil && len(containers) > 1 {
			return true
		}

		retries++
	}

	containers, err = page.Elements(containerSelector)
	if err == nil && len(containers) > 1 {
		return true
	}

	return retries >= 20
}

func Crawler(page *rod.Page) {
	page.Timeout(15 * time.Second).WaitStable(1 * time.Second)

	for {
		hasButton, button, err := page.HasX(buttonSelector)
		if err != nil || !hasButton {
			break
		}

		shouldStop := pressButton(button, page)
		if shouldStop {
			break
		}
	}
}
