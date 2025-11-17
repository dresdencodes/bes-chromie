package javascript

func JSTestBase64Images() string {
	return `
		(() => {
			const imgs = Array.from(document.querySelectorAll("img"))
				.filter(i => i.src.startsWith("data:image/"));

			if (imgs.length === 0) return "OK";

			return Promise.all(imgs.map(img => {
				return new Promise(resolve => {
					const t = new Image();

					t.onload = () => resolve(true);
					t.onerror = () => resolve(false);

					// a small timeout ensures we don't hit sync load edge cases
					setTimeout(() => { t.src = img.src; }, 0);
				});
			})).then(all => all.every(Boolean) ? "OK" : "FAIL");
		})();
	`
}

func JSEnsure() string {
	return `

		if (waitForQueueAsync===undefined) {
			function waitForQueueAsync() {
				return Promise.resolve();
			}
		}

		(async () => {
			console.log("Start");

			await waitForQueueAsync();

			console.log("Runs after stack is clear");
		})();

		"OK"
	`
}

func JSSetFrame(f string) string {

	return `
		window.context.setFrame(`+f+`);
		"OK";
	` 

}